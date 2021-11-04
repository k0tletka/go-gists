package linknetwork

type LinkAddr struct {
    MACAddress []byte
}

func (l LinkAddr) Network() string {
    return "link"
}

func (l LinkAddr) String() string {
    return string(l.MACAddress)
}
