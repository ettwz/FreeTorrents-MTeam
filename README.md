## FreeTorrents-MTeam
**M-Team站免费种获取工具**

页面层级结构参考自 [ptscripts](https://github.com/cshenxi/ptscripts) 项目

1. 支持筛选免费时长
2. 支持筛选文件大小

### 食用指南
1. 配置文件缺省值为`./conf.yaml`
2. 适配国内网络环境，使用`kp`前缀域名
3. 日志文件为`./freeTorrent.log`

#### 编译

参考 [Golang 在 Mac、Linux、Windows 下如何交叉编译
](https://blog.csdn.net/panshiqu/article/details/53788067)

#### 使用
#### Linux
./FreeTorrents-MTeam -c 配置文件路径

#### Windows
.\FreeTorrents-MTeam.exe -c 配置文件路径


#### 配置文件
|字段|值|
|---|---|
|siteCookie|从web页面登陆后自行获取，cookie中tp字段的值|
|passKey|web页面中个人设置-设定首页-密钥字段的值|
|userAgent|从web页面登陆后自行获取，从网络请求中粘贴|
|torrentPath|torrent文件存储位置|
|freeDays|免费限时，单位天(大于配置阈值或长期免费)|
|freeSize|免费种文件大小，单位GB(小于配置阈值)|

## 坑

1. Windows自带的定时任务不能获取用户挂载nfs的映射盘符，需要以`\\ip\folder`形式配置在`torrentPath`字段中
2. `UserAgent`字段必须和获取cookie的设备信息匹配，不然会被拦截到google动态码的两步认证页面
3. 偶发拉取到非免费资源(怀疑是站方扩散行为)

## TODO

1. 增加文件大小下限范围
2. 根据资源名设置ignored清单
3. 调查非免费种污染原因
4. 通过计算上传下载比判断资源是否值得下载