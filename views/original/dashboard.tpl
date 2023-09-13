{{ template "base.tpl" . }}

{{define "content"}}
<h1>Welcome, {{.name}}!</h1>
<h2>My Owe List</h2>
{{ range .Owelist }}
<div>
    Creditor Name: {{ .Creditor.Name }}
    Items: {{ .Items }}
    Money: {{ .Money }}
    Unit: {{ .Unit }}
    <a class="btn btn-primary" href="/updateowe/{{.Id}}" role="button">Update</a>
    {{if eq .Finish 1}}
    <button class="btn btn-success change-status-btn" data-id="{{.Id}}">確認</button>
    {{end}}
</div>
{{end}}
<a class="btn btn-primary" href="/newowe" role="button">New Owe</a>

<script>
    $(document).ready(function () {
        // 當按鈕被點擊時觸發 AJAX 請求
        $('.change-status-btn').click(function () {
            var id = $(this).data('id');

            $.ajax({
                url: '/updateowe/' + id + '/changestatus',
                type: 'POST',
                dataType: 'json',
                success: function (data) {
                    // 處理伺服器返回的 JSON 資料
                    if (data.status === 'success') {
                        // 更改狀態成功，更新前端畫面
                        alert('狀態更改成功！');
                        // 這裡你可以根據需要更新前端畫面或重新載入頁面
                        location.reload();
                    } else {
                        // 更改狀態失敗，顯示錯誤訊息
                        alert('狀態更改失敗：' + data.message);
                    }
                },
                error: function (xhr, status, error) {
                    // 處理 AJAX 錯誤
                    alert('發生錯誤：' + error);
                }
            });
        });
    });
</script>
{{end}}