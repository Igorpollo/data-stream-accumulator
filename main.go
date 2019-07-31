package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"data-stream/models"
	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

var jsonChn = make(chan models.DataPackage, 500)

const (
	DTYPE_JSON = iota
	DTYPE_STRING
	DTYPE_ARRAY_STRING
	DTYPE_IMAGE
	DTYPE_AUDIO
	DTYPE_OTHER
)

// PutRecord recieve a single record
func PutRecord(ctx *fasthttp.RequestCtx) {

	dataBody := ctx.PostBody()

	dataStructure := models.DataPackage{
		DateTimeReceived: time.Now(),
		DataSize:         binary.Size(dataBody),
		DataType:         DTYPE_JSON,
		Data:             base64.StdEncoding.EncodeToString(dataBody),
		OwnerID:          1,
		IP:               ctx.RemoteIP(),
		UUID:             uuid.New(),
	}
	jsonChn <- dataStructure

	ctx.WriteString("Welcome!")

}

func writeJSONWorker() {
	for job := range jsonChn {
		fmt.Println("recebi um job")
		f, err := os.OpenFile("test.json", os.O_APPEND|os.O_WRONLY,os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonData, _ := json.Marshal(job)
		jsonStr := string(jsonData)+"\n"
		_, err = f.WriteString(jsonStr)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}
}

func createWriteJSONWorkers(noOfWorkers int) {
	for i := 0; i < noOfWorkers; i++ {
		go writeJSONWorker()
	}
}

func main() {
	go createWriteJSONWorkers(1)
	r := router.New()
	r.POST("/", PutRecord)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
