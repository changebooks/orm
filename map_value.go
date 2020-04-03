package orm

import "errors"

func MapValue(attributes map[string]interface{}) (columns []string, values []interface{}, err error) {
	if attributes == nil {
		return nil, nil, errors.New("attributes can't be nil")
	}

	size := len(attributes)
	if size == 0 {
		return nil, nil, errors.New("attributes can't be empty")
	}

	for k, v := range attributes {
		columns = append(columns, k)
		values = append(values, v)
	}

	return columns, values, nil
}
