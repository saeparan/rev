package controllers

import (
    "net"
    "time"
	"log"
	
    "github.com/revel/revel"
    "github.com/jgraham909/revmgo"
    "bitbucket.org/pappy007/rev/app/models"
)

type Board struct {
    *revel.Controller
    revmgo.MongoController
}

func (c Board) Index() revel.Result {
    return c.Render()
}

func (c Board) List(paging *models.Paging, article *models.Article) revel.Result {
	var BoardID string
	var page int
	c.Params.Bind(&BoardID, "id")
	c.Params.Bind(&page, "page")
	
	paging.GetPagination(BoardID, page, c.MongoSession)
	log.Println(paging)
	
	articles := article.GetArticles(c.MongoSession, BoardID, paging)
	
	return c.Render(BoardID, paging, articles)
}

func (c Board) View(article *models.Article) revel.Result {
	var BoardID, ArticleID string
	c.Params.Bind(&BoardID, "id")
	c.Params.Bind(&ArticleID, "ArticleID")
	
	articleItem := article.GetArticle(c.MongoSession, BoardID, ArticleID)
	
	return c.Render(BoardID, articleItem)
}

func (c Board) Write() revel.Result {
    var BoardID string
    c.Params.Bind(&BoardID, "id")
    return c.Render(BoardID)
}

func (c Board) Regist(article *models.Article) revel.Result {
    var BoardID, title, content string
    c.Params.Bind(&title, "title")
    c.Params.Bind(&content, "content")
    c.Params.Bind(&BoardID, "id")

    article.BoardID = BoardID
    article.Title = title
    article.Content = content
    article.NickName = c.Session["nickname"]
    article.Hit = 0
    article.Vote = 0
    article.IP = getIPAddress()
    article.RegDate = time.Now()

    article.Save(c.MongoSession)
    return c.Redirect("/board/list/"+BoardID)
}

func getIPAddress() string {
    var ip string
    interfa, err := net.InterfaceAddrs()
    if err != nil {

    }

    for _, address := range interfa {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ip = ipnet.IP.String()
                break
            }
        }
    }

    return ip
}
