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

## 设置登录账号资料

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

> 提示
>
> 该 API 无需参数

**响应数据:**

响应内容为 json 数组, 每个元素和上面的 `GetGroupInfo` 接口相同。

### 

## 获取群成员信息

**方法:**



**传参:**

| 字段名     | 数据类型 | 默认值  | 说明                                                 |
| ---------- | -------- | ------- | ---------------------------------------------------- |
| `group_id` | int64    | -       | 群号                                                 |
| `user_id`  | int64    | -       | QQ 号                                                |
| `no_cache` | boolean  | `false` | 是否不使用缓存（使用缓存可能更新不及时, 但响应更快） |

**响应数据:**

| 字段名              | 数据类型 | 说明                                  |
| ------------------- | -------- | ------------------------------------- |
| `group_id`          | int64    | 群号                                  |
| `user_id`           | int64    | QQ 号                                 |
| `nickname`          | string   | 昵称                                  |
| `card`              | string   | 群名片／备注                          |
| `sex`               | string   | 性别, `male` 或 `female` 或 `unknown` |
| `age`               | int32    | 年龄                                  |
| `area`              | string   | 地区                                  |
| `join_time`         | int32    | 加群时间戳                            |
| `last_sent_time`    | int32    | 最后发言时间戳                        |
| `level`             | string   | 成员等级                              |
| `role`              | string   | 角色, `owner` 或 `admin` 或 `member`  |
| `unfriendly`        | boolean  | 是否不良记录成员                      |
| `title`             | string   | 专属头衔                              |
| `title_expire_time` | int64    | 专属头衔过期时间戳                    |
| `card_changeable`   | boolean  | 是否允许修改群名片                    |
| `shut_up_timestamp` | int64    | 禁言到期时间                          |

## 获取群成员列表

**方法:**



**传参:**

| 字段名     | 数据类型 | 默认值 | 说明 |
| ---------- | -------- | ------ | ---- |
| `group_id` | int64    | -      | 群号 |

**响应数据:**

响应内容为 json 数组, 每个元素的内容和上面的 `get_group_member_info` 接口相同, 但对于同一个群组的同一个成员, 获取列表时和获取单独的成员信息时, 某些字段可能有所不同, 例如 `area`、`title` 等字段在获取列表时无法获得, 具体应以单独的成员信息为准。

## 获取群荣誉信息

**方法:**



**传参:**

| 字段名     | 数据类型 | 默认值 | 说明                                                         |
| ---------- | -------- | ------ | ------------------------------------------------------------ |
| `group_id` | int64    | -      | 群号                                                         |
| `type`     | string   | -      | 要获取的群荣誉类型, 可传入 `talkative` `performer` `legend` `strong_newbie` `emotion` 以分别获取单个类型的群荣誉数据, 或传入 `all` 获取所有数据 |

**响应数据:**

| 字段名               | 数据类型 | 说明                                                       |
| -------------------- | -------- | ---------------------------------------------------------- |
| `group_id`           | int64    | 群号                                                       |
| `current_talkative`  | object   | 当前龙王, 仅 `type` 为 `talkative` 或 `all` 时有数据       |
| `talkative_list`     | array    | 历史龙王, 仅 `type` 为 `talkative` 或 `all` 时有数据       |
| `performer_list`     | array    | 群聊之火, 仅 `type` 为 `performer` 或 `all` 时有数据       |
| `legend_list`        | array    | 群聊炽焰, 仅 `type` 为 `legend` 或 `all` 时有数据          |
| `strong_newbie_list` | array    | 冒尖小春笋, 仅 `type` 为 `strong_newbie` 或 `all` 时有数据 |
| `emotion_list`       | array    | 快乐之源, 仅 `type` 为 `emotion` 或 `all` 时有数据         |

其中 `current_talkative` 字段的内容如下：

| 字段名      | 数据类型 | 说明     |
| ----------- | -------- | -------- |
| `user_id`   | int64    | QQ 号    |
| `nickname`  | string   | 昵称     |
| `avatar`    | string   | 头像 URL |
| `day_count` | int32    | 持续天数 |

其它各 `*_list` 的每个元素是一个 json 对象, 内容如下：

| 字段名        | 数据类型 | 说明     |
| ------------- | -------- | -------- |
| `user_id`     | int64    | QQ 号    |
| `nickname`    | string   | 昵称     |
| `avatar`      | string   | 头像 URL |
| `description` | string   | 荣誉描述 |

## 获取Cookies

**方法:**

> 注意
>
> 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 [提交 pr](https://github.com/Mrs4s/go-cqhttp/compare)

**传参:**

| 字段名   | 数据类型 | 默认值 | 说明                    |
| -------- | -------- | ------ | ----------------------- |
| `domain` | string   | 空     | 需要获取 cookies 的域名 |

**响应数据:**

| 字段名    | 数据类型 | 说明    |
| --------- | -------- | ------- |
| `cookies` | string   | Cookies |

## 获取 CSRF Token

**方法:**

> 注意
>
> 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 [提交 propen in new window](https://github.com/Mrs4s/go-cqhttp/compare)

**传参:**

提示

该 API 无需参数

**响应数据:**

| 字段名  | 数据类型 | 说明       |
| ------- | -------- | ---------- |
| `token` | int32    | CSRF Token |

## 获取 QQ 相关接口凭证

**方法:**

> 注意
>
> 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 [提交 propen in new window](https://github.com/Mrs4s/go-cqhttp/compare)
>
> 提示
>
> 即上面两个接口的合并

**传参:**

| 字段名   | 数据类型 | 默认值 | 说明                    |
| -------- | -------- | ------ | ----------------------- |
| `domain` | string   | 空     | 需要获取 cookies 的域名 |

**响应数据:**

| 字段名       | 数据类型 | 说明       |
| ------------ | -------- | ---------- |
| `cookies`    | string   | Cookies    |
| `csrf_token` | int32    | CSRF Token |

## 获取语音

**方法:**

> 注意
>
> 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 [提交 pr](https://github.com/Mrs4s/go-cqhttp/compare)
>
> 提示
>
> 要使用此接口, 通常需要安装 ffmpeg, 请参考 OneBot 实现的相关说明。

**传参:**

| 字段名       | 数据类型 | 默认值 | 说明                                                         |
| ------------ | -------- | ------ | ------------------------------------------------------------ |
| `file`       | string   | -      | 收到的语音文件名（消息段的 `file` 参数）, 如 `0B38145AA44505000B38145AA4450500.silk` |
| `out_format` | string   | -      | 要转换到的格式, 目前支持 `mp3`、`amr`、`wma`、`m4a`、`spx`、`ogg`、`wav`、`flac` |

**响应数据:**

| 字段名 | 数据类型 | 说明                                                         |
| ------ | -------- | ------------------------------------------------------------ |
| `file` | string   | 转换后的语音文件路径, 如 `/home/somebody/cqhttp/data/record/0B38145AA44505000B38145AA4450500.mp3` |

## 检查是否可以发送图片

**方法:**



**传参:**

提示

该 API 无需参数

**响应数据:**

| 字段名 | 数据类型 | 说明   |
| ------ | -------- | ------ |
| `yes`  | boolean  | 是或否 |

## 检查是否可以发送语音

**方法:**



**传参:**

> 提示
>
> 该 API 无需参数

**响应数据:**

| 字段名 | 数据类型 | 说明   |
| ------ | -------- | ------ |
| `yes`  | boolean  | 是或否 |

## 获取版本信息

**方法:**



**传参:**

> 提示
>
> 该 API 无需参数

**响应数据:**

| 字段名                       | 数据类型 | 默认值       | 说明                            |
| ---------------------------- | -------- | ------------ | ------------------------------- |
| `app_name`                   | string   | `go-cqhttp`  | 应用标识, 如 `go-cqhttp` 固定值 |
| `app_version`                | string   |              | 应用版本, 如 `v0.9.40-fix4`     |
| `app_full_name`              | string   |              | 应用完整名称                    |
| `protocol_version`           | string   | `v11`        | OneBot 标准版本 固定值          |
| `coolq_edition`              | string   | `pro`        | 原Coolq版本 固定值              |
| `coolq_directory`            | string   |              |                                 |
| `go-cqhttp`                  | bool     | true         | 是否为go-cqhttp 固定值          |
| `plugin_version`             | string   | `4.15.0`     | 固定值                          |
| `plugin_build_number`        | int      | 99           | 固定值                          |
| `plugin_build_configuration` | string   | `release`    | 固定值                          |
| `runtime_version`            | string   |              |                                 |
| `runtime_os`                 | string   |              |                                 |
| `version`                    | string   |              | 应用版本, 如 `v0.9.40-fix4`     |
| `protocol`                   | int      | `0/1/2/3/-1` | 当前登陆使用协议类型            |

## 重启 go-cqhttp

**方法:**

由于重启 go-cqhttp 实现同时需要重启 API 服务, 这意味着当前的 API 请求会被中断, 因此需要异步地重启, 接口返回的 `status` 是 `async`。

**传参:**

| 字段名  | 数据类型 | 默认值 | 说明                                                         |
| ------- | -------- | ------ | ------------------------------------------------------------ |
| `delay` | number   | `0`    | 要延迟的毫秒数, 如果默认情况下无法重启, 可以尝试设置延迟为 2000 左右 |

**响应数据:**

> 提示
>
> 该 API 无响应数据

## 清理缓存

**方法:**

> 注意
>
> 该 API 暂未被 go-cqhttp 支持, 您可以提交 pr 以使该 API 被支持 [提交 pr](https://github.com/Mrs4s/go-cqhttp/compare)

用于清理积攒了太多的缓存文件。

> 提示
>
> 该 API 无需参数也没有响应数据



## 设置群头像

**方法:**



**传参:**

| 字段       | 类型   | 说明                     |
| ---------- | ------ | ------------------------ |
| `group_id` | int64  | 群号                     |
| `file`     | string | 图片文件名               |
| `cache`    | int    | 表示是否使用已缓存的文件 |

[1] `file` **参数**支持以下几种格式：

- 绝对路径, 例如 `file:///C:\\Users\Richard\Pictures\1.png`, 格式使用 [`file` URIopen in new window](https://tools.ietf.org/html/rfc8089)
- 网络 URL, 例如 `http://i1.piimg.com/567571/fdd6e7b6d93f1ef0.jpg`
- Base64 编码, 例如 `base64://iVBORw0KGgoAAAANSUhEUgAAABQAAAAVCAIAAADJt1n/AAAAKElEQVQ4EWPk5+RmIBcwkasRpG9UM4mhNxpgowFGMARGEwnBIEJVAAAdBgBNAZf+QAAAAABJRU5ErkJggg==`

[2] `cache`**参数**: 通过网络 URL 发送时有效, `1`表示使用缓存, `0`关闭关闭缓存, 默认 为`1`

[3] 目前这个API在登录一段时间后因cookie失效而失效, 请考虑后使用



## 图片 OCR

**方法:**

注意

目前图片OCR接口仅支持接受的图片

ocr_image API移除了实验模式, 目前版本 .ocr_image 和 ocr_image 均能访问, 后期将只保留后者.

[go-cqhttp-v0.9.34更新日志](https://github.com/Mrs4s/go-cqhttp/releases/tag/v0.9.34)

**传参:**

| 字段    | 类型   | 说明   |
| ------- | ------ | ------ |
| `image` | string | 图片ID |

**响应数据:**

| 字段       | 类型            | 说明    |
| ---------- | --------------- | ------- |
| `texts`    | TextDetection[] | OCR结果 |
| `language` | string          | 语言    |

**TextDetection**

| 字段          | 类型    | 说明   |
| ------------- | ------- | ------ |
| `text`        | string  | 文本   |
| `confidence`  | int32   | 置信度 |
| `coordinates` | vector2 | 坐标   |

## 获取群系统消息

**方法:**



**传参:**

| 字段               | 类型             | 说明         |
| ------------------ | ---------------- | ------------ |
| `invited_requests` | InvitedRequest[] | 邀请消息列表 |
| `join_requests`    | JoinRequest[]    | 进群消息列表 |

> 注意
>
> 如果列表不存在任何消息, 将返回 `null`

**响应数据:**

| 字段               | 类型             | 说明         |
| ------------------ | ---------------- | ------------ |
| `invited_requests` | InvitedRequest[] | 邀请消息列表 |
| `join_requests`    | JoinRequest[]    | 进群消息列表 |

注意

如果列表不存在任何消息, 将返回 `null`

**InvitedRequest**

| 字段           | 类型   | 说明              |
| -------------- | ------ | ----------------- |
| `request_id`   | int64  | 请求ID            |
| `invitor_uin`  | int64  | 邀请者            |
| `invitor_nick` | string | 邀请者昵称        |
| `group_id`     | int64  | 群号              |
| `group_name`   | string | 群名              |
| `checked`      | bool   | 是否已被处理      |
| `actor`        | int64  | 处理者, 未处理为0 |

**JoinRequest**

| 字段             | 类型   | 说明              |
| ---------------- | ------ | ----------------- |
| `request_id`     | int64  | 请求ID            |
| `requester_uin`  | int64  | 请求者ID          |
| `requester_nick` | string | 请求者昵称        |
| `message`        | string | 验证消息          |
| `group_id`       | int64  | 群号              |
| `group_name`     | string | 群名              |
| `checked`        | bool   | 是否已被处理      |
| `actor`          | int64  | 处理者, 未处理为0 |

> 注意
>
> 在 `go-cqhttp-v0.9.40` 之前的版本中，无法获取被过滤的群系统消息

## 上传私聊文件

**方法:**



**传参:**

| 字段      | 类型   | 说明         |
| --------- | ------ | ------------ |
| `user_id` | int64  | 对方 QQ 号   |
| `file`    | string | 本地文件路径 |
| `name`    | string | 文件名称     |

注意

只能上传本地文件, 需要上传 `http` 文件的话请先调用 [`download_file` API](https://docs.go-cqhttp.org/api/#下载文件到缓存目录)下载

## 上传群文件

**方法:**



**传参:**

| 字段       | 类型   | 说明         |
| ---------- | ------ | ------------ |
| `group_id` | int64  | 群号         |
| `file`     | string | 本地文件路径 |
| `name`     | string | 储存名称     |
| `folder`   | string | 父目录ID     |

注意

在不提供 `folder` 参数的情况下默认上传到根目录

只能上传本地文件, 需要上传 `http` 文件的话请先调用 [`download_file` API](https://docs.go-cqhttp.org/api/#下载文件到缓存目录)下载



## 获取群文件系统信息

**方法:**



**传参:**

| 字段       | 类型  | 说明 |
| ---------- | ----- | ---- |
| `group_id` | int64 | 群号 |

**响应数据:**

| 字段          | 类型  | 说明       |
| ------------- | ----- | ---------- |
| `file_count`  | int32 | 文件总数   |
| `limit_count` | int32 | 文件上限   |
| `used_space`  | int64 | 已使用空间 |
| `total_space` | int64 | 空间上限   |

## 获取群根目录文件列表

> 提示
>
> `File` 和 `Folder` 对象信息请参考最下方

**方法:**



**传参:**

| 字段       | 类型  | 说明 |
| ---------- | ----- | ---- |
| `group_id` | int64 | 群号 |

**响应数据:**

| 字段      | 类型     | 说明       |
| --------- | -------- | ---------- |
| `files`   | File[]   | 文件列表   |
| `folders` | Folder[] | 文件夹列表 |



## 获取群子目录文件列表

> 提示
>
> `File` 和 `Folder` 对象信息请参考最下方

**方法:**



**传参:**

| 字段        | 类型   | 说明                        |
| ----------- | ------ | --------------------------- |
| `group_id`  | int64  | 群号                        |
| `folder_id` | string | 文件夹ID 参考 `Folder` 对象 |

**响应数据:**

| 字段      | 类型     | 说明       |
| --------- | -------- | ---------- |
| `files`   | File[]   | 文件列表   |
| `folders` | Folder[] | 文件夹列表 |

## 创建群文件文件夹

> 注意
>
> 仅能在根目录创建文件夹

**方法:**



**传参:**

| 字段        | 类型   | 说明       |
| ----------- | ------ | ---------- |
| `group_id`  | int64  | 群号       |
| `name`      | string | 文件夹名称 |
| `parent_id` | string | 仅能为 `/` |

**响应数据:**

提示

该 API 无响应数据

## 删除群文件夹

> 提示
>
> `Folder` 对象信息请参考最下方

**方法:**

**传参:**

| 字段        | 类型   | 说明                        |
| ----------- | ------ | --------------------------- |
| `group_id`  | int64  | 群号                        |
| `folder_id` | string | 文件夹ID 参考 `Folder` 对象 |

**响应数据:**

提示

该 API 无响应数据

## 删除群文件

> 提示
>
> `File` 对象信息请参考最下方

**方法:**



**传参:**

| 字段       | 类型   | 说明                      |
| ---------- | ------ | ------------------------- |
| `group_id` | int64  | 群号                      |
| `file_id`  | string | 文件ID 参考 `File` 对象   |
| `busid`    | int32  | 文件类型 参考 `File` 对象 |

**响应数据:**

> 提示
>
> 该 API 无响应数据

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

