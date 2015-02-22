package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	ana "github.com/ChimeraCoder/anaconda"
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
		usernames, err := fetchUsernames()
		if err != nil {
			fmt.Println(err)
		}

		for entry := range bugs.Iter() {
			if entry.Status == "Fix Released" || entry.Status == "Fix Committed" {

				if entry.FixDate.After(lastChecked) {
					tweet(twitterClient, createTweet(entry.Title, entry.Assignee, usernames))

					fmt.Println("Tweeted about bug " + entry.Title)
				}
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

func createTweet(bugTitle, bugAssignee string, usernames map[string]string) string {
	bugDesc := strings.Replace(bugTitle, "Bug #", "pad.lv/", 1)

	firstHalf := ""
	if bugAssignee == "elementary Devs" {
		firstHalf = "We fixed bug "
	} else {
		uname := usernameTransform(bugAssignee, usernames)
		firstHalf = uname + " fixed bug "
	}
	return firstHalf + bugDesc
}

func usernameTransform(username string, twitterHandles map[string]string) string {
	if twitterHandles != nil {
		if val, ok := twitterHandles[username]; ok {
			return getRandomEmoji() + val
		}
	}
	return "pad.lv/~" + username
}

func floattostr(input_num float64) string {
	return strconv.FormatFloat(input_num, 'g', 1, 64)
}
