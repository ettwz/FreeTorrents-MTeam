# OpenAPI definition


**简介**:OpenAPI definition


**HOST**:http://test2.m-team.cc/api


**联系人**:


**Version**:v0


**接口路径**:/api/v3/api-docs/normal


[TOC]






# 菠菜


## addBetgameOpt


**接口地址**:`/api/bet/addBetgameOpt`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "createdDate": "",
  "lastModifiedDate": "",
  "id": 0,
  "gameid": 0,
  "text": "",
  "odds": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|betoptions|Betoptions|body|true|Betoptions|Betoptions|
|&emsp;&emsp;createdDate|||false|string(date-time)||
|&emsp;&emsp;lastModifiedDate|||false|string(date-time)||
|&emsp;&emsp;id|||false|integer(int32)||
|&emsp;&emsp;gameid|||false|integer(int32)||
|&emsp;&emsp;text|||false|string||
|&emsp;&emsp;odds|||false|number(double)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## betgameDetailLog


**接口地址**:`/api/bet/betgameDetailLog`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|search||query|true|PageForm|PageForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## betgameOdds


**接口地址**:`/api/bet/betgameOdds`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|optId||query|true|integer(int32)||
|bonus||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## bonusTopList


**接口地址**:`/api/bet/bonusTopList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|win||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## createOrUpdate


**接口地址**:`/api/bet/createOrUpdate`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "id": 0,
  "heading": "",
  "undertext": "",
  "endtime": "",
  "sort": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|betgamesForm|BetgamesForm|body|true|BetgamesForm|BetgamesForm|
|&emsp;&emsp;id|||false|integer(int32)||
|&emsp;&emsp;heading|||false|string||
|&emsp;&emsp;undertext|||false|string||
|&emsp;&emsp;endtime|||false|string(date-time)||
|&emsp;&emsp;sort|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## delBetgame


**接口地址**:`/api/bet/delBetgame`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|gameid||query|true|integer(int32)||
|reason||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## delBetgameOpt


**接口地址**:`/api/bet/delBetgameOpt`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|betoptId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## findBetgameList


**接口地址**:`/api/bet/findBetgameList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|active|可用值:PENDING,LIVE,OVER,FINISH|query|true|string||
|fix||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## gamefinish


**接口地址**:`/api/bet/gamefinish`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|optId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getDetail


**接口地址**:`/api/bet/getDetail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|gameId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getDetailBetList


**接口地址**:`/api/bet/getDetailBetList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|gameId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## myCouponLog


**接口地址**:`/api/bet/myCouponLog`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## nullBetgame


**接口地址**:`/api/bet/nullBetgame`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|gameid||query|true|integer(int32)||
|shite||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## 狀態資訊


**接口地址**:`/api/bet/state`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## updateBetgameStatus


**接口地址**:`/api/bet/updateBetgameStatus`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|gameId||query|true|integer(int32)||
|active|可用值:PENDING,LIVE,OVER,FINISH|query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 導航菜單


## list


**接口地址**:`/api/menu/list`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 管理組信箱


## post


**接口地址**:`/api/staffbox/post`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "title": "",
  "context": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|staffMessagePostForm|StaffMessagePostForm|body|true|StaffMessagePostForm|StaffMessagePostForm|
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;context|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 好友


## addBlock


**接口地址**:`/api/friends/addBlock`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|blockId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## 添加好友


**接口地址**:`/api/friends/addFriend`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|friendId|對方uid|query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getBlocks


**接口地址**:`/api/friends/getBlocks`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## 好友列表


**接口地址**:`/api/friends/getFriends`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## removeBlock


**接口地址**:`/api/friends/removeBlock`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|blockId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## removeFriend


**接口地址**:`/api/friends/removeFriend`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|friendId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 積分流水


## logs


**接口地址**:`/api/credit/logs`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "uid": 0,
  "type": "",
  "relatedId": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|creditQueryForm|CreditQueryForm|body|true|CreditQueryForm|CreditQueryForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;uid|||true|integer(int32)||
|&emsp;&emsp;type|可用值:TRACKER,POLLVOTE,FUNREWARD,SUBTITLE,TORRENT,SEEK,ARTIFICIAL,DMM,ORDERS,BET,EXCHANGE,OTHER,FORUM,CHARITY,ALBUM||false|string||
|&emsp;&emsp;relatedId|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 積分商店


## exchange


**接口地址**:`/api/mall/exchange`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|ExchangeForm|ExchangeForm|
|&emsp;&emsp;goodsId|||true|integer(int32)||
|&emsp;&emsp;num|||true|integer(int32)||
|&emsp;&emsp;title|||false|string||
|&emsp;&emsp;bonuscharity|||false|integer(int32)||
|&emsp;&emsp;ratiocharity|||false|number(float)||
|&emsp;&emsp;torrentId|||false|integer(int32)||
|&emsp;&emsp;bonusFreeCharity|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getGlobalFreeSingleList


**接口地址**:`/api/mall/getGlobalFreeSingleList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getGlobalFreeSinglePrice


**接口地址**:`/api/mall/getGlobalFreeSinglePrice`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|goodsId||query|true|integer(int32)||
|torrentId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## globalFreeSingleAuction


**接口地址**:`/api/mall/globalFreeSingleAuction`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|queueId||query|true|integer(int32)||
|price||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## list


**接口地址**:`/api/mall/list`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 舉報


## report


**接口地址**:`/api/report/report`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "type": "",
  "reportid": 0,
  "reason": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|reportForm|ReportForm|body|true|ReportForm|ReportForm|
|&emsp;&emsp;type|可用值:torrent,user,offer,request,post,comment,subtitle,complete,album,albumTorrent||true|string||
|&emsp;&emsp;reportid|||true|integer(int32)||
|&emsp;&emsp;reason|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 考核


## getMyActiveList


**接口地址**:`/api/examine/getMyActiveList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 聯盟小組


## apply


**接口地址**:`/api/team/apply`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "name": "",
  "email": "",
  "reason": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|teamApplyForm|TeamApplyForm|body|true|TeamApplyForm|TeamApplyForm|
|&emsp;&emsp;name|||true|string||
|&emsp;&emsp;email|||true|string||
|&emsp;&emsp;reason|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## myTeams


**接口地址**:`/api/team/myTeams`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## updateMembers


**接口地址**:`/api/team/updateMembers`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|teamId||query|true|integer(int32)||
|members||query|true|array|integer|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 論壇


## forums


**接口地址**:`/api/forum/forums`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## postDel


**接口地址**:`/api/forum/post/del`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|pid||query|true|integer(int32)||
|reason||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## postDetail


**接口地址**:`/api/forum/post/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|pid||query|true|integer(int32)||
|origin||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## postEdit


**接口地址**:`/api/forum/post/edit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "fid": 0,
  "tid": 0,
  "pid": 0,
  "body": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicReplyForm|TopicReplyForm|body|true|TopicReplyForm|TopicReplyForm|
|&emsp;&emsp;fid|||true|integer(int32)||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;pid|||false|integer(int32)||
|&emsp;&emsp;body|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## postNew


**接口地址**:`/api/forum/post/new`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "fid": 0,
  "tid": 0,
  "pid": 0,
  "body": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicReplyForm|TopicReplyForm|body|true|TopicReplyForm|TopicReplyForm|
|&emsp;&emsp;fid|||true|integer(int32)||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;pid|||false|integer(int32)||
|&emsp;&emsp;body|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## postSearch


**接口地址**:`/api/forum/post/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "fid": 0,
  "author": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicPostSearch|TopicPostSearch|body|true|TopicPostSearch|TopicPostSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;fid|||false|integer(int32)||
|&emsp;&emsp;author|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## topicDel


**接口地址**:`/api/forum/topic/del`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|tid||query|true|integer(int32)||
|reason||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## topicDetail


**接口地址**:`/api/forum/topic/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "tid": 0,
  "authorId": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicDetailForm|TopicDetailForm|body|true|TopicDetailForm|TopicDetailForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;authorId|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## topicEdit


**接口地址**:`/api/forum/topic/edit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "fid": 0,
  "tid": 0,
  "subject": "",
  "body": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicForm|TopicForm|body|true|TopicForm|TopicForm|
|&emsp;&emsp;fid|||true|integer(int32)||
|&emsp;&emsp;tid|||false|integer(int32)||
|&emsp;&emsp;subject|||true|string||
|&emsp;&emsp;body|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## topicMod


**接口地址**:`/api/forum/topic/mod`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "fid": 0,
  "tid": 0,
  "targetFid": 0,
  "sticky": true,
  "hlcolor": "",
  "locked": true
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicModForm|TopicModForm|body|true|TopicModForm|TopicModForm|
|&emsp;&emsp;fid|||true|integer(int32)||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;targetFid|||false|integer(int32)||
|&emsp;&emsp;sticky|||false|boolean||
|&emsp;&emsp;hlcolor|||false|string||
|&emsp;&emsp;locked|||false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## topicPost


**接口地址**:`/api/forum/topic/post`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "fid": 0,
  "tid": 0,
  "subject": "",
  "body": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicForm|TopicForm|body|true|TopicForm|TopicForm|
|&emsp;&emsp;fid|||true|integer(int32)||
|&emsp;&emsp;tid|||false|integer(int32)||
|&emsp;&emsp;subject|||true|string||
|&emsp;&emsp;body|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## redirect


**接口地址**:`/api/forum/topic/redirect`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|ForumRedirectForm|ForumRedirectForm|
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;action|可用值:findpost,lastpost||true|string||
|&emsp;&emsp;pid|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|400|Bad Request||


**响应参数**:


暂无


**响应示例**:
```javascript

```


## redirectV2


**接口地址**:`/api/forum/topic/redirectV2`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|ForumRedirectForm|ForumRedirectForm|
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;action|可用值:findpost,lastpost||true|string||
|&emsp;&emsp;pid|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## topicSearch


**接口地址**:`/api/forum/topic/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "fid": 0,
  "author": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|topicSearch|TopicSearch|body|true|TopicSearch|TopicSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;fid|||false|integer(int32)||
|&emsp;&emsp;author|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## topicEdit


**接口地址**:`/api/forum/topic/viewHits`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|tid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 評論


## del


**接口地址**:`/api/comment/del`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## detail


**接口地址**:`/api/comment/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||
|origin||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## edit


**接口地址**:`/api/comment/edit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||
|text||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## fetchList


**接口地址**:`/api/comment/fetchList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "type": "",
  "relationId": 0,
  "authorId": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|commentListForm|CommentListForm|body|true|CommentListForm|CommentListForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;type|可用值:TORRENT,SEEK||false|string||
|&emsp;&emsp;relationId|||false|integer(int32)||
|&emsp;&emsp;authorId|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## post


**接口地址**:`/api/comment/post`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "type": "",
  "relationId": 0,
  "context": "",
  "refType": "",
  "refId": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|commentForm|CommentForm|body|true|CommentForm|CommentForm|
|&emsp;&emsp;type|可用值:TORRENT,SEEK||true|string||
|&emsp;&emsp;relationId|||true|integer(int32)||
|&emsp;&emsp;context|||true|string||
|&emsp;&emsp;refType|可用值:normal,quote,reply||true|string||
|&emsp;&emsp;refId|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## redirect


**接口地址**:`/api/comment/redirect`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|CommentRedirectForm|CommentRedirectForm|
|&emsp;&emsp;type|可用值:TORRENT,SEEK||true|string||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;action|可用值:findpost,lastpost||true|string||
|&emsp;&emsp;pid|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|400|Bad Request||


**响应参数**:


暂无


**响应示例**:
```javascript

```


## redirectV2


**接口地址**:`/api/comment/redirectV2`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|CommentRedirectForm|CommentRedirectForm|
|&emsp;&emsp;type|可用值:TORRENT,SEEK||true|string||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;action|可用值:findpost,lastpost||true|string||
|&emsp;&emsp;pid|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 求種


## addto


**接口地址**:`/api/seek/addto`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekId||query|true|integer(int32)||
|reward||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## create


**接口地址**:`/api/seek/create`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "seekId": 0,
  "title": "",
  "category": 0,
  "source": 0,
  "standard": 0,
  "imdb": "",
  "douban": "",
  "dmmCode": "",
  "reward": 0,
  "intro": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekEditForm|SeekEditForm|body|true|SeekEditForm|SeekEditForm|
|&emsp;&emsp;seekId|||false|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;category|||true|integer(int32)||
|&emsp;&emsp;source|||false|integer(int32)||
|&emsp;&emsp;standard|||false|integer(int32)||
|&emsp;&emsp;imdb|||false|string||
|&emsp;&emsp;douban|||false|string||
|&emsp;&emsp;dmmCode|||false|string||
|&emsp;&emsp;reward|||true|integer(int32)||
|&emsp;&emsp;intro|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## detail


**接口地址**:`/api/seek/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekId||query|true|integer(int32)||
|origin||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## edit


**接口地址**:`/api/seek/edit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "seekId": 0,
  "title": "",
  "category": 0,
  "source": 0,
  "standard": 0,
  "imdb": "",
  "douban": "",
  "dmmCode": "",
  "reward": 0,
  "intro": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekEditForm|SeekEditForm|body|true|SeekEditForm|SeekEditForm|
|&emsp;&emsp;seekId|||false|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;category|||true|integer(int32)||
|&emsp;&emsp;source|||false|integer(int32)||
|&emsp;&emsp;standard|||false|integer(int32)||
|&emsp;&emsp;imdb|||false|string||
|&emsp;&emsp;douban|||false|string||
|&emsp;&emsp;dmmCode|||false|string||
|&emsp;&emsp;reward|||true|integer(int32)||
|&emsp;&emsp;intro|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## recovery


**接口地址**:`/api/seek/recovery`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekId||query|true|integer(int32)||
|reason||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## search


**接口地址**:`/api/seek/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "sortField": "",
  "sortDirection": "",
  "cid": 0,
  "take": true
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekSearchFrom|SeekSearchFrom|body|true|SeekSearchFrom|SeekSearchFrom|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;sortField|||false|string||
|&emsp;&emsp;sortDirection|可用值:ASC,DESC||false|string||
|&emsp;&emsp;cid|||false|integer(int32)||
|&emsp;&emsp;take|||false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## submit


**接口地址**:`/api/seek/submit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekId||query|true|integer(int32)||
|torrentId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## take


**接口地址**:`/api/seek/take`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|seekId||query|true|integer(int32)||
|takeIds||query|true|array|integer|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 實驗室


## funcState


**接口地址**:`/api/laboratory/funcState`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## telegramGetBindToken


**接口地址**:`/api/laboratory/telegram/getBindToken`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## telegramUnBind


**接口地址**:`/api/laboratory/telegram/unBind`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## tiggerFunc


**接口地址**:`/api/laboratory/tiggerFunc`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 首頁趣味盒


## detail


**接口地址**:`/api/fun/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## edit


**接口地址**:`/api/fun/edit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "funid": 0,
  "title": "",
  "body": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|funForm|FunForm|body|true|FunForm|FunForm|
|&emsp;&emsp;funid|||false|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;body|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## first


**接口地址**:`/api/fun/first`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## post


**接口地址**:`/api/fun/post`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "funid": 0,
  "title": "",
  "body": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|funForm|FunForm|body|true|FunForm|FunForm|
|&emsp;&emsp;funid|||false|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;body|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## vote


**接口地址**:`/api/fun/vote`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|funid||query|true|integer(int32)||
|opinion||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 首頁投票


## first


**接口地址**:`/api/poll/first`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## vote


**接口地址**:`/api/poll/vote`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|pollid||query|true|integer(int32)||
|selection||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 首頁新聞


## list


**接口地址**:`/api/news/list`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 系統


## banlogs


**接口地址**:`/api/system/banlogs`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|SearchForm|SearchForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## countryList


**接口地址**:`/api/system/countryList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getConf


**接口地址**:`/api/system/getConf`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|items||query|true|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## hello


**接口地址**:`/api/system/hello`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## ip


**接口地址**:`/api/system/ip`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ip||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## ip


**接口地址**:`/api/system/ip`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ip||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## ipASN


**接口地址**:`/api/system/ipASN`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ip||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## ipASN


**接口地址**:`/api/system/ipASN`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ip||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## ip


**接口地址**:`/api/system/ips`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ips||query|true|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## ip


**接口地址**:`/api/system/ips`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ips||query|true|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## iscn


**接口地址**:`/api/system/iscn`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ip||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## iscn


**接口地址**:`/api/system/iscn`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|ip||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## langs


**接口地址**:`/api/system/langs`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## news


**接口地址**:`/api/system/news`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## promotionRules


**接口地址**:`/api/system/promotion/rules`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## staff


**接口地址**:`/api/system/staff`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## state


**接口地址**:`/api/system/state`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sysConf


**接口地址**:`/api/system/sysConf`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## top


**接口地址**:`/api/system/top`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|type||query|true|string||
|subTypes||query|true|array|string|
|pageSize||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 邀請


## getUserInviteHistory


**接口地址**:`/api/invite/getUserInviteHistory`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getUserInviteInfo


**接口地址**:`/api/invite/getUserInviteInfo`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getUserInviteSendHistory


**接口地址**:`/api/invite/getUserInviteSendHistory`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sendInvite


**接口地址**:`/api/invite/sendInvite`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|email||query|true|string||
|sms||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 用戶


## base


**接口地址**:`/api/member/base`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## bases


**接口地址**:`/api/member/bases`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "ids": []
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|userBaseBatchQuery|UserBaseBatchQuery|body|true|UserBaseBatchQuery|UserBaseBatchQuery|
|&emsp;&emsp;ids|||true|array|integer(int32)|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## bindOTP


**接口地址**:`/api/member/bindOTP`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|code||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## checkInviteCode


**接口地址**:`/api/member/checkInviteCode`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|inviteCode||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## forgotPwd


**接口地址**:`/api/member/forgotPwd`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|email||query|true|string||
|captchaId||query|true|string||
|captcha||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## forgotPwdTow


**接口地址**:`/api/member/forgotPwdTow`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|email||query|true|string||
|captcha||query|true|string||
|newPassword||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## genOTPUrl


**接口地址**:`/api/member/genOTPUrl`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getCrimeRecords


**接口地址**:`/api/member/getCrimeRecords`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getSessionList


**接口地址**:`/api/member/getSessionList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## getUserTorrentList


**接口地址**:`/api/member/getUserTorrentList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "userid": 0,
  "type": "",
  "official": true
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|userTorrentSearch|UserTorrentSearch|body|true|UserTorrentSearch|UserTorrentSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;userid|||true|integer(int32)||
|&emsp;&emsp;type|可用值:UPLOADED,SEEDING,LEECHING,COMPLETED,INCOMPLETE,SEEK||true|string||
|&emsp;&emsp;official|||false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## logout


**接口地址**:`/api/member/logout`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## profile


**接口地址**:`/api/member/profile`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## queryUserLoginHistory


**接口地址**:`/api/member/queryUserLoginHistory`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|userId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## register


**接口地址**:`/api/member/register`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|from||query|true|RegisterForm|RegisterForm|
|&emsp;&emsp;email|||false|string||
|&emsp;&emsp;username|||true|string||
|&emsp;&emsp;password|||true|string||
|&emsp;&emsp;country|||false|integer(int32)||
|&emsp;&emsp;gender|可用值:MALE,FEMALE,OTHER||true|string||
|&emsp;&emsp;inviteCode|||false|string||
|&emsp;&emsp;has13YearsOld|||true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## revokeSession


**接口地址**:`/api/member/revokeSession`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sendEmailCode


**接口地址**:`/api/member/sendEmailCode`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|zone|可用值:AccountVerifiy,ForgotPwd,TwoStepSuspend,Login,DONATE_USER,ChangeEmail,Close2FA|query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sendEmailVerifyCode


**接口地址**:`/api/member/sendEmailVerifyCode`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|email||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sendLoginEmailVerifyCode


**接口地址**:`/api/member/sendLoginEmailVerifyCode`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|username||query|true|string||
|email||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sendPasskey


**接口地址**:`/api/member/sendPasskey`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sysRoleList


**接口地址**:`/api/member/sysRoleList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## unbindOTP


**接口地址**:`/api/member/unbindOTP`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|code||query|true|string||
|emailCode||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## updateLastBrowse


**接口地址**:`/api/member/updateLastBrowse`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## updateProfile


**接口地址**:`/api/member/updateProfile`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "parked": true,
  "gender": "",
  "country": 0,
  "anonymous": true,
  "acceptpms": "",
  "commentpm": true,
  "deletepms": true,
  "magicgivingpm": true,
  "savepms": true,
  "avatarUrl": "",
  "info": "",
  "isp": 0,
  "downloadSpeed": 0,
  "uploadSpeed": 0,
  "config": {
    "trackerDomain": "",
    "downloadDomain": "",
    "rssDomain": "",
    "blockCategories": [],
    "hideFun": true,
    "showThumbnail": true,
    "timeType": "",
    "anonymous": true,
    "trackerDisableSeedbox": true
  }
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|memberFrofileSelfUpdateForm|MemberFrofileSelfUpdateForm|body|true|MemberFrofileSelfUpdateForm|MemberFrofileSelfUpdateForm|
|&emsp;&emsp;parked|||true|boolean||
|&emsp;&emsp;gender|可用值:MALE,FEMALE,OTHER||false|string||
|&emsp;&emsp;country|||false|integer(int32)||
|&emsp;&emsp;anonymous|||false|boolean||
|&emsp;&emsp;acceptpms|可用值:yes,no,friends||false|string||
|&emsp;&emsp;commentpm|||false|boolean||
|&emsp;&emsp;deletepms|||false|boolean||
|&emsp;&emsp;magicgivingpm|||false|boolean||
|&emsp;&emsp;savepms|||false|boolean||
|&emsp;&emsp;avatarUrl|||false|string||
|&emsp;&emsp;info|||false|string||
|&emsp;&emsp;isp|||false|integer(int32)||
|&emsp;&emsp;downloadSpeed|||false|integer(int32)||
|&emsp;&emsp;uploadSpeed|||false|integer(int32)||
|&emsp;&emsp;config|||false|MemberConfigVo|MemberConfigVo|
|&emsp;&emsp;&emsp;&emsp;trackerDomain|||false|string||
|&emsp;&emsp;&emsp;&emsp;downloadDomain|||false|string||
|&emsp;&emsp;&emsp;&emsp;rssDomain|||false|string||
|&emsp;&emsp;&emsp;&emsp;blockCategories|||false|array|integer|
|&emsp;&emsp;&emsp;&emsp;hideFun|||false|boolean||
|&emsp;&emsp;&emsp;&emsp;showThumbnail|||false|boolean||
|&emsp;&emsp;&emsp;&emsp;timeType|可用值:timeAdded,timeAlive||false|string||
|&emsp;&emsp;&emsp;&emsp;anonymous|||false|boolean||
|&emsp;&emsp;&emsp;&emsp;trackerDisableSeedbox|||false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## updateSecurity


**接口地址**:`/api/member/updateSecurity`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|securityUpdateForm||query|true|MemberSecurityUpdateForm|MemberSecurityUpdateForm|
|&emsp;&emsp;resetpasskey|||false|boolean||
|&emsp;&emsp;chpassword|||false|string||
|&emsp;&emsp;oldPwd|||false|string||
|&emsp;&emsp;privacy|可用值:STRONG,NORMAL,LOW||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## verifyAccount


**接口地址**:`/api/member/verifyAccount`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|email||query|true|string||
|code||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## verifyAccountByUser


**接口地址**:`/api/member/verifyAccountByUser`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 友情連結


## apply


**接口地址**:`/api/links/apply`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "linkName": "",
  "url": "",
  "title": "",
  "admin": "",
  "email": "",
  "reason": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|linksApplyForm|LinksApplyForm|body|true|LinksApplyForm|LinksApplyForm|
|&emsp;&emsp;linkName|||true|string||
|&emsp;&emsp;url|||true|string||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;admin|||true|string||
|&emsp;&emsp;email|||true|string||
|&emsp;&emsp;reason|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## view


**接口地址**:`/api/links/view`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 站內信


## boxList


**接口地址**:`/api/msg/boxList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## delBox


**接口地址**:`/api/msg/delBox`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|boxId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## delete


**接口地址**:`/api/msg/delete`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|msgIds||query|true|array|integer|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## forward


**接口地址**:`/api/msg/forward`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|MessageForm|MessageForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;refId|||false|integer(int32)||
|&emsp;&emsp;to|||true|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;content|||true|string||
|&emsp;&emsp;options|||false|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## markRead


**接口地址**:`/api/msg/markRead`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|msgIds||query|true|array|integer|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## move


**接口地址**:`/api/msg/move`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|boxId||query|true|integer(int32)||
|msgIds||query|true|array|integer|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## newBox


**接口地址**:`/api/msg/newBox`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|boxName||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## notifyStatistic


**接口地址**:`/api/msg/notify/statistic`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## read


**接口地址**:`/api/msg/read`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|msgId||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## reply


**接口地址**:`/api/msg/reply`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|MessageForm|MessageForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;refId|||false|integer(int32)||
|&emsp;&emsp;to|||true|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;content|||true|string||
|&emsp;&emsp;options|||false|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## search


**接口地址**:`/api/msg/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|search||query|true|MessageSearch|MessageSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;box|||false|integer(int32)||
|&emsp;&emsp;place|||false|string||
|&emsp;&emsp;unread|||false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## send


**接口地址**:`/api/msg/send`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|MessageForm|MessageForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;refId|||false|integer(int32)||
|&emsp;&emsp;to|||true|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;content|||true|string||
|&emsp;&emsp;options|||false|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## statistic


**接口地址**:`/api/msg/statistic`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## updateBoxName


**接口地址**:`/api/msg/updateBoxName`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|boxId||query|true|integer(int32)||
|boxName||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 種子


## doubanElessarInfo


**接口地址**:`/api/media/douban/elessar`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|code||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## doubanInfo


**接口地址**:`/api/media/douban/info`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|code||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## imdbInfo


**接口地址**:`/api/media/imdb/info`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|code||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## audioCodecList


**接口地址**:`/api/torrent/audioCodecList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## categoryList


**接口地址**:`/api/torrent/categoryList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## chearCollection


**接口地址**:`/api/torrent/chearCollection`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collection


**接口地址**:`/api/torrent/collection`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||
|make||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## createOredit


**接口地址**:`/api/torrent/createOredit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|TorrentUploadForm|TorrentUploadForm|
|&emsp;&emsp;torrent|||false|integer(int32)||
|&emsp;&emsp;offer|||false|integer(int32)||
|&emsp;&emsp;name|||true|string||
|&emsp;&emsp;smallDescr|||false|string||
|&emsp;&emsp;descr|||true|string||
|&emsp;&emsp;category|||true|integer(int32)||
|&emsp;&emsp;source|||false|integer(int32)||
|&emsp;&emsp;medium|||false|integer(int32)||
|&emsp;&emsp;standard|||false|integer(int32)||
|&emsp;&emsp;videoCodec|||false|integer(int32)||
|&emsp;&emsp;audioCodec|||false|integer(int32)||
|&emsp;&emsp;team|||false|integer(int32)||
|&emsp;&emsp;processing|||false|integer(int32)||
|&emsp;&emsp;countries|||false|string||
|&emsp;&emsp;imdb|||false|string||
|&emsp;&emsp;douban|||false|string||
|&emsp;&emsp;dmmCode|||false|string||
|&emsp;&emsp;cids|||false|string||
|&emsp;&emsp;aids|||false|string||
|&emsp;&emsp;anonymous|||true|boolean||
|&emsp;&emsp;labels|||false|integer(int32)||
|&emsp;&emsp;tags|||false|string||
|&emsp;&emsp;file|||false|string(binary)||
|&emsp;&emsp;nfo|||false|string(binary)||
|&emsp;&emsp;mediainfo|||false|string||
|&emsp;&emsp;labelsNew|||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## detail


**接口地址**:`/api/torrent/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||
|origin||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## doubanInfo


**接口地址**:`/api/torrent/doubanInfo`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|code||query|true|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## files


**接口地址**:`/api/torrent/files`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## genDlToken


**接口地址**:`/api/torrent/genDlToken`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integffer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## imdbInfo


**接口地址**:`/api/torrent/imdbInfo`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|code||query|true|array|string|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## mediaInfo


**接口地址**:`/api/torrent/mediaInfo`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## mediumList


**接口地址**:`/api/torrent/mediumList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## peers


**接口地址**:`/api/torrent/peers`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## processingList


**接口地址**:`/api/torrent/processingList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## queryTorrentTrackerHistory


**接口地址**:`/api/torrent/queryTorrentTrackerHistory`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|from||query|true|TorrentTrackerUserHistoryFom|TorrentTrackerUserHistoryFom|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;torrent|||true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## requestReseed


**接口地址**:`/api/torrent/requestReseed`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## rewardStatus


**接口地址**:`/api/torrent/rewardStatus`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sayThank


**接口地址**:`/api/torrent/sayThank`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## search


**接口地址**:`/api/torrent/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "categories": [],
  "imdb": "",
  "douban": "",
  "dmmCode": "",
  "author": 0,
  "sources": [],
  "mediums": [],
  "standards": [],
  "videoCodecs": [],
  "audioCodecs": [],
  "teams": [],
  "processings": [],
  "countries": [],
  "labels": 0,
  "uploadDateStart": "",
  "uploadDateEnd": "",
  "visible": 0,
  "onlyFav": true,
  "offer": true,
  "hot": true,
  "mode": "",
  "dmmField": "",
  "dmmKeyword": "",
  "discount": "",
  "labelsNew": [],
  "sortField": "",
  "sortDirection": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|torrentSearch|TorrentSearch|body|true|TorrentSearch|TorrentSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;categories|||false|array|integer(int32)|
|&emsp;&emsp;imdb|||false|string||
|&emsp;&emsp;douban|||false|string||
|&emsp;&emsp;dmmCode|||false|string||
|&emsp;&emsp;author|||false|integer(int32)||
|&emsp;&emsp;sources|||false|array|integer(int32)|
|&emsp;&emsp;mediums|||false|array|integer(int32)|
|&emsp;&emsp;standards|||false|array|integer(int32)|
|&emsp;&emsp;videoCodecs|||false|array|integer(int32)|
|&emsp;&emsp;audioCodecs|||false|array|integer(int32)|
|&emsp;&emsp;teams|||false|array|integer(int32)|
|&emsp;&emsp;processings|||false|array|integer(int32)|
|&emsp;&emsp;countries|||false|array|integer(int32)|
|&emsp;&emsp;labels|||false|integer(int32)||
|&emsp;&emsp;uploadDateStart|||false|string(date-time)||
|&emsp;&emsp;uploadDateEnd|||false|string(date-time)||
|&emsp;&emsp;visible|||false|integer(int32)||
|&emsp;&emsp;onlyFav|||false|boolean||
|&emsp;&emsp;offer|||false|boolean||
|&emsp;&emsp;hot|||false|boolean||
|&emsp;&emsp;mode|可用值:normal,adult,movie,music,tvshow,waterfall,rss,rankings||false|string||
|&emsp;&emsp;dmmField|可用值:kid,director,series,maker,label,product_number||false|string||
|&emsp;&emsp;dmmKeyword|||false|string||
|&emsp;&emsp;discount|可用值:NORMAL,PERCENT_70,PERCENT_50,FREE,_2X_FREE,_2X,_2X_PERCENT_50||false|string||
|&emsp;&emsp;labelsNew|||false|array|string|
|&emsp;&emsp;sortField|可用值:CREATED_DATE,SIZE,SEEDERS,LEECHERS,TIMES_COMPLETED,NAME||false|string||
|&emsp;&emsp;sortDirection|可用值:ASC,DESC||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sendReward


**接口地址**:`/api/torrent/sendReward`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||
|reward||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## sourceList


**接口地址**:`/api/torrent/sourceList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## standardList


**接口地址**:`/api/torrent/standardList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## teamList


**接口地址**:`/api/torrent/teamList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## thanksStatus


**接口地址**:`/api/torrent/thanksStatus`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## videoCodecList


**接口地址**:`/api/torrent/videoCodecList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## viewHits


**接口地址**:`/api/torrent/viewHits`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 種子候選


## config


**接口地址**:`/api/offer/config`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## vote


**接口地址**:`/api/offer/vote`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|torrentId||query|true|integer(int32)||
|vote|可用值:yeah,against|query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 专辑


## albumCollect


**接口地址**:`/api/album/albumCollect`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumId||query|true|integer(int32)||
|mark||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumCreate


**接口地址**:`/api/album/albumCreate`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "id": 0,
  "title": "",
  "intro": "",
  "adult": true
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumForm|AlbumForm|body|true|AlbumForm|AlbumForm|
|&emsp;&emsp;id|||false|integer(int32)||
|&emsp;&emsp;title|||false|string||
|&emsp;&emsp;intro|||false|string||
|&emsp;&emsp;adult|||true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumEdit


**接口地址**:`/api/album/albumEdit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "id": 0,
  "title": "",
  "intro": "",
  "adult": true
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumForm|AlbumForm|body|true|AlbumForm|AlbumForm|
|&emsp;&emsp;id|||false|integer(int32)||
|&emsp;&emsp;title|||false|string||
|&emsp;&emsp;intro|||false|string||
|&emsp;&emsp;adult|||true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumHistoryRanking


**接口地址**:`/api/album/albumHistoryRanking`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|year||query|true|integer(int32)||
|week||query|true|integer(int32)||
|adult||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumReward


**接口地址**:`/api/album/albumReward`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumId||query|true|integer(int32)||
|reward||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumSearch


**接口地址**:`/api/album/albumSearch`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "adult": true,
  "status": "",
  "authorId": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumSearch|AlbumSearch|body|true|AlbumSearch|AlbumSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;adult|||false|boolean||
|&emsp;&emsp;status|可用值:NORMAL,PENDING,FREEZE||true|string||
|&emsp;&emsp;authorId|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumTorrentAuditing


**接口地址**:`/api/album/albumTorrentAuditing`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "albumId": 0,
  "tids": [],
  "result": "",
  "reason": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumTorrentAuditingForm|AlbumTorrentAuditingForm|body|true|AlbumTorrentAuditingForm|AlbumTorrentAuditingForm|
|&emsp;&emsp;albumId|||true|integer(int32)||
|&emsp;&emsp;tids|||true|array|integer(int32)|
|&emsp;&emsp;result|可用值:NORMAL,PENDING,FREEZE||true|string||
|&emsp;&emsp;reason|||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumTorrentJoin


**接口地址**:`/api/album/albumTorrentJoin`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "albumId": 0,
  "tid": 0,
  "slogan": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumTorrentJoinForm|AlbumTorrentJoinForm|body|true|AlbumTorrentJoinForm|AlbumTorrentJoinForm|
|&emsp;&emsp;albumId|||true|integer(int32)||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;slogan|||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumTorrentRemove


**接口地址**:`/api/album/albumTorrentRemove`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "tids": [],
  "albumId": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumTorrentRemoveForm|AlbumTorrentRemoveForm|body|true|AlbumTorrentRemoveForm|AlbumTorrentRemoveForm|
|&emsp;&emsp;tids|||true|array|integer(int32)|
|&emsp;&emsp;albumId|||true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumTorrentSearch


**接口地址**:`/api/album/albumTorrentSearch`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "albumId": 0,
  "status": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumTorrentSearch|AlbumTorrentSearch|body|true|AlbumTorrentSearch|AlbumTorrentSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;albumId|||true|integer(int32)||
|&emsp;&emsp;status|可用值:NORMAL,PENDING,FREEZE||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## albumEdit


**接口地址**:`/api/album/albumUpdateJoinType`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|albumId||query|true|integer(int32)||
|joinType|可用值:Self,Public|query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## myAlbumList


**接口地址**:`/api/album/myAlbumList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# 字幕


## dl


**接口地址**:`/api/subtitle/dl`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|400|Bad Request||


**响应参数**:


暂无


**响应示例**:
```javascript

```


## dlV2


**接口地址**:`/api/subtitle/dlV2`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|credential||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|400|Bad Request||


**响应参数**:


暂无


**响应示例**:
```javascript

```


## genlink


**接口地址**:`/api/subtitle/genlink`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## langs


**接口地址**:`/api/subtitle/langs`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## list


**接口地址**:`/api/subtitle/list`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## search


**接口地址**:`/api/subtitle/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "langId": 0
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|subtitleSearch|SubtitleSearch|body|true|SubtitleSearch|SubtitleSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;langId|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## upload


**接口地址**:`/api/subtitle/upload`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|TorrenSubtitleForm|TorrenSubtitleForm|
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;torrent|||true|integer(int32)||
|&emsp;&emsp;anonymous|||true|boolean||
|&emsp;&emsp;lang|||true|integer(int32)||
|&emsp;&emsp;file|||true|string(binary)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# dmm


## collagesAddToList


**接口地址**:`/api/dmm/collages/addToList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|cid||query|true|integer(int32)||
|tid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesClearFav


**接口地址**:`/api/dmm/collages/clearFav`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesCreate


**接口地址**:`/api/dmm/collages/createOrEdit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|DmmCollagesForm|DmmCollagesForm|
|&emsp;&emsp;cid|||false|integer(int32)||
|&emsp;&emsp;typeId|||true|integer(int32)||
|&emsp;&emsp;title|||true|string||
|&emsp;&emsp;note|||true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesDelete


**接口地址**:`/api/dmm/collages/delete`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|cid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesDetail


**接口地址**:`/api/dmm/collages/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|cid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesFav


**接口地址**:`/api/dmm/collages/fav`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|cid||query|true|integer(int32)||
|subscribe||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesFetchList


**接口地址**:`/api/dmm/collages/fetchList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||
|page||query|true|integer(int32)||
|pageSize||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesSearch


**接口地址**:`/api/dmm/collages/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "typeId": 0,
  "action": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|dmmCollageSearch|DmmCollageSearch|body|true|DmmCollageSearch|DmmCollageSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;typeId|||false|integer(int32)||
|&emsp;&emsp;action|可用值:my,fav||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## collagesViewHits


**接口地址**:`/api/dmm/collages/viewHits`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## dmmInfo


**接口地址**:`/api/dmm/dmmInfo`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|dmmCode||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## dmmSeearch


**接口地址**:`/api/dmm/dmmSeearch`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|keyword||query|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseAddToList


**接口地址**:`/api/dmm/showcase/addToList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|aid||query|true|integer(int32)||
|tid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseClearFav


**接口地址**:`/api/dmm/showcase/clearFav`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseCreate


**接口地址**:`/api/dmm/showcase/create`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|DmmShowcaseForm|DmmShowcaseForm|
|&emsp;&emsp;aid|||false|integer(int32)||
|&emsp;&emsp;cntitle|||true|string||
|&emsp;&emsp;entitle|||true|string||
|&emsp;&emsp;note|||true|string||
|&emsp;&emsp;pic|||false|string(binary)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseDelete


**接口地址**:`/api/dmm/showcase/delete`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|aid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseDetail


**接口地址**:`/api/dmm/showcase/detail`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|aid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseEdit


**接口地址**:`/api/dmm/showcase/edit`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|form||query|true|DmmShowcaseForm|DmmShowcaseForm|
|&emsp;&emsp;aid|||false|integer(int32)||
|&emsp;&emsp;cntitle|||true|string||
|&emsp;&emsp;entitle|||true|string||
|&emsp;&emsp;note|||true|string||
|&emsp;&emsp;pic|||false|string(binary)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseFav


**接口地址**:`/api/dmm/showcase/fav`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|aid||query|true|integer(int32)||
|subscribe||query|true|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseFetchList


**接口地址**:`/api/dmm/showcase/fetchList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||
|page||query|true|integer(int32)||
|pageSize||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseSearch


**接口地址**:`/api/dmm/showcase/search`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "searchType": "",
  "action": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|dmmShowCaseSearch|DmmShowCaseSearch|body|true|DmmShowCaseSearch|DmmShowCaseSearch|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;searchType|可用值:ens,cns||false|string||
|&emsp;&emsp;action|可用值:my,fav||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## showcaseViewHits


**接口地址**:`/api/dmm/showcase/viewHits`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# RSS


## dlv2


**接口地址**:`/api/rss/dlv2`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|&emsp;&emsp;uid|||true|integer(int32)||
|&emsp;&emsp;tid|||true|integer(int32)||
|&emsp;&emsp;t|||true|integer(int32)||
|&emsp;&emsp;sign|||true|string||
|&emsp;&emsp;useHttps|||false|boolean||
|&emsp;&emsp;type|||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|400|Bad Request||
|404|Not Found||


**响应参数**:


暂无


**响应示例**:
```javascript

```


## fetch


**接口地址**:`/api/rss/fetch`


**请求方式**:`GET`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`application/xml;charset=utf-8,*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|params||query|true|RssFetchForm|RssFetchForm|
|&emsp;&emsp;uid|||true|integer(int32)||
|&emsp;&emsp;t|||true|integer(int64)||
|&emsp;&emsp;sign|||true|string||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;categories|||false|string||
|&emsp;&emsp;sources|||false|string||
|&emsp;&emsp;mediums|||false|string||
|&emsp;&emsp;standards|||false|string||
|&emsp;&emsp;videoCodecs|||false|string||
|&emsp;&emsp;audioCodecs|||false|string||
|&emsp;&emsp;teams|||false|string||
|&emsp;&emsp;processings|||false|string||
|&emsp;&emsp;countries|||false|string||
|&emsp;&emsp;labels|||false|integer(int32)||
|&emsp;&emsp;tkeys|||false|string||
|&emsp;&emsp;imdb|||false|string||
|&emsp;&emsp;douban|||false|string||
|&emsp;&emsp;dmmCode|||false|string||
|&emsp;&emsp;onlyFav|||false|boolean||
|&emsp;&emsp;dl|||false|boolean||
|&emsp;&emsp;pageSize|||false|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|400|Bad Request||
|404|Not Found||


**响应参数**:


暂无


**响应示例**:
```javascript

```


## genlink


**接口地址**:`/api/rss/genlink`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "pageNumber": 0,
  "pageSize": 0,
  "lastId": 0,
  "keyword": "",
  "tkeys": [],
  "categories": [],
  "sources": [],
  "mediums": [],
  "standards": [],
  "videoCodecs": [],
  "audioCodecs": [],
  "teams": [],
  "processings": [],
  "countries": [],
  "labels": 0,
  "imdb": "",
  "douban": "",
  "dmmCode": "",
  "onlyFav": true
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|rssGenForm|RssGenForm|body|true|RssGenForm|RssGenForm|
|&emsp;&emsp;pageNumber|||false|integer(int32)||
|&emsp;&emsp;pageSize|||false|integer(int32)||
|&emsp;&emsp;lastId|||false|integer(int32)||
|&emsp;&emsp;keyword|||false|string||
|&emsp;&emsp;tkeys|||false|array|string|
|&emsp;&emsp;categories|||false|array|integer(int32)|
|&emsp;&emsp;sources|||false|array|integer(int32)|
|&emsp;&emsp;mediums|||false|array|integer(int32)|
|&emsp;&emsp;standards|||false|array|integer(int32)|
|&emsp;&emsp;videoCodecs|||false|array|integer(int32)|
|&emsp;&emsp;audioCodecs|||false|array|integer(int32)|
|&emsp;&emsp;teams|||false|array|integer(int32)|
|&emsp;&emsp;processings|||false|array|integer(int32)|
|&emsp;&emsp;countries|||false|array|integer(int32)|
|&emsp;&emsp;labels|||false|integer(int32)||
|&emsp;&emsp;imdb|||false|string||
|&emsp;&emsp;douban|||false|string||
|&emsp;&emsp;dmmCode|||false|string||
|&emsp;&emsp;onlyFav|||false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


# tracker


## clientList


**接口地址**:`/api/tracker/clientList`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## flush


**接口地址**:`/api/tracker/flush`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## mybonus


**接口地址**:`/api/tracker/mybonus`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## myStatistics


**接口地址**:`/api/tracker/myPeerStatistics`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|uid||query|true|integer(int32)||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## myPeerStatus


**接口地址**:`/api/tracker/myPeerStatus`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


暂无


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```


## queryHistory


**接口地址**:`/api/tracker/queryHistory`


**请求方式**:`POST`


**请求数据类型**:`application/x-www-form-urlencoded,application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```javascript
{
  "tids": []
}
```


**请求参数**:


| 参数名称 | 参数说明 | 请求类型    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|trackerHistoryQuery|TrackerHistoryQuery|body|true|TrackerHistoryQuery|TrackerHistoryQuery|
|&emsp;&emsp;tids|||false|array|integer(int32)|


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Result|
|400|Bad Request||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|message||string||
|data||object||
|code||integer(int32)|integer(int32)|


**响应示例**:
```javascript
{
	"message": "",
	"data": {},
	"code": 0
}
```