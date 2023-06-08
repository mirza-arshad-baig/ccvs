package libs

import (
	"encoding/json"
	"sync"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (

	//ConfigFile is config file name
	ConfigFile  string
	onceWatcher sync.Once
)

// InitConfig will init server config
func InitConfig(ConfigPath string) {
	v, err := ReadConfig(ConfigPath, ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Infoln("Initialized Configurations")
	configs, err := json.MarshalIndent(v.AllSettings(), "", "")
	if err != nil {
		log.Infoln("Error:", err)
	}
	log.Infoln(string(configs))
}

// ReadConfig will read config files
func ReadConfig(dir string, filename string) (*viper.Viper, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(dir)

	viper.AutomaticEnv()
	err := viper.MergeInConfig() //viper.ReadInConfig()
	if err == nil {
		onceWatcher.Do(func() {
			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				//InitConfig()
			})
		})
	}
	return viper.GetViper(), err
}
