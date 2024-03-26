package service

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:24

type IUserService interface {
}

type userService struct {
}

func NewUserService() IUserService {
	return &userService{}
}
