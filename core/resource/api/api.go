package api

import (
	"fmt"
)

type Handle interface{}


// Api
type Api struct {
	// http method
	Method string

	// uri path
	Path string

	// unique key
	Key string
}


// String
func (api *Api) String() string {
	return fmt.Sprintf("Method:%s,Path:%s,Key:%s", api.Method, api.Path, api.Key)
}
