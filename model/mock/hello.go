package mock

import "TDD/model"

func Hello(name string) string {
	return model.ENGLISHELLOPREFIX + name
}
