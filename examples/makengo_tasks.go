package main

import (
	. "makengo"
	"os"
	"fmt"
	"path"
)

var arch = map[string]string { "amd64": "6", "386": "8", "arm": "5" }[os.Getenv("GOARCH")]
var basePath, _ = os.Getwd()

func init() {

	Desc("Compile makengo")
	Task("compile", func() {

		os.Chdir("src")
		System(fmt.Sprintf("%sg -o makengo.%s *.go", arch, arch))
		System(fmt.Sprintf("gopack grc makengo.a makengo.%s", arch))

	})

	Desc("Package makengo")
	Task("package", func() {

		os.Chdir("src")
		System(fmt.Sprintf("gopack grc makengo.a makengo.%s", arch))

	}).DependsOn("compile")

	Desc("Cleanup folders")
	Task("clean", func() {

		os.Chdir("src")
		System("rm *.[a68]")
		System("rm -rf _test/")

	})

	Desc("Run specs")
	Task("spec", func() {

		os.Chdir(path.Join(basePath, "spec"))
		System("specify *.go")

	}).DependsOn("testpackage")

	Task("testpackage", func() {

		System(fmt.Sprintf("gopack grc testmakengo.a testmakengo.%s", arch))

	}).DependsOn("testpackage.arch")

	Task("testpackage.arch", func() {

		for _, src := range FileList(path.Join(basePath, "src"), "src/[a-z]+\\.go$").ToSlice() {
			_, basename := path.Split(src)
			System(fmt.Sprintf("sed -e 's/package makengo/package testmakengo/' < %s > _test/%s", src, basename))
		}
		System(fmt.Sprintf("%sg -o testmakengo.%s _test/*.go", arch, arch))

	}).DependsOn("_test")

	Task("_test", func() {
		os.Chdir(path.Join(basePath, "src"))
		os.Mkdir("_test", 0777)
	})

        Default("spec")
}

