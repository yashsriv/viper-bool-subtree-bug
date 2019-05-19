package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// auto read config values from environment
	viper.SetEnvPrefix("spf")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	databaseConfig := viper.Sub("pg.database")

	fmt.Printf("from subtree: %v\n", databaseConfig.GetBool("ssl"))
	fmt.Printf("from actual: %v\n", viper.GetBool("pg.database.ssl"))
}
