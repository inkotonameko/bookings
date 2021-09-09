package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

//Appconfig holds the application config
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	InProduction bool
	Session *scs.SessionManager
}