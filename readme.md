## 如果您喜欢这个开源项目，不妨给它点个⭐️⭐️⭐️

# 项目地址

api接口地址: https://www.showdoc.cc/2466392510613440?page_id=10951794335914772

前端项目地址: https://github.com/ChenMiaoQiu/go-cloud-disk-front

项目在线地址: http://114.55.234.33/ (域名备案中)

# 介绍

**go-cloud-disk**是一个使用Go语言实现的在线网盘系统，采用前后端分离技术，Go语言提供api进行数据支撑，前端页面采用vue3+element-plus进行编写。
支持文件上传，分享，创建文件夹等多种功能。支持管理员对用户权限动态修改，对用户分享文件一键封禁。

![1710427409318](images/readme/1710427409318.png)

![1710427425259](images/readme/1710427425259.png)

![1710427386984](images/readme/1710427386984.png)

![1710427452714](images/readme/1710427452714.png)

![1710427476445](images/readme/1710427476445.png)

![1710427509775](images/readme/1710427509775.png)

## 采用技术介绍

前端：
使用element-plus+vue快速构建前端页面
使用pinia对用户信息进行存储，保证全局状态一致性
使用pinia-plugin-persistedstate将token存储至localStorage，实现用户刷新后无需再次登录，优化用户体验
使用vue-router实现多级路由跳转，完成单页面应用;对axios进行二次封装，完成前后端数据通讯
利用路由守卫完成路由鉴权，确保非登录状态下无权访问预约及个人中心页面
后端
使用JWT进行身份验证
权限管理使用 CASBIN，实现基于 RBAC模式 的权限管理
使用腾讯云cos进行文件存储，并使用接口对相关功能进行统一，方便后续扩展更多云服务器平台
使用Redis的Zset数据结构和cron实现每日排行榜功能
使用Redis对当日高频次访问的分享进行存储，提高接口响应速度，每1w次访问平均耗时从200ms下降至10ms
使用go-email+redis实现邮箱验证码功能，使用户在注册时使用邮箱注册
其他
采用 Restful 风格的 API
前后端分离部署，使用nginx进行反向代理，优化服务器安全性
使用docker+portainer 将项目部署在腾讯云，阿里云平台

## 具体技术栈

前端技术栈: 使用npm包进行管理

- 基于TypeScript
- Vue3
- Pinia
- Vue Router
- Axios
- Element-plus
- pinia-plugin-persistedstate
- ...

后端技术栈:

- Gin
- Gorm
- go-redis
- cron
- casbin
- Docker
- Nginx 部署静态资源 + 反向代理
- ...

其他：

- 腾讯云Cos
- 阿里云服务器

# 部署

按照.env.example进行环境配置，创建.env文件，创建docker时配置环境变量都可
