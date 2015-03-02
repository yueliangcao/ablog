{{range .tags}}    
  <a href="/tag/{{.Name}}" class="btn btn-primary">{{.Name}} <span class="badge">{{.Count}}</span></a>
{{end}}