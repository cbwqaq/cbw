package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
	"os"
	"time"
)

func init() {
	setup("")
}

var (
	configFilePaths = [3]string{
		"config/config.yaml",
		"../config/config.yaml",
		"../../config/config.yaml",
	}
	//EnvConfig 环境配置，指定是什么环境，如开发环境，测试环境，本地等
	EnvConfig *envConfig
	//HTTPSetting http相关配置
	HTTPSetting *httpSetting
	//RemoteSetting 远程调用配置
	RemoteSetting *remoteSetting
)

type envConfig struct {
	Title      string                  `yaml:"title"`
	ReleaseEnv string                  `yaml:"releaseEnv"`
	Version    string                  `yaml:"version"`
	Server     map[string]ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Domain string `yaml:"domain"`
	//AppConfig     appConfig     `yaml:"appConfig"`
	HTTPSetting   httpSetting   `yaml:"httpSetting"`
	RemoteSetting remoteSetting `yaml:"remoteSetting"`
	/*NotifySetting notifySetting `yaml:"notifySetting"`
	RpcConfig     rpcConfig     `yaml:"rpcConfig"`
	LogConfig     logConfig     `yaml:"logConfig"`*/
}

type httpSetting struct {
	HTTPPort     string        `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	RunMode      string        `yaml:"runMode"`
}

type remoteSetting struct {
	MysqlAddr     string `yaml:"mysqlAddr"`
	MysqlHost     string `yaml:"mysqlHost"`
	MysqlUserName string `yaml:"mysqlUserName"`
	MysqlPassWord string `yaml:"mysqlPassWord"`
	MysqlPort     string `yaml:"mysqlPort"`
	MysqlDatabase string `yaml:"mysqlDatabase"`
}

func setup(path string) {
	loadConfig(path)
}

func loadConfig(path string) {
	if path != "" {
		if notExists(path) {
			panic(fmt.Errorf("config file paht: %s not exists", path))
		}

		loadConfigWithPath(path)
		return
	}
	for _, configPath := range configFilePaths {
		//从不同层级目录初始化环境配置， 直到有一次初始化成功后退出
		currentDir, _ := os.Getwd()
		klog.Infof("current directory: %s, load config: %s", currentDir, configPath)
		if notExists(configPath) {
			continue
		}
		loadConfigWithPath(configPath)
		break
	}
	if EnvConfig == nil {
		panic("envConfig init fail")
	}
}

func loadConfigWithPath(configPath string) {
	config, err := os.ReadFile(configPath)
	if err != nil {
		klog.Error("read config file err:", err)
		panic(err)
	}
	err = yaml.Unmarshal(config, &EnvConfig)
	if err != nil {
		panic(err)
	}
	klog.Infof("read Config: %s", configPath)

	releaseEnv := os.Getenv("releaseEnv")
	fmt.Println(releaseEnv)
	if releaseEnv == "" {
		releaseEnv = EnvConfig.ReleaseEnv
	}

	serverConfig := EnvConfig.Server[releaseEnv]
	HTTPSetting = &serverConfig.HTTPSetting
	RemoteSetting = &serverConfig.RemoteSetting
	HTTPSetting.ReadTimeout = HTTPSetting.ReadTimeout * time.Second
	HTTPSetting.WriteTimeout = HTTPSetting.WriteTimeout * time.Second

}

func notExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	return err != nil
}
