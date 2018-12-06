// Copyright © 2018 Hugo Torres <hugo.torres1993@gmail.com>
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
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode STRING",
	Short: "URL-encode an input string",
	Long:  `That's it. URL-encode a string.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := ""
		if len(args) < 1 {
			in := bufio.NewReader(os.Stdin)
			input, _ = in.ReadString('\n')
		} else {
			input = args[0]
		}

		fmt.Println(url.QueryEscape(input))
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
