package main

import (
	. "specify"
	"time"
	t "../src/testmakengo"
)

func init() {
	Describe("makengo.Task", func() {
		
		It("should add a new task", func(e Example) {
			t.Task("NewTask", func() {})
		})

		It("should add dependencies to task", func(e Example) {
			t.Task("TestTask1", func() {}).DependsOn("TestTask2")
		})

		It("should invoke a task", func(e Example) {
			var value int = 0

			t.Task("TestTask", func() { value = 1 }).Invoke()			
			e.Value(value).Should(Be(1))			
		})

		It("should invoke a task with all its prerequisites", func(e Example) {
			var value1 int = 0
			var value2 int = 1

			t.Task("TestTask1", func() { value1++ })
			t.Task("TestTask2", func() { value2++ }).DependsOn("TestTask1").Invoke()

			e.Value(value1).Should(Be(1))
			e.Value(value2).Should(Be(2))
		})

		It("should invoke a task with all its prerequisites recursively", func(e Example) {
			var value1 int = 0
			var value2 int = 1
			var value3 int = 2

			t.Task("TestTask1", func() { value1++ })
			t.Task("TestTask2", func() { value2++ }).DependsOn("TestTask1")
			t.Task("TestTask3", func() { value3++ }).DependsOn("TestTask2").Invoke()

			e.Value(value1).Should(Be(1))
			e.Value(value2).Should(Be(2))
			e.Value(value3).Should(Be(3))
		})

		It("should run tasks concurrently", func(e Example) {
			var value string

			t.Task("TestTask1", func() { 
				time.Sleep(1 * 1e9)
				value += "foo"
			})

			t.Task("TestTask2", func() { 
				time.Sleep(0.5 * 1e9)
				value += "bar" 
			})
			
			t.TaskManager.InvokeByName([]string { "TestTask1", "TestTask2" })

			e.Value(value).Should(Be("barfoo"))
		})

		It("should wait for dependencies to finish", func(e Example) {
			var value string

			t.Task("TestTask2", func() { 
				time.Sleep(1 * 1e9)
				value += "foo"
			})

			t.Task("TestTask3", func() { 
				time.Sleep(0.5 * 1e9)
				value += "bar" 
			})

			t.Task("TestTask1", func() { value += "biz" }).DependsOn("TestTask2", "TestTask3").Invoke()
			e.Value(value).Should(Be("barfoobiz"))
		})

		It("should not invoke the task if the block argument is nil", func(e Example) {
			t.Task("TestTask", nil).Invoke()
		})

		
	})

	Describe("makengo.Describe", func() {
		
		It("should associate the given description to the task", func(e Example) {

			t.Describe("A nice task")
			e.Value(t.Task("NewTask", func() {}).Description).Should(Be("A nice task"))

			t.Describe("Yet another nice task")
			e.Value(t.Task("NewTask 2", func() {}).Description).Should(Be("Yet another nice task"))

			e.Value(t.Task("NewTask 3", func() {}).Description).Should(Be(""))

		})

	})

	Describe("makengo.Default", func() {
		
		It("should define a default task", func(e Example) {
			t.Task("NewTask", func() {})
			t.Default("NewTask")
		})

	})


}

