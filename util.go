package gopkg

import (
	"encoding/json"
)

func Convert[E1, E2 any](src E1, des E2) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, des)
}

func Transform[E1, E2 any](src E1) (E2, error) {
	var des E2
	return des, Convert(src, &des)
}
