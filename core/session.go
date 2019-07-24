package core

import (
	"time"
)

type Session struct {

	// represent the session on FS
	Path string

	// List of modules
	Modules []Module

	// Logs
	Logs []Log

	MainGui *Broxygui
	Config  *Config

	LogC chan Log
}

func NewSession(path string) *Session {
	return &Session{
		Path:    path,
		MainGui: NewBroxygui(nil, 0),
		Config: &Config{
			Address: "127.0.0.1",
			Port:    8080,
		},
		LogC: make(chan Log),
	}
}

func LoadSession(path string) *Session {
	// Load session from file
	return nil
}

func (s *Session) LoadModule(m Module, g GuiModule) {
	s.Modules = append(s.Modules, m)
	s.MainGui.AddGuiModule(g)
}

func (s *Session) Info(mod string, message string) {
	t := time.Now()
	l := Log{Type: "I", ModuleName: mod, Time: t.Format("2006-01-02 15:04:05"), Message: message}
	s.Logs = append(s.Logs, l)
	go func() { s.LogC <- l }()
}

func (s *Session) Debug(mod string, message string) {
	t := time.Now()
	l := Log{Type: "I", ModuleName: mod, Time: t.Format("2006-01-02 15:04:05"), Message: message}
	s.Logs = append(s.Logs, l)
	go func() { s.LogC <- l }()
}

func (s *Session) Err(mod string, message string) {
	t := time.Now()
	l := Log{Type: "E", ModuleName: mod, Time: t.Format("2006-01-02 15:04:05"), Message: message}
	s.Logs = append(s.Logs, l)
	go func() { s.LogC <- l }()
}