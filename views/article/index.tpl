<!DOCTYPE html>
<html>
<head>
  <title>{{.article.Title}}</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-theme.min.css"> -->
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
  <script type="text/javascript" src="/static/js/jquery-1.9.1.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.form.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script type="text/javascript">
    (function(){
      var appid = 'cyrCM8in5',
      conf = 'prod_6d8824628b0c8855951252a0382f1e63';
      var doc = document,
      s = doc.createElement('script'),
      h = doc.getElementsByTagName('head')[0] || doc.head || doc.documentElement;
      s.type = 'text/javascript';
      s.charset = 'utf-8';
      s.src =  'http://assets.changyan.sohu.com/upload/changyan.js?conf='+ conf +'&appid=' + appid;
      h.insertBefore(s,h.firstChild);
      window.SCS_NO_IFRAME = true;
    })()
  </script>
</head>
<body>
<div class="page-container">
  <div class="page-sidebar">
    <ul class="nav nav-stacked nav-blog">
      <li class="active">
          <a href="/">
            <i class="glyphicon glyphicon-home"></i> 
            <span class="title">Home</span>
          </a>
      </li>
      <li>
          <a href="/timeline">
            <i class="glyphicon glyphicon-time"></i> 
            <span class="title">Time Line</span>
          </a>
      </li>
      <li>
          <a href="/tags">
            <i class="glyphicon glyphicon-time"></i> 
            <span class="title">Tags</span>
          </a>
      </li>
      <li>
          <a href="/aboutme">
            <i class="glyphicon glyphicon-eye-open"></i> 
            <span class="title">About Me</span>
          </a>
      </li>
      <li>
          <a href="#" data-toggle="modal" data-target=".login-modal">
            <i class="glyphicon glyphicon-log-in"></i> 
            <span class="title">Login</span>
          </a>
      </li>
    </ul>
  </div>
  <div class="page-content">
  <div class="container-fluid">
  <div class="row">
  <div class="col-xs-12">
  {{with .article}}
  <div class="article">
    <h1>{{.Title}}</h1>
    <div>{{str2html .Content}}</div>
  </div>
  <div class="comment">
    <div id="SOHUCS" sid="{{.Id}}"></div>
  </div>
  {{end}}
  <div class="footer">
    <p>Copyright © 2014-2015 - <a href="#" title="yueliang">Yueliang</a> - <span class="credit">Powered by <a href="/">ABlog</a></span></p>
  </div>
  </div>
  </div>
  </div>
  </div>
  <style>
    #SOHUCS #SOHU_MAIN #powerby_sohu {display: none!important;}/*去除畅言by*/
  </style>  
</body>
</html>
