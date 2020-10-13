// +build windows

package gohelper

import (
	"golang.org/x/sys/windows/registry"
)

// Cwindows ...Edit registry to suppory ANSII
func Cwindows() error {
	reg, err := registry.OpenKey(registry.CURRENT_USER, `Console`, registry.WRITE)
	if err != nil {
		return err
	}
	err = reg.SetDWordValue("VirtualTerminalLevel", 0x00000001)
	if err != nil {
		return err
	}
	reg.Close()
	return nil
}
