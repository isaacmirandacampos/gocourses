package helpers

import "fmt"

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("the parameter %s is required and should be a %s", name, typ)
}
