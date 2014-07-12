package main

import (
	"fmt"
	ana "github.com/ChimeraCoder/anaconda"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		fmt.Println("Could not load config file: bot.cfg")
		return
	}

	sleepDuration, err := time.ParseDuration(cfg.Bot.Frequency)
	if err != nil {
		fmt.Println("Bot frequency is incorrect, try 10s")
		return
	}

	twitterClient := getTwitterClient(cfg)

	lastChecked := time.Now()

	for {
		bugs, err := GetBugs(lastChecked)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for entry := range bugs.Iter() {
			if entry.Status == "Fix Released" || entry.Status == "Fix Committed" {
				tweet(twitterClient, createTweet(entry.Title, entry.Assignee))

				fmt.Println("Tweeted about bug " + entry.Title)
			}
		}

		fmt.Println("Total Bugs: " + floattostr(bugs.TotalSize))
		fmt.Println("===")

		lastChecked = time.Now()
		time.Sleep(sleepDuration)
	}
}

func getTwitterClient(cfg Config) (api *ana.TwitterApi) {
	ana.SetConsumerKey(cfg.Twitter.ConsumerKey)
	ana.SetConsumerSecret(cfg.Twitter.ConsumerSecret)

	api = ana.NewTwitterApi(cfg.Twitter.AccessTokenKey, cfg.Twitter.AccessTokenSecret)

	return
}

func tweet(api *ana.TwitterApi, tweetStr string) {
	params := url.Values{}

	api.PostTweet(tweetStr, params)
}

func createTweet(bugTitle, bugAssignee string) string {
	bugDesc := strings.Replace(bugTitle, "Bug #", "pad.lv/", 1)

	if bugAssignee == "elementary Devs" {
		return "We fixed bug " + bugDesc
	} else {
		return "pad.lv/~" + bugAssignee + " fixed bug " + bugDesc
	}
}

func floattostr(input_num float64) string {
	return strconv.FormatFloat(input_num, 'g', 1, 64)
}
