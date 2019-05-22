// Copyright © 2019 Eric Freeman Zhou <neochau@qq.com>
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
	"github.com/dejavuzhou/felix/ginbro"
	"github.com/spf13/cobra"
	"path"
)

// ginbinCmd represents the ginbin command
var ginbinCmd = &cobra.Command{
	Use:   "ginbin",
	Short: "Ginbin allows you to embed a directory of static files into your Go binary to be later served from github.com/gin-goin/gin,used like gin middleware",
	Long: `Is this a crazy idea? No, not necessarily.
If you're building a tool that has a Web component,
you typically want to serve some images, CSS and JavaScript.
You like the comfort of distributing a single binary,
so you don't want to mess with deploying them elsewhere.
If your static files are not large in size and will be browsed by a few people,
ginbin is a solution you are looking for
`,
	Run: func(cmd *cobra.Command, args []string) {
		ginbro.RunGinStatic(flagSrc, flagDest, flagTags, flagPkg, flagPkgCmt, flagNoMtime, flagNoCompress, flagForce)
	},
}
var (
	flagSrc, flagDest, flagTags, flagPkg, flagPkgCmt string
	flagNoMtime, flagNoCompress, flagForce           bool
)

func init() {
	rootCmd.AddCommand(ginbinCmd)

	ginbinCmd.Flags().StringVarP(&flagSrc, "src", "s", path.Join(".", "dist"), "The path of the source directory.")
	ginbinCmd.Flags().StringVarP(&flagDest, "dest", "d", ".", "The destination path of the generated package.")
	ginbinCmd.Flags().StringVarP(&flagTags, "tags", "t", "", "The golang tags.")
	ginbinCmd.Flags().StringVarP(&flagPkg, "package", "p", "felixbin", "The destination path of the generated package.")
	ginbinCmd.Flags().StringVarP(&flagPkgCmt, "comment", "c", "", "The package comment. An empty value disables this comment.")
	ginbinCmd.Flags().BoolVarP(&flagNoCompress, "zip", "z", false, "Do not use compression to shrink the files.")
	ginbinCmd.Flags().BoolVarP(&flagNoMtime, "mtime", "m", false, "Ignore modification times on files.")
	ginbinCmd.Flags().BoolVarP(&flagForce, "force", "f", true, "Overwrite destination file if it already exists.")
}
