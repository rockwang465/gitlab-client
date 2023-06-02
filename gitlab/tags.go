package gitlab

import "encoding/json"

const (
	ProjectTagApiPath = "/projects/:id/repository/tags/:tag"
)

func (g *Gitlab) ProjectTag(projectId, tagName string) (*Tag, *ResponseMeta, error) {
	u := g.ResourceUrl(ProjectTagApiPath, map[string]string{
		":id":  projectId,
		":tag": tagName,
	})

	tag := new(Tag)

	contents, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err == nil {
		err = json.Unmarshal(contents, &tag)
	}

	return tag, meta, err
}
