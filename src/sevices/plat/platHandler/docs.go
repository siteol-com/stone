package platHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
		c.JSON(404, nil)
	} else {
		c.Data(http.StatusOK, contextType, data)
	}
	return
}
