package module

import (
	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/web-module/internal/controllers/web"
)

type WebModule struct {
	registry      core.IHandlersRegistry
	loggerFactory core.ILoggerFactory
	executor      core.IExecutor
	configuration core.IConfiguration
}

func NewWebModule() *WebModule {
	return &WebModule{}
}

func (m *WebModule) Name() string {
	return "WEB Module"
}

func (m *WebModule) Description() string {
	return ""
}

func (m *WebModule) Version() string {
	return "v0.1.0"
}

func (m *WebModule) Start() error {
	ctr := web.NewController(m.configuration, m.registry)

	ctr.AddRoutes(m.registry)

	return nil
}

func (m *WebModule) AddHandlersRegistry(registry core.IHandlersRegistry) {
	m.registry = registry
}

func (m *WebModule) AddLoggerFactory(loggerFactory core.ILoggerFactory) {
	m.loggerFactory = loggerFactory
}

func (m *WebModule) AddExecutor(executor core.IExecutor) {
	m.executor = executor
}

func (m *WebModule) AddConfiguration(configuration core.IConfiguration) {
	m.configuration = configuration
}
