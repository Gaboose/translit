package translit

import (
	"errors"

	"github.com/valyala/fastjson"
)

type StringTuple [2]string

// UnmarshalJSON allows us to avoid the encoding/json library.
// The standard json library makes tinygo panic
// (https://github.com/tinygo-org/tinygo/issues/447).
func (t *StringTuple) UnmarshalJSON(b []byte) error {
	v, err := fastjson.ParseBytes(b)
	if err != nil {
		return err
	}

	arr, err := v.Array()
	if err != nil {
		return err
	}

	if len(arr) != 2 {
		return errors.New("tuple size not 2")
	}

	*t = [2]string{arr[0].String(), arr[1].String()}

	return nil
}
