{{ template "base.tpl" . }}
{{define "content"}}
    <h1>創建帳號</h1>
    <form action="" method="post">
    <label for="name">Name:<input type="text" id="name" name="name"></label>
    <label for="account">User:<input type="text" id="account" name="account"></label>
    <label for="password">Password:<input type="password" id="password" name="password"></label>
    <input type="submit" value="創建">
</form>
{{end}}