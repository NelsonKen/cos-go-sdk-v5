// Copyright © 2018 ken Kang <gengkang@live.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"

	"github.com/fpay/cos-go-sdk-v5/cli/internal/services"
	"github.com/fpay/cos-go-sdk-v5/cos"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if src == "" {
			fmt.Println("--src option is required")
			cmd.Help()
			return
		}
		if dst == "" {
			fmt.Println("--dst option is required")
			cmd.Help()
			return
		}
		if bucket == "" {
			fmt.Println("--bucket option is required")
			cmd.Help()
			return
		}

		o := bootstrap()
		client := cos.New(&o.Cos)
		us, err := services.NewUploadService(client, bucket)
		if err != nil {
			log.Fatalf("init upload service failed %v", err)
		}

		err = us.Upload(src, dst)
		if err != nil {
			log.Fatalf("upload failed %v", err)
		}

		log.Println("success")
	},
}

func init() {
	uploadCmd.PersistentFlags().StringVar(&src, "src", "", "filepath at local")
	uploadCmd.PersistentFlags().StringVar(&dst, "dst", "", "filename at remote")
	uploadCmd.PersistentFlags().StringVar(&bucket, "bucket", "", "bucket at remote")
	RootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
