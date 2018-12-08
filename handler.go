package main

import (
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
    Version string  `json:"version"`
    Body    ResBody `json:"response"`
}

type ResBody struct {
    OutputSpeech     Payload `json:"outputSpeech,omitempty"`
    ShouldEndSession bool    `json:"shouldEndSession"`
}

type Payload struct {
    Type string `json:"type,omitempty"`
    Text string `json:"text,omitempty"`
}

func NewResponse(speech string) Response {
    fmt.Println("Generating NewResponse")
    resp := Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: Payload{
				Type: "PlainText",
				Text: speech,
			},
			ShouldEndSession: true,
		},
	}
    fmt.Println(resp)
    return resp
}

func Handler() (Response, error) {
    fmt.Println("Handler")
    return NewResponse("I'm your guy!"), nil
}

func main() {
    fmt.Println("main()")
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(Handler)
    //Handler()
    //fmt.Println(NewResponse("Hello"))
}
