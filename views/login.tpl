{{ template "base.tpl" . }}
{{define "content"}}
    <h1>Welcome to the Home Page</h1>
    <form action="" method="post">
    <label for="account">User:<input type="text" id="account" name="account"></label>
    <label for="password">Password:<input type="password" id="password" name="password"></label>
    <input type="submit" value="登入">
</form>
{{end}}