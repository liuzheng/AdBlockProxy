package config

import (
	"flag"
	"github.com/liuzheng/golog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const lname = "config"

var (
	config = flag.String("c", "config.yaml", "Print the version and exit")
)

type Configer struct {
	//Module      []Module `yaml:"module"`
	yaml []byte
	//Output      map[string]interface{}
	//Modules     map[string]interface{}

	outputs     map[string]interface{}

	Log struct {
		ConsoleLevel string `yaml:consoleLevel`
		Level        string `yaml:"level"`
		Path         string `yaml:"path"`
	} `yaml:"log"`
	Blacklist map[string][]struct {
		UriRe  string `yaml:uri`
		Action string `yaml:action`
	} `yaml:blacklist`
}

var Config *Configer = &Configer{
	//Output:make(map[string]interface{}),
	//Module: make([]Module, 20),
	//Modules: make(map[string]interface{}),
	//Outputs: make(map[string]interface{}),
}

func LoadConfig() {
	golog.Info("LoadConfig", "Load config %v", *config)
	filename, err := filepath.Abs(*config)
	if err != nil {
		golog.Error("filepath", "filepath.Abs error: %v", err)
		panic(err)
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		golog.Error("ReadFile", "error: %v", err)
		return
	}
	Config.yaml = yamlFile
	//zoo.Load(&Config.Yaml)

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		golog.Error("yaml", "yaml.Unmarshal : %v", err)
		panic(err)
	}

	//log := struct {
	//	Log struct {
	//		Level string `yaml:"level"`
	//		Path  string  `yaml:"path"`
	//	} `yaml:"log"`
	//}{}
	//log.Log.Level = "INFO"
	//log.Log.Path = ""
	//err = yaml.Unmarshal(yamlFile, &log)
	//if err != nil {
	//	golog.Error("yaml", "yaml.Unmarshal : %v", err)
	//	panic(err)
	//}

	//Config.LogLevelB = log.Log.Level
	//Config.LogFilePath = log.Log.Path
	//golog.Info("Config", "%v", log)
	golog.Logs(Config.Log.Path, Config.Log.ConsoleLevel, Config.Log.Level)

	//for name, f := range Config.Outputs {
	//    golog.Info("LoadConfig", "Loading %s config", name)
	//    f.(func())()
	//}
}
func DumpConfig() (err error) {
	d, err := yaml.Marshal(&Config)
	if err != nil {
		golog.Error(lname, "yaml.Marshal: %v", err)
	}
	err = ioutil.WriteFile("config.yaml", d, 0644)
	if err != nil {
		golog.Error(lname, "ioutil.WriteFile: %v", err)
	}
	return
}
