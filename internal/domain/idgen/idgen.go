package idgen

type IDGenerator interface {
	Generate() (int64, error)
}

var generator IDGenerator

func InitIDGen(g IDGenerator) {
	generator = g
}

func GenerateID() (int64, error) {
	return generator.Generate()
}
