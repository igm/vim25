package main

import (
	"fmt"
	"log"
)

func init() {
	commands["rpList"] = rpList
}

func rpList() {
	res, err := moList("ResourcePool")
	if err != nil {
		log.Fatal(err)
	}
	for _, rep := range res.Objects {
		fmt.Println(rep.Obj.Type, rep.Obj.Value)
		for _, prop := range rep.PropSet {
			fmt.Printf("\t %s %s\n", prop.Name, prop.Val.Value)
		}
	}
}
