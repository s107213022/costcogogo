{{ template "base.tpl" . }}
{{define "content"}}
<h1>Update Owe</h1>
<form method="post" action="">
    <div>
        <!-- 以 input 元素顯示修改的欄位，使用 value 屬性填入預設值 -->
        <label>Items: </label>
        <input type="text" name="debtor" value="{{ .Owelist.Debtor.Name }}">
        <br>
        <label>Items: </label>
        <input type="text" name="items" value="{{ .Owelist.Items }}">
        <br>
        <label>Money: </label>
        <input type="text" name="money" value="{{ .Owelist.Money }}">
        <br>
        <label>Unit: </label>
        <input type="text" name="unit" value="{{ .Owelist.Unit }}">
        <br>
        <!-- 以此類推，列出其他需要修改的欄位，並填入預設值 -->
    </div>
    <input type="submit" value="Save Changes">
</form>
{{end}}