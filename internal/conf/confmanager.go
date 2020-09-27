package conf

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// config - all configs are stored
var config *viper.Viper

// defaultConfigFileDirs - configuration file directories, file will be searched in the directories below.
var defaultConfigFileDirs = [...]string{"", "app/", "configs/app/"}

// defaultConfigFileName - default configuration file name
var defaultConfigFileName = "config.yaml"

// ConfigFileNameKey - configuration file name (it will be searched in configFileDirs)
const ConfigFileNameKey string = "CONFIG_FILE"

// ConfigFilePathKey - configuration file's full path
const ConfigFilePathKey string = "CONFIG_FILE_PATH"

// ConfigImportConfKey - importable configurations in the same path(if starts with '/' full path will be used)
const ConfigImportConfKey string = "config.import"

// ConfigEnvDisabledKey - if true env var will not applied
const ConfigEnvDisabledKey string = "config.env.disabled"

// ConfigEnvPrefixKey - if set all env vars with this prefix applied and overrides current ones
const ConfigEnvPrefixKey string = "config.env.prefix"

func generate() {
	config = viper.New()
	confFilePath := os.Getenv(ConfigFilePathKey)
	if confFilePath == "" {
		confFilePath = searchForConfigFile(config, confFilePath)
	} else {
		if fileExists(confFilePath) {
			loadConfFromFile(config, confFilePath, false)
		} else {
			confFilePath = searchForConfigFile(config, confFilePath)
		}
	}

	if config.IsSet(ConfigImportConfKey) {
		for _, filePath := range config.GetStringSlice(ConfigImportConfKey) {
			if strings.HasPrefix(filePath, "file:") {
				filePath = filePath[len("file:"):]
				if !strings.HasPrefix(filePath, "/") {
					filePath = path.Join(path.Dir(confFilePath), filePath)
				}

				if fileExists(filePath) {
					loadConfFromFile(config, filePath, true)
				} else {
					log.Printf("File is not found or not accessible. Please check the file(%s)", filePath)
				}
			}
		}
	}

	if !config.GetBool(ConfigEnvDisabledKey) {
		if config.IsSet(ConfigEnvPrefixKey) {
			config.SetEnvPrefix(config.GetString(ConfigEnvPrefixKey))
		}
		config.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
		config.AutomaticEnv()
	}
}

func searchForConfigFile(config *viper.Viper, confFilePath string) string {
	confFileName := os.Getenv(ConfigFileNameKey)
	if confFileName == "" {
		confFileName = defaultConfigFileName
	}

	for _, basePath := range defaultConfigFileDirs {
		configPath := path.Join(confFilePath, basePath, confFileName)
		if fileExists(configPath) {
			confFilePath = configPath
			break
		}
	}

	if confFilePath == "" {
		log.Printf("Configuration file not found.")
	} else if(!fileExists(confFilePath)) {
		errTxt := fmt.Sprintf("configuration file is not valid (file:%s).", confFilePath)
		log.Printf(errTxt)
		panic(errTxt)
	} else {
		loadConfFromFile(config, confFilePath, false)
	}
	return confFilePath
}

func loadConfFromFile(config *viper.Viper, confFilePath string, merge bool) {
	log.Printf("Configuration file found. Path: %s (merge: %v)", confFilePath, merge)
	config.SetConfigFile(confFilePath)
	var err error
	if merge {
		err = config.MergeInConfig()
	} else {
		err = config.ReadInConfig()
	}
	if err != nil {
		panic(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// GetString - get string by key
func GetString(key string) string {
	return config.GetString(key)
}

// InitConfigs - initialing configuration manager
func InitConfigs() {
	log.Printf("Configurations will be initialized.")
	generate()
}