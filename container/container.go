package container

type Container interface {
	GetEnv() string
}

type container struct {
	env string
}

func NewContainer(env string) Container {
	return &container{
		env: env,
	}
}

func (c *container) GetEnv() string {
	return ""
}
