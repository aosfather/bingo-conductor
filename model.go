package bingo_conductor

type TTaskStatus byte    //任务状态
type TRetryLogic byte    //重试逻辑
type TTimeoutPolicy byte //超时策略
type TTaskType byte      //任务类型
const (
	TS_Schedule  = 1  //计划
	TS_Completed = 2  //完成
	TS_Cancel    = 4  //取消
	TS_Pause     = 3  //暂停
	TS_Progress  = 11 //进行中

)

type Flow struct {
	Name             string //工作流程的名称，唯一
	Description      string //流程描述
	Version          int    //版本号
	OutputParameters string
	InputParameters  string //输入参数列表，用于记录工作流所需要的输入，可选
	Tasks            []Task //任务定义
}

type Task struct {
	Name            string    //任务类型
	TaskName        string    //任务名称，用于在流程中使用，必须唯一
	Type            TTaskType //任务类型
	Description     string    //任务描述
	Optional        bool      //是否可忽略，如果是true，当失败后，流程将继续.任务的状态反映为COMPLETED_WITH_ERRORS	默认为 false
	InputParameters string    //任务输入定义
}

//任务设置定义
type TaskDefine struct {
	Name                   string         //任务类型，唯一
	RetryCount             int            //重试次数
	RetryLogic             TRetryLogic    //重试机制
	TimeoutSeconds         int64          //超时间 单位毫秒ms
	TimeoutPolicy          TTimeoutPolicy //超时策略
	ResponseTimeoutSeconds int64          //返回超时时间
	OutputKeys             []string       //任务输出字段
}

//流程实例
type FlowInstance struct {
	Id      string //实例唯一ID
	Name    string //对应的流程名称
	Version int    //对应的流程版本
}

//任务实例
type TaskInstance struct {
	FlowInstance string //流程的实例Id
	Id           string //任务的实例Id
	TaskName     string //任务名称

}
