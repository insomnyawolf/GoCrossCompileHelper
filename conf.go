package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	//config Target Platform Values
	config *myConf
)

type myConf struct {
	//AppConfig
	Path string `description:"AppConfig:"`
	Name string
	Ver  string
	Arch Architectures    `description:"Architectures:"`
	OS   OperatingSystems `description:"Operating Systems:"`
}

// ConfigPath Defines where the configuration will be saved/written
const ConfigPath = "./BuildConfig.conf"

func loadConf() {
	defaultConf := myConf{
		//AppConfig
		Path: "bin/",
		Name: getCurrentDirName(),
		Ver:  "1.0.0",
		//Arch
		Arch: Architectures{
			I386:  true,
			Amd64: true,
		},
		//OS
		OS: OperatingSystems{
			Linux:   true,
			Windows: true,
		},
	}

	println("Please configure the app and start it again.")
	config = &myConf{}
	load(config, defaultConf)
}

func load(output interface{}, defaultValues interface{}) error {
	_, err := os.Stat(ConfigPath)
	if err == nil {
		file, err := ioutil.ReadFile(ConfigPath)
		CheckErrFatal(err)
		err = json.Unmarshal(file, &output)
		CheckErrFatal(err)
		return nil
	}
	if os.IsNotExist(err) {
		config, err := json.MarshalIndent(defaultValues, "", "  ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(ConfigPath, config, 777)
		if err != nil {
			return err
		}
	}
	return err
}
