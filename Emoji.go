package main

import (
	"math/rand"
)

func getEmojis() []string {
	// from http://apps.timwhitlock.info/emoji/tables/unicode
	emojis := []string{"\xF0\x9F\x98\x81", //😁
		"\xF0\x9F\x98\x82", //😂
		"\xF0\x9F\x98\x83", //😃
		"\xF0\x9F\x98\x8B", //😋
		"\xF0\x9F\x98\x9A", //😚
		"\xF0\x9F\x98\x98", //😘
		"\xF0\x9F\x98\xB3", //😳
		"\xF0\x9F\x98\xB8", //😸
		"\xF0\x9F\x98\xB9", //😹
		"\xF0\x9F\x99\x86", //🙆
		"\xE2\x9C\x8C",     //✌
		"\xE2\x98\x9D",     //☝
		"\xF0\x9F\x92\xAA", //💪
		"\xF0\x9F\x92\x99", //💙
	}

	return emojis
}

func getRandomEmoji() string {
	rand.Seed(42) // what is life??

	emojis := getEmojis()

	return emojis[rand.Intn(len(emojis))]
}
