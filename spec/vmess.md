# VMess 设计
## 摘要
* 版本：1

## 数据请求
版本部分：
* 1 字节：版本号，目前为 0x1
认证部分：
* 16 字节：md5(用户 VID + 'ASK')
指令部分：
* 1 字节：随机填充长度 M (M <= 32)
* M 字节：随机填充内容
* 1 字节：保留，总是 0x00
* 16 字节：请求数据 IV
* 16 字节：响应数据 IV
* 4 字节：认证信息 V
* 1 字节：指令
  * 0x00：保留
  * 0x01：TCP 请求
  * 0x02：UDP 请求
* 2 字节：目标端口
* 1 字节：目标类型
  * 0x01：IPv4
  * 0x02：域名
  * 0x03：IPv6
* 目标地址：
  * 4 字节：IPv4
  * 1 字节长度 + 域名
  * 16 字节：IPv6
数据部分
  * N 字节：请求数据

其中指令部分经过 AES-128 加密，Key 为用户 VID；数据部分使用块密码加密

## 数据应答
认证部分：
* 4 字节：认证信息 V
数据部分
* N 字节：应答数据

其中数据部分使用块密码加密