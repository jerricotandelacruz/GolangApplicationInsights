package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jerricodelacruz/goappinsights/appinsights_wrapper"
)

func f(num int) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		RunFirst(num, i)
	}
}

func main() {
	appinsights_wrapper.Init(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	for i := 0; i < 3; i++ {
		go f(i)
	}

	f(4)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func RunFirst(parentNum, num int) {
	client := appinsights_wrapper.Client()

	client.StartOperation(fmt.Sprintf("OPERATION CORRELATION %d:::%d", parentNum, num))

	client.TrackEvent(fmt.Sprintf("FIRST NO:%d:::%d", parentNum, num))

	RunSecond(parentNum, num)

	RunThird(parentNum, num)

	client.EndOperation()
}

func RunSecond(parentNum, num int) {
	client := appinsights_wrapper.Client()

	client.TrackEvent(fmt.Sprintf("SECOND NO:%d:::%d", parentNum, num))

	RunSecondFirst(parentNum, num)
}

func RunSecondFirst(parentNum, num int) {
	client := appinsights_wrapper.Client()

	client.TrackEvent(fmt.Sprintf("SECOND FIRST NO:%d:::%d", parentNum, num))
}

func RunThird(parentNum, num int) {
	client := appinsights_wrapper.Client()

	client.TrackEvent(fmt.Sprintf("THIRD NO:%d:::%d", parentNum, num))
}
