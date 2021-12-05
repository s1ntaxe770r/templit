/*
Copyright © 2021 NAME  hello@jubril.me

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
	"github.com/s1ntaxe770r/templit/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	TemplateName string
	addfile         string

)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new template",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		dest := utils.GetTemplateDir()+ addfile
		err := utils.CopyTemplate(addfile, dest)
		if err != nil {
			fmt.Println(color.RedString(err.Error()))
			os.Exit(1)
		}
		viper.Set(TemplateName, dest)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(color.RedString(err.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&TemplateName, "template-name", "n", "", "assigns a name to the the file being added as a template")
	addCmd.Flags().StringVarP(&addfile, "file", "f", "", "path to the file you are trying to copy")
	//addCmd.Flags().StringVarP(&remote,"remote","f","","url of the remote template you wish to add")
	addCmd.MarkFlagRequired("file")
	addCmd.MarkFlagRequired("template-name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
https://api.github.com/repos/s1ntaxe770r/pawxi/releases/assets/33429151
https://api.github.com/repos/s1ntaxe770r/pawxi/releases/39769313/assets
https://github.com/meshery/meshery/blob/master/mesheryctl/internal/cli/root/pattern/list.go

*/
