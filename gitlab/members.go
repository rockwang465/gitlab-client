package gitlab

import (
	"encoding/json"
	"io"
)

const (
	MambersApiPath      = "/:type/:id/members"             // List group or project team members
	MamberApiPath       = "/:type/:id/members/:user_id"    // Get group or project team member
	AddProjectMember    = "/projects/:id/members"          // add project member
	AddGroupMember      = "/groups/:id/members"            // add group member
	UpdateProjectMember = "/projects/:id/members/:user_id" // update project member
	UpdateGroupMember   = "/groups/:id/members/:user_id"   // update group member

	MemberActiveState = "active"
	MemberBlockState  = "blocked"
)

// AddProjectMemberOptions represents the available AddProjectMember() options.
// GitLab API docs:
// https://docs.gitlab.com/ce/api/members.html#add-a-member-to-a-group-or-project
type MemberOptions struct {
	UserID      int    `url:"user_id,omitempty" json:"user_id,omitempty"`
	AccessLevel int    `url:"access_level,omitempty" json:"access_level,omitempty"`
	ExpiresAt   string `url:"expires_at,omitempty" json:"expires_at"` // A date string in the format YEAR-MONTH-DAY, 2022-12-30
}

type Member struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	State       string `json:"state"`
	AvatarUrl   string `json:"avatar_url"`
	WebUrl      string `json:"web_url"`
	AccessLevel int    `json:"access_level"`
	ExpiresAt   string `json:"expires_at"`
}

type MemberCollection struct {
	Items []*Member
}

type MembersOptions struct {
	PaginationOptions

	Query string `url:"query,omitempty"`
}

func (m *Member) RenderJson(w io.Writer) error {
	return renderJson(w, m)
}

func (m *Member) RenderYaml(w io.Writer) error {
	return renderYaml(w, m)
}

func (c *MemberCollection) RenderJson(w io.Writer) error {
	return renderJson(w, c.Items)
}

func (c *MemberCollection) RenderYaml(w io.Writer) error {
	return renderYaml(w, c.Items)
}

func (g *Gitlab) getResourceMembers(resourceType, projectId string, o *MembersOptions) (*MemberCollection, *ResponseMeta, error) {
	u := g.ResourceUrlQ(MambersApiPath, map[string]string{
		":type": resourceType,
		":id":   projectId,
	}, o)

	collection := new(MemberCollection)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &collection.Items)
	}

	return collection, meta, err
}

func (g *Gitlab) ProjectMembers(projectId string, o *MembersOptions) (*MemberCollection, *ResponseMeta, error) {
	return g.getResourceMembers("projects", projectId, o)
}

func (g *Gitlab) GroupMembers(groupId string, o *MembersOptions) (*MemberCollection, *ResponseMeta, error) {
	return g.getResourceMembers("groups", groupId, o)
}

func (g *Gitlab) AddProjectMember(projectId string, opt *MemberOptions) (*Member, *ResponseMeta, error) {
	u := g.ResourceUrl(AddProjectMember, map[string]string{
		":id": projectId,
	})

	memberOpt, err := json.Marshal(opt)
	if err != nil {
		return nil, nil, err
	}

	contents, meta, err := g.buildAndExecRequest("POST", u.String(), memberOpt)
	if err != nil {
		return nil, nil, err
	}

	member := new(Member)
	err = json.Unmarshal(contents, member)
	if err != nil {
		return nil, nil, err
	}
	return member, meta, nil
}

func (g *Gitlab) UpdateProjectMember(projectId, userId string, opt *MemberOptions) (*Member, *ResponseMeta, error) {
	u := g.ResourceUrl(UpdateProjectMember, map[string]string{
		":id":      projectId,
		":user_id": userId,
	})

	memberOpt, err := json.Marshal(opt)
	if err != nil {
		return nil, nil, err
	}

	contents, meta, err := g.buildAndExecRequest("PUT", u.String(), memberOpt)
	if err != nil {
		return nil, nil, err
	}

	member := new(Member)
	err = json.Unmarshal(contents, member)
	if err != nil {
		return nil, nil, err
	}
	return member, meta, nil
}

func (g *Gitlab) AddGroupMember(projectId string, opt *MemberOptions) (*Member, *ResponseMeta, error) {
	u := g.ResourceUrl(AddGroupMember, map[string]string{
		":id": projectId,
	})

	memberOpt, err := json.Marshal(opt)
	if err != nil {
		return nil, nil, err
	}

	contents, meta, err := g.buildAndExecRequest("POST", u.String(), memberOpt)
	if err != nil {
		return nil, nil, err
	}

	member := new(Member)
	err = json.Unmarshal(contents, member)
	if err != nil {
		return nil, nil, err
	}
	return member, meta, nil
}

func (g *Gitlab) UpdateGroupMember(projectId, userId string, opt *MemberOptions) (*Member, *ResponseMeta, error) {
	u := g.ResourceUrl(UpdateGroupMember, map[string]string{
		":id":      projectId,
		":user_id": userId,
	})

	memberOpt, err := json.Marshal(opt)
	if err != nil {
		return nil, nil, err
	}

	contents, meta, err := g.buildAndExecRequest("PUT", u.String(), memberOpt)
	if err != nil {
		return nil, nil, err
	}

	member := new(Member)
	err = json.Unmarshal(contents, member)
	if err != nil {
		return nil, nil, err
	}
	return member, meta, nil
}
