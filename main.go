package main

import (
    "fmt"
    "time"
    "net/url"
    "strings"
    "strconv"
    ana "github.com/ChimeraCoder/anaconda"
)

func main() {
    twitterClient := getTwitterClient()
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
        time.Sleep(10 * time.Second)
    }
}

func getTwitterClient() *ana.TwitterApi {
    consumer_key := "Q1yiFkN6ednf1SXhXhvsYQ"
    consumer_secret := "IFMrGJAdjf3uxAmku19KbBit7w0D2XXgdd7f2i89ZA"
    access_token_key := "1308685658-9TQKkZdrywUa6cbad9OQUemZwHfzteMwg7um92h"
    access_token_secret := "4Rarhzy3nn2E8RU7IaUWqr259MBGlGA0hIoeqTr0"

    ana.SetConsumerKey(consumer_key)
    ana.SetConsumerSecret(consumer_secret)

    api := ana.NewTwitterApi(access_token_key, access_token_secret)

    return api
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
