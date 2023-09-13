{{ template "base.tpl" . }}
{{define "content"}}
    <h1>還錢</h1>
    <form action="" method="post">
    <label for="debtor">人:<input type="text" id="debtor" name="debtor"></label>
    <label for="items">商品:<input type="text" id="items" name="items"></label>
    <label for="money">價格:<input type="text" id="money" name="money"></label>
    <label for="unit">數量:<input type="text" id="unit" name="unit"></label>
    
</form>
<input type="submit" value="創建">
{{end}}