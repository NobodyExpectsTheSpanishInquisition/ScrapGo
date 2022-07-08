package main

import (
	"notifier/src"
)

func main() {
	var err error

	err = src.LoadEnv()

	n, err := src.CreateSmptNotifier()

	if nil != err {
		panic(err)
	}

	err = n.SendEmail(src.CreateEmail(src.GetSenderEmail(), src.GetReceiverEmail(), "TEST", "BODY"))

	if err != nil {
		panic(err)
	}
}
