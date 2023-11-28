package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jerricodelacruz/goappinsights/appinsights_wrapper"
)

func main() {
	appinsights_wrapper.Init(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	for i := 0; i < 10; i++ {
		RunFirst(i)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func RunFirst(num int) {
	client := appinsights_wrapper.Client()

	client.StartOperation(fmt.Sprintf("OPERATION CORRELATION:%d", num))

	client.TrackEvent(fmt.Sprintf("FIRST NO:%d", num))

	RunSecond(num)

	RunThird(num)

	client.EndOperation()
}

func RunSecond(num int) {
	client := appinsights_wrapper.Client()

	client.TrackEvent(fmt.Sprintf("SECOND NO:%d", num))

	RunSecondFirst(num)
}

func RunSecondFirst(num int) {
	client := appinsights_wrapper.Client()

	client.TrackEvent(fmt.Sprintf("SECOND FIRST NO:%d", num))
}

func RunThird(num int) {
	client := appinsights_wrapper.Client()

	client.TrackEvent(fmt.Sprintf("THIRD NO:%d", num))
}
