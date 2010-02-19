package makengo

type taskmanager map [string] *task

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
