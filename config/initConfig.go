package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func Init(c interface{}) {
	// 获取可执行文件相对于当前工作目录的相对路径
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, fileName := filepath.Split(file)

	// fmt.Println("可执行文件路径: ", path)

	var confPath string
	if len(os.Args) > 1 && os.Args[1] == "dev" {
		confPath = "config.yaml"
	} else {
		confPath = filepath.Join(path, fileName[0:strings.Index(fileName, ".")], "config.yaml")
	}

	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
}
