package utils

func HasPrivilege(p PrivilegeType, slice []PrivilegeType) bool {
	for _, v := range slice {
		if v == p {
			return true
		}
	}
	return false
}
