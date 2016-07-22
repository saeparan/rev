package controllers

import (
    "bitbucket.org/pappy007/rev/app/models"
    "golang.org/x/oauth2"
    "encoding/json"
    "github.com/revel/revel"
    "net/http"
    "net/url"
    "log"
)

type OAuth struct {
    *revel.Controller
}

var FACEBOOK = &oauth2.Config{
    ClientID:     "363354330407306",
    ClientSecret: "13c7289354c349a0ba4751d4fcccf0c6",
    AuthURL:      "https://graph.facebook.com/oauth/authorize",
    //TokenURL:     "https://graph.facebook.com/oauth/access_token",
    RedirectURL:  "http://test.saeparan.com:9000/user/oauth/auth",
}

func (c OAuth) Index() revel.Result {
    u := c.connected()
    me := map[string]interface{}{}
    if u != nil && u.AccessToken != "" {
        resp, _ := http.Get("https://graph.facebook.com/me?access_token=" +
        url.QueryEscape(u.AccessToken))
        defer resp.Body.Close()
        if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
            revel.ERROR.Println(err)
        }
        revel.INFO.Println(me)
    }

    authUrl := FACEBOOK.AuthCodeURL("foo")
    return c.Render(me, authUrl)
}

func (c OAuth) Auth(code string) revel.Result {
    t := &oauth2.Transport{Config: FACEBOOK}
    tok, err := t.Exchange(code)
    if err != nil {
        revel.ERROR.Println(err)
        return c.Redirect(c.Index)
    }

//    user := c.connected()
    log.Println(tok.AccessToken, tok.Expiry)
    return c.Redirect("/user/oauth/")
}

func setuser(c *revel.Controller) revel.Result {
    var user *models.User
    if _, ok := c.Session["email"]; ok {

    } else {

    }
    c.RenderArgs["user"] = user
    return nil
}

func init() {
    revel.InterceptFunc(setuser, revel.BEFORE, &OAuth{})
}

func (c OAuth) connected() *models.User {
    return c.RenderArgs["user"].(*models.User)
}