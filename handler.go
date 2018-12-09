package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/arienmalec/alexa-go"
)

func Handler(request alexa.Request) (alexa.Response, error) {
    var response alexa.Response
    var err error = nil

	switch request.Body.Intent.Name {
	case "JokeIntent":
		response = handleJoke()
	case alexa.HelpIntent:
    default:
		response = handleHelp()
	}

	return response, err
    //return alexa.NewSimpleResponse("Saying Hello", "I'm your guy!"), nil
}

func handleHelp() alexa.Response {
	return alexa.NewSimpleResponse("Help for My Guy", "I'm your guy. Ask me for a joke.")
}

func handleJoke() alexa.Response {
    return alexa.NewSimpleResponse("Joke Response", "My grandfather has the heart of a lion, and lifetime ban at the zoo.")
}

func main() {
    // Make the handler available for Remote Procedure Call by AWS Lambda
    lambda.Start(Handler)
}
