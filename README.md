# New视频
<br />
<div style="text-align:center;">
  <a href="https://github.com/chajiuqqq/qiniu-1024/" style="display: inline-block; vertical-align: middle;">
    <img src="doc/img/logo.png" alt="Logo" width="50" height="50">
  </a>
  <h2 style="display: inline-block; vertical-align: middle; margin: 0;">New视频</h2>
</div>

New视频是一款使用七牛云OSS实现视频存储的Web端短视频网站，用户可以自由注册登录，浏览不同分类的视频，也可以上传自己的视频成为博主哦。

# 特点
 
- 视频上传，自动生成封面
- 视频播放，暂停，进度拖动
- 视频点赞，收藏
- 用户注册、登录
- 视频分类浏览

<div style="text-align: center;">
        <img src="doc/img/main.png" alt="Logo" style="width: 70%; max-width: 100%; height: auto; display: inline-block;">
</div>

# 上手指南
下载源码：
```
    git clone https://github.com/chajiuqqq/qiniu-1024.git
```
依赖：

- go >= 1.21.0
- node.js >= 18
- MongoDB

开发环境启动步骤：

- Go：
    - 进入server/cmd/api，修改local.yaml，配置好mongodb
    - 当前目录下运行 `go run .`，默认使用 local.yaml作为配置文件
- Web：
    - 进入web/
    - 运行npm i，安装依赖
    - 进入 app/lib/api/api-client.js，修改go后端Endpoint地址
    - 在web下运行`npm run dev`启动项目，在:3000端口查看界面

## 界面展示

<div style="text-align: center;">
    <div>
        <img src="doc/img/my.png" alt="Logo" style="width: 70%; max-width: 100%; height: auto; display: inline-block;">
        <p>个人详情</p>   
        <img src="doc/img/publish.png" alt="Logo" style="width: 70%; max-width: 100%; height: auto; display: inline-block;">
        <p>发布</p>   
    </div>
    <div>
        <img src="doc/img/food.png" alt="Logo" style="width: 70%; max-width: 100%; height: auto; display: inline-block;">
        <p>分类页</p>   
    </div>
    <div>
        <img src="doc/img/login.png" alt="Logo" style="width: 48%; max-width: 100%; height: auto; display: inline-block;">
        <img src="doc/img/register.png" alt="Logo" style="width: 48%; max-width: 100%; height: auto; display: inline-block;">
        <p>注册、登录</p>   
    </div>
</div>

## 演示视频

[Watch the video](http://47.106.228.5:8080/new-video.mp4)

# 项目结构

## 目录结构

```
root
|──server //go项目
|   ├─cmd
|    │  └─api // 程序入口
|    ├─model // 数据库模型
|    ├─service //业务层
|    ├─types //前端交互类型
|    └─utils // 工具类
|        ├─oss 
|        ├─shared
|        ├─xecho
|        ├─xerr
|        ├─xlog
|        └─xmongo
|——web // 前端项目
    ├─app // nextjs的页面入口，包含各类页面
    └─public // 静态文件目录
```

## 系统结构

![Alt text](doc/img/struct.png)

## 模块设计

| 模块 | 说明 | 工时 |
| --- | --- | --- |
| 视频模块 | 1. 视频列表拉取(视频状态的获取，用户对视频的互动记录获取（是否点过赞、是否收藏、是否关注)) 2. 视频播放、暂停、切换 3. 上传视频（文案） | 1+1  |
| 分类模块 | 热门、体育。。。 | 0.5 |
| 用户模块 | 1. 注册（头像，名称，手机号,描述） 2. 登陆 3. 查看我的视频  | 1+1  |
| 互动模块 | 1. 点赞 2. 收藏 | 1+1 |
| 中心化框架  | mongodb & OSS SDK ; echo的中间件（JWT、错误处理、日志zap） ; Api和web的部署 | 2天 |

## 技术栈

后端Web框架：Echo

数据库：MongoDB

前端：
- Next.js
- TypeScript
- TailWind CSS
- Plyr 播放组件