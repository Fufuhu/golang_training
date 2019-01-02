package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/gonids"
)

func main() {
	buf, err := ioutil.ReadFile("./rules/snort3-community.rules")

	if err != nil {
		panic(err)
	}

	var rules []*gonids.Rule
	for _, v := range strings.Split(string(buf), "\n") {
		//fmt.Println(v)
		r, err := gonids.ParseRule(v)
		if err != nil {
			//fmt.Println("parse error!")
			continue
		} else {
			if r.Action != "" {
				rules = append(rules, r)
			}
		}
	}

	for _, v := range rules {
		//fmt.Printf("%v\n", v)
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n", v.Description, v.Action, v.Protocol, v.Source, v.Destination)
	}
}
