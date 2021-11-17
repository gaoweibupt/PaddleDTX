// Copyright (c) 2021 PaddlePaddle Authors. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodes

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	httpclient "github.com/PaddlePaddle/PaddleDTX/xdb/client/http"
	"github.com/PaddlePaddle/PaddleDTX/xdb/pkgs/crypto/ecdsa"
)

var (
	description string
)

// addNodeCmd adds new node
var addNodeCmd = &cobra.Command{
	Use:   "add",
	Short: "add node into xuper db",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := httpclient.New(host)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		privkey, err := ecdsa.DecodePrivateKeyFromString(privateKey)
		if err != nil {
			fmt.Printf("failed to decode private key, err: %v\n", err)
			return
		}
		opt := httpclient.AddNodeOptions{
			Name:       name,
			Address:    address,
			Online:     true,
			PrivateKey: privkey,
		}

		if err := client.AddNode(context.Background(), opt); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Println("node added")
	},
}

func init() {
	rootCmd.AddCommand(addNodeCmd)

	addNodeCmd.Flags().StringVarP(&privateKey, "privateKey", "k", "", "private key")
	addNodeCmd.Flags().StringVarP(&address, "address", "a", "", "address")
	addNodeCmd.Flags().StringVarP(&name, "name", "n", "", "node name")
	addNodeCmd.Flags().StringVarP(&description, "description", "d", "", "description")

	addNodeCmd.MarkFlagRequired("privateKey")
	addNodeCmd.MarkFlagRequired("name")
	addNodeCmd.MarkFlagRequired("address")
	addNodeCmd.MarkFlagRequired("host")
}