package main

import "net/http"
import "os"
import "os/exec"
import "os/signal"

func main() {

	go func() {
		sigchan := make(chan os.Signal, 2)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		cmd := exec.Command("rm", "tempserver.go")
		cmd.Run()
		os.Exit(0)
	}()

	panic(http.ListenAndServe(":8000", http.FileServer(http.Dir("."))))

}
