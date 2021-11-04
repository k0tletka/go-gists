package linknetwork

import (
    "os"
    "net"
    "time"
    "errors"

    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
)

const (
    mtu = 1514
    linkLayerLength = 14

    LinkMTU = mtu - linkLayerLength
    MinLinkMTU = 46
)

var (
    ErrPacketMayTruncated = errors.New("parsed packet may be truncated")
)

type LinkConn struct {
    localAddr *LinkAddr
    remoteAddr *LinkAddr

    listenConn bool

    fileFd *os.File
    fd int
}

func (l LinkConn) Read(b []byte) (int, error) {
    var err error

    if (len(b) < LinkMTU) {
        err = ErrPacketMayTruncated
    }

    readbuf := make([]byte, mtu)

    if n, err := l.fileFd.Read(readbuf); err != nil {
        return n, err
    }

    packet := gopacket.NewPacket(readbuf, layers.LayerTypeEthernet, gopacket.NoCopy)
    ethernetPacket, ok := packet.Layer(layers.LayerTypeEthernet).(*layers.Ethernet)

    if !ok {
        panic("Invalid layer has been parsed")
    }

    // Save remote MAC address, so we could send response back
    if l.listenConn {
        l.remoteAddr = &LinkAddr{MACAddress: []byte(ethernetPacket.SrcMAC)}
    }

    copy(b, ethernetPacket.Payload)
    return len(ethernetPacket.Payload), err
}

func (l LinkConn) Write(b []byte) (int, error) {
    var payloadLength int

    if len(b) < MinLinkMTU {
        payloadLength = MinLinkMTU
    } else {
        payloadLength = len(b)
    }

    writeBufferSerializer := gopacket.NewSerializeBuffer()

    payloadData, _ := writeBufferSerializer.AppendBytes(payloadLength)
    copy(payloadData, b)

    // Fill out ethernet data
    ethernetOutPacket := &layers.Ethernet{
        EthernetType: layers.EthernetType(networkProtocol),
        DstMAC: net.HardwareAddr(l.remoteAddr.MACAddress),
        SrcMAC: net.HardwareAddr(l.localAddr.MACAddress),
    }

    if err := ethernetOutPacket.SerializeTo(writeBufferSerializer, gopacket.SerializeOptions{}); err != nil {
        return 0, err
    }

    return l.fileFd.Write(writeBufferSerializer.Bytes())
}

func (l LinkConn) Close() error {
    return l.fileFd.Close()
}

func (l LinkConn) LocalAddr() net.Addr {
    return *l.localAddr
}

func (l LinkConn) RemoteAddr() net.Addr {
    if l.remoteAddr == nil {
        return LinkAddr{}
    }

    return *l.remoteAddr
}

func (l LinkConn) SetDeadline(t time.Time) error {
    return l.fileFd.SetDeadline(t)
}

func (l LinkConn) SetReadDeadline(t time.Time) error {
    return l.fileFd.SetReadDeadline(t)
}

func (l LinkConn) SetWriteDeadline(t time.Time) error {
    return l.fileFd.SetWriteDeadline(t)
}
