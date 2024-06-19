package main

import (
	"log"

	"github.com/jerricodelacruz/goappinsights/init_sample"
	"github.com/jerricodelacruz/goappinsights/init_sample_second"
)

func init() {
	log.Println("INIT")
}

func main() {
	log.Println("MAIN")

	init_sample.HelloWorld()
	init_sample_second.HelloWorld()
	// // Set environment variables
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Print(err.Error())
	// }

	// muxRouter := mux.NewRouter()

	// appinsights_wrapper.Init(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

	// muxRouter.HandleFunc("/appinsights/{operationName}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)

	// 	operationName := vars["operationName"]

	// 	fmt.Fprintf(w, "CURRENT OPERATION %s! : SECRETS : %s", operationName)
	// 	RunEvent(operationName)
	// })

	// http.Handle("/", muxRouter)

	// http.ListenAndServe(":8080", muxRouter)
}
