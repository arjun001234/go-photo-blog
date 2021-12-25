package middleware

import (
	"github.com/julienschmidt/httprouter"
)

type Adapter func(httprouter.Handle) httprouter.Handle

func Adapt(h httprouter.Handle, adaptors ...Adapter) httprouter.Handle {
	for _, a := range adaptors {
		h = a(h)
	}
	return h
}
