<!DOCTYPE html>
<html>
<head>
  <title>ablog-login</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-theme.min.css">
  <script type="text/javascript" src="/static/js/jquery-1.9.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <style type="text/css">
  body {
    color: #797979;
  }
  h1 {
    font-size: 24px;
    margin-top: 0px;
    margin-bottom: 20px;
  }
  body.login {
    background-color: rgb(190, 190, 190);
  }
  .login-page-container {
    width: 300px;
    margin: 100px auto 0 auto;
    background-color: white;
  }
  .page-content {
    background-color: white;
    margin-left: 200px;
  }
  .nav-blog {
    margin-top: 200px;
  }
  .nav-blog >li+li {
    margin-top: 0;
  }
  .nav-blog > li > a {
    color: white;
    border-top-style: groove;
    border-top-color: white;
    border-width: 1px;
  }
  .nav-blog > li.active > a,.nav-blog > li > a:hover{
    background-color: white;
    color: #797979;
  }
  .nav-blog span.title {
    margin-left: 10px;
  }
  .blog-section {
    margin-top: 20px;
  }
  </style>
</head>
<body class="login">
<div class="login-page-container">
  <form class="form-vertical login-form" action="index.html">
      <h3 class="form-title">Login to your account</h3>
      <div class="alert alert-error hide">
        <button class="close" data-dismiss="alert"></button>
        <span>Enter any username and password.</span>
      </div>
      <div class="control-group">
 
        <!--ie8, ie9 does not support html5 placeholder, so we just show field title for that-->

        <label class="control-label visible-ie8 visible-ie9">Username</label>

        <div class="controls">

          <div class="input-icon left">

            <i class="icon-user"></i>

            <input class="m-wrap placeholder-no-fix" type="text" placeholder="Username" name="username">

          </div>

        </div>

      </div>

      <div class="control-group">

        <label class="control-label visible-ie8 visible-ie9">Password</label>

        <div class="controls">

          <div class="input-icon left">

            <i class="icon-lock"></i>

            <input class="m-wrap placeholder-no-fix" type="password" placeholder="Password" name="password">

          </div>

        </div>

      </div>

      <div class="form-actions">

        <label class="checkbox">

        <div class="checker"><span><input type="checkbox" name="remember" value="1"></span></div> Remember me

        </label>

        <button type="submit" class="btn green pull-right">

        Login <i class="m-icon-swapright m-icon-white"></i>

        </button>            

      </div>

      <div class="forget-password">

        <h4>Forgot your password ?</h4>

        <p>

          no worries, click <a href="javascript:;" class="" id="forget-password">here</a>

          to reset your password.

        </p>

      </div>

      <div class="create-account">

        <p>

          Don't have an account yet ?&nbsp; 

          <a href="javascript:;" id="register-btn" class="">Create an account</a>

        </p>

      </div>

    </form>
</div>
</body>
</html>
