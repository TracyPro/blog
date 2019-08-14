package models

import (
	"encoding/json"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	CONF = "/Users/tracy/go/" +
		"src/tracy_blog/blog/" +
		"conf/app.conf"
	DATA = "/Users/tracy/go/" +
		"src/tracy_blog/blog/" +
		"conf/recommend.json"
	DEFAULT_RANK_CNT  = 10
	MAX_RECOMMEND     = 5
	UPLOAD_IMG_PATH   = "static/upload/images/"
	UPLOAD_VIDEO_PATH = "static/upload/videos/"
	SUCCESS           = 1
	FAILED            = 0
)

// 文章
type Article struct {
	Id       int       // 编号
	Title    string    // 标题
	Url      string    // 图片链接
	Content  string    `orm:"type(text)"` // 转为html格式的字符串内容
	Meta     string    `orm:"type(text)"` // 用户编辑的原始内容
	Category *Category `orm:"rel(fk)"`    // 设置一对多关系，一个分类多篇文章
	Click    int       // 点击数
	Created  time.Time `orm:"index"` // 发布时间
	Updated  time.Time `orm:"index"` // 更新时间
	Author   string    // 作者
	Source   string    // 来源
	User     string    // 文章对应的登陆用户
}

// 分类
type Category struct {
	Id      int       // 编号
	Title   string    `orm:"unique"` // 类别
	Created time.Time `orm:"index"`  // 创建时间
	Updated time.Time `orm:"index"`  // 更新时间
}

// 关注我
type Mine struct {
	Id     int    // 编号
	Phone  string // 电话
	Email  string // 邮箱
	QQ     string // QQ号
	Prizes string // 奖项
}

// 图文推荐
type Recommend struct {
	Id       int       // 编号
	Title    string    // 推荐文章标题
	Url      string    // 文章第三方链接，将推荐文章设置为第三方文章
	Image    string    // 图片链接
	Date     time.Time `orm:"index"`   // 链接上线日期
	Category *Category `orm:"rel(fk)"` // 推荐文章类别
}

// 点击排行
type Rank struct {
	Id      int      // 编号
	Article *Article `orm:"rel(one)"` // 文章
}

// 相册图片
type JpgList struct {
	Id  int    // 编号
	Url string // 链接
}

// 相册视频
type VideoList struct {
	Id  int    // 编号
	Url string // 链接
}

func RegisterDB() {
	orm.RegisterModel(new(Article), new(Category),
		new(Mine), new(Recommend), new(Rank), new(JpgList), new(VideoList))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/blog?charset=utf8")
	orm.Debug = true
	orm.RunSyncdb("default", false, false)
}

func AddCategory(title string) *Category {
	o := orm.NewOrm()
	category := &Category{
		Title:   title,
		Created: time.Now(),
		Updated: time.Now(),
	}
	qs := o.QueryTable("category")
	if err := qs.Filter("title", title).One(category); err != nil {
		o.Insert(category)
	}
	return category
}

func AddArticleContent(title, url, html_content, meta, author, category, source string) {
	o := orm.NewOrm()
	category_ptr := AddCategory(category)
	article := &Article{
		Title:    title,
		Url:      url,
		Content:  html_content,
		Meta:     meta,
		Category: category_ptr,
		Created:  time.Now(),
		Updated:  time.Now(),
		Author:   author,
		Source:   source,
	}
	_, err := o.Insert(article)
	if err != nil {
		log.Fatal(err)
	}
}

func TotalArticles() (int, []*Article) {
	o := orm.NewOrm()
	var arts []*Article
	qs := o.QueryTable("article")
	qs.OrderBy("-Updated").All(&arts)
	article_total := len(arts) //	文章总数
	for i := 0; i < article_total; i++ {
		arts[i].Content = TrimHtml(arts[i].Content)
	}
	return article_total, arts
}

func FilterArticles(per, current int, arts []*Article) []*Article {
	// current
	// 1 1-5
	// 2 6-10
	// 3 11-15
	// [(current-1)*per,current*per]
	end := current * per
	if end < len(arts) {
		return arts[(current-1)*per : end]
	}
	return arts[(current-1)*per : len(arts)]
}

func TotalCategories() []*Category {
	o := orm.NewOrm()
	var cates []*Category
	qs := o.QueryTable("category")
	qs.All(&cates)
	return cates
}

func GetCategory(id int) string {
	o := orm.NewOrm()
	category := Category{Id: id}
	o.Read(&category)
	return category.Title
}

var DetailArt Article

func ViewDetail(article_id int) {
	o := orm.NewOrm()
	DetailArt = Article{Id: article_id}

	o.Read(&DetailArt)
	DetailArt.Click += 1
	_, err := o.Update(&DetailArt, "Click")
	if err != nil {
		log.Fatal(err)
	}
}

func InsertMine() *Mine {
	iniconf, err := config.NewConfig("ini", CONF)
	if err != nil {
		log.Fatal(err)
	}
	mine := &Mine{
		Phone:  iniconf.String("phone"),
		Email:  iniconf.String("email"),
		QQ:     iniconf.String("qq"),
		Prizes: "",
	}
	o := orm.NewOrm()
	qs := o.QueryTable("mine")
	err = qs.Filter("phone", iniconf.String("phone")).One(mine)
	if err != nil {
		o.Insert(mine)
	}
	return mine
}

func ViewsRank() []*Article {
	o := orm.NewOrm()
	var arts []*Article
	qs := o.QueryTable("article")
	qs.OrderBy("-click").All(&arts)
	if len(arts) <= DEFAULT_RANK_CNT {
		return arts
	}
	return arts[0:DEFAULT_RANK_CNT]
}

var Arts []*Article
var Final []*Article

func StorClickCate(cid string) {
	o := orm.NewOrm()
	cid_int, _ := strconv.Atoi(cid)
	qs := o.QueryTable("article").Filter("category__id", cid_int)
	qs.All(&Arts)
	for i := 0; i < len(Arts); i++ {
		Arts[i].Content = TrimHtml(Arts[i].Content)
	}
}

func Search(content string) {
	arts := Arts
	text := strings.TrimSpace(content)
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	qs.All(&arts)
	for i := 0; i < len(arts); i++ {
		if strings.Contains(arts[i].Content, text) {
			Final = append(Final, arts[i])
		} else {
			continue
		}
	}
	for i := 0; i < len(Final); i++ {
		Final[i].Content = TrimHtml(Final[i].Content)
	}
}

type RecommendGroup struct {
	Group    string       `json:"group"`
	Articles []*Recommend `json:"articles"`
}

func GetRecommendData() []*Recommend {
	bytes, _ := ioutil.ReadFile(DATA)
	data := &RecommendGroup{}
	recommend := &Recommend{}
	json.Unmarshal(bytes, &data)
	o := orm.NewOrm()
	qs := o.QueryTable("recommend")

	for i := 0; i < len(data.Articles); i++ {
		err := qs.Filter("id", data.Articles[i].Id).One(recommend)
		if err != nil {
			data.Articles[i].Date = time.Now()
			o.Insert(data.Articles[i])
		} else {
			break
		}
	}
	if len(data.Articles) <= 5 {
		return data.Articles
	}
	return data.Articles[0:MAX_RECOMMEND]
}

func StorImgs(url string) {
	o := orm.NewOrm()
	jpg := &JpgList{
		Url: url,
	}
	qs := o.QueryTable("jpg_list")
	err := qs.Filter("url", url).One(jpg)
	if err != nil {
		o.Insert(jpg)
	}
}

func StorVideos(url string) {
	o := orm.NewOrm()
	video := &VideoList{
		Url: url,
	}
	qs := o.QueryTable("video_list")
	err := qs.Filter("url", url).One(video)
	if err != nil {
		o.Insert(video)
	}
}

func GetImgs() ([]*JpgList, int) {
	var Jpgs []*JpgList
	o := orm.NewOrm()
	qs := o.QueryTable("jpg_list")
	qs.All(&Jpgs)
	return Jpgs, len(Jpgs)
}

func GetVideos() ([]*VideoList, int) {
	var videos []*VideoList
	o := orm.NewOrm()
	qs := o.QueryTable("video_list")
	qs.All(&videos)
	return videos, len(videos)
}

func CheckPath(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}

func FilterImgs(per, current int, jpgs []*JpgList) []*JpgList {
	// current
	// 1 1-5
	// 2 6-10
	// 3 11-15
	// [(current-1)*per,current*per]
	end := current * per
	if end < len(jpgs) {
		return jpgs[(current-1)*per : end]
	}
	return jpgs[(current-1)*per : len(jpgs)]
}

func FilterVideos(per, current int, videos []*VideoList) []*VideoList {
	// current
	// 1 1-5
	// 2 6-10
	// 3 11-15
	// [(current-1)*per,current*per]
	end := current * per
	if end < len(videos) {
		return videos[(current-1)*per : end]
	}
	return videos[(current-1)*per : len(videos)]
}

func TrimHtml(src string) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	src = strings.Replace(src, "&quot;", "", -1)
	return src
}

func GetTechnologyArts(arts []*Article) ([]*Article, int) {
	var tech_arts []*Article
	o := orm.NewOrm()
	for i := 0; i < len(arts); i++ {
		cate := Category{Id: arts[i].Category.Id}
		_ = o.Read(&cate)
		if strings.Contains(cate.Title, "技术") ||
			strings.Contains(cate.Title, "golang") ||
			strings.Contains(cate.Title, "python") ||
			strings.Contains(cate.Title, "beego") {
			tech_arts = append(tech_arts, arts[i])
		} else {
			continue
		}
	}
	return tech_arts, len(tech_arts)
}

func DelArticle(art_id string) int {
	var art Article
	o := orm.NewOrm()
	del_id, _ := strconv.Atoi(art_id)
	art.Id = del_id
	_, err := o.Delete(&art)
	if err != nil {
		return FAILED
	}
	return SUCCESS
}

var EditArt *Article

func GetEditArticle(art_id string) {
	edit_id, _ := strconv.Atoi(art_id)
	o := orm.NewOrm()
	EditArt = &Article{Id: edit_id}
	o.Read(EditArt)
}

func UpdateArticle(u_id, title, url, html_content,
	meta, author, category, source string) {
	o := orm.NewOrm()
	u_id_int, _ := strconv.Atoi(u_id)
	category_ptr := UpdateCategory(category)
	article := &Article{
		Id:       u_id_int,
		Title:    title,
		Url:      url,
		Content:  html_content,
		Meta:     meta,
		Category: category_ptr,
		Updated:  time.Now(),
		Author:   author,
		Source:   source,
	}
	_, err := o.Update(article, "Title", "Url", "Content",
		"Meta", "Category", "Updated", "Author", "Source")
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateCategory(title string) *Category {
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	cate := &Category{
		Title:   title,
		Updated: time.Now(),
	}
	err := qs.Filter("title", title).One(cate)

	new_update_cate := &Category{
		Id:      cate.Id,
		Title:   title,
		Updated: time.Now(),
	}

	if err != nil {
		new_create_cate := &Category{
			Title:   title,
			Created: time.Now(),
			Updated: time.Now(),
		}
		o.Insert(new_create_cate)
		return new_create_cate
	}

	_, err = o.Update(new_update_cate, "Title", "Updated")
	if err != nil {
		log.Fatal(err)
	}
	return new_update_cate
}
