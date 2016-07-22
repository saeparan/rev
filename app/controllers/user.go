package controllers

import (
	"bitbucket.org/pappy007/rev/app/models"

    "github.com/revel/revel"
    "github.com/jgraham909/revmgo"
    "golang.org/x/crypto/bcrypt"

    "regexp"
)

type User struct {
    *revel.Controller
    revmgo.MongoController
}

func (c User) Index() revel.Result {
    return c.Render()
}

func (c User) LoginSign(user *models.User) revel.Result {
    var email, password string
    check := true
    c.Params.Bind(&email, "email")
    c.Params.Bind(&password, "password")

    c.Validation.Required(email).Message("이메일주소를 입력해주세요.")
    c.Validation.Match(email, regexp.MustCompile("\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3}")).Message("올바른 이메일주소를 입력해주세요.")
    c.Validation.Required(password).Message("암호를 입력해주세요.")
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        return c.Redirect("/")
    }

    u := user.GetUserByEmail(c.MongoSession, email)

    err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
    if err != nil {
        check = false
    }

    c.Validation.Required(check).Message("이메일 혹은 패스워드가 틀립니다.")
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        return c.Redirect("/")
    }

    c.Session["email"] = email
    c.Session["nickname"] = u.Nickname
    return c.Redirect("/")
}

func (c User) Logout() revel.Result {
    if c.Session["email"] != "" {
        c.Session["email"] = ""
        c.Session["nickname"] = ""
    }
    return c.Redirect("/")
}

func (c User) Regist() revel.Result {
    if c.Session["email"] != "" {
        return c.Redirect("/")
    }
    return c.Render()
}

func (c User) Login() revel.Result {
    return c.Render()
}

func (c User) Add(user *models.User) revel.Result {
    var email, nickname, password, password_confirm string
    c.Params.Bind(&email, "email")
    c.Params.Bind(&password, "password")
    c.Params.Bind(&password_confirm, "password_conform")
    c.Params.Bind(&nickname, "nickname")

    c.Validation.Required(email).Message("이메일주소를 입력해주세요.")
    c.Validation.Match(email, regexp.MustCompile("\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3}")).Message("올바른 이메일주소를 입력해주세요.")
    c.Validation.Required(password).Message("암호를 입력해주세요.")
    c.Validation.MinSize(password, 6).Message("암호는 6자리 이상 입력해주세요.")
    c.Validation.MaxSize(password, 20).Message("암호는 20자 이하 입력해주세요.")
    c.Validation.MaxSize(password_confirm, 20).Message("암호는 20자 이하 입력해주세요.")
    c.Validation.Required(password != password_confirm).Message("암호, 암호확인란이 일치하지 않습니다.")
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        return c.Redirect("/user/regist")
    }

    hashedPassword := hashPassword(password)

    user.Email = email
    user.Password = hashedPassword
    user.Nickname = nickname
    user.Save(c.MongoSession)

    return c.Render()
}

func hashPassword(password string) []byte {
    pwdByte := []byte(password)
    hashedPassword, err := bcrypt.GenerateFromPassword(pwdByte, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    return hashedPassword
}