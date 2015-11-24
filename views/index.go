package views

import (
	"TWLibrary/app"
)

type IndexOptions struct {
	Title string
}

const Index = `	<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html lang="en">
<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
	<title>{{.Title}}</title>
	<link rel="icon" type="image/x-icon" href="public/images/favicon.ico">
	<link rel="stylesheet" type="text/css" href="public/styles/main.css">
	<link type="image" href="public/images/wall.jpg">
</head>

<body>
	<span>
		<h1 id="ww">` + app.AppName + `</h1>
	</span>
	<div id="root"></div>
	<script src="public/javascript/main.js"></script>
</body>
</html>`
