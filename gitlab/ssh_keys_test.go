package gitlab

import (
	"testing"

	"github.com/rockwang465/go-gitlab-client/test"
	"github.com/stretchr/testify/assert"
)

func TestGitlab_CurrentUserSshKeys(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"ssh_keys/current_user_ssh_keys",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	keys, meta, err := gitlab.CurrentUserSshKeys(nil)

	assert.NoError(t, err)

	assert.Equal(t, 3, len(keys.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}

func TestGitlab_UserSshKeys(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"ssh_keys/user_1_ssh_keys",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	keys, meta, err := gitlab.UserSshKeys(1, nil)

	assert.NoError(t, err)

	assert.Equal(t, 3, len(keys.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
