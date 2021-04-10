<!DOCTYPE html>
<html lang="zh">
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta charset="UTF-8">

    <title>{{.Title}}</title>
    <meta name="keywords" content="{{.Keywords}}"/>
    <meta name="description" content="{{.Description}}"/>

    <link href="/resource/css/style.css" type="text/css" rel="stylesheet" />
</head>
<body>
<div class="root-container">
    {{include "module/menu.tpl" .}}
    {{include .MainTpl .}}
</div>
</body>
</html>