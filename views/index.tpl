<!DOCTYPE html>
<html>
<head>
  <title>ablog</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
  <!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-theme.min.css"> -->
  <link rel="stylesheet" type="text/css" href="/static/css/style.css">
  <script type="text/javascript" src="/static/js/jquery-1.9.1.min.js"></script>
  <script type="text/javascript" src="/static/js/jquery.form.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<script type="text/javascript">
  $(function(){
    $(".form-login").ajaxForm(function(data){
      if (data.succ)
        location.href = "/"
      else
        alert(data.msg)
    });
  });
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
    {{range .articles}}    
      <div class="blog-section">
        <div class="row">
          <div class="col-md-10"><h1><a href="/article/{{.Id}}">{{.Title}}</a>
            <div class="btn-group">
  <a type="button" class="btn btn-default btn-xs" href="/article/edit/{{.Id}}">编辑</a>
  <a type="button" class="btn btn-default btn-xs dropdown-toggle" data-toggle="dropdown">
    <span class="caret"></span>
    <span class="sr-only">Toggle Dropdown</span>
  </a>
  <ul class="dropdown-menu">
    <li><a href="/article/del/{{.Id}}">删除</a></li>
    <li><a href="/article/edit/{{.Id}}">修改分类</a></li>
    <li><a href="/article/top/{{.Id}}">置顶</a></li>
  </ul>
</div></h1></div>
          <div class="col-md-2">
            <div class="meta">
              <div class="date">2015/10/10</div>
              <div class="tags">Go</div>
            </div>
          </div>
        </div>
        <hr>
        <p>
          {{.Content}}
        </p>
        <hr>
      </div>
      {{end}}
  </div>
  <div class="modal fade login-modal" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel"aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-body">
          <form class="form-login" action="/login" method="post">
            <h2 class="form-login-heading">请登录</h2>
            <input name="Usn" type="text" class="form-control usn" placeholder="Username" autofocus="">
            <input name="Pwd" type="text" class="form-control pwd" placeholder="Password">
            <div class="checkbox">
              <label>
                <input type="checkbox" value="remember-me"> 记住我
              </label>
            </div>
            <button class="btn btn-lg btn-primary btn-block" type="submit">Login</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
