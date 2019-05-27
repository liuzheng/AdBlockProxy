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

type URI struct {
	UriRe  string `yaml:"uri"`
	Action string `yaml:"action"`
	Code   int    `yaml:"code"`
	Proxy  string `yaml:"proxy"`
}
type Configer struct {
	//Module      []Module `yaml:"module"`
	yaml []byte
	//Output      map[string]interface{}
	//Modules     map[string]interface{}

	outputs   map[string]interface{}
	Server    string           `yaml:"server"`
	Listen    uint16           `yaml:"listen"`
	Blacklist map[string][]URI `yaml:"blacklist"`
	ProxyList struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Urls     []struct {
			Name string `yaml:"name"`
			Url  string `yaml:"url"`
		} `yaml:"url"`
	} `yaml:"proxy"`
	Log struct {
		ConsoleLevel string `yaml:"consoleLevel"`
		Level        string `yaml:"level"`
		Path         string `yaml:"path"`
	} `yaml:"log"`
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
	if Config.Listen < 80 {
		Config.Listen = 8080
	}
	if Config.Server == "" {
		Config.Server = "127.0.0.1"
	}
	if Config.Log.ConsoleLevel == "" {
		Config.Log.ConsoleLevel = "INFO"
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
