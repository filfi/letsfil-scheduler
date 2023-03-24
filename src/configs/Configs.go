package configs

import (
	"bytes"
	_ "embed"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/global/constant"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fileUtils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
)

var config = new(Config)

type Config struct {
	Server struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	MySQL struct {
		Username     string `yaml:"username" required:"false"`
		Password     string `yaml:"password" required:"false"`
		Host         string `yaml:"host" required:"false"`
		Port         string `yaml:"port" required:"false"`
		Db           string `yaml:"db" required:"false"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	} `yaml:"mysql"`

	Redis struct {
		Host     string `yaml:"host" required:"false"`
		Port     string `yaml:"port" required:"false"`
		Database int    `yaml:"database"`
		Password string `yaml:"password" required:"false"`
		PoolSize int    `yaml:"poolSize"`
		Timeout  string `yaml:"timeout"`
	} `yaml:"redis"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`

	Influxdb struct {
		Url    string `yaml:"url" required:"false"`
		Token  string `yaml:"token" required:"false"`
		Bucket string `yaml:"bucket"`
		Org    string `yaml:"org" required:"false"`
	} `yaml:"influxdb"`

	Lotus struct {
		Url       string `yaml:"url" required:"false"`
		AuthToken string `yaml:"authToken" required:"false"`
		NameSpace string `yaml:"nameSpace"`
	} `yaml:"lotus"`

	Glif struct {
		Url       string `yaml:"url" required:"false"`
		AuthToken string `yaml:"authToken" required:"false"`
	} `yaml:"glif"`

	FilDate struct {
		Url       string `yaml:"url" required:"false"`
		AuthToken string `yaml:"authToken" required:"false"`
	} `yaml:"fildata"`

	FilFevm struct {
		HttpsUrl        string `yaml:"httpsUrl" required:"false"`
		WsUrl           string `yaml:"wsUrl" required:"false"`
		FactoryContract string `yaml:"factoryContract" required:"false"`
	} `yaml:"filfevm"`
}

var (

	//go:embed application.yaml
	defaultConfigs []byte

	//go:embed application_leo.yaml
	leoConfigs []byte

	//go:embed application_local.yaml
	localConfigs []byte

	//go:embed application_dev.yaml
	devConfigs []byte

	//go:embed application_prod.yaml
	prodConfigs []byte
)

func InitConfig() {
	var r io.Reader
	var filename string
	switch os.Getenv(constant.ENV_MODE) {
	case "leo":
		filename = "application_leo"
		r = bytes.NewReader(leoConfigs)
	case "local":
		filename = "application_local"
		r = bytes.NewReader(localConfigs)
	case "dev":
		filename = "application_dev"
		r = bytes.NewReader(devConfigs)
	case "prod":
		filename = "application_prod"
		r = bytes.NewReader(prodConfigs)
	default:
		filename = "application"
		r = bytes.NewReader(defaultConfigs)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName(filename)
	viper.AddConfigPath("./src/configs")

	configFile := "./src/configs/" + filename + ".yaml"
	_, ok := fileUtils.IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}

func GetValue(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	} else {
		return defaultValue
	}
}
