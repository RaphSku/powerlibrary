package handler

import (
	"fmt"
)

type ArgumentMissing struct {
	field string
}

func (a ArgumentMissing) Error() string {
	return fmt.Sprintf("no argument for the %s was specified", a.field)
}
