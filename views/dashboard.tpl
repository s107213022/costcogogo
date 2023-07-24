{{ template "base.tpl" . }}
{{define "content"}}
    <h1>{{.name}}</h1>
    {{ range .Owelist }}
    <div>
        items: {{ .item }}
        <!-- 其他字段... -->
    </div>
    {{end}}
{{end}}