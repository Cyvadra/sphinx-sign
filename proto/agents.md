# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/agents.proto](#npool/agents.proto)
    - [AccountBalance](#sphinx.v1.AccountBalance)
    - [AccountTxJSON](#sphinx.v1.AccountTxJSON)
    - [BroadcastScriptRequest](#sphinx.v1.BroadcastScriptRequest)
    - [BroadcastScriptResponse](#sphinx.v1.BroadcastScriptResponse)
    - [CreateAccountRequest](#sphinx.v1.CreateAccountRequest)
    - [CreateAccountResponse](#sphinx.v1.CreateAccountResponse)
    - [GetBalanceRequest](#sphinx.v1.GetBalanceRequest)
    - [GetSignInfoRequest](#sphinx.v1.GetSignInfoRequest)
    - [GetTxJSONRequest](#sphinx.v1.GetTxJSONRequest)
    - [GetTxStatusRequest](#sphinx.v1.GetTxStatusRequest)
    - [GetTxStatusResponse](#sphinx.v1.GetTxStatusResponse)
    - [SignInfo](#sphinx.v1.SignInfo)
    - [SignScriptRequest](#sphinx.v1.SignScriptRequest)
    - [SignScriptResponse](#sphinx.v1.SignScriptResponse)
  
    - [Plugin](#sphinx.v1.Plugin)
    - [Sign](#sphinx.v1.Sign)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/agents.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/agents.proto



<a name="sphinx.v1.AccountBalance"></a>

### AccountBalance
GetBalance 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 查询的钱包地址 |
| timestamp_utc | [int64](#int64) |  | 长整型时间戳 |
| amount_int | [int64](#int64) |  | 金额整数 |
| amount_digits | [int32](#int32) |  | 金额*了10的^n，默认为9 |
| amount_string | [string](#string) |  | 金额字符串，&#34;123.45678901&#34; |






<a name="sphinx.v1.AccountTxJSON"></a>

### AccountTxJSON
GetTxJSONRequest 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json | [string](#string) |  |  |






<a name="sphinx.v1.BroadcastScriptRequest"></a>

### BroadcastScriptRequest
BroadcastScript 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_script | [string](#string) |  |  |






<a name="sphinx.v1.BroadcastScriptResponse"></a>

### BroadcastScriptResponse
BroadcastScript 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id_chain | [string](#string) |  |  |






<a name="sphinx.v1.CreateAccountRequest"></a>

### CreateAccountRequest
CreateAccount 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| uuid | [string](#string) |  | 数据库内账户私钥结构：createtime_utc, salt, prikey(encrypted) prikey(encrypted) = encrypt(prikey,string(uuid&#43;salt)) 对称加密 上线后，通知用户后台管理员做好用户数据（uuid）的备份，salt由离线签名服务端随机生成；需两把密钥解锁私钥 低安全性模式中/测试环境，复制uuid为salt，以防uuid丢失 设计思路：uuid为2FA所生成的用户私钥，保存在谷歌验证器中，但服务端必须做备份，以防用户丢失谷歌验证器后进行重置 理想情况下，2FA私钥仅保存在离线签名端，全链仅传输时效性验证码；这样可以实现安全策略 |






<a name="sphinx.v1.CreateAccountResponse"></a>

### CreateAccountResponse
CreateAccount 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 创建的钱包地址 |
| uuid | [string](#string) |  | user_id或与其绑定的唯一标识符，用于加密私钥提高安全性 |






<a name="sphinx.v1.GetBalanceRequest"></a>

### GetBalanceRequest
GetBalance 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| address | [string](#string) |  | 查询的钱包地址 |
| timestamp_utc | [int64](#int64) |  | 长整型时间戳 |






<a name="sphinx.v1.GetSignInfoRequest"></a>

### GetSignInfoRequest
GetSignInfo 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | 发送方钱包地址 |






<a name="sphinx.v1.GetTxJSONRequest"></a>

### GetTxJSONRequest
GetTxJSONRequest 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  | 继承钱包节点基础功能，预留 |
| address | [string](#string) |  | 要查询的钱包地址 |
| timefrom_utc | [uint64](#uint64) |  | 开始时间 |
| timetill_utc | [uint64](#uint64) |  | 结束时间 |
| limit_n | [int32](#int32) |  | 服务端限制返回条数 |






<a name="sphinx.v1.GetTxStatusRequest"></a>

### GetTxStatusRequest
GetTxStatus 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_id_chain | [string](#string) |  |  |






<a name="sphinx.v1.GetTxStatusResponse"></a>

### GetTxStatusResponse
GetTxStatus 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount_int | [int64](#int64) |  | 放大后的金额整数 |
| amount_digits | [int32](#int32) |  | amount_int == amount*10^n |
| amount_string | [string](#string) |  | 便于验证，数据库里不存 |
| address_from | [string](#string) |  | 发送方 |
| address_to | [string](#string) |  | 接收方 |
| transaction_id_chain | [string](#string) |  | 公链交易ID |
| createtime_utc | [int64](#int64) |  | 创建时间 |
| updatetime_utc | [int64](#int64) |  | 上次更新时间 |
| is_success | [bool](#bool) |  | 便于调用方判断 |
| is_failed | [bool](#bool) |  | 不success不fail就是pending了 |
| num_blocks_confirmed | [int32](#int32) |  | 已确认区块数 |






<a name="sphinx.v1.SignInfo"></a>

### SignInfo
GetSignInfo 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json | [string](#string) |  | 需要的预签名信息 |






<a name="sphinx.v1.SignScriptRequest"></a>

### SignScriptRequest
SignScript 参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| transaction_id_insite | [string](#string) |  | 站内交易ID，缓存去重 |
| address_from | [string](#string) |  | 发送方 |
| address_to | [string](#string) |  | 接收方 |
| amount_int | [int64](#int64) |  | 金额整数 |
| amount_digits | [int32](#int32) |  | 默认为9 |
| amount_string | [string](#string) |  | 字符串格式数据，便于确认 |
| uuid_signature | [string](#string) |  | 2FA签名 |
| createtime_utc | [int64](#int64) |  | 用户发起提现的时间，与2FA绑定 |






<a name="sphinx.v1.SignScriptResponse"></a>

### SignScriptResponse
SignScript 返回


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coin_id | [int32](#int32) |  |  |
| transaction_id_insite | [string](#string) |  | 站内交易ID，缓存去重 |
| address_from | [string](#string) |  | 发送方 |
| address_to | [string](#string) |  | 接收方 |
| amount_int | [int64](#int64) |  | 金额整数 |
| amount_digits | [int32](#int32) |  | 默认为9 |
| amount_string | [string](#string) |  | 字符串格式数据，便于确认 |
| script_json | [string](#string) |  | 可用于广播的交易script对象 |





 

 

 


<a name="sphinx.v1.Plugin"></a>

### Plugin
钱包代理插件

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSignInfo | [GetSignInfoRequest](#sphinx.v1.GetSignInfoRequest) | [SignInfo](#sphinx.v1.SignInfo) | 获取预签名信息 |
| GetBalance | [GetBalanceRequest](#sphinx.v1.GetBalanceRequest) | [AccountBalance](#sphinx.v1.AccountBalance) | 余额查询 |
| BroadcastScript | [BroadcastScriptRequest](#sphinx.v1.BroadcastScriptRequest) | [BroadcastScriptResponse](#sphinx.v1.BroadcastScriptResponse) | 广播交易 |
| GetTxStatus | [GetTxStatusRequest](#sphinx.v1.GetTxStatusRequest) | [GetTxStatusResponse](#sphinx.v1.GetTxStatusResponse) | 交易状态查询 |
| GetTxJSON | [GetTxJSONRequest](#sphinx.v1.GetTxJSONRequest) | [AccountTxJSON](#sphinx.v1.AccountTxJSON) | 账户交易查询 |


<a name="sphinx.v1.Sign"></a>

### Sign
离线签名服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateAccount | [CreateAccountRequest](#sphinx.v1.CreateAccountRequest) | [CreateAccountResponse](#sphinx.v1.CreateAccountResponse) | 创建账户 |
| SignScript | [SignScriptRequest](#sphinx.v1.SignScriptRequest) | [SignScriptResponse](#sphinx.v1.SignScriptResponse) | 进行签名 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

