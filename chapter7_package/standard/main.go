package main

import "syscall"

const (
	LINUX_REBOOT_MAGIC1      uintptr = 0xfee1dead
	LINUX_REBOOT_MAGIC2      uintptr = 0xbadf00d
	LINUX_REBOOT_CMD_RESTART uintptr = 0x12345678
)

func main() {
	syscall.SyscallN(
		//该命令是linux下的命令
		//syscall.SYS_REBOOT,
		LINUX_REBOOT_MAGIC1,
		LINUX_REBOOT_MAGIC2,
		LINUX_REBOOT_CMD_RESTART,
	)

}
