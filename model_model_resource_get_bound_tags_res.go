/*
 * tag-api OpenAPI Doc
 *
 * tag-api OpenAPI doc
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TagClient

type ModelResourceGetBoundTagsRes struct {
	CustomTags []ModelResourceGetBoundTagsData `json:"custom_tags,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ProjectId string `json:"project_id,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}
