package orm

import "errors"

func NewEmptyTag() error {
	return errors.New(EmptyTag)
}

func IsEmptyTag(err error) bool {
	if err == nil {
		return false
	} else {
		return EmptyTag == err.Error()
	}
}
