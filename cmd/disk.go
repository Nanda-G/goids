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

	"github.com/obviyus/goids/fetch"
	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "List basic information about disk",
	Long:  `List Total, Used and Free space on all disks.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := fetch.DiskInfo()
		if err != nil {
			fmt.Println("Calling DiskInfo failed: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)
}
