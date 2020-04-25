// Copyright 2020 COS Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/jmckind/cos/pkg/cos"
	"github.com/spf13/cobra"
)

var from = ""
var to = ""
var version = false

// New will return a new Command for the application.
func New() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "cos",
		Short: "cos is a Cloud Object Storage utility.",
		Long:  `cos is a Cloud Object Storage utility. It is used to copy objects to and from the cloud.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if version {
				cmd.Println(Version)
				return nil
			}
			return cos.Execute(&cos.Options{
				From: from,
				To:   to,
			})
		},
	}

	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "display version")
	rootCmd.Flags().StringVarP(&from, "from", "f", "", "copy from file")
	rootCmd.MarkFlagFilename("from")

	rootCmd.Flags().StringVarP(&to, "to", "t", "", "copy to file")
	rootCmd.MarkFlagFilename("to")

	rootCmd.SetOutput(os.Stdout)

	return rootCmd
}
