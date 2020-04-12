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
	"github.com/obviyus/goids/loginfo"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log info about all currently running processes",
	Long: `Log the following about every process: Name, PID, background, 
CPU Percent, Running Time, Memory Percent, Status and CPU times in a JSON file
per process. All logs are stored in ./logs/`,
	Run: func(cmd *cobra.Command, args []string) {
		loginfo.LogProcessInfo()
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
