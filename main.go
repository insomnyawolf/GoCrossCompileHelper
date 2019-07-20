package main

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	loadConf()
}

func main() {
	for _, ops := range config.OS.getOSList() {
		for _, arch := range config.Arch.getArchList() {
			ext := ""
			if ops == "windows" {
				ext = ".exe"
			}
			build(ops, arch, ext)
		}
	}
}

func build(ops, arch, ext string) {
	cmd := exec.Command("go", "build", "-i", "-o", fmt.Sprintf("%v/%v/%v-%v-%v-%v%v", config.Path, config.Ver, config.Name, config.Ver, ops, arch, ext))
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%v", ops))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%v", arch))
	fmt.Printf("Building for %v %v", ops, arch)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))
}
