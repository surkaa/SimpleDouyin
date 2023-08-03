# 仿抖音后端(Douyin, Chinses TikTok)

本仓库旨在写一个简易版的抖音后台

## 目标功能:

1. 正常刷视频功能
    * 发布视频
    * 视频点赞
    * 视频评论
    * 视频收藏
    * 评论点赞
    * 评论回复
2. 用户社交
    * 关注用户
    * 用户聊天
    * 查看其他用户视频

## 项目结构

### 技术栈

* golang
* gin
* gorm
* mysql
* redis
* nginx

### 目录树

```
.
+---config                          - 各种配置  用于项目中涉及的所以配置
|       ConfigurationType.go            - 配置格式
|       loadConfig.go                   - 加载配置文件
|
+---controller                      - 控制器层
|       UserController.go               - 用户控制器
|
+---dao                             - 数据访问层
|   |   UserDao.go                      - 用户数据访问文件
|   |
|   \---impl                        - 数据访问实现层
|           UserDaoImpl.go              - 用户数据访问实现文件
|  
+---db                              - 数据库连接层
|       Mysql.go                        - MySQL连接
|       Redis.go                        - Redis连接
|
+---module                          - 模型层
|       Entity.go                       - 通用的 super 的模型
|       User.go                         - 用户模型
|
+---request                         - 请求体
|       UserInfoBody.go                 - 用户信息请求体
|       UserLoginBody.go                - 用户登录请求体
|       UserRegisterBody.go             - 用户注册请求体
|   
+---response                        - 响应体
|       Response.go                     - 通用相应体
|       UserInfoResponse.go             - 用户信息响应体
|       UserLoginResponse.go            - 用户登录响应体
|       UserRegisterResponse.go         - 用户注册响应体
|
+---router                          - 路由
|       Router.go                       - 路由
|
+---service                         - 服务层
|   |   UserService.go                  - 用户服务
|   |
|   \---impl                        - 服务实现层
|           UserServiceImpl.go          - 用户服务
|
+---sql                             - 数据库
|       initial.sql                     - 数据库建表语句
|
+---test                            - 测试包
|       user_dao_test.go                - 测试用户数据访问层
|
\---util                            - 工具包
        Snowflake.go                    - 雪花生成id
        Time.go                         - 时间工具包
        Token.go                        - jwt生成token
```
