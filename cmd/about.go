/*
Copyright © 2020 Ayaan Zaidi <biblioklept@icloud.com>

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
	"github.com/obviyus/goids/app"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var aboutCmd = &cobra.Command{
	Use:     "about",
	Aliases: []string{},
	Short:   fmt.Sprintf("about %s", app.NAME),
	Long:    `The about command shows some basic information on this tool.`,
	Run: func(*cobra.Command, []string) {
		color.Set(color.FgHiWhite)
		fmt.Print(app.PRETTY_SHORTNAME + " *** ")
		color.Set(color.FgHiGreen)
		fmt.Printf(app.PRETTY_FULLNAME)
		color.Unset()
		fmt.Printf(`
version:         %s
config:          %s
`,
			color.HiGreenString(viper.GetString("version")),
			color.HiGreenString(func() string {
				if len(viper.ConfigFileUsed()) > 0 {
					return viper.ConfigFileUsed()
				}
				return "builtin"
			}()))

		color.Unset()

		fmt.Printf(`
code:            %s
maintainer:      %s
`,
			color.HiBlueString("https://github.com/obviyus/goids"),
			color.HiBlueString("Ayaan Zaidi <biblioklept@icloud.com>"))
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
