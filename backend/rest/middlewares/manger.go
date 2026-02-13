package middlewares

import "net/http"

// declaring a type based on middleware function's input datatype and output datatype
type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware // here we declared a array of functions
}

// creating a function which will return object
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

// creating builder function for cleanliness
func (mngr *Manager) Use(middlewares ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return mngr
}

// creating a reciever function for Manager
// middlewares ...Middleware here middlewares have turned into array. (... = this is a sphread operator)

// includes additional middleware provided while calling handler.
func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n
}

// includes global middlewares where cors and preflight handled
func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
	n := next
	// going through middlwares through loop
	// globalMiddleware=[logger, middleware2, CorsWithPreflight]
	// CorsPreflight(middleware2(logger(http.handlerfunc(test))))
	for _, middleware := range mngr.globalMiddlewares {
		n = middleware(n)
	}
	return n
}
