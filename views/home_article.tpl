{{with .article}}
<div class="article">
  <h1>{{.Title}}</h1>
  <div>{{str2html .Content}}</div>
</div>
<div class="comment">
  <div id="SOHUCS" sid="{{.Id}}"></div>
</div>
{{end}}

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

<style>
  #SOHUCS #SOHU_MAIN #powerby_sohu {display: none!important;}/*去除畅言by*/
</style>