package repositories

import "fmt"

type GetMappedResultError struct {
	Previous error
}

func (e *GetMappedResultError) Error() string {
	return fmt.Sprint("Could not retrieve entity: ", e.Previous.Error())
}
