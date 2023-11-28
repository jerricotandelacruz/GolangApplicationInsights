package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/jerricodelacruz/goappinsights/appinsights_wrapper"
)

func f(num int) {
	for i := 0; i < 5; i++ {
		RunFirst(num, i)

		source := rand.NewSource(time.Now().UnixNano())
		randomGenerator := rand.New(source)
		randomNumber := randomGenerator.Intn(5) + 1
		duration := time.Duration(randomNumber) * time.Second
		time.Sleep(duration)
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

	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	randomNumber := randomGenerator.Intn(5) + 1
	duration := time.Duration(randomNumber) * time.Second
	time.Sleep(duration)

	RunSecond(parentNum, num)

	RunThird(parentNum, num)

	client.EndOperation()
}

func RunSecond(parentNum, num int) {
	client := appinsights_wrapper.Client()

	client.TrackEvent(fmt.Sprintf("SECOND NO:%d:::%d", parentNum, num))

	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	randomNumber := randomGenerator.Intn(5) + 1
	duration := time.Duration(randomNumber) * time.Second
	time.Sleep(duration)

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
