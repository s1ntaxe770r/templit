/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "templit",
	Short: "boring template manager",
	Long:  `templit is a command line application that helps you manage various templates and config files so you don't have too`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/templit/templit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	homedir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".templit" (without extension).

		viper.AddConfigPath(homedir+"/.config/templit")
		viper.SetConfigType("yaml")
		viper.SetConfigName("templit")
		if err := viper.SafeWriteConfigAs(homedir+"/.config/templit/templit.yaml"); err != nil {
			if os.IsNotExist(err) {
				err = viper.WriteConfigAs(homedir+"/.config/templit/templit.yaml")
				if err != nil {
					fmt.Println(color.RedString(err.Error()))
				}
			}
		}

	}

	if _, err := os.Stat(homedir+"/.config/templit/templates"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println(color.HiYellowString("{info:} template directory does not exist. Creating one at" + homedir + "/.config/templit/templates"))
			err = os.MkdirAll(homedir+"/.config/templit/templates", 0755)
			if err != nil {
				fmt.Println(color.RedString(err.Error()))
				os.Exit(1)
			}

		} else {
			fmt.Println(color.RedString("failed to create templates directory, " + err.Error()))
		}
		fmt.Println(color.GreenString("done ✔"))
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, err2 := fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err2 != nil {
			return 
		}
	}

}
