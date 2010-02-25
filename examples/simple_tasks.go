package main

import (
	. "makengo"
	"fmt"
	"time"
)

func init() {

	Desc("Print Hello")
	Task("hello", func() { fmt.Print("Hello ") })

	Desc("Print Hello Joe!")
	Task("joe", func() { fmt.Println("Joe!") }).DependsOn("hello")

	Task("task_1", func() {
		time.Sleep(5 * 1e9)
		fmt.Println("task_1 has finished")
	})

	Task("task_2", func() {
		time.Sleep(1 * 1e9)
		fmt.Println("task_2 has finished")
	})

        Default("joe")
}


