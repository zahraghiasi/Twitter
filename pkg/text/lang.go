package text

import "strings"

type ServerLang uint8

const (
	En ServerLang = iota
	Fa
)

func (sl ServerLang) String() string {
	switch sl {
	case En:
		return "en"
	case Fa:
		return "fa"
	}
	return "fa"
}

func (sl ServerLang) FromString(language string) ServerLang {
	switch strings.ToLower(language) {
	case "en":
		return En
	case "fa":
		return Fa
	}
	return Fa
}
