package sqshandler

import (
	"github.com/aws/aws-lambda-go/events"
	"gopkg.in/go-playground/validator.v9"
)

type Status string

const (
	Failure Status = "FAILURE"
	Retry   Status = "RETRY"
	Skip    Status = "SKIP"
	Success Status = "SUCCESS"
)

type Result struct {
	Message *events.SQSMessage `validate:"required"`
	Status  Status             `validate:"oneof=FAILURE RETRY SKIP SUCCESS"`
	Error   error
}

func (r *Result) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
