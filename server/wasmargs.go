package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/perlin-network/life/exec"
)

var ErrShortBuffer = errors.New("short buffer")

type unsafeBytes []byte

func wasmArgs(vm *exec.VirtualMachine, vals ...interface{}) error {
	var nonFatal error
	args := vm.GetCurrentFrame().Locals
	for _, iVal := range vals {
		if len(args) < 1 {
			return fmt.Errorf("not enough args")
		}
		a0 := args[0]
		consumed := 1
		switch val := iVal.(type) {
		case *bool:
			*val = (a0 != 0)

		case *int8:
			*val = int8(a0)
		case *int16:
			*val = int16(a0)
		case *int32:
			*val = int32(a0)
		case *int64:
			*val = a0

		case *uint8:
			*val = uint8(a0)
		case *uint16:
			*val = uint16(a0)
		case *uint32:
			*val = uint32(a0)
		case *uint64:
			*val = uint64(a0)

		// TODO: test these
		case *float32:
			*val = math.Float32frombits(uint32(a0))
		case *float64:
			*val = math.Float64frombits(uint64(a0))

		case []byte, *[]byte, *unsafeBytes, *string:
			if len(args) < 2 {
				return fmt.Errorf("not enough args for ptr+len")
			}
			a1 := args[1]
			consumed = 2

			if int(a0+a1) > len(vm.Memory) {
				return fmt.Errorf("ptr+len out of bounds")
			}
			buf := vm.Memory[a0 : a0+a1]

			switch v := val.(type) {
			case []byte:
				n := copy(v, buf)
				if n < len(buf) {
					nonFatal = ErrShortBuffer
				}
			case *[]byte:
				*v = make([]byte, len(buf))
				copy(*v, buf)
			case *unsafeBytes:
				*v = unsafeBytes(buf)
			case *string:
				*v = string(buf)
			default:
				panic("unreachable")
			}

		default:
			return fmt.Errorf("cannot handle type %T", val)
		}
		args = args[consumed:]
	}
	if len(args) != 0 {
		return fmt.Errorf("too many args")
	}
	return nonFatal
}
