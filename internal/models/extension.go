package models

type Extension string

const (
	GO     = Extension(".go")
	KOTLIN = Extension(".kt")
	PYTHON = Extension(".py")
)

func (ex Extension) GetByString(value string) Extension {
	switch value {
	case "kotlin":
		ex = KOTLIN
	case "python":
		ex = PYTHON
	default:
		ex = GO
	}
	return ex
}
