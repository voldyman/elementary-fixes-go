package main

import (
	"math/rand"
)

func getEmojis() []string {
	// from http://apps.timwhitlock.info/emoji/tables/unicode
	emojis := []string{"\xF0\x9F\x98\x81", //ð
		"\xF0\x9F\x98\x83", //ð
		"\xF0\x9F\x98\x8D", //ð
		"\xF0\x9F\x98\xB8", //ð
		"\xF0\x9F\x98\xBA", //ðº
		"\xF0\x9F\x99\x86", //ð
		"\xF0\x9F\x99\x8C", //ð
		"\xE2\x9C\x8C",     //â
		"\xE2\x9C\x94",     //â
		"\xE2\x98\xBA",     //âº
		"\xE2\x98\x9D",     //â
		"\xE2\x99\xA0",     //â
		"\xE2\x9A\xA1",     //â¡
		"\xE2\xAD\x90",     //­
		"\xF0\x9F\x8D\x95", //🍕
        "\xF0\x9F\x8D\x97", //🍗
	}

	return emojis
}

func getRandomEmoji() string {
	rand.Seed(42) // what is life??

	emojis := getEmojis()

	return emojis[rand.Intn(len(emojis))]
}