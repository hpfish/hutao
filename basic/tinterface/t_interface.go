package main

import "fmt"

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{
		Code:    code,
		Message: msg,
	}
}

func AsError(err error) error {
	return err
}

func main() {
	var rpcErr error = NewRPCError(400, "unknown err")
	err := AsError(rpcErr)
	println(err)

}
