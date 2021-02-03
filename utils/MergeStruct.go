package utils

import (
	"github.com/imdario/mergo"
)

//Merge : merge all exported fields of two instances of struct
//filling empty fields of destination object by equal fields of
//merge object
func Merge(dest interface{}, src interface{}) error {
	err := mergo.Merge(&dest, &src)
	if err != nil {
		return err
	}
	return nil
}
