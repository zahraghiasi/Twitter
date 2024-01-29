package config

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

var Conf ConfYaml

var defaultConf = []byte(`
core:
  mode: "debug" # release, debug, test
  server_name: "social-network"
  port: 9090
log:
  level: debug
  discard_output: false
mysql:
  host: "127.0.0.1"
  port: "3306"
  db: "service"
  user: "root"
  pass: "ghiasi"
mongo:
  address: "localhost:27017"
  username: "ghiasi"
  password: "ghiasi"
`)

type ConfYaml struct {
	Core  SectionCore  `yaml:"core"`
	Mysql SectionMysql `yaml:"mysql"`
	Log   SectionLog   `yaml:"log"`
	Mongo SectionMongo `yaml:"mongo"`
}

type SectionLog struct {
	Level         string `yaml:"level"`
	DiscardOutput bool   `yaml:"discard_output"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Mode string `yaml:"mode"`
	Port int    `yaml:"port"`
}

type SectionMongo struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// SectionMysql is sub section of config.
type SectionMysql struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB   string `yaml:"db"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()            // read in environment variables that match
	viper.SetEnvPrefix("GO_SOCIAL") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			log.Errorf("FileRepo does not exist : %s", confPath)
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".pkg" (without extension).
		viper.AddConfigPath("/etc/go-social/")
		viper.AddConfigPath("$HOME/.go-social")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Core
	conf.Core.Mode = viper.GetString("core.mode")
	conf.Core.Port = viper.GetInt("core.port")

	// Mysql
	conf.Mysql.Host = viper.GetString("mysql.host")
	conf.Mysql.Port = viper.GetString("mysql.port")
	conf.Mysql.DB = viper.GetString("mysql.db")
	conf.Mysql.User = viper.GetString("mysql.user")
	conf.Mysql.Pass = viper.GetString("mysql.pass")

	//Log
	conf.Log.Level = viper.GetString("log.level")
	conf.Log.DiscardOutput = viper.GetBool("log.discard_output")

	// Mongo
	conf.Mongo.Address = viper.GetString("mongo.address")
	conf.Mongo.Username = viper.GetString("mongo.username")
	conf.Mongo.Password = viper.GetString("mongo.password")

	return conf, nil
}

func Initialize(path string) {
	var err error
	Conf, err = LoadConf(path)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return
	}
}
