package main

func main() {
	done := make(chan struct{})
	go func(chan struct{}) <-chan struct{} {
		work()
		close(done)
		return (done)
	}(done)
	<-done
}
