package core

// Log 日志
type Log interface {
	Critical(string, ...interface{})
	Error(string, ...interface{})
	Warning(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

// Kernel 默认核心类
type Kernel struct {
	tokenServer AccessTokenServer
}

var log Log

// NewKernel 初始化
func NewKernel(certificate map[string]string) *Kernel {
	return &Kernel{
		tokenServer: NewDefaultTokenServer(certificate),
	}
}

// SetLogInst 设置全局日志实例
func (t *Kernel) SetLogInst(l Log) {
	log = l
}

// StartTokenServer 启动 AccessTokenServer
func (t *Kernel) StartTokenServer() {
	t.tokenServer.Start()
}

// GetAccessToken 获取最新 access_token
func (t *Kernel) GetAccessToken() string {
	return t.tokenServer.Get()
}
