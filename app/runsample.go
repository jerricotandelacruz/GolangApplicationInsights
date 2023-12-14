package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jerricodelacruz/goappinsights/appinsights_wrapper"
	"github.com/jerricodelacruz/goappinsights/dummy"
)

func f(num int) {
	for i := 0; i < 5; i++ {
		RunFirst(num, i)

		// source := rand.NewSource(time.Now().UnixNano())
		// randomGenerator := rand.New(source)
		// randomNumber := randomGenerator.Intn(5) + 1
		// duration := time.Duration(randomNumber) * time.Second
		// time.Sleep(duration)
	}
}

func RunTimeFirstSample() {
	fmt.Println("==============================")

	skip := 0

	for {
		pc, file, line, ok := runtime.Caller(skip)

		if !ok {
			break
		}

		fmt.Println("PC : ", pc)
		fmt.Println("FILE : ", file)
		fmt.Println("LINE : ", line)
		fmt.Println("OK : ", ok)
		fmt.Println("FUNC :", runtime.FuncForPC(pc).Name())

		skip++
	}

	RunTimeSecondSample()

	dummy.RunTimeOutFirstSample()
}

func RunTimeSecondSample() {
	fmt.Println("SECOND")
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println("PC : ", pc)
	fmt.Println("FILE : ", file)
	fmt.Println("LINE : ", line)
	fmt.Println("OK : ", ok)
}

func RunEvent(operationName string) {
	client := appinsights_wrapper.NewClient()

	client.StartOperation(fmt.Sprintf("OPERATION NAME : %s", operationName))

	client.TrackEvent(fmt.Sprintf("EVENT FIRST : %s", operationName))

	RunSecondEvent(client, operationName)

	client.EndOperation()
}

func RunSecondEvent(client *appinsights_wrapper.TelemetryClient, operationName string) {
	time.Sleep(5 * time.Second)
	client.TrackEvent(fmt.Sprintf("EVENT SECOND : %s", operationName))
}

func RunFirst(parentNum, num int) {
	client := appinsights_wrapper.NewClient()

	client.StartOperation(fmt.Sprintf("OPERATION CORRELATION %d:::%d", parentNum, num))

	client.TrackEvent(fmt.Sprintf("FIRST NO:%d:::%d", parentNum, num))

	// source := rand.NewSource(time.Now().UnixNano())
	// randomGenerator := rand.New(source)
	// randomNumber := randomGenerator.Intn(5) + 1
	// duration := time.Duration(randomNumber) * time.Second
	// time.Sleep(duration)

	RunSecond(parentNum, num)

	RunThird(parentNum, num)

	client.EndOperation()
}

func RunSecond(parentNum, num int) {
	client := appinsights_wrapper.NewClient()

	client.TrackEvent(fmt.Sprintf("SECOND NO:%d:::%d", parentNum, num))

	// source := rand.NewSource(time.Now().UnixNano())
	// randomGenerator := rand.New(source)
	// randomNumber := randomGenerator.Intn(5) + 1
	// duration := time.Duration(randomNumber) * time.Second
	// time.Sleep(duration)

	RunSecondFirst(parentNum, num)
}

func RunSecondFirst(parentNum, num int) {
	client := appinsights_wrapper.NewClient()

	client.TrackEvent(fmt.Sprintf("SECOND FIRST NO:%d:::%d", parentNum, num))
}

func RunThird(parentNum, num int) {
	client := appinsights_wrapper.NewClient()

	client.TrackEvent(fmt.Sprintf("THIRD NO:%d:::%d", parentNum, num))
}
