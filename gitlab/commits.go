package gitlab

import (
	"encoding/json"
	"time"
)

// GitLab API docs: https://docs.gitlab.com/ee/api/commits.html

const (
	ProjectCommitStatusesApiPath = "/projects/:id/repository/commits/:sha/statuses"
	CreateCommitApiPath          = "/projects/:id/repository/commits"
)

type CommitStatus struct {
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	StartedAt    *time.Time `json:"started_at"`
	Name         string     `json:"name"`
	AllowFailure bool       `json:"allow_failure"`
	Author       User       `json:"author"`
	Description  *string    `json:"description"`
	Sha          string     `json:"sha"`
	TargetURL    string     `json:"target_url"`
	FinishedAt   *time.Time `json:"finished_at"`
	ID           int        `json:"id"`
	Ref          string     `json:"ref"`
}

// CreateCommitOptions represents the available options for a new commit.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#create-a-commit-with-multiple-files-and-actions
type CreateCommitOptions struct {
	Branch        *string                `url:"branch,omitempty" json:"branch,omitempty"`
	CommitMessage *string                `url:"commit_message,omitempty" json:"commit_message,omitempty"`
	StartBranch   *string                `url:"start_branch,omitempty" json:"start_branch,omitempty"`
	StartSHA      *string                `url:"start_sha,omitempty" json:"start_sha,omitempty"`
	StartProject  *string                `url:"start_project,omitempty" json:"start_project,omitempty"`
	Actions       []*CommitActionOptions `url:"actions" json:"actions"`
	AuthorEmail   *string                `url:"author_email,omitempty" json:"author_email,omitempty"`
	AuthorName    *string                `url:"author_name,omitempty" json:"author_name,omitempty"`
	Stats         *bool                  `url:"stats,omitempty" json:"stats,omitempty"`
	Force         *bool                  `url:"force,omitempty" json:"force,omitempty"`
}

// CommitActionOptions represents the available options for a new single
// file action.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/commits.html#create-a-commit-with-multiple-files-and-actions
type CommitActionOptions struct {
	Action          *string `url:"action,omitempty" json:"action,omitempty"`               // The action to perform: create, delete, move, update, chmod
	FilePath        *string `url:"file_path,omitempty" json:"file_path,omitempty"`         // Full path to the file. Ex. lib/class.rb
	PreviousPath    *string `url:"previous_path,omitempty" json:"previous_path,omitempty"` // Original full path to the file being moved. Ex. lib/class1.rb. Only considered for move action.
	Content         *string `url:"content,omitempty" json:"content,omitempty"`             // File content, required for all except delete, chmod, and move. Move actions that do not specify content preserve the existing file content, and any other value of content overwrites the file content.
	Encoding        *string `url:"encoding,omitempty" json:"encoding,omitempty"`
	LastCommitID    *string `url:"last_commit_id,omitempty" json:"last_commit_id,omitempty"`
	ExecuteFilemode *bool   `url:"execute_filemode,omitempty" json:"execute_filemode,omitempty"` // When true/false enables/disables the execute flag on the file. Only considered for chmod action.
}

func (g *Gitlab) ProjectCommitStatuses(id, sha1 string) ([]*CommitStatus, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectCommitStatusesApiPath, map[string]string{
		":id":  id,
		":sha": sha1,
	})

	statuses := make([]*CommitStatus, 0)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return statuses, meta, err
	}

	err = json.Unmarshal(contents, &statuses)

	return statuses, meta, err
}

// Create a commit and push with multiple files and actions
func (g *Gitlab) CreateCommit(id string, commitOpts *CreateCommitOptions) (*Commit, *ResponseMeta, error) {
	// get http url
	u := g.ResourceUrl(CreateCommitApiPath, map[string]string{":id": id})

	commitJson, err := json.Marshal(commitOpts)
	if err != nil {
		return nil, nil, err
	}

	var createCommit *Commit
	contents, meta, err := g.buildAndExecRequest("POST", u.String(), commitJson)
	if err == nil {
		err = json.Unmarshal(contents, &createCommit)
	}

	return createCommit, meta, err
}
