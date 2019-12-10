/**
 *
 * @author nghiatc
 * @since Dec 10, 2019
 */

package nconf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var config *viper.Viper
var env string

func Init() {
	//// init Configuration
	environment := flag.String("e", "development", "run project with mode [-e development | test | production]")
	flag.Usage = func() {
		fmt.Println("Usage: ./[appname] -e development | test | production")
		os.Exit(1)
	}
	flag.Parse()
	log.Printf("============== NConf environment: %s", *environment)
	InitEnv(*environment)
}

// Init is an exported method that takes the environment starts the viper (external lib) and
// returns the configuration struct.
func InitEnv(environment string) {
	log.Printf("============== NConfig Init Environment: %s ==============", environment)
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(environment)
	v.AddConfigPath("../conf/")
	v.AddConfigPath("conf/")
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config = v
	env = environment
}

func RelativePath(basedir string, path *string) {
	p := *path
	if p != "" && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}

func GetEnv() string {
	return env
}
