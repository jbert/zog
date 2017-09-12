package zog

import "fmt"

type Memory struct {
	buf       []byte
	debug     bool
	watches   Regions
	watchFunc func(addr uint16, old byte, new byte)
	readonly  Regions
}

func NewMemory(size uint16) *Memory {
	intSize := int(size)
	if intSize == 0 {
		intSize = 64 * 1024
	}
	m := &Memory{
		buf:      make([]byte, intSize),
		readonly: make([]Region, 0),
	}
	return m
}

func (m *Memory) SetDebug(debug bool) {
	m.debug = debug
}

func (m *Memory) SetWatchFunc(wf func(uint16, byte, byte)) {
	m.watchFunc = wf
}

func (m *Memory) Len() int {
	return len(m.buf)
}

func (m *Memory) Peek(addr uint16) (byte, error) {
	if int(addr) >= m.Len() {
		return 0, fmt.Errorf("Out of bounds memory read: %d", addr)
	}
	n := m.buf[addr]
	//	if m.debug || m.watches.contains(addr) {
	//		fmt.Printf("MEM: %04X -> %02X\n", addr, n)
	//	}
	return n, nil
}

func (m *Memory) Poke(addr uint16, n byte) error {
	// Poke to ROM is a NOP
	if m.readonly.contains(addr) {
		fmt.Printf("Deny RO POKE [%04X] [%02X]\n", addr, n)
		return nil
	}

	if int(addr) >= m.Len() {
		return fmt.Errorf("Out of bounds memory write: %d (%d)", addr, n)
	}
	if m.debug || m.watches.contains(addr) {
		//		fmt.Printf("MEM: %04X <- %02X\n", addr, n)
		m.watchFunc(addr, m.buf[addr], n)
	}
	m.buf[addr] = n
	return nil
}

func (m *Memory) Peek16(addr uint16) (uint16, error) {
	l, err := m.Peek(addr)
	if err != nil {
		return 0, err
	}
	h, err := m.Peek(addr + 1) // overflow correct
	if err != nil {
		return 0, err
	}
	return uint16(h)<<8 | uint16(l), nil
}

func (m *Memory) Poke16(addr uint16, nn uint16) error {
	if m.readonly.contains(addr) {
		return nil
	}

	l := byte(nn)
	h := byte(nn >> 8)
	err := m.Poke(addr, l)
	if err != nil {
		return err
	}
	err = m.Poke(addr+1, h)
	if err != nil {
		return err
	}
	return nil
}

func (m *Memory) Clear() {
	m.buf = make([]byte, int(m.Len()))
}

func (m *Memory) Copy(addr uint16, buf []byte) error {
	if int(addr)+len(buf) >= int(m.Len()) {
		panic(fmt.Sprintf("Can't load - base addr %04X length %04X memsize %04X", addr, len(buf), m.Len()))
	}
	for i := 0; i < len(buf); i++ {
		m.buf[addr+uint16(i)] = buf[i]
	}
	return nil
}

// Fetch a chunk of memory. Error if overflows end.
// Don't write to this please.
func (m *Memory) PeekBuf(addr uint16, size int) ([]byte, error) {
	if size <= 0 || size > 64*1024*1024 {
		return nil, fmt.Errorf("PeekBuf invalid size: %d", size)
	}
	if size+int(addr) > len(m.buf) {
		return nil, fmt.Errorf("PeekBuf invalid size+addr: %04x - %04x", addr, size)
	}
	return m.buf[addr : int(addr)+size], nil
}
