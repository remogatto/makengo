package main

import (
	. "specify"
	t "../src/testmakengo"
)

func init() {
	Describe("makengo.FileList", func() {

		It("should create a FileList instance", func(e Example) {
			t.FileList("./", "go")
		})

		It("should return a slice", func(e Example) {
			e.Value(len(t.FileList("./", "go").ToSlice()) > 0).Should(Be(true))
		})

		It("should catch errors in Errors field", func(e Example) {
			fl := t.FileList("doesntexist", "go")
			fl.ToSlice()
			if _, ok := <-fl.Errors; ok {}			
		})


	})

}
