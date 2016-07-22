package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//"log"
	"math"
	"time"
    "html/template"
)

type Article struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	BoardID       string        `bson:"BoardID"`
	Title         string        `bson:"Title"`
	Content       string        `bson:"Content"`
    ContentHTML   template.HTML
	NickName      string        `bson:"NickName"`
	RegDate       time.Time     `bson:"RegDate"`
	RegDateString string
	Hit           int    `bson:"Hit"`
	Vote          int    `bson:"Vote"`
	IP            string `bson:"IP"`
}

type Paging struct {
	TotalCount int
	TotalPage  int
	Page       int
	LimitStart int
	LimitEnd   int
	PageStart  int
	PageEnd    int
}

func (b *Article) Save(s *mgo.Session) error {
	coll := Collection(s, "article")
	err := coll.Insert(b)
	return err
}

func (a *Article) GetArticles(s *mgo.Session, BoardID string, p *Paging) []*Article {
	var result []*Article
	coll := Collection(s, "article")
	coll.Find(bson.M{}).Skip(p.LimitStart).Limit(p.LimitEnd).Sort("-_id").All(&result)

	for i, item := range result {
		item.RegDateString = item.RegDate.Format("2006-01-02 15:04:05")
		result[i] = item
	}

	return result
}

func (a *Article) GetArticle(s *mgo.Session, BoardID string, ArticleID string) *Article {
	var result *Article
	coll := Collection(s, "article")
	coll.Find(bson.M{"BoardID": BoardID, "_id": bson.ObjectIdHex(ArticleID)}).One(&result)
    result.ContentHTML = template.HTML(result.Content)
	return result
}

func (a *Article) GetTotalCount(s *mgo.Session, BoardID string) int {
	coll := Collection(s, "article")
	count, _ := coll.Find(bson.M{"BoardID": BoardID}).Count()
	return count
}

func (p *Paging) GetPagination(BoardID string, page int, s *mgo.Session) {
	var article *Article
	var totalCount, totalPage, limit, limitStart, limitEnd, pageStart, pageEnd int

	limit = 15
	totalCount = article.GetTotalCount(s, BoardID)
	totalPage = int(math.Ceil(float64(totalCount) / float64(limit)))

	if page < 1 || page > totalPage {
		page = 1
	}

	limitStart = (page - 1) * limit
	limitEnd = page * limit

	pageEnd = int(math.Ceil(float64(page)/float64(10))) * 10
	pageStart = pageEnd - 9

	if pageEnd > totalPage {
		pageEnd = totalPage
	}

	p.TotalCount = totalCount
	p.TotalPage = totalPage
	p.Page = page
	p.LimitStart = limitStart
	p.LimitEnd = limitEnd
	p.PageStart = pageStart
	p.PageEnd = pageEnd
}
