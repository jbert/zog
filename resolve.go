package zog

func (i *InstU8) Resolve(a *Assembly) error {
	return a.ResolveLoc8(i.l)
}

func (i *InstBin8) Resolve(a *Assembly) error {
	err := a.ResolveLoc8(i.src)
	if err != nil {
		return err
	}
	return a.ResolveLoc8(i.dst)
}

func (i *InstU16) Resolve(a *Assembly) error {
	return a.ResolveLoc16(i.l)
}

func (i *InstBin16) Resolve(a *Assembly) error {
	err := a.ResolveLoc16(i.src)
	if err != nil {
		return err
	}
	return a.ResolveLoc16(i.dst)
}
