package utils

import "gin/models"

func hasPrivilege(user models.User, privilege PrivilegeType) bool {
	for _, v := range user.Privileges {
		if PrivilegeType(v.Name) == privilege {
			return true
		}
	}
	return false
}
