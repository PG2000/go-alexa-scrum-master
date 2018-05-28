package alexa

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
)

var a *alexa.Alexa

const cardTitle = "ScrumMaster"

// ScrumMaster handles reqeusts from the ScrumMaster skill.
type ScrumMaster struct{}

// Handle processes calls from Lambda
func Handle(ctx context.Context, requestEnv *alexa.RequestEnvelope) (interface{}, error) {
	return a.ProcessRequest(ctx, requestEnv)
}

// OnSessionStarted called when a new session is created.
func (h *ScrumMaster) OnSessionStarted(ctx context.Context, request *alexa.Request, session *alexa.Session, ctx_ptr *alexa.Context, response *alexa.Response) error {

	log.Printf("OnSessionStarted requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}

// OnLaunch called with a reqeust is received of type LaunchRequest
func (h *ScrumMaster) OnLaunch(ctx context.Context, request *alexa.Request, session *alexa.Session, ctx_ptr *alexa.Context, response *alexa.Response) error {
	speechText := "Welcome to the Alexa Skills Kit, you can say hello"

	log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	response.SetSimpleCard(cardTitle, speechText)
	response.SetOutputText(speechText)
	response.SetRepromptText(speechText)

	response.ShouldSessionEnd = true

	return nil
}

// OnIntent called with a reqeust is received of type IntentRequest
func (h *ScrumMaster) OnIntent(ctx context.Context, request *alexa.Request, session *alexa.Session, ctx_ptr *alexa.Context, response *alexa.Response) error {

	log.Printf("OnIntent requestId=%s, sessionId=%s, intent=%s", request.RequestID, session.SessionID, request.Intent.Name)

	switch request.Intent.Name {
	case "HelloWorldIntent":
		log.Println("HelloWorldIntent triggered")
		speechText := "Hello World"

		response.SetSimpleCard(cardTitle, speechText)
		response.SetOutputText(speechText)

		log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
	case "AMAZON.HelpIntent":
		log.Println("AMAZON.HelpIntent triggered")
		speechText := "You can say hello to me!"

		response.SetSimpleCard("ScrumMaster", speechText)
		response.SetOutputText(speechText)
		response.SetRepromptText(speechText)
	case "GetSprintSummary":
		log.Println("GetSprintSummary triggered")
		speechText := "Thanks for asking"
		response.SetSimpleCard("ScrumMaster", speechText)
		response.SetOutputText(speechText)
	default:
		return errors.New("Invalid Intent")
	}

	return nil
}

// OnSessionEnded called with a reqeust is received of type SessionEndedRequest
func (h *ScrumMaster) OnSessionEnded(ctx context.Context, request *alexa.Request, session *alexa.Session, ctx_ptr *alexa.Context, response *alexa.Response) error {

	log.Printf("OnSessionEnded requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	return nil
}

func Start(arn string) {
	a = &alexa.Alexa{
		ApplicationID:       arn,
		RequestHandler:      &ScrumMaster{},
		IgnoreApplicationID: true,
		IgnoreTimestamp:     true,
	}
	log.Println("Alexa Handling started")
	lambda.Start(Handle)
}
