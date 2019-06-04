package utils

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)
const (
	LoadPath       = "./conf/config.yml"
)
type YamlFile struct{
	Mysql struct {
		Database     string `yaml:"Database"`
	} `yaml:"Mysql"`
}

func InitConfig()(YamlFile,error){
	configInstance := &YamlFile{}
	err := Load(configInstance)
	if err != nil{
		logrus.WithError(err).Error("load local yml config file failed!")
		return *configInstance,err
	}
	logrus.WithField("dsn",configInstance.Mysql.Database).Infof("connect success")
	return *configInstance,nil
}




func Load(config interface{})error{
	env := ENV()
	pathName := getConfigFilePath(LoadPath, env)
	return configor.New(&configor.Config{Verbose: true}).Load(config, pathName)
}
func getConfigFilePath(rawPath string, env string) string {
	pathName, err := getConfigFileWithFileNameSuffix(rawPath, env)
	if err != nil {
		logrus.Infof("path not exist with specified env=%v, err=%v\n", env, err)
		return rawPath
	} else {
		fmt.Printf("load conf from file: %v\n", pathName)
		return pathName
	}
}


func getConfigFileWithFileNameSuffix(file, suffix string) (string, error) {
	var (
		envFile string
		extname = path.Ext(file)
	)

	if extname == "" {
		envFile = fmt.Sprintf("%v.%v", file, suffix)
	} else {
		envFile = fmt.Sprintf("%v.%v%v", strings.TrimSuffix(file, extname), suffix, extname)
	}

	if fileInfo, err := os.Stat(envFile); err == nil && fileInfo.Mode().IsRegular() {
		return envFile, nil
	}
	return "", fmt.Errorf("failed to find file %v", envFile)
}