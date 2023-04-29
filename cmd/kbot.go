/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

var (
	Teletoken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Aliases: []string{"start"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)

		kbot, err := tele.NewBot(tele.Settings{
			URL:    "",
			Token:  Teletoken,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(tele.OnText, func(ctx tele.Context) error {
			log.Printf(ctx.Message().Payload, ctx.Text())
			payload := ctx.Message().Payload

			switch payload {
			case "hello":
				err = ctx.Send(fmt.Sprintf("Hello I am telebot %s !", appVersion))
			
			case "info":
				err = ctx.Send(fmt.Sprintf("This is bot for education. Bot version %s !", appVersion))
				
			}

			return err
		})

		kbot.Start()

	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
