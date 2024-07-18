package main

import (
	"encoding/json"
	"fmt"

	"github.com/jarekjaryszew/goxtree"
)

type MyButton struct {
	me any `tag:"button" text:"backend" class:"btn" id:"btn1"`
}

type MyRoot struct {
	me any `tag:"div" id:"myroot"`
	_  any `tag:"h1" text:"Hello World!" class:"header" id:"head1"`
	_  any `tag:"div" id:"secondDiv"`
	_  struct {
		_ any `tag:"button" text:"Hi" class:"btn" id:"btn2"`
		_ any `tag:"button" text:"Bye" class:"btn" id:"btn3"`
	} `tag:"div" id:"innerDiv"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	rootNode, _ := goxtree.DressDomTree(&MyRoot{})
	buttonNode, _ := goxtree.DressDomTree(&MyButton{})
	rootNode.AddChildToElementWithId("secondDiv", buttonNode)

	cbf := func(res string) {
		fmt.Print("fetch callback ", res)
		var r Response
		err := json.Unmarshal([]byte(res), &r)
		if err != nil {
			fmt.Println("error", err)
			return
		}
		rootNode.SetTextToElementWithId("head1", r.Message)
		rootNode.Render()
	}

	cb1 := func() {
		goxtree.Fetch("something.json", cbf, nil)
	}

	buttonNode.AddEventListenerToElementWithId("btn1", "click", cb1)

	cb2 := func() {
		go func() {
			rootNode.SetTextToElementWithId("head1", "Hi")
			rootNode.Render()
		}()
	}
	rootNode.AddEventListenerToElementWithId("btn2", "click", cb2)

	cb3 := func() {
		go func() {
			rootNode.SetTextToElementWithId("head1", "Bye")
			rootNode.Render()
		}()
	}
	rootNode.AddEventListenerToElementWithId("btn3", "click", cb3)

	rootNode.MountToNode("root")

	select {}
}
