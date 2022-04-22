package main

import "math/rand"

var tags = []string {
	"good",
	"man",
	"woman",
	"nice",
	"rich",
	"coder",
	"model",
}

// return tag1,tag2
func (s *sBench) createTags (num int)(res string){
	var te = make([]string,num)
	for i := 0;i<len(te);i ++ {
		res += tags[rand.Intn(len(tags))] + ","
	}
	return
}