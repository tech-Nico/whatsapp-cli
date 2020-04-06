/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"sort"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tech-nico/whatsapp-cli/client"
)

// getChatsCmd represents the getChats command
var getChatsCmd = &cobra.Command{
	Use:   "chats",
	Short: "Retrieve the list of chats",
	Long:  `Retrieve the list of chats (1-1 or groups) currently opened`,
	Run:   getChats,
}

func getChats(cmd *cobra.Command, args []string) {
	log.Debug("getChats called")
	wc, err := client.NewClient()
	if err != nil {
		log.Errorf("Error while initializing Whatsapp client: %s", err)
	}
	chats := wc.GetChats()
	type orderedChat struct {
		Name string
		Time int64
	}
	ordered := make([]orderedChat, 0)

	log.Debugf("Chats is %v", chats)
	for _, v := range chats {

		time, err := strconv.ParseInt(v.LastMessageTime, 10, 64)
		if err != nil {
			log.Warnf("Error while converting a timestamp to integer: %s. Skipping this chat...", err)
			continue
		}

		newChat := orderedChat{
			Name: v.Name,
			Time: time,
		}

		ordered = append(ordered, newChat)
	}

	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Time < ordered[j].Time
	})

	for k := range ordered {
		fmt.Printf("%s - %s\n", time.Unix(ordered[k].Time, 0), ordered[k].Name)
	}
}

func init() {
	getCmd.AddCommand(getChatsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getChatsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getChatsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
