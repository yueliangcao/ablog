<!DOCTYPE html>
<html>
<head>
  <title>{{.title}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-theme.min.css"> -->
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
  <script type="text/javascript" src="/static/js/jquery-1.9.1.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.form.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</head>
<body>
<div class="page-container">
  <div class="page-sidebar">
    <div class="settings">
      <img src="/static/img/head-pic.jpg">
      <h1>羊咩咩</h1>
      <p>“羊年喜洋洋，新人新气象”</p>
    </div>
    <ul class="nav nav-stacked nav-blog">
      <li {{if eq .active "home"}} class="active" {{end}}>
          <a href="/">
            <i class="glyphicon glyphicon-home"></i> 
            <span class="title">首页</span>
          </a>
      </li>
      <li {{if eq .active "archives"}} class="active" {{end}}>
          <a href="/archives">
            <i class="glyphicon glyphicon-time"></i> 
            <span class="title">时间轴</span>
          </a>
      </li>
      <li {{if eq .active "tags"}} class="active" {{end}}>
          <a href="/tags">
            <i class="glyphicon glyphicon-tags"></i> 
            <span class="title">标签贴</span>
          </a>
      </li>
      <li {{if eq .active "aboutme"}} class="active" {{end}}>
          <a href="/aboutme">
            <i class="glyphicon glyphicon-eye-open"></i> 
            <span class="title">关于我</span>
          </a>
      </li>
    </ul>
  </div>
  <div class="page-content">
    <div class="container-fluid">
    <div class="row">
      <div class="col-xs-12">
        {{.LayoutContent}}
        <div class="footer">
          <p>Copyright © 2014-2015 - <a href="#" title="yueliang">YueLiang</a> - <span class="credit">Powered by <a href="/">ABlog</a></span></p>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
