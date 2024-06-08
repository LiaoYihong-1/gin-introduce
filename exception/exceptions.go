package exception

type UserNotFoundError struct {
	Info string
}

func (e UserNotFoundError) Error() string {
	return "User not found"
}

type BadRequestError struct {
	info string
}

func (e BadRequestError) Error() string {
	return "Bad request"
}
