<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <!-- 加入Bootstrap CSS CDN -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">
</head>
<body>
    <!-- 頁面內容區域 -->
    {{template "content" .}}

    <!-- 加入Bootstrap JS CDN -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>