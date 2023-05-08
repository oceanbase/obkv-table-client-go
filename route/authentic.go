package route

type ObUserAuth struct {
	userName string
	password string
}

func (a *ObUserAuth) Password() string {
	return a.password
}

func (a *ObUserAuth) UserName() string {
	return a.userName
}

func NewObUserAuth(userName string, password string) *ObUserAuth {
	return &ObUserAuth{userName, password}
}

func (a *ObUserAuth) String() string {
	return "ObUserAuth{" +
		"userName:" + a.userName + ", " +
		"password:" + a.password +
		"}"
}
