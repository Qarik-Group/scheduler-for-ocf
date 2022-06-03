package ops

import (
	"time"

	"github.com/starkandwayne/scheduler-for-ocf/core"
)

func blank(candidate string) bool {
	return len(candidate) == 0
}

func dummyCall(call *core.Call) *core.Call {
	if call == nil {
		call = &core.Call{}
	}

	now := time.Now().UTC()
	call.CreatedAt = now
	call.UpdatedAt = now

	if blank(call.GUID) {
		call.GUID, _ = core.GenGUID()
	}

	if blank(call.Name) {
		call.Name = "dummy-call"
	}

	if blank(call.URL) {
		call.URL = "http://example.com"
	}

	if blank(call.AuthHeader) {
		call.AuthHeader = "dummy"
	}

	if blank(call.AppGUID) {
		call.AppGUID, _ = core.GenGUID()
	}

	if blank(call.SpaceGUID) {
		call.SpaceGUID, _ = core.GenGUID()
	}

	return call
}

func dummyJob(job *core.Job) *core.Job {
	if job == nil {
		job = &core.Job{}
	}

	now := time.Now().UTC()
	job.CreatedAt = now
	job.UpdatedAt = now

	if blank(job.GUID) {
		job.GUID, _ = core.GenGUID()
	}

	if blank(job.Name) {
		job.Name = "dummy-job"
	}

	if blank(job.Command) {
		job.Command = "echo 'I sure am a dummy job'"
	}

	if blank(job.AppGUID) {
		job.AppGUID, _ = core.GenGUID()
	}

	if blank(job.SpaceGUID) {
		job.SpaceGUID, _ = core.GenGUID()
	}

	return job
}
