package makengo

import (
	"fmt"
	"container/vector"
)

type task struct {
	Name        string
	Description string

	block func()
	deps  vector.StringVector
}

var TaskManager taskmanager
var lastDescription string

func (self *task) String() string {

	return fmt.Sprintf("Task<name: %s, block: %s, deps: %s", self.Name, self.block, self.deps)
}

func (self *task) DependsOn(args ...string) *task {

	for _, taskname := range args {
		self.deps.Push(taskname)
	}

	return self
}

func (self *task) Invoke() *task {

	ok := make(chan bool)

	self.deps.Do(func(el string) {
		go func() {
			TaskManager.InvokeByName([]string{el})
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

func Task(name string, block func()) (t *task) {
	t = &task{Name: name, block: block, Description: lastDescription}
	TaskManager[name] = t
	lastDescription = ""
	return
}

func Default(name string) { Task("Default", nil).DependsOn(name) }
