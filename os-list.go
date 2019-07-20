package main

import (
	"reflect"
	"strings"
)

// OperatingSystems Posible operating system targets
type OperatingSystems struct {
	Android   bool
	Darwin    bool
	Dragonfly bool
	Freebsd   bool
	Linux     bool
	Nacl      bool
	Netbsd    bool
	Openbsd   bool
	Plan9     bool
	Solaris   bool
	Windows   bool
	Zos       bool
}

func (c *OperatingSystems) getOSList() []string {
	var OSList []string
	value := reflect.ValueOf(*c)
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Bool() {
			typeOfS := value.Type()
			name := strings.ToLower(typeOfS.Field(i).Name)
			OSList = append(OSList, name)
		}
	}
	return OSList
}
