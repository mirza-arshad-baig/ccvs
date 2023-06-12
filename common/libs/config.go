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

// InitConfig initializes the server configuration.
// It reads the config file specified by ConfigPath and ConfigFile,
// logs the initialized configurations, and sets up a config watcher.
func InitConfig(ConfigPath string) {
	v, err := ReadConfig(ConfigPath, ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Infoln("Initialized Configurations")

	// Marshal the configurations into a formatted JSON string
	configs, err := json.MarshalIndent(v.AllSettings(), "", "")
	if err != nil {
		log.Infoln("Error:", err)
	}
	log.Infoln(string(configs))
}

// ReadConfig reads the config files located in the specified directory (dir) with the given filename.
// It sets up the config file name and path using viper, enables automatic environment variable binding,
// merges the config file into the viper instance, and sets up a watcher for config changes.
// The function returns the viper instance and any error that occurred during the process.
func ReadConfig(dir string, filename string) (*viper.Viper, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(dir)

	viper.AutomaticEnv()
	err := viper.MergeInConfig() //viper.ReadInConfig()
	if err == nil {
		onceWatcher.Do(func() {
			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				// Handle config changes if needed
			})
		})
	}
	return viper.GetViper(), err
}
