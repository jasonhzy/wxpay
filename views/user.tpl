<!Doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>数据展示</title>
</head>
<body>
    <table border="1" align="center" cellspacing=0>
        <tr><th>ID</th><th>用户名</th><th>昵称</th><th>年龄</th><th>创建时间</th><th>更新时间</th></tr>
        {{ range $k, $v := .data }}
           <tr>
           <td>{{$v.Id}}</td>
           <td>{{$v.Username}}</td>
           <td>{{$v.Nickname}}</td>
           <td>{{$v.Age}}</td>
           <td>{{dateformat $v.Create_time "2006-01-02 15:04:05"}}</td>
           <td>{{dateformat $v.Update_time "2006-01-02 15:04:05"}}</td>
           </tr>
        {{else}}
            <tr><td colspan="9">无记录</td></tr>
        {{end}}
        {{ if .count }}
            <tr><td colspan="1">总数:</td><td colspan="5"> {{ .count }}</td></tr>
        {{end}}
    </table>
</body>
</html>