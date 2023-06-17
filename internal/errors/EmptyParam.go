package apperrors

import "fmt"

type EmptyParamError struct {
	FuncName    string
	ParamName   string
}

func (e *EmptyParamError) Error() string {
	return fmt.Sprintf("Empty param %s in %s", e.ParamName, e.FuncName)
}
