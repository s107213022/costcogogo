<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <!-- 加入Bootstrap CSS CDN -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">
    <link rel="shortcut icon" href="static/images/favicon.svg" type="image/x-icon" />
    <link rel="stylesheet" href="static/css/bootstrap.min.css" />
    <link rel="stylesheet" href="static/css/lineicons.css" />
    <link rel="stylesheet" href="static/css/materialdesignicons.min.css" />
    <link rel="stylesheet" href="static/css/fullcalendar.css" />
    <link rel="stylesheet" href="static/css/main.css" />
</head>
<body>
    <!-- 頁面內容區域 -->
    {{ block "content" .}}{{end}}

    
    <!-- 加入Bootstrap JS CDN -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/5.3.0/js/bootstrap.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.10.2/dist/umd/popper.min.js" integrity="sha384-7+zCNj/IqJ95wo16oMtfsKbZ9ccEh31eOz1HGyDuCk5Ck3yVpEg5f+2maTj5macd" crossorigin="anonymous"></script>
    <!-- <script src="static/js/bootstrap.bundle.min.js"></script> -->
    <script src="static/js/Chart.min.js"></script>
    <script src="static/js/dynamic-pie-chart.js"></script>
    <script src="static/js/moment.min.js"></script>
    <script src="static/js/fullcalendar.js"></script>
    <script src="static/js/jvectormap.min.js"></script>
    <script src="static/js/world-merc.js"></script>
    <script src="static/js/polyfill.js"></script>
    <script src="static/js/main.js"></script>
    {{ block "scripts" .}}{{end}}

</body>
</html>