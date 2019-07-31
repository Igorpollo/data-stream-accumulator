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
	logger "github.com/Igorpollo/go-custom-log"
)

var jsonChn = make(chan models.DataPackage, 5000)

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

func writeJSONWorker(f *os.File) {
	for job := range jsonChn {
		//fmt.Println("recebi um job")
		
		jsonData, _ := json.Marshal(job)
		jsonStr := string(jsonData)+"\n"
		_, err := f.WriteString(jsonStr)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}
}

func createWriteJSONWorkers(noOfWorkers int) {
	f, err := os.OpenFile("test.json", os.O_APPEND|os.O_WRONLY,os.ModePerm)
	if err != nil {
		panic(err)
	}
	for i := 0; i < noOfWorkers; i++ {
		go writeJSONWorker(f)
	}
}

func main() {
	go createWriteJSONWorkers(100)
	r := router.New()
	r.POST("/", PutRecord)
	logger.Info("Started at port 8081")
	log.Fatal(fasthttp.ListenAndServe(":8081", r.Handler))
}
