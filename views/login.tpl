{{ template "base.tpl" . }}
{{define "content"}}
    <h1>Welcome to the Home Page</h1>
    <form action="" method="post">
    <label for="name">User:<input type="text" id="name" name="name"></label>
    <label for="password">Password:<input type="password" id="password" name="password"></label>
    <input type="submit" value="登入">
</form>
{{end}}