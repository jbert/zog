package zog

import "fmt"

type Memory struct {
	buf   []byte
	debug bool
}

func NewMemory(size uint16) *Memory {
	m := &Memory{
		buf: make([]byte, size),
	}
	return m
}

func (m *Memory) SetDebug(debug bool) {
	m.debug = debug
}

func (m *Memory) Len() int {
	return len(m.buf)
}

func (m *Memory) Peek(addr uint16) (byte, error) {
	if int(addr) >= m.Len() {
		return 0, fmt.Errorf("Out of bounds memory read: %d", addr)
	}
	n := m.buf[addr]
	if m.debug {
		fmt.Printf("MEM: RD %04X -> %02X\n", addr, n)
	}
	return n, nil
}

func (m *Memory) Poke(addr uint16, n byte) error {
	if int(addr) >= m.Len() {
		return fmt.Errorf("Out of bounds memory write: %d (%d)", addr, n)
	}
	m.buf[addr] = n
	if m.debug {
		fmt.Printf("MEM: WR %04X <- %02X\n", addr, n)
	}
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
	m.buf = make([]byte, m.Len())
}

func (m *Memory) Copy(addr uint16, buf []byte) error {
	if int(addr)+len(buf) >= m.Len() {
		panic(fmt.Sprintf("Can't load - base addr %04X length %04X memsize %04X", addr, len(buf), m.Len()))
	}
	for i := 0; i < len(buf); i++ {
		m.buf[addr+uint16(i)] = buf[i]
	}
	return nil
}
