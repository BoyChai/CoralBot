# 动作

## 发送私聊消息

**方法:**

```
func (h Handle) SendPrivateMsg(m Msg) (map[string]interface{}, error) 
```

**Msg结构体:**

| 字段名       | 数据类型 | 默认值  | 说明                                                         |
| ------------ | -------- | ------- | ------------------------------------------------------------ |
| `UserId`     | int64    | -       | 对方 QQ 号                                                   |
| `GroupId`    | int64    | -       | 主动发起临时会话群号(机器人本身必须是管理员/群主)            |
| `Message`    | message  | -       | 要发送的内容                                                 |
| `AutoEscape` | boolean  | `false` | 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 `message` 字段是字符串时有效 |

**响应数据:**

| 字段名       | 数据类型 | 说明    |
| ------------ | -------- | ------- |
| `message_id` | int32    | 消息 ID |

## 发送群聊消息

**方法:**

```
func (h Handle) SendGroupMsg(m Msg) (map[string]interface{}, error)
```

**Msg结构体:**

| 字段名       | 数据类型 | 默认值  | 说明                                                         |
| ------------ | -------- | ------- | ------------------------------------------------------------ |
| `GroupId`    | int64    | -       | 群号                                                         |
| `Message`    | message  | -       | 要发送的内容                                                 |
| `AutoEscape` | boolean  | `false` | 消息内容是否作为纯文本发送 ( 即不解析 CQ 码) , 只在 `message` 字段是字符串时有效 |

**响应数据:**

| 字段名       | 数据类型 | 说明    |
| ------------ | -------- | ------- |
| `message_id` | int32    | 消息 ID |

## 发送消息

**方法:**

```
func (h Handle) SendMsg(m Msg) (map[string]interface{}, error)
```

**传参:**

| 字段名        | 数据类型 | 默认值  | 说明                                                         |
| ------------- | -------- | ------- | ------------------------------------------------------------ |
| `MessageType` | string   | -       | 消息类型, 支持 `private`、`group` , 分别对应私聊、群组, 如不传入, 则根据传入的 `*_id` 参数判断 |
| `UserId`      | int64    | -       | 对方 QQ 号 ( 消息类型为 `private` 时需要 )                   |
| `GroupId`     | int64    | -       | 群号 ( 消息类型为 `group` 时需要 )                           |
| `Message`     | message  | -       | 要发送的内容                                                 |
| `AutoEscape`  | boolean  | `false` | 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 `message` 字段是字符串时有效 |

**响应数据:**

| 字段名       | 数据类型 | 说明    |
| ------------ | -------- | ------- |
| `message_id` | int32    | 消息 ID |

## 撤回消息

**方法:**

```
func (h Handle) DeleteMsg(messageId int32) (map[string]interface{}, error) 
```

**传参:**

| 字段名      | 数据类型 | 默认值 | 说明    |
| ----------- | -------- | ------ | ------- |
| `messageId` | int32    | -      | 消息 ID |

**响应数据:**

> 提示
>
> 该 API 无响应数据

## 获取信息

**方法:**

```
func (h Handle) GetMsg(messageId int32) (map[string]interface{}, error)
```

**传参:**

| 字段        | 类型  | 说明   |
| ----------- | ----- | ------ |
| `messageId` | int32 | 消息id |

**响应数据:**

| 字段          | 类型    | 说明         |
| ------------- | ------- | ------------ |
| `message_id`  | int32   | 消息id       |
| `real_id`     | int32   | 消息真实id   |
| `sender`      | object  | 发送者       |
| `time`        | int32   | 发送时间     |
| `message`     | message | 消息内容     |
| `raw_message` | message | 原始消息内容 |

## 获取合并转发内容

**方法:**

```
func (h Handle) GetForwardMsg(messageId int32) (map[string]interface{}, error)
```

**传参:**

| 字段        | 类型   | 说明   |
| ----------- | ------ | ------ |
| `messageId` | string | 消息id |

**响应数据:**

```json
{
    "data": {
        "messages": [
            {
                "content": "合并转发1",
                "sender": {
                    "nickname": "发送者A",
                    "user_id": 10086
                },
                "time": 1595694374
            },
            {
                "content": "合并转发2[CQ:image,file=xxxx,url=xxxx]",
                "sender": {
                    "nickname": "发送者B",
                    "user_id": 10087
                },
                "time": 1595694393
            }
        ]
    },
    "retcode": 0,
    "status": "ok"
}
```

## 获取图片信息

**方法:**

```
func (h Handle) GetImage(file string) (map[string]interface{}, error)
```

**传参:**

| 字段   | 类型   | 说明           |
| ------ | ------ | -------------- |
| `file` | string | 图片缓存文件名 |

**响应数据:**

| 字段       | 类型   | 说明           |
| ---------- | ------ | -------------- |
| `size`     | int32  | 图片源文件大小 |
| `filename` | string | 图片文件原名   |
| `url`      | string | 图片下载地址   |

## 标记消息已读

**方法:**

```
func (h Handle) MarkMsgAsRead(messageId int32) (map[string]interface{}, error) 
```

**传参:**

| 字段        | 类型  | 说明   |
| ----------- | ----- | ------ |
| `messageId` | int32 | 消息id |

**响应数据:**

>提示
>
>该 API 无响应数据	

## 群组踢人

**方法:**

```
func (h Handle) SetGroupKick(groupId int64, userId int64, rejectAddRequest bool) (map[string]interface{}, error)
```

**传参:**

| 字段名     | 数据类型 | 默认值    | 说明                             |
| ---------- | -------- | --------- | -------------------------------- |
| `groupId`  | int64    | -         | 群号                             |
| `userId`   | int64    | -         | 要禁言的 QQ 号                   |
| `duration` | number   | `30 * 60` | 禁言时长, 单位秒, 0 表示取消禁言 |

**响应数据:**

>提示
>
>该 API 无响应数据

## 群组单人禁言

**方法:**

```
func (h Handle) SetGroupBan(groupId int64, userId int64, duration string) (map[string]interface{}, error)
```

**传参:**

| 字段名     | 数据类型 | 默认值    | 说明                             |
| ---------- | -------- | --------- | -------------------------------- |
| `group_id` | int64    | -         | 群号                             |
| `user_id`  | int64    | -         | 要禁言的 QQ 号                   |
| `duration` | number   | `30 * 60` | 禁言时长, 单位秒, 0 表示取消禁言 |

**响应数据:**

> 提示
>
> 该 API 无响应数据

## 群组匿名禁言

**方法:**

```
func (h Handle) SetGroupAnonymousBan(groupId int64, flag string, duration string) (map[string]interface{}, error)
```

**传参:**

| 字段名     | 数据类型 | 默认值    | 说明                                                        |
| ---------- | -------- | --------- | ----------------------------------------------------------- |
| `groupId`  | int64    | -         | 群号                                                        |
| `flag`     | string   | -         | 可选, 要禁言的匿名用户的 flag（需从群消息上报的数据中获得） |
| `duration` | number   | `30 * 60` | 禁言时长, 单位秒, 无法取消匿名用户禁言                      |

**响应数据:**

>提示
>
>该 API 无响应数据



## 群组全员禁言

**方法:**

```
func (h Handle) SetGroupWholeBan(groupId int64, enable bool) (map[string]interface{}, error)
```

**传参:**

| 字段名    | 数据类型 | 默认值 | 说明     |
| --------- | -------- | ------ | -------- |
| `groupId` | int64    | -      | 群号     |
| `enable`  | boolean  | `true` | 是否禁言 |

**响应数据:**

>提示
>
>该 API 无响应数据

## 群组设置管理员

**方法:**

```
func (h Handle) SetGroupAdmin(groupId int64, userId int64, enable bool) (map[string]interface{}, error)
```

**传参:**

| 字段名    | 数据类型 | 默认值 | 说明                      |
| --------- | -------- | ------ | ------------------------- |
| `groupId` | int64    | -      | 群号                      |
| `userId`  | int64    | -      | 要设置管理员的 QQ 号      |
| `enable`  | boolean  | `true` | true 为设置, false 为取消 |

**响应数据:**

> 提示
>
> 该 API 无响应数据

## 群组匿名设置

**方法:**

```
func (h Handle) SetGroupAnonymous(groupId int64, enable bool) (map[string]interface{}, error) 
```

**传参:**

| 字段名    | 数据类型 | 默认值 | 说明             |
| --------- | -------- | ------ | ---------------- |
| `groupId` | int64    | -      | 群号             |
| `enable`  | boolean  | `true` | 是否允许匿名聊天 |

**响应数据:**

> 提示
>
> 该 API 无响应数据

## 修改群聊名称

**方法:**

```
func (h Handle) SetGroupName(groupId int64, groupName string) (map[string]interface{}, error) 
```

**传参:**

**响应数据:**

| 字段名       | 数据类型 | 说明   |
| ------------ | -------- | ------ |
| `group_id`   | int64    | 群号   |
| `group_name` | string   | 新群名 |

> 提示
>
> 该 API 无响应数据

## 退出群聊

**方法:**

```
func (h Handle) SetGroupLeave(groupId int64, isDismiss bool) (map[string]interface{}, error)
```

**传参:**

| 字段名      | 数据类型 | 默认值  | 说明                                                     |
| ----------- | -------- | ------- | -------------------------------------------------------- |
| `groupId`   | int64    | -       | 群号                                                     |
| `isDismiss` | boolean  | `false` | 是否解散, 如果登录号是群主, 则仅在此项为 true 时能够解散 |

**响应数据:**

>  tip 提示
> 该 API 无响应数据

## 设置群组专属头衔

**方法:**

```
func (h Handle) SetGroupSpecialTitle(groupId int64, userId int64, specialTitle string, duration string) (map[string]interface{}, error)
```

**传参:**

| 字段名         | 数据类型 | 默认值 | 说明                                                         |
| -------------- | -------- | ------ | ------------------------------------------------------------ |
| `groupId`      | int64    | -      | 群号                                                         |
| `userId`       | int64    | -      | 要设置的 QQ 号                                               |
| `specialTitle` | string   | 空     | 专属头衔, 不填或空字符串表示删除专属头衔                     |
| `duration`     | number   | `-1`   | 专属头衔有效期, 单位秒, -1 表示永久, 不过此项似乎没有效果, 可能是只有某些特殊的时间长度有效, 有待测试 |

**响应数据:**

> tip 提示
> 该 API 无响应数据

## 群打卡

**方法:**

```
func (h Handle) SendGroupSign(groupId int64) (map[string]interface{}, error) 
```

**传参:**

| 字段名    | 数据类型 | 说明 |
| --------- | -------- | ---- |
| `groupId` | int64    | 群号 |

**响应数据:**

> tip 提示
> 该 API 无响应数据

## 处理加好友请求

**方法:**

```
func (h Handle) SetFriendAddRequest(flag string, approve bool, remark string) (map[string]interface{}, error) 
```

**传参:**

| 字段名    | 数据类型 | 默认值 | 说明                                      |
| --------- | -------- | ------ | ----------------------------------------- |
| `flag`    | string   | -      | 加好友请求的 flag（需从上报的数据中获得） |
| `approve` | boolean  | `true` | 是否同意请求                              |
| `remark`  | string   | 空     | 添加后的好友备注（仅在同意时有效）        |

**响应数据:**

> tip 提示
> 该 API 无响应数据

## 处理加群请求/邀请

**方法:**

```
func (h Handle) SetGroupAddRequest(flag string, subType string, approve bool, reason string) (map[string]interface{}, error) 
```

**传参:**

| 字段名     | 数据类型 | 默认值 | 说明                                                         |
| ---------- | -------- | ------ | ------------------------------------------------------------ |
| `flag`     | string   | -      | 加群请求的 flag（需从上报的数据中获得）                      |
| `sub_type` | string   | -      | `add` 或 `invite`, 请求类型（需要和上报消息中的 `sub_type` 字段相符） |
| `approve`  | boolean  | `true` | 是否同意请求／邀请                                           |
| `reason`   | string   | 空     | 拒绝理由（仅在拒绝时有效）                                   |

**响应数据:**

> tip 提示
> 该 API 无响应数据

## 获取登录号信息

**方法:**

```
func (h Handle) GetLoginInfo() (map[string]interface{}, error) 
```

**传参:**

> tip 提示
> 该 API 无需参数

**响应数据:**

| 字段名     | 数据类型 | 说明    |
| ---------- | -------- | ------- |
| `userId`   | int64    | QQ 号   |
| `nickname` | string   | QQ 昵称 |

## 获取企点账号信息

**方法:**

```
func (h Handle) QiDianGetAccountInfo() (map[string]interface{}, error) {
```

> tip 注意
> 该API只有企点协议可用

**传参:**

> tip 提示
> 该 API 无需参数

**响应数据:**

| 字段          | 类型   | 说明         |
| ------------- | ------ | ------------ |
| `master_id`   | int64  | 父账号ID     |
| `ext_name`    | string | 用户昵称     |
| `create_time` | int64  | 账号创建时间 |

## 设置账号资料

**方法:**



**传参:**

| 字段名          | 数据类型 | 默认值 | 说明     |
| --------------- | -------- | ------ | -------- |
| `nickname`      | string   | -      | 名称     |
| `company`       | string   | -      | 公司     |
| `email`         | string   | -      | 邮箱     |
| `college`       | string   | -      | 学校     |
| `personal_note` | string   | -      | 个人说明 |

**响应数据:**



## 获取陌生人信息

**方法:**



**传参:**

| 字段名     | 数据类型 | 默认值  | 说明                                                 |
| ---------- | -------- | ------- | ---------------------------------------------------- |
| `user_id`  | int64    | -       | QQ 号                                                |
| `no_cache` | boolean  | `false` | 是否不使用缓存（使用缓存可能更新不及时, 但响应更快） |

**响应数据:**

| 字段名       | 数据类型 | 说明                                  |
| ------------ | -------- | ------------------------------------- |
| `user_id`    | int64    | QQ 号                                 |
| `nickname`   | string   | 昵称                                  |
| `sex`        | string   | 性别, `male` 或 `female` 或 `unknown` |
| `age`        | int32    | 年龄                                  |
| `qid`        | string   | qid ID身份卡                          |
| `level`      | int32    | 等级                                  |
| `login_days` | int32    | 等级                                  |

## 获取好友列表

**方法:**



**传参:**

> tip 提示
> 该 API 无需参数

**响应数据:**

响应内容为 json 数组, 每个元素如下：

| 字段名     | 数据类型 | 说明   |
| ---------- | -------- | ------ |
| `user_id`  | int64    | QQ 号  |
| `nickname` | string   | 昵称   |
| `remark`   | string   | 备注名 |

## 获取单向好友列表

**方法:**



**传参:**

> tip 提示
> 该 API 无需参数

**响应数据:**

响应内容为 json 数组, 每个元素如下：

| 字段名     | 数据类型 | 说明  |
| ---------- | -------- | ----- |
| `user_id`  | int64    | QQ 号 |
| `nickname` | string   | 昵称  |
| `source`   | string   | 来源  |

## 删除好友

**方法:**



**传参:**

| 字段名     | 数据类型 | 默认值 | 说明       |
| ---------- | -------- | ------ | ---------- |
| `friendId` | int64    | -      | 好友 QQ 号 |

**响应数据:**

> tip 提示
> 该 API 无响应数据

## 获取群信息

**方法:**



**传参:**

| 字段名     | 数据类型 | 默认值  | 说明                                                 |
| ---------- | -------- | ------- | ---------------------------------------------------- |
| `group_id` | int64    | -       | 群号                                                 |
| `no_cache` | boolean  | `false` | 是否不使用缓存（使用缓存可能更新不及时, 但响应更快） |

**响应数据:**

> tip 提示
> 如果机器人尚未加入群, `group_create_time`, `group_level`, `max_member_count` 和 `member_count` 将会为0

## 获取群列表

**方法:**



**传参:**



**响应数据:**



## 获取群成员信息

**方法:**



**传参:**



**响应数据:**



## 获取群成员列表

**方法:**



**传参:**



**响应数据:**



## 获取群荣誉信息

**方法:**



**传参:**



**响应数据:**



## 获取Cookies

**方法:**



**传参:**



**响应数据:**



## 获取 CSRF Token

**方法:**



**传参:**



**响应数据:**



## 获取 QQ 相关接口凭证

**方法:**



**传参:**



**响应数据:**



## 获取语音

**方法:**



**传参:**



**响应数据:**



## 检查是否可以发送图片

**方法:**



**传参:**



**响应数据:**



## 检查是否可以发送语音

**方法:**



**传参:**



**响应数据:**



## 获取版本信息

**方法:**



**传参:**



**响应数据:**



## 重启 go-cqhttp

**方法:**



**传参:**



**响应数据:**



## 清理缓存

**方法:**



**传参:**



**响应数据:**



## 设置群头像

**方法:**



**传参:**



**响应数据:**



## OcrImage 图片 OCR

**方法:**



**传参:**



**响应数据:**



## 获取群系统消息

**方法:**



**传参:**



**响应数据:**



## 上传私聊文件

**方法:**



**传参:**



**响应数据:**



## 上传群文件

**方法:**



**传参:**



**响应数据:**



## 获取群文件系统信息

**方法:**



**传参:**



**响应数据:**



## 获取群根目录文件列表

**方法:**



**传参:**



**响应数据:**



## 获取群子目录文件列表

**方法:**



**传参:**



**响应数据:**



## 创建群文件文件夹

**方法:**



**传参:**



**响应数据:**



## 删除群文件夹

**方法:**



**传参:**



**响应数据:**



## 删除群文件

**方法:**



**传参:**



**响应数据:**



## 获取群文件资源链接

**方法:**



**传参:**



**响应数据:**



## 获取状态

**方法:**



**传参:**



**响应数据:**



## 获取群@全体成员剩余次数

**方法:**



**传参:**



**响应数据:**



## 发送群公告

**方法:**



**传参:**



**响应数据:**



## 获取群公告

**方法:**



**传参:**



**响应数据:**



## 重载事件过滤器

**方法:**



**传参:**



**响应数据:**



## 下载文件到缓存目录

**方法:**



**传参:**



**响应数据:**



## 获取当前账号在线客户端列表

**方法:**



**传参:**



**响应数据:**



## 获取群消息历史记录

**方法:**



**传参:**



**响应数据:**



## 设置精华消息

**方法:**



**传参:**



**响应数据:**



## 移除精华消息

**方法:**



**传参:**



**响应数据:**



## 获取精华消息列表

**方法:**



**传参:**



**响应数据:**



## 检查链接安全性

**方法:**



**传参:**



**响应数据:**



## 获取在线机型

**方法:**



**传参:**



**响应数据:**



## 设置在线机型

**方法:**



**传参:**



**响应数据:**



# PS

目前[CoralBot](https://github.com/BoyChai/CoralBot)对于[go-cqhttp](https://docs.go-cqhttp.org/)的**发送合并转发 ( 群 )**和**隐藏API**还不支持

