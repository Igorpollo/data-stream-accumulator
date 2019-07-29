package main

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type dataRecieve struct {
	body interface{}
}

func Index(ctx *fasthttp.RequestCtx) {
	var result interface{}
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(ctx.PostBody(), &result)
	str := fmt.Sprintf("%v", result)
	fmt.Println(str)
	ctx.WriteString("Welcome!")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	r := router.New()
	r.POST("/", Index)
	r.GET("/hello/:name", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}