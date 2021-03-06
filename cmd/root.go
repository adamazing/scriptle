package cmd

import (
	"fmt"
  "log"
  "os"

	homedir "github.com/mitchellh/go-homedir"
  "github.com/adamazing/scriptle/pkg/scriptle"
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
            // fmt.Println("In the command")

            flagStruct := scriptle.Flags {
              Foo: filePath,
            }

            stat, err := os.Stdin.Stat()
            if err != nil {
              log.Fatal(err)
            }

            if (stat.Mode() & os.ModeCharDevice) == 0 {
              err := scriptle.FromStdin(length, &flagStruct)
              if err != nil {
                log.Println("Error: Could not read stdin.", err)
                os.Exit(1)
              }


            } else if filePath != "" {

            } else {
              log.Println("Error: Could not use stdin or file")
            }


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
  if err != nil {
    fmt.Fprintln(os.Stderr, "Using config file: ", viper.ConfigFileUsed())
  }
}
