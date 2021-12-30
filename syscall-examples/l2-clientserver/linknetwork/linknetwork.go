package linknetwork

import (
    "net"
    "fmt"
    "os"
    "encoding/binary"

    "golang.org/x/sys/unix"
)

const (
    networkProtocol int = 0x88b5
)

func Listen(interfaceName string) (net.Conn, error) {
    resConn := &LinkConn{}

    fd, interf, err := createSocket(interfaceName)

    if err != nil {
        return nil, err
    }

    resConn.localAddr = &LinkAddr{interf.HardwareAddr}
    resConn.fileFd = os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))
    resConn.fd = fd
    resConn.listenConn = true

    return resConn, nil
}

func DialLink(interfaceName string, remoteAddr LinkAddr) (net.Conn, error) {
    resConn := &LinkConn{}

    fd, interf, err := createSocket(interfaceName)

    if err != nil {
        return nil, err
    }

    resConn.localAddr = &LinkAddr{interf.HardwareAddr}
    resConn.remoteAddr = &remoteAddr

    resConn.fileFd = os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))
    resConn.fd = fd

    return resConn, nil
}

func createSocket(interfaceName string) (fd int, interf *net.Interface, err error) {
    fd, err = unix.Socket(unix.AF_PACKET, unix.SOCK_RAW | unix.SOCK_CLOEXEC, htons(3, 16))

    if err != nil {
        return
    }

    interf, err = net.InterfaceByName(interfaceName)

    if err != nil {
        return
    }

    // Create sockaddr struct and bind socket
    sockaddr := unix.SockaddrLinklayer {
        Protocol: uint16(htons(networkProtocol, 16)),
        Ifindex: interf.Index,
    }

    err = unix.Bind(fd, &sockaddr)
    return
}

func htons(value int, valueLen int) int {
    if valueLen > 64 && valueLen < 1 {
        panic("Invalid valueLen arg")
    }

    tmpBuff := make([]byte, 8)

    binary.BigEndian.PutUint64(tmpBuff, uint64(value))
    return int(binary.LittleEndian.Uint64(tmpBuff) >> (64 - valueLen))
}
