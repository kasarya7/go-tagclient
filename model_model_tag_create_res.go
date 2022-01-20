/*
 * tag-api OpenAPI Doc
 *
 * tag-api OpenAPI doc
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tagClient

type ModelTagCreateRes struct {
	// 标签颜色
	Color string `json:"color,omitempty"`
	// 标签创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 标签部门id
	DomainId string `json:"domain_id,omitempty"`
	// 标签部门名字
	DomainName string `json:"domain_name,omitempty"`
	// 标签id
	Id string `json:"id,omitempty"`
	// 标签项目id
	ProjectId string `json:"project_id,omitempty"`
	// 标签项目名字
	ProjectName string `json:"project_name,omitempty"`
	// 标签绑定资源数
	ResNum int32 `json:"res_num,omitempty"`
	// 标签名
	Tag string `json:"tag,omitempty"`
}