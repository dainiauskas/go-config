package config

type Logger struct {
	Path          string
	MaxFiles      int
	FilesToDelete int
	MaxSize       int64
}
