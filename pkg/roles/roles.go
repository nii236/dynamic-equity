package roles

import (
	"github.com/zpatrick/rbac"
)

const ActionChunkRequest = "ChunkRequest"
const ActionChunkApprove = "ChunkApprove"
const ActionChunkReject = "ChunkReject"
const ActionChunkList = "ChunkList"

func New() []*rbac.Role {
	roles := []*rbac.Role{
		{
			RoleID: "Owner",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("*", "*"),
			},
		},
		{
			RoleID: "Manager",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission(ActionChunkRequest, "*"),
				rbac.NewGlobPermission(ActionChunkApprove, "*"),
				rbac.NewGlobPermission(ActionChunkReject, "*"),
				rbac.NewGlobPermission(ActionChunkList, "*"),
			},
		},
		{
			RoleID: "Grunt",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission(ActionChunkRequest, "self"),
				rbac.NewGlobPermission(ActionChunkList, "self"),
			},
		},
	}
	return roles
}
