<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>网站监测</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css"
          integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <script src="https://code.jquery.com/jquery-3.3.1.min.js"
            integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8=" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js"
            integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl"
            crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js"
            integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q"
            crossorigin="anonymous"></script>
    <style>
        html, body {
            height: 100%;
            width: 100%;
        }

        .container-fluid {
            height: 100%;
        }

        .container-fluid .max-row {
            height: 100%;
        }

        .pre-scrollable {
            max-height: 95%;
        }

        .jumbotron span {
            margin-right: 10px;
        }
    </style>
</head>
<body>
<div class="container-fluid">
    <div class="row clearfix max-row">
        <div class="col-md-5 column rounded border border-light">
            <form role="form" action='{{urlfor "MainController.post"}}' method="post">
                <div class="row">
                    <div class="column col-md-6">
                        <div class="form-group">
                            <label for="sites">需要监测的网站</label>
                            <textarea class="form-control" id="sites" name="sites" rows="12"
                                      placeholder="一行一个地址">{{.sites}}</textarea>
                        </div>
                    </div>
                    <div class="column col-md-6">
                        <div class="form-group">
                            <label for="keywords">监测关键词</label>
                            <textarea class="form-control" id="keywords" name="keywords" rows="12"
                                      placeholder="一行一个关键词">{{.keywords}}</textarea>
                        </div>
                    </div>
                    <div class="form-group">
                        <div class="col-sm-offset-2 col-sm-10">
                            <button type="submit" class="btn btn-primary">保存配置</button>
                        </div>
                    </div>
            </form>
        </div>
        <div class="row">
            <div class="column col-md-6">
                <div class="form-group">
                    <label for="pertimes">每个网站监测次数</label>
                    <select id="pertimes" class="form-control">
                        <option>1</option>
                        <option selected="selected">2</option>
                        <option>3</option>
                        <option>4</option>
                        <option>5</option>
                        <option>6</option>
                        <option>7</option>
                        <option>8</option>
                        <option>9</option>
                        <option>10</option>
                        <option>11</option>
                        <option>12</option>
                        <option>13</option>
                        <option>14</option>
                        <option>15</option>
                    </select>
                </div>
            </div>
            <div class="column col-md-6">
                <div class="form-group">
                    <label for="iptype">监测IP选择</label>
                    <select id="iptype" class="form-control">
                        <option value="1" selected="selected">本机IP</option>
                        <option value="2">代理IP群</option>
                    </select>
                </div>
            </div>
            <div class="column col-md-12">
                <div class="form-group">
                    <label for="proxyapi">代理地址(当监测IP为代理IP群时必须设置)</label>
                    <input type="text" class="form-control" id="proxyapi"
                           value="http://120.25.150.39:8081/index.php/api/entry?method=proxyServer.generate_api_url&packid=1&fa=0&qty=1&time=1&pro=&city=&port=1&format=json&ss=1&css=&ipport=1&et=1&pi=1&co=1&dt=1">
                </div>
            </div>
            <div class="btn-group btn-group-lg col-sm-offset-2 col-sm-10">
                <button type="button" class="btn btn-success" id="btnstart">开始</button>
                <!--<button type="button" class="btn btn-warning" id="btnstop">停止</button>-->
                <!--<button type="button" class="btn btn-info">导出</button>-->
            </div>
            <div class="panel panel-info col-sm-offset-2 col-sm-10">
                <div class="panel-heading">
                    使用说明
                </div>
                <div class="panel-body" style="font-size: 12px;">
                    <span class="text-muted">① 正常：无跳转或http永久跳转https</span><br>
                    <span class="text-muted">② 跳转：结果为被劫持域名及劫持域名</span><br>
                    <span class="text-muted">③ 异常：无跳转，但访问出现错误，需分析</span><br>
                    <span class="text-muted">④ 警告：不满足的监测关键字</span><br>
                    <span class="text-muted">建议使用谷歌浏览器</span>
                </div>
            </div>
        </div>
    </div>
    <div class="col-md-7 column border border-light">
        <div>监测结果</div>
        <div class="pre-scrollable">
            <div class="jumbotron">
                <p>总：<span class="badge badge-secondary" id="totalcount">0</span>异常：<span class="badge badge-danger"
                                                                                          id="errcount">0</span>跳转：<span
                            class="badge badge-danger" id="redirectcount">0</span>正常：<span class="badge badge-light"
                                                                                           id="normalcount">0</span></p>
                <hr class="my-4">
                <p id="logs"></p>
            </div>
        </div>
    </div>
</div>
</div>
<div class="modal fade" id="modal-tip" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-body">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                <span id="tipmsg"></span>
            </div>
        </div>
    </div>
</div>
<!-- loading -->
<div class="modal fade" id="loading" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" data-backdrop='static'>
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title" id="myModalLabel">监测开始</h4>
            </div>
            <div class="modal-body">
                监测中……<span id="result"></span>
            </div>
        </div>
    </div>
</div>
<script>
    jQuery(document).ready(function () {
        var id;

        function getLog() {
            jQuery.ajax({
                url: '/getlog',
                type: 'get',
                dataType: 'json',
                success: function (data) {
                    if (data) {
                        if (data.code == 2) {
                            clearInterval(id);
                            jQuery("#logs").prepend('<span class="text-white bg-secondary">↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓监测结束↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓</span><br>');
                        } else if (data.code == 1 && !!data.data) {
                            var m = data.data;
                            // 统计信息
                            jQuery("#totalcount").text(parseInt(jQuery("#totalcount").text()) + 1);
                            switch (m[i].Status) {
                                case 1:
                                    jQuery("#normalcount").text(parseInt(jQuery("#normalcount").text()) + 1);
                                    break;
                                case 2:
                                    jQuery("#redirectcount").text(parseInt(jQuery("#redirectcount").text()) + 1);
                                    break;
                                default:
                                    jQuery("#errcount").text(parseInt(jQuery("#errcount").text()) + 1);
                            }
                            // 具体日志
                            var html = '<span class="text-primary">' + m[i].SiteName + '</span>'
                                + '<span class="text-muted">状态：<span class="badge ' + (m[i].Status == 1 ? 'badge-light' : 'badge-danger') + '">'
                                + (m[i].Status == 1 ? '正常' : (m[i].Status == 2 ? '跳转' : '异常')) + '</span>' + m[i].StatusCode + '</span>'
                                + '--&gt;<span class="text-muted">' + m[i].Location + '</span>'
                                + '<span class="text-muted">Ip：' + m[i].ip + '</span><span class="text-muted">' + m[i].Loc + '</span><br>';
                            if (!!m[i].err && m[i].err.length > 0) {
                                var list = m[i].err;
                                html += '<p class="text-danger">异常：';
                                for (var i = 0; i < list.length; i++) {
                                    html += list[i] + "<br>";
                                }
                                html += '</p>';
                            }
                            if (!!m[i].warn && m[i].warn.length > 0) {
                                var list = m[i].warn;
                                html += '<p class="text-warning">警告：';
                                for (var i = 0; i < list.length; i++) {
                                    html += list[i] + "<br>";
                                }
                                html += '</p>';
                            }
                            if (!!m[i].info && m[i].info.length > 0) {
                                var list = m[i].info;
                                html += '<p class="text-info">消息：';
                                for (var i = 0; i < list.length; i++) {
                                    html += list[i] + "<br>";
                                }
                                html += '</p>';
                            }
                            jQuery("#logs").prepend(html);
                        }
                    }
                },
                error: function (xhr, Status, error) {
                    alert("异常！");
                }
            });
        }

        function stop(isReflesh) {
            jQuery.ajax({
                url: '/stop',
                type: 'get',
                dataType: 'json',
                success: function (data) {
                    clearInterval(id);
                    if (!isReflesh) {
                        jQuery("#tipmsg").text(data.msg);
                        jQuery('#modal-tip').modal('show');
                        jQuery("#logs").prepend('<span class="text-white bg-secondary">↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓监测结束↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓</span><br>');
                    }
                },
                error: function (xhr, Status, error) {
                    if (!isReflesh) {
                        alert("异常！");
                    }
                }
            });
        }

        jQuery('#btnstart').on('click', function () {
            jQuery.ajax({
                url: '/start?pertimes=' + jQuery("#pertimes").val() + "&iptype=" + jQuery("#iptype").val() + "&proxyapi=" + encodeURIComponent(jQuery("#proxyapi").val()),
                type: 'post',
                dataType: 'json',
                beforeSend: function () {
                    //提示框
                    $('#loading').modal('show');
                },
                success: function (info) {
                    $('#loading').modal('hide');
                    jQuery("#tipmsg").text(info.msg);
                    jQuery('#modal-tip').modal('show');
                    if (info.code === 1) {
                        var myDate = new Date;
                        var year = myDate.getFullYear(); //获取当前年
                        var mon = myDate.getMonth() + 1; //获取当前月
                        var date = myDate.getDate(); //获取当前日
                        var h = myDate.getHours();//获取当前小时数(0-23)
                        var mm = myDate.getMinutes();//获取当前分钟数(0-59)
                        var s = myDate.getSeconds();//获取当前秒
                        var week = myDate.getDay();
                        jQuery("#logs").prepend('<span class="text-white bg-secondary">' + h + '时' + mm + '分' + s + '秒-↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓监测结束↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓</span><br>');
                        var m = info.data.dict;
                        for (var i = 0; i < m.length; i++) {
                            // 统计信息
                            jQuery("#totalcount").text(parseInt(jQuery("#totalcount").text()) + 1);
                            switch (m[i].Status) {
                                case "1":
                                    jQuery("#normalcount").text(parseInt(jQuery("#normalcount").text()) + 1);
                                    break;
                                case "2":
                                    jQuery("#redirectcount").text(parseInt(jQuery("#redirectcount").text()) + 1);
                                    break;
                                default:
                                    jQuery("#errcount").text(parseInt(jQuery("#errcount").text()) + 1);
                            }
                            // 具体日志
                            var html = '<span class="text-primary">' + m[i].SiteName + '</span>'
                                + '<span class="text-muted">状态：<span class="badge ' + (m[i].Status == 1 ? 'badge-light' : 'badge-danger') + '">'
                                + (m[i].Status == 1 ? '正常' : (m[i].Status == 2 ? '跳转' : '异常')) + '</span>' + m[i].StatusCode + '</span>'
                                + '--&gt;<span class="text-muted">' + m[i].Location + '</span>'
                                + '<span class="text-muted">Ip：' + m[i].IP + '</span><span class="text-muted">' + m[i].Loc + '</span><br>';
                            if (m[i].Err !== "") {
                                html += '<p class="text-danger">异常：' + m[i].Err;
                                html += '</p>';
                            }
                            if (m[i].Warn !== "") {
                                html += '<p class="text-warning">警告：' + m[i].Warn;
                                html += '</p>';
                            }
                            if (m[i].Info !== "") {
                                html += '<p class="text-info">消息：' + m[i].Info;
                                html += '</p>';
                            }
                            jQuery("#logs").prepend(html);
                        }
                        jQuery("#logs").prepend('<span class="text-white bg-secondary">' + h + '时' + mm + '分' + s + '秒↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑监测启动↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑</span><br>');
                    }
                },
                error: function (xhr, Status, error) {
                    console.log(error);
                    alert("异常！");
                }
            });
        });
        jQuery('#btnstop').on('click', function () {
            stop(false);
        });
        stop(true);
        window.onbeforeunload = function (event) {
            stop(true);
        };
    });
</script>
</body>
</html>