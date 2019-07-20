package main

import (
	"reflect"
	"strings"
)

// Architectures posible processor architecture
type Architectures struct {
	I386        bool
	Amd64       bool
	Amd64p32    bool
	Arm         bool
	Armbe       bool
	Arm64       bool
	Arm64be     bool
	Ppc         bool
	Ppc64       bool
	Ppc64le     bool
	Mips        bool
	Mipsle      bool
	Mips64      bool
	Mips64le    bool
	Mips64p32   bool
	Mips64p32le bool
	S390        bool
	S390x       bool
	Sparc       bool
	Sparc64     bool
}

func (c *Architectures) getArchList() []string {
	var ArchList []string
	value := reflect.ValueOf(*c)
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).Bool() {
			typeOfS := value.Type()
			name := typeOfS.Field(i).Name
			name = strings.ToLower(name)
			if name == "i386" {
				name = "386"
			}
			ArchList = append(ArchList, name)
		}
	}
	return ArchList
}
