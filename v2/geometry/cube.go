package geometry

import "errors"

// CubeVolume : returns  s*s*s, error
func CubeVolume(s int) (int, error) {
	if s <= 0 {
		return 0, errors.New("edge must be greater zero")
	}
	return s * s * s, nil
}
