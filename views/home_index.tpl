{{range .articles}}    
  <div class="blog-section">
    <div class="row">
      <div class="col-md-10"><h1 class="title"><a href="/article/{{.Id}}">{{.Title}}</a></h1></div>
      <div class="col-md-2">
        <div class="meta">
          <div class="date">{{.CreateOn.Format "1 月 02 日 · 2006"}}</div>
          <div class="tags">{{range $k,$v:=.TagNames}}<span {{if eq $k 0}}style="display:none"{{end}}>,</span><a href="/tag/{{$v}}">{{$v}}</a>{{end}}</div>
        </div>
      </div>
    </div>
    <p>
      {{str2html .Content}}
    </p>
    <hr>
  </div>
{{end}}