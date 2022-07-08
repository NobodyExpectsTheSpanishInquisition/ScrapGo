package src

import (
	"errors"
	"net/smtp"
)

func CreateSmptNotifier() (*SmptNotifier, error) {
	var err error
	var notifier *SmptNotifier
	env := GetEnvironment()

	switch env {
	case LOCAL:
		return NewSmptNotifier(GetNotifierHost(), GetNotifierPort(), nil), err
	case PROD:
		return NewSmptNotifier(
				GetNotifierHost(),
				GetNotifierPort(),
				smtp.PlainAuth(GetNotifierIdentity(), GetSenderEmail(), GetSenderPassword(), GetNotifierHost()),
			),
			err

	}

	return notifier, errors.New("invalid environment")
}

const (
	LOCAL = "LOCAL"
	PROD  = "PRODUCTION"
)
