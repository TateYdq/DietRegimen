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
	StaticUploadPath = "static/image/"
	StaticVoicePath = "static/voice/"
)


type YamlFile struct{
	Mysql struct {
		Database     string `yaml:"Database"`
	} `yaml:"Mysql"`
	Redis struct{
		Addr string `yaml:"Addr"`
		Password string `yaml:"Password"`
		DbToken int `yaml:"DbToken"`
	}`yaml:"Redis"`
	Wechat struct{
		AppID string `yaml:"AppID"`
		AppSecret string `yaml:"AppSecret"`
	}`yaml:"Wechat"`
	Ai struct{
		ApiKey string `yaml:"ApiKey"`
		SecretKey string `yaml:"SecretKey"`
	}`yaml:"Ai"`
}

var(
	ConfigInstance *YamlFile
)


func InitConfig()(YamlFile,error){
	ConfigInstance = &YamlFile{}
	err := Load(ConfigInstance)
	if err != nil{
		logrus.WithError(err).Error("load local yml config file_interact failed!")
		return *ConfigInstance,err
	}
	logrus.WithField("dsn",ConfigInstance.Mysql.Database).Infof("connect success")
	return *ConfigInstance,nil
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
		fmt.Printf("load conf from file_interact: %v\n", pathName)
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
	return "", fmt.Errorf("failed to find file_interact %v", envFile)
}