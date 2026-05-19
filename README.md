## FreeTorrents-MTeam
**M-Team站免费种获取工具**

1. 支持筛选免费时长
2. 支持筛选文件大小
3. 支持标题黑名单和分类黑名单筛选
4. 支持新版馒头

### 食用指南
1. 配置文件缺省值为`./conf.yaml`
2. 适配馒头重构，使用`api2`前缀域名
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
|apiKey|从web页面登陆后自行获取，控制台->实验室->存取令牌|
|torrentPath|torrent文件存储位置|
|pageSize|每页捞取数量，未配置或小于等于 `0` 时默认 `200`|
|freeDays|免费限时，单位天(大于配置阈值或长期免费)|
|freeSizeMin|免费种文件大小下限，单位GB，小于配置阈值就不会下载，`0` 表示不启用|
|freeSize|免费种文件大小上限，单位GB，大于配置阈值就不会下载，`0` 表示不启用|
|blockList|屏蔽列表，匹配到就不会下载|
|categoryBlockList|分类ID屏蔽列表，精确匹配到就不会下载，例如 `442`|

#### 筛选顺序
捞取结果会按以下顺序筛选：

1. 免费状态和免费剩余天数
2. 文件大小下限 `freeSizeMin` 和上限 `freeSize`
3. 分类黑名单 `categoryBlockList`
4. 标题黑名单 `blockList`

任一步命中跳过条件后，都会直接跳过下载。

## 坑

1. Windows自带的定时任务不能获取用户挂载nfs的映射盘符，需要以`\\ip\folder`形式配置在`torrentPath`字段中
2. 有些限时很短的免费资源是骗rss做种的，建议筛选时长最少3天以上

## TODO

1. ~~增加文件大小下限范围~~
2. ~~根据资源名设置ignored清单~~
3. ~~调查非免费种污染原因~~
4. 通过计算上传下载比判断资源是否值得下载
