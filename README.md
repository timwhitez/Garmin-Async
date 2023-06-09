# Garmin-Async
自动化使用dailysync.vyzt.dev定期同步佳明国内与国际版账号



## Usage
- 在dailysync.vyzt.dev登陆并且添加国内与国际版账号
- 成功进行同步并抓包 uri:"/api/garmin/syncWorkoutData"
- 将获取到的Cookie与UserId填写到main.go
- go build 编译
- 执行./garminAsync.exe 或者后面跟上请求间隔时间(小时) ./garminAsync.exe 6
- 不建议将频率设置太高，默认为12h一次



## Reference
一切功劳归功于dailysync.vyzt.dev by Yan Zhitao 
