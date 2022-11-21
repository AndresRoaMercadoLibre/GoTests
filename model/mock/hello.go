package mock

import (
	"TDD/model"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	if language == model.SPANISH {
		return model.SPANISHHELLOPREFIX
	}
	if language == model.FRENCH {
		return model.FRENCHHHELLOPREFIX
	}

	return model.ENGLISHELLOPREFIX
}
