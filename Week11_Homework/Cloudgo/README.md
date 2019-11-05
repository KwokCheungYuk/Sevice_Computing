### 框架选择 Martini
优点（官方）：

- 使用极其简单.
- 无侵入式的设计.
- 很好的与其他的Go语言包协同使用.
- 超赞的路径匹配和路由.
- 模块化的设计 - 容易插入功能件，也容易将其拔出来.
- 已有很多的中间件可以直接使用.（ 更多的中间件和功能组件, 请查看代码仓库: [martini-contrib](https://github.com/martini-contrib)）
- 框架内已拥有很好的开箱即用的功能支持.
- **完全兼容[http.HandlerFunc](http://godoc.org/net/http#HandlerFunc)接口.**



### 使用curl测试

```
gzx@gzx-VirtualBox:~$ curl -v http://localhost:4869
* Rebuilt URL to: http://localhost:4869/
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 4869 (#0)
> GET / HTTP/1.1
> Host: localhost:4869
> User-Agent: curl/7.58.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: text/html; charset=UTF-8
< Date: Tue, 05 Nov 2019 08:00:32 GMT
< Content-Length: 657
< 
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>CloudGo</title>
</head>

<body>
    <div id="image">
        <img src="1.jpg" height="100%" width="100%" />
    </div>

    <div>
        <p color = red >Hello, Please log in!</p>
        <p class="content">Now is 2019-11-05 16:00:32</p>
    </div>

    <div id="the_form">
        <form method="post" action="/">
            <p>Username:</p>
            <input type="text" name="username"><br />
            <p>Password:</p>
            <input type="password" name="password"><br />
            <input type="submit" value="登录" id="submit">
        </form>
    </div>
</body>

* Connection #0 to host localhost left intact
</html>
```



### 使用ab执行压力测试

发送10000个请求，并发数为1000个

```
gzx@gzx-VirtualBox:~$ ab -n 10000 -c 1000 http://localhost:4869/
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            4869

Document Path:          /
Document Length:        657 bytes

Concurrency Level:      1000
Time taken for tests:   8.506 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      7740000 bytes
HTML transferred:       6570000 bytes
Requests per second:    1175.66 [#/sec] (mean)
Time per request:       850.583 [ms] (mean)
Time per request:       0.851 [ms] (mean, across all concurrent requests)
Transfer rate:          888.64 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   42 193.4      0    1034
Processing:     2  785 125.0    816    1081
Waiting:        1  785 125.6    816    1081
Total:          2  827 247.1    817    2102

Percentage of the requests served within a certain time (ms)
  50%    817
  66%    840
  75%    852
  80%    866
  90%    891
  95%    924
  98%   1880
  99%   2054
 100%   2102 (longest request)
```



参数解释在博客中。