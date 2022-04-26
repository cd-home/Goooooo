package v1

type UserLogic struct {
}

func NewUserRegisterLogicContext() *UserLogic {
	return &UserLogic{}
}

func (logic *UserLogic) Register() error {
	return nil
}

func (logic *UserLogic) Login() error {
	return nil
}
