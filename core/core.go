package core

import "context"

type CronService interface {
	Start()
	Stop() context.Context
	Add(Runnable) error
	Delete(Runnable) error
	Count() int
}

type Executable interface {
	Type() string
	ToCall() (*Call, error)
	ToJob() (*Job, error)
}

type Services struct {
	Environment EnvironmentInfoService
	Jobs        JobService
	Calls       CallService
	Schedules   ScheduleService
	Workers     WorkerService
	Runner      RunService
	Executions  ExecutionService
	Cron        CronService
	Logger      LogService
}
