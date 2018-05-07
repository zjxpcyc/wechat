package core

// APIInfo API 接口信息
type APIInfo struct {
	Method       string
	URI          string
	ResponseType string
}

// Kernel 默认核心类
type Kernel struct {
	scheduleServer ScheduleTask
	certificate    map[string]string
}

var log Log

// NewKernel 初始化
func NewKernel() *Kernel {
	return &Kernel{}
}

// SetScheduleTask 设置定时任务
func (t *Kernel) SetScheduleTask(scheduleServer ScheduleTask) {
	t.scheduleServer = scheduleServer
}

// SetTask 设置任务
func (t *Kernel) SetTask(task Task) {
	t.scheduleServer = NewDefaultScheduleServer()
	taskServer, _ := t.scheduleServer.(*DefaultScheduleServer)
	taskServer.SetTask(task)
}

// StartTokenServer 启动 AccessTokenServer
func (t *Kernel) StartTokenServer() {
	t.scheduleServer.Start()
}

// SetLogInst 设置全局日志实例
func SetLogInst(l Log) {
	log = l
}

func init() {
	log = &DefaultLogger{}
}
