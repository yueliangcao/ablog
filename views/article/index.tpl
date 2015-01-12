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
  {{with .article}}
    <h1>{{.Title}}</h1>
    <div>{{.Content}}</div>
  {{end}}
  </div>
</body>
</html>
