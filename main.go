package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

const (
	DTYPE_JSON = iota
	DTYPE_STRING
	DTYPE_ARRAY_STRING
	DTYPE_IMAGE
	DTYPE_AUDIO
	DTYPE_OTHER
)

// TODO: add tag to json
type dataPackage struct {
	uuid             string
	label            string
	dateTimeReceived time.Time
	dataSize         uintptr
	dataType         int
	data             string
	ownerID          int
	ip               net.IP
}

// asdfasdf
func Index(ctx *fasthttp.RequestCtx) {

	dataBody := ctx.PostBody()

	dataStructure := dataPackage{
		dateTimeReceived: time.Now(),
		dataSize:         unsafe.Sizeof(dataBody),
		dataType:         DTYPE_JSON,
		data:             base64.StdEncoding.EncodeToString(dataBody),
		ownerID:          1,
		ip:               ctx.RemoteIP(),
		uuid:             "asdfçlkasjdf",
	}

	fmt.Println("The JSON data is:")

	spew.Dump(dataStructure)
	jsonStr := string(dataStructure)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

	f, err := os.Create("test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(jsonStr)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	ctx.WriteString("Welcome!")

	ctx.WriteString("Welcome!")
}

// asdf
func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	r := router.New()
	r.POST("/", Index)
	r.GET("/hello/:name", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
