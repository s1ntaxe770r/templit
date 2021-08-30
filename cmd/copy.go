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
	"github.com/fatih/color"
	"github.com/s1ntaxe770r/templit/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy a template",
	Long: `copy a template to the specified directory`,
	Example: "templit copy [template name] [destination directory]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2{
			fmt.Println(color.RedString(" two arguments are required. Run templit list to see a list of available templates"))
			os.Exit(1)
		}
		file := args[0]
		dest := args[1]
		directory := viper.Get(file)
		fmt.Println(directory)
		err := utils.CopyTemplate(directory.(string),dest)
		if err != nil {
			fmt.Println(color.RedString(err.Error()))
			os.Exit(1)
		}
		fmt.Println(color.GreenString("successfully copied "+file +" to "+dest))
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
