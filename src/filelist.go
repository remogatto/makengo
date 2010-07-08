package makengo

import (
	"path"
	"os"
	"regexp"
	"container/vector"
)

type fileList struct {
	basePath  string
	pattern   string
	filenames vector.StringVector
	Errors    chan os.Error
}

func FileList(basePath, pattern string) *fileList {
	return &fileList{basePath: basePath, pattern: pattern, Errors: make(chan os.Error, 64)}
}

func (self *fileList) ToSlice() []string {
	path.Walk(self.basePath, self, self.Errors)
	return self.filenames.Data()
}

func (self *fileList) VisitDir(currpath string, d *os.FileInfo) bool {
	return true
}

func (self *fileList) VisitFile(currPath string, d *os.FileInfo) {
	match, _ := regexp.MatchString(self.pattern, currPath)
	if match {
		self.filenames.Push(currPath)
	}
}
