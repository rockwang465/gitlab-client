package gitlab

import (
	"testing"

	"github.com/rockwang465/gitlab-client/test"
	"github.com/stretchr/testify/assert"
)

func TestGitlab_Runners(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"runners/runners",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	runners, meta, err := gitlab.Runners(nil)

	assert.NoError(t, err)

	assert.NotNil(t, runners)
	assert.Equal(t, meta.StatusCode, 200)
	assert.Equal(t, len(runners.Items), 2)

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

/*
func TestRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/show.json")
	runner, _, err := gitlab.Runner(6)

	assert.NoError(t, err)
	assert.IsType(t, new(RunnerWithDetails), runner)
	assert.Equal(t, runner.Id, 6)
	assert.Equal(t, runner.IsShared, false)
	assert.Equal(t, runner.Description, "test-1-20150125")
	assert.Equal(t, runner.Token, "205086a8e3b9a2b818ffac9b89d102")
	assert.Equal(t, len(runner.TagList), 2)
	assert.Equal(t, runner.ContactedAt, "2016-01-25T16:39:48.066Z")
	defer ts.Close()
}

func TestUpdateRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/update.json")

	runner := Runner{
		Description: "New Runner Description",
	}

	resp, _, err := gitlab.UpdateRunner(6, &runner)
	assert.NoError(t, err)
	assert.IsType(t, new(Runner), resp)
	assert.Equal(t, resp.Description, "New Runner Description")
	defer ts.Close()
}

func TestDeleteRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/delete.json")
	resp, _, err := gitlab.DeleteRunner(6)

	assert.NoError(t, err)
	assert.IsType(t, new(Runner), resp)
	assert.IsType(t, resp.Id, 6)
	defer ts.Close()
}

func TestProjectRunners(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/projects/index.json")
	runners, _, err := gitlab.ProjectRunners("1", 0, 2)

	assert.NoError(t, err)
	assert.Equal(t, len(runners), 2)
	assert.Equal(t, runners[0].IsShared, false)
	defer ts.Close()
}

func TestEnableProjectRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/projects/enable.json")
	runner, _, err := gitlab.EnableProjectRunner("1", 9)

	assert.NoError(t, err)
	assert.IsType(t, new(Runner), runner)
	assert.Equal(t, runner.Id, 9)
	defer ts.Close()
}

func TestDisableProjectRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/projects/disable.json")
	runner, _, err := gitlab.DisableProjectRunner("1", 9)

	assert.NoError(t, err)
	assert.IsType(t, new(Runner), runner)
	assert.Equal(t, runner.Id, 9)
	defer ts.Close()
}
*/
