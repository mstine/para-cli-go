/*
Copyright © 2022 Matt Stine

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/mstine/para-cli/root"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:          "add <name> <path>",
	Aliases:      []string{"a"},
	Short:        "Add a new PARA Root",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var prl root.ParaRootList
		if err := viper.Unmarshal(&prl); err != nil {
			return err
		}
		if err := addAction(os.Stdout, &prl, args); err != nil {
			return err
		}
		viper.Set("roots", prl.Roots)
		return viper.WriteConfig()
	},
}

func addAction(out io.Writer, prl *root.ParaRootList, args []string) error {
	root := root.ParaRoot{
		Name: args[0],
		Path: args[1],
	}
	if err := prl.Add(root); err != nil {
		return err
	}
	fmt.Fprintln(out, "Added PARA Root: ", root)
	return nil
}

func init() {
	paraRootCmd.AddCommand(addCmd)
}
