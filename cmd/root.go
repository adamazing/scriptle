package cmd

import (
	"fmt"
  "os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	filePath string
)

var rootCmd = &cobra.Command{
  Use:    "scriptle",
  Short:  "Learn yo' lines",
  Long:   "A tool to help you learn lines, cues, and stage directions from the command line.",
  Run:    func(cmd *cobra.Command, args []string){
            fmt.Println("In the command")
          },
}

func Execute() {
	// fmt.Println("Hello world!")
  cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// fmt.Println("Init function called")
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "","config file (default is $HOME/.scriptle.yaml)")
}

func initConfig() {
	// fmt.Println("Init config")

	if cfgFile != "" {
		// fmt.Println("Init config, cfgFile not set")
	} else {
    home, err := homedir.Dir()
    cobra.CheckErr(err)
    viper.AddConfigPath(home)
    viper.SetConfigName(".scriptle")
		// fmt.Println("Init config, cfgFile found")
	}

  viper.AutomaticEnv()

  err := viper.ReadInConfig()
  if err == nil {
    // fmt.Fprintln(os.Stderr, "Using config file: ", viper.ConfigFileUsed())
  }
}
