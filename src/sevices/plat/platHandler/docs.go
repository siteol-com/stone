package platHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"siteOl.com/stone/server/src/data/constant"
	"siteOl.com/stone/server/src/data/model"
	"siteOl.com/stone/server/src/data/resp"
	"siteOl.com/stone/server/src/sevices"
	"strings"
)

var docHtml = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="renderer" content="webkit">
    <meta name="viewport" content="width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <title>Stone - API文档 - ReDoc版本</title>
	<meta name="description" content="Stone 物联基石项目 API文档 ReDoc版本"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="icon" href="/docs/sc/icon.png"> 
  	<link rel="shortcut icon" href="/docs/sc/icon.png"> 
  	<link rel="apple-touch-icon-precomposed" href="/docs/sc/icon.png"> 
    <style>
        body {
            margin: 0;
            padding: 0;
        }
    </style>
</head>
<body>
<redoc spec-url='/docs/sc/swagger.yaml' expand-responses="200,400,401,403,500" pagination="section"></redoc>
<script src="/docs/sc/redoc.standalone.js"> </script>
</body>
</html>`

var swaggerHtml = `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="renderer" content="webkit">
  <meta name="viewport" content="width=device-width,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <title>Stone - API文档 - SwaggerUI版本</title>
  <meta name="description" content="Stone 物联基石项目 API文档 SwaggerUI版本"/>
  <link rel="icon" href="/docs/sc/icon.png"> 
  <link rel="shortcut icon" href="/docs/sc/icon.png"> 
  <link rel="apple-touch-icon-precomposed" href="/docs/sc/icon.png"> 
  <link rel="stylesheet" href="/docs/sc/swagger.ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="/docs/sc/swagger.ui.bundle.js" crossorigin></script>
<script>
  window.onload = () => {
    window.ui = SwaggerUIBundle({
      url: '/docs/sc/swagger.json',
      dom_id: '#swagger-ui',
    });
  };
</script>
</body>
</html>`

// Sample godoc
// @id			 Sample示例
// @Summary      通用API示例
// @Description  系统API基本示例
// @Router       /docs/sample [post]
// @Tags         开放接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.DemoReq true "示例请求"
// @Success      200 {object} resp.DemoOk "业务受理成功"
// @Failure      400 {object} resp.DemoVail "数据校验失败"
// @Failure      401 {object} resp.DemoAuthLg "当前尚未登陆"
// @Failure      403 {object} resp.DemoAuthNg "权限校验失败"
// @Failure      500 {object} resp.DemoErr "服务系统异常"
func Sample(c *gin.Context) {
	_, req, err := sevices.ValidateReqObj(c, &model.DemoReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.DemoVail{Code: constant.ValidErr, Msg: "参数非法"})
		return
	}
	demoReq := req.(*model.DemoReq)
	switch demoReq.HttpCode {
	case http.StatusInternalServerError: // 500
		c.JSON(http.StatusInternalServerError, resp.DemoErr{Code: constant.SysErr, Msg: "系统异常"})
	case http.StatusBadRequest: // 400
		c.JSON(http.StatusBadRequest, resp.DemoVail{Code: constant.ValidErr, Msg: "参数非法"})
	case http.StatusUnauthorized: // 401
		c.JSON(http.StatusBadRequest, resp.DemoAuthLg{Code: constant.LoginErr, Msg: "当前尚未登陆"})
	case http.StatusForbidden: // 403
		c.JSON(http.StatusBadRequest, resp.DemoAuthNg{Code: constant.AuthErr, Msg: "禁止访问"})
	default:
		c.JSON(http.StatusOK, resp.DemoOk{Code: constant.Success, Msg: "业务请求成功"})
	}
	return
}

// ReDoc HTML加载
func ReDoc(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte(docHtml))
	return
}

// SwaggerDoc HTML加载
func SwaggerDoc(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte(swaggerHtml))
	return
}

// ScFile 静态文件
func ScFile(c *gin.Context) {
	url := c.Request.URL.Path
	fileInfo := url[strings.LastIndex(url, "/"):]
	fileEnd := fileInfo[strings.LastIndex(fileInfo, "."):]
	contextType := "application/octet-stream"
	switch fileEnd {
	case ".png":
		contextType = "image/png"
	case ".js":
		contextType = "application/javascript"
	case ".css":
		contextType = "text/css"
	}
	data, err := os.ReadFile("docs" + fileInfo)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.Data(http.StatusOK, contextType, data)
	}
	return
}
