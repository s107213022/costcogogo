{{ template "base.tpl" . }}
{{define "content"}}
    <h1>{{.name}}</h1>
    {{ range .Owelist }}
    <div>
        <!-- 印出 Creditor 的名稱 -->
        Creditor Name: {{ .Creditor.Name }}
        <!-- 印出其他欄位，例如 Items、Money、Unit 等 -->
        Items: {{ .Items }}
        Money: {{ .Money }}
        Unit: {{ .Unit }}
        <!-- 以此類推，列出其他需要顯示的欄位 -->
    </div>
    <a class="btn btn-primary" href="/newowe" role="button">新欠款</a>
    {{end}}
{{end}}