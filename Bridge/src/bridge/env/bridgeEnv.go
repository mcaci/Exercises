package env

type Env struct {
	m *Motorbike
	b *Bridge
}

func Init(motoInitPos, motoInitSpeed, roadLen, gapLen, platformLen int) *Env {
	env := new(Env)
	env.m = NewMotorbike(motoInitPos, motoInitSpeed)
	env.b = NewBridge(roadLen, gapLen, platformLen)
	return env
}

func (env *Env) GetBridge() *Bridge {
	return env.b
}

func (env *Env) GetMotorbike() *Motorbike {
	return env.m
}
