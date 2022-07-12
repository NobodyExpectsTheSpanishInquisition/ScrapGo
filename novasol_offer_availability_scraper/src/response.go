package src

type response struct {
	IsAvailable bool
}

func newResponse(isAvailable bool) response {
	return response{IsAvailable: isAvailable}
}
