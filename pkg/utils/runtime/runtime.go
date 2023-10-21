package runtime

import (
	"log"
	"net/http"
	"runtime"
)

type PanicHanler func(interface{})

var PanicHanlers = []PanicHanler{logPanic}

func HandlePanic(additionalHanlers ...PanicHanler) {
	if r := recover(); r != nil {
		for _, handler := range PanicHanlers {
			handler(r)
		}

		for _, handler := range additionalHanlers {
			handler(r)
		}

	}
}

func logPanic(r interface{}) {
	if r == http.ErrAbortHandler {
		return
	}

	const size = 64 << 10
	stacktrace := make([]byte, size)
	stacktrace = stacktrace[:runtime.Stack(stacktrace, false)]
	if _, ok := r.(string); ok {
		log.Printf("Observed a panic: %s\n%s", r, stacktrace)
	} else {
		log.Printf("Observed a panic: %#v (%v)\n%s", r, r, stacktrace)
	}
}
