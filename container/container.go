package container

type Container interface {
	GetEnv() string
	GetLogger() string
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

func (c *container) GetLogger() string {
	return ""
}
