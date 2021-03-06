// Copyright © 2016 Dropbox, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/files"
	"github.com/spf13/cobra"
)

func restore(cmd *cobra.Command, args []string) (err error) {
	path, err := validatePath(args[0])
	if err != nil {
		return
	}

	rev := args[1]

	arg := files.NewRestoreArg()
	arg.Path = path
	arg.Rev = rev

	if _, err = dbx.Restore(arg); err != nil {
		return
	}

	return
}

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore files",
	RunE:  restore,
}

func init() {
	RootCmd.AddCommand(restoreCmd)
}
