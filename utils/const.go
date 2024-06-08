package utils

type PrivilegeType string

const (
	CREATE PrivilegeType = "create"
	DELETE PrivilegeType = "delete"
	GRANT  PrivilegeType = "grant"
	GET    PrivilegeType = "get"
	UPDATE PrivilegeType = "update"
)
