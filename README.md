# stone

## 介绍

基座工程，物联网项目基石。SaaS授权管理中心，多租户、权限集。计划提供Golang和Java两个版本。

## 当前说明

工程尚在构建阶段，仓库未提供数据库初始化脚本。

请耐心等待1.0版本完成。

# SwaggerAPI文档

轻度依赖集成

## 生成说明

- 首次集成

go install github.com/swaggo/swag/cmd/swag@latest

- 初始化swagger.yaml文件

swag init

- 如果像本工程一样依赖了md文件

swag init --md .

- 删除多余的生成

rm .\docs\docs.go

删除docs.go是因为本项目中并未采用下述依赖来集成。

而是通过HTML模板+JS+YAML引入集成，移除docs.go用于避免依赖编译报错。

- github.com/swaggo/swag
- github.com/swaggo/gin-swagger
- github.com/swaggo/files

源帮助页：[https://github.com/swaggo/swag/blob/master/README_zh-CN.md](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

## API文档地址

该接口文档提供Swagger[支持调试]和ReDoc[阅读增强]两个版本。

[Swagger[支持调试]：http://localhost:8000/docs/swagger/index.html](http://localhost:8000/docs/swagger/index.html) 

[ReDoc[阅读增强]：http://localhost:8000/docs/redoc/index.html](http://localhost:8000/docs/redoc/index.html)