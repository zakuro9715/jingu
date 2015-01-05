package core

type Actor interface {
  Config() Config
  SetConfig(Config)
}

type Config struct {
}
