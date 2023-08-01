<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <!-- 加入Bootstrap CSS CDN -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">
    <link rel="shortcut icon" href="static/images/favicon.svg" type="image/x-icon" />
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <link rel="stylesheet" href="static/css/lineicons.css" rel="stylesheet" type="text/css" />
    <link rel="stylesheet" href="static/css/materialdesignicons.min.css" rel="stylesheet" type="text/css" />
    <link rel="stylesheet" href="static/css/fullcalendar.css" />
    <link rel="stylesheet" href="static/css/main.css" />
</head>
<body>
    <!-- 頁面內容區域 -->
    {{template "content" .}}

    <!-- 加入Bootstrap JS CDN -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>