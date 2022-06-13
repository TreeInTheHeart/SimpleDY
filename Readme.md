### 环境依赖

go 1.8.0+ 

mysql 8.0+

ffmpeg

gin框架 gorm,viper,jwt库

### 目录结构描述

```
│  config.yaml   //配置本地数据库用户名密码、服务器端口号等
│  go.mod       
│  go.sum   
│  main.go          
│  readme.md     //help   
├─api            // 接口函数
│      favorite.go      //点赞操作接口
│      feed.go          //视频流接口
│      follow.go        //关注他人，查看关注列表，粉丝列表接口
│      publish.go       //发布视频接口
│      user.go          //用户相关接口 包括注册、登录、查询用户信息
│
├─config                //数据库参数配置
│      config.go        
│
├─data                  //存储用户上传视频及相关封面
│  ├─cover
│  └─video
├─global                //全局变量
│      golbal.go
│
├─handler               //gin handler
│      handler.go
│
├─initial               //初始化
│      load.go          //从config.yaml中加载配置信息
│      mysql.go         //数据库初始化
│
├─middleware            //中间件
│      jwt.go           //jwt鉴权
│
├─pojo                  //存入数据库的表以及给前端返回的数据结构
│      response.go
│      user.go
│      user_like_video.go
│      video.go
│      follow.go
│      feed.go
│
├─service               //操作数据库
│      user.go          //操作User表的相关函数
│      video.go         //操作video表的相关函数
│      following.go     //操作follow表的相关函数
│
├─status               //错误码以及对应消息
│      code.go
│      msg.go
│
├─utils                //使用到的工具类函数         
│      getCoverFromVideo.go     //从上传的视频中截取一帧作为封面
│      makeCoverPathById.go     //通过视频ID生成封面地址
│      makeVideoPathById.go     //通过视频ID生成视频地址
```

### V1.0.0 功能介绍

1. 完成注册功能
   
    用户输入用户名，密码完成注册功能并返回token,用户名不能重复存在。用户信息存储在users表中,密码进行了Md5加密
    
2. 完成登录功能
   
    用户输入正确的用户名和密码完成登录，并返回token。
    
3. 完成用户信息功能

    返回登录后的用户的用户相关信息，需要正确的token信息
    
4. 完成视频投稿功能
   
    登录用户上传视频，保存到本地并截图作为封面。保存到本地时会以视频主键id作为文件名。
    
5. 完成发布列表功能

    查看登录用户发布的所有视频。
    
6. 完成视频流功能

    完成视频流功能，能够返回某一时刻以前的视频列表，时间倒叙，列表最多返回30条视频。
    
7. 完成关注列表和粉丝列表

    完成关注列表和粉丝列表，可以看见自己的关注列表数数量和粉丝列表数量，以及详情。并且也可以看见别人的关注和粉丝