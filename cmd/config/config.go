package config

import (
	"os"
	"strings"

	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// InitConfig reads in config file and ENV variables if set.
func InitConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		// Search config in config directory with name ".slatomate" (without extension).
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("slatomate")
		viper.SetDefault("auth_token", "")
		viper.SetDefault("service_host", "127.0.0.1:8081")
		viper.SafeWriteConfig()

	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Debug("Using config file:", viper.ConfigFileUsed())
	}
}
