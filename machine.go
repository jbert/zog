package zog

type Machine interface {
	LoadAddr() uint16
	RunAddr() uint16
	Name() string
	Load(z *Zog) error
}
