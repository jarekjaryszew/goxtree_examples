package main

import (
	"fmt"

	"github.com/jarekjaryszew/goxtree"
)

type MyRoot struct {
	me any `tag:"div" id:"myroot"`
	_  any `tag:"h1" text:"Todo list"`
	_  any `tag:"p" text:"Add a task"`
	_  any `tag:"input" id:"task_input"`
	_  any `tag:"button" text:"Add" id:"add_button"`
	_  any `tag:"ul" id:"task_list"`
}

type TaskComponent struct {
	me any `tag:"li" id:"task_item"`
	_  any `tag:"span"`
	_  any `tag:"button" text:"Delete" id:"delete_button"`
}

type Task struct {
	Ord  string
	Text string
}

var taskList = make([]Task, 0)

func createTask(task Task, cn *goxtree.CoreNode) {
	taskNode, _ := goxtree.DressDomTree(&TaskComponent{}, string(task.Ord))
	taskNode.SetTextToElementWithId(taskNode.Id, task.Text)

	delete := func() {
		fmt.Println("delete task", task.Ord)
		cn.RemoveChildFromElementWithId("task_list", taskNode.Id)
		for i, t := range taskList {
			if t.Ord == task.Ord {
				taskList = append(taskList[:i], taskList[i+1:]...)
				break
			}
		}
		cn.RenderFromElementWithId("task_list")
	}
	taskNode.AddEventListenerToElementWithId("delete_button"+string(task.Ord), "click", delete)
	cn.AddChildToElementWithId("task_list", taskNode)
	cn.RenderFromElementWithId("task_list")
}

func main() {
	rootNode, _ := goxtree.DressDomTree(&MyRoot{}, "")
	var tid = 1
	tidAddr := &tid
	addTask := func() {
		text := rootNode.ReadValueFromElementWithId("task_input")
		fmt.Println("add task", tid, text)
		createTask(Task{Ord: fmt.Sprint(tid), Text: text}, rootNode)
		(*tidAddr)++
	}
	rootNode.AddEventListenerToElementWithId("add_button", "click", addTask)
	rootNode.MountToNode("root")
	select {}
}
