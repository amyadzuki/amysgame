// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amyadzuki/amygolib/dirs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	VendorName      = "Amy" // TODO: I need a vendor name
	ApplicationName = "amysgame" // TODO: I need an app name
	DefaultConfType = "yaml"
)

var (
	cfgFile string
	Dirs *dirs.Dirs
	Ofs int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   ApplicationName,
	Short: "The client application for games by amyadzuki",
	Long: `The client application for games by amyadzuki`,
	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	Dirs = dirs.New(VendorName, ApplicationName)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	configFile := filepath.Join(Dirs.Config(), "config." + DefaultConfType)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is " + configFile + ")")
	rootCmd.PersistentFlags().IntVar(&Ofs, "ofs", 0, "offset")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in dirs.Config() directory with name "ApplicationName" (without extension).
		viper.AddConfigPath(Dirs.Config())
		viper.SetConfigName(ApplicationName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
