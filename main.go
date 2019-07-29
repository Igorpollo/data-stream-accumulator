package main

import (
	"fmt"
	"log"
	"encoding/json"
	"os"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type dataRecieve struct {
	body interface{}
}

func Index(ctx *fasthttp.RequestCtx) {
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(ctx.PostBody()), &result)
	
	empData, err := json.Marshal(result)   
    if err != nil {
        fmt.Println(err.Error())
        return
	}

	jsonStr := string(empData)
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