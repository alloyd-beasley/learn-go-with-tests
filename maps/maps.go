package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(val string) (string, error) {
	def, ok := d[val]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = val

	case nil:
		return ErrWordExists

	default:
		return err
	}

	return nil
}
