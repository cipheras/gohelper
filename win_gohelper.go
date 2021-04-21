// +build windows

package gohelper

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

// Cwindows ...Edit registry to suppory ANSII
func Cwindows() error {
	reg, err := registry.OpenKey(registry.CURRENT_USER, `Console`, registry.QUERY_VALUE)
	if err != nil {
		return err
	}
	s, _, err := reg.GetIntegerValue("VirtualTerminalLevel")
	reg.Close()
	if s == 0 {
		reg, err = registry.OpenKey(registry.CURRENT_USER, `Console`, registry.WRITE)
		if err != nil {
			return err
		}
		err = reg.SetDWordValue("VirtualTerminalLevel", 0x00000001)
		if err != nil {
			return err
		}
		err = reg.SetDWordValue("VirtualTerminalLevel", 0x00000000)
		if err != nil {
			return err
		}
		reg.Close()
		cmd := exec.Command("cmd")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
		fmt.Printf("\033[2J\033[H")

		// if err := undo(); err != nil {
		// 	return err
	}
	return nil
}
