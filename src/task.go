package makengo

import (
	"fmt"
	"reflect"
	"container/vector"
)

type task struct {
	Name string
	block func()
	deps vector.Vector
	Description string
}

type taskmanager map [string] *task

var TaskManager taskmanager
var lastDescription string

func init() {
	TaskManager = make(taskmanager) 
}

func (self taskmanager) InvokeByName(tasknames []string) {
	ok := make(chan bool)
	for _, taskname := range tasknames {
		go func(taskname string) {
			self[taskname].Invoke()
			ok <- true
		}(taskname)
	}
	for i := 0; i < len(tasknames); i++ {
		<-ok
	}
}

func Task(name string, block func()) (*task) {
	
	t := &task{ Name: name, block: block, Description: lastDescription }
	TaskManager[name] = t
	lastDescription = ""

	return t
}

func (self *task) String() (string) {

	return fmt.Sprintf("Task<name: %s, block: %s, deps: %s", self.Name, self.block, self.deps)
}

func (self *task) DependsOn(args ...) (*task) {

	v := reflect.NewValue(args).(*reflect.StructValue)
	for i := 0; i < v.NumField(); i++ { self.deps.Push((v.Field(i)).(*reflect.StringValue).Get()) }

	return self
}

func (self *task) Invoke() (*task) {

	ok := make(chan bool)

	self.deps.Do(func (el interface{}) { 
		go func() {
			TaskManager.InvokeByName([]string { el.(string) }) 
			ok <- true
		}()
	})

	for i := 0; i < len(self.deps); i++ { <-ok }

	self.block()

	return self
}


func Describe(desc string) {
	lastDescription = desc
}
