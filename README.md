# 基于Golang和Beego的博客项目

## 开发目标
熟悉go语言使用，练习Beego框架，将博客作为个人工作记录以及学习笔记存储或复习的容器，今后有时间还会购买域名搭建在云服务器，供大家一起使用。

## 技术要点
- Go语言
- Beego框架
- Mysql数据库
- Bootstrap前端框架
- JavaScript使用

## 功能概述

#### 博客首页
博客首页主要包含下面几个模块：
- 导航栏
- 文章摘要
- 作者信息
- 图文推荐
- 阅读排行

##### 首页截图如下：
![首页](https://github.com/TracyPro/blog/blob/master/blog截图/首页.png)

##### 点击阅读全文，显示文章详情
![文章详情](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E6%96%87%E7%AB%A0%E8%AF%A6%E6%83%85.png)

##### 点击推荐文章，跳转
![图文推荐](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E5%9B%BE%E6%96%87%E6%8E%A8%E8%8D%90.png)

##### 阅读排行为点击数前十的文章，点击后同样显示文章详情
![排行榜文章详情](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E6%8E%92%E8%A1%8C%E6%A6%9C%E6%96%87%E7%AB%A0%E8%AF%A6%E6%83%85.png)

#### 生活记录
生活记录包含下面几个模块：
- RNG荣誉介绍（本人软泥怪一枚😄）
- 我的相册
- 我的视频

##### RNG荣誉
![RNG荣誉](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/RNG.png)

##### 相册
![相册](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E7%9B%B8%E5%86%8C.png)

##### 支持批量上传图片
可拖拽，可从目录选择。
![上传图片](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E5%9B%BE%E7%89%87%E4%B8%8A%E4%BC%A0.png)

##### 视频列表
可拖拽，可从目录选择。
![视频列表](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E8%A7%86%E9%A2%91.png)

##### 支持批量上传视频
![上传视频](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E8%A7%86%E9%A2%91%E4%B8%8A%E4%BC%A0.png)

#### 技术文章
技术文章模块展示的内容跟首页差不多，唯一的区别就是文章部分都是介绍技术的。
![技术文章]()

#### 文章列表
文章列表模块展示了当前登陆用户所发布的所有文章信息，用户可以随时修改或删除文章。管理员账户可以查看所有用户的文章。
![文章列表](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E6%96%87%E7%AB%A0%E5%88%97%E8%A1%A8.png)

#### 写文章
登陆到系统内的用户可以编写并发布自己的文章，支持markdown语法，并且可以随时浏览编辑后的样式。支持文章导出html格式和文本格式。
![写文章](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E5%86%99%E6%96%87%E7%AB%A0.png)

#### 搜索
支持关键字搜索：
![关键字查询](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E5%85%B3%E9%94%AE%E5%AD%97%E6%9F%A5%E8%AF%A2.png)

支持点击标签筛选：
![点击标签筛选](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E6%A0%87%E7%AD%BE%E6%9F%A5%E8%AF%A2.png)

#### 注册
![注册](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E6%B3%A8%E5%86%8C.png)

#### 登陆
![登陆](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E7%99%BB%E9%99%86.png)
登陆后，显示退出
![退出](https://github.com/TracyPro/blog/blob/master/blog%E6%88%AA%E5%9B%BE/%E9%80%80%E5%87%BA.png)
退出后，显示登陆
![登陆](https://github.com/TracyPro/blog/blob/master/blog截图/退出后显示登陆.png)
