package cmd

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var (
	// Used for flags.
	cfgFile     string

	rootCmd = &cobra.Command{
		Use:   "cdf",
		Short: "CloudEvents Emitter for Continuous Delivery Events",
		Long: `CDF - Continuous Delivery Flow is a simple framework and CLI to emit CloudEvents related to 
Continuous Delivery cycle.`,
	}
)

var CDF_SINK = os.Getenv("CDF_SINK")

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}


func init() {

	if CDF_SINK == "" {
		CDF_SINK = "http://localhost:8080"
	}

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cdf.yaml)")
	//rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		configPath := filepath.Join(home, ".cdf.yaml")

		_, err = os.Stat(configPath)
		if !os.IsExist(err) {
			if _, err := os.Create(configPath); err != nil { // perm 0666
				log.Printf("failed to create file: %s\n", err )
			}
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cdf")
		viper.SetConfigType("yaml")
		viper.SetConfigPermissions(0666)



	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.AutomaticEnv()


	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	log.Printf("> cdf.project.name =  %s", viper.GetString("cdf.project.name") )
	log.Printf("> cdf.module.name = %s", viper.GetString("cdf.module.name") )
	log.Printf("> all keys = %s", viper.AllKeys() )

}