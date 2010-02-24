package makengo

import (
	"fmt"
	"reflect"
	"container/vector"
)

type task struct {
	Name        string
	Description string

	block func()
	deps  vector.Vector
}

var TaskManager taskmanager
var lastDescription string

func (self *task) String() string {

	return fmt.Sprintf("Task<name: %s, block: %s, deps: %s", self.Name, self.block, self.deps)
}

func (self *task) DependsOn(args ...) *task {

	v := reflect.NewValue(args).(*reflect.StructValue)
	for i := 0; i < v.NumField(); i++ {
		self.deps.Push((v.Field(i)).(*reflect.StringValue).Get())
	}

	return self
}

func (self *task) Invoke() *task {

	ok := make(chan bool)

	self.deps.Do(func(el interface{}) {
		go func() {
			TaskManager.InvokeByName([]string{el.(string)})
			ok <- true
		}()
	})

	for i := 0; i < len(self.deps); i++ {
		<-ok
	}

	if self.block != nil {
		self.block()
	}

	return self
}

func init() { TaskManager = make(taskmanager) }

func Desc(desc string) { lastDescription = desc }

func Task(name string, block func()) *task {

	t := &task{Name: name, block: block, Description: lastDescription}
	TaskManager[name] = t
	lastDescription = ""

	return t
}

func Default(name string) { Task("Default", nil).DependsOn(name) }
