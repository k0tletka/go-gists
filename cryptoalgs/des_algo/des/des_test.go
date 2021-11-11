package des

import (
    "fmt"
    "bytes"
    "testing"
)

func TestConvertByteSliceToUint64(t *testing.T) {
    var input []byte
    var excepted uint64

    input, excepted = []byte{0xFF, 0x00, 0xEE, 0x00, 0xDD, 0x00, 0xCC, 0x00}, 0xFF00EE00DD00CC00

    if got := convertByteSliceToUint64(input); got != excepted {
        t.Errorf("%s: invalid value, excepted %x, got %x\n", t.Name(), excepted, got)
    }
}

func TestConvertUint64ToByteSlice(t *testing.T) {
    var input uint64
    var excepted []byte

    input, excepted = 0xFF00EE00DD00CC00, []byte{0xFF, 0x00, 0xEE, 0x00, 0xDD, 0x00, 0xCC, 0x00}

    if got := convertUint64ToByteSlice(input); !bytes.Equal(got, excepted) {
        t.Errorf("%s: invalid value, excepted %v, got %v\n", t.Name(), excepted, got)
    }
}

func TestIPStraightPermutation(t *testing.T) {
    var input uint64
    var excepted uint64

    input, excepted = 0x5555555555555555, 0xFFFFFFFF00000000

    if got := performIPStraightPermutation(input); got != excepted {
        t.Errorf("%s: invalid value, excepted %x, got %x\n", t.Name(), excepted, got)
    }
}

func TestIPReversePermutation(t *testing.T) {
    var input uint64
    var excepted uint64

    input, excepted = 0x0F0F0F0F0F0F0F0F, 0xFFFFFFFF00000000

    if got := performIPReversePermutation(input); got != excepted {
        t.Errorf("%s: invalid value, excepted %x, got %x\n", t.Name(), excepted, got)
    }
}

func TestExpandEFunction(t *testing.T) {
    var input uint32
    var excepted uint64

    input, excepted = 0xFFFF8001, 0xFFFFFFC00003

    if got := performEExpandFunction(input); got != excepted {
        t.Errorf("%s: invalid value, excepted %x, got %x\n", t.Name(), excepted, got)
    }
}

func TestPPermutation(t *testing.T) {
    var input uint32
    var excepted uint32

    input, excepted = 0x8A53DA5A, 0xFFFF0000
    fmt.Printf("%b\n", input)

    if got := performPPermutation(input); got != excepted {
        t.Errorf("%s: invalid value, excepted %x, got %x\n", t.Name(), excepted, got)
    }
}

func TestSPermutation(t *testing.T) {
    var input uint64
    var excepted uint32

    input, excepted = 0x00000C031E583203, 0xFFFFFFFF

    if got := performSPermutation(input); got != excepted {
        t.Errorf("%s: invalid value, excepted %x, got %x\n", t.Name(), excepted, got)
    }
}
