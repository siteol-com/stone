# 概述

物联网基座Stone，提供一个多层级SaaS化的基础开箱即用中台管理服务。

支持响应码自定义国际化、租户赋权、高粒度权限分配、接口路由规划和实时生效，支持加盟商、业务客户多层级租户嵌套。

本系统的全部接口采用【POST】【application/json】方式传输数据。

除开放接口以外的其他接口均需要通过【ApiKeyAuth:请求头[Token]】完成鉴权。

系统技术栈：Golang、VueNext、MySQL、Redis、Gin、Arco

# 接口文档

该接口文档提供Swagger[支持调试]和ReDoc[阅读增强]两个版本。

[Swagger[支持调试]：http://localhost:8000/docs/swagger/index.html](http://localhost:8000/docs/swagger/index.html)

[ReDoc[阅读增强]：http://localhost:8000/docs/redoc/index.html](http://localhost:8000/docs/redoc/index.html)

## 生成说明

// 首次集成

go install github.com/swaggo/swag/cmd/swag@latest

// 初始化swagger.yaml文件

swag init

// 如果像本工程一样依赖了md文件

swag init --md .

// 删除多余的生成

rm .\docs\docs.go

删除docs.go是因为本项目中并未采用下述依赖来集成。

而是通过HTML模板+JS+YAML引入集成，移除docs.go用于避免依赖编译报错。

- github.com/swaggo/swag
- github.com/swaggo/gin-swagger
- github.com/swaggo/files

源帮助页：[https://github.com/swaggo/swag/blob/master/README_zh-CN.md](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)
