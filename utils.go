package main

import (
	"fmt"
)

type (
	Error struct {
		Code    int
		Message string
	}
	Wrapper struct {
		Code         int         `json:"code"`
		ErrorMessage string      `json:"error,omitempty"`
		Content      interface{} `json:"content,omitempty"`
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}