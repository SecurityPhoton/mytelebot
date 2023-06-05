/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

func kelvinToCelsius(kStr string) string {
	k, err := strconv.ParseFloat(kStr, 64)
	if err != nil {
		return ""
	}
	c := k - 273.15
	return fmt.Sprintf("%.2f", c)
}

var (
	Teletoken = os.Getenv("TELE_TOKEN")
	WAPI      = os.Getenv("WEATHER_API")
	pattern   = regexp.MustCompile("^[a-zA-Z]{2,12}$")
	city      = "/get Kyiv"
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "this command starts a bot",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:



Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("kbot %s started", appVersion)
		fmt.Printf("TAPI= %s", Teletoken)
		fmt.Printf("WAPI= %s", WAPI)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  Teletoken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf(m.Message().Payload, m.Text())
			payload := m.Message().Payload
			log.Print(payload)

			if pattern.MatchString(payload) {
				city = payload
			}

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I am telebot %s !", appVersion))

			case "info":
				err = m.Send(fmt.Sprintln("This is code for mytelebot. It is writen in Go. It uses 'TELE_TOKEN' env var for telegram bot API token.!"))

			case city:
				cityName := strings.TrimSpace(strings.Replace(m.Text(), "/get", "", 1))
				log.Printf("City = %s", cityName)
				request := gorequest.New()
				_, body, errs := request.Get("http://api.openweathermap.org/data/2.5/weather?q=" + cityName + "&appid=" + WAPI).End()
				if errs != nil {
					log.Println("Error getting weather information:", errs)
					return err
				}
				fmt.Printf("Output %s", body)
				if strings.Contains(body, "404") {
					fmt.Println("The body contains a 404 error")
					response := "Error finding city - maybe you misstyped"
					fmt.Printf("The response is %s", response)
					kbot.Send(m.Sender(), response)
				} else {
					fmt.Println("The body does not contain a 404 error")

					temp := strings.TrimSpace(strings.Split(strings.Split(body, "\"temp\":")[1], ",")[0])
					temp = strings.TrimSuffix(temp, ".")
					celsius := kelvinToCelsius(temp)

					response := "The current temperature in " + cityName + " is " + celsius + " C."
					fmt.Printf("The response is %s", response)
					kbot.Send(m.Sender(), response)
				}

			case "help":
				response := "To see temperature use > get <city name> / get help to see this / get info to see info "
				fmt.Printf("The response is %s", response)
				kbot.Send(m.Sender(), response)
			default:
				response := "Cant understand you type - help - to see help"
				fmt.Printf("The response is %s", response)
				kbot.Send(m.Sender(), response)
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
