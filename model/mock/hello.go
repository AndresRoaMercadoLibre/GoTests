package mock

import "TDD/model"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return model.ENGLISHELLOPREFIX + name
}
