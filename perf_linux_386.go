package perf

import (
	"os"
	"syscall"
	"unsafe"
)

func sys_perf_counter_open(attr *Attr, pid, cpu, group_fd int, flags uint64) (counter *os.File, err error) {
	attr.Size = _ATTR_SIZE

	r1, _, e := syscall.Syscall6(_SYS_PERF_OPEN,
		uintptr(unsafe.Pointer(attr)),
		uintptr(pid),
		uintptr(cpu),
		uintptr(group_fd),
		uintptr(flags), uintptr(flags>>32))

	if e != 0 {
		return nil, os.NewSyscallError("perf counter open", e)
	}

	return os.NewFile(r1, "<perf counter>"), nil
}
