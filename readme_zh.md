# CUAgain —— 一个高可自定义度的Hololy开源服务端

---
Readme: 简体中文  [English](/readme.md)  [日本語](/readme_ja.md)
## 使用说明
1. 使用Reqable
2. 使用mitmproxy脚本(WIP)

### 使用[Reqable](https://reqable.com)
1. 在移动设备上安装Hololy
1. 按照 [Reqable文档](https://reqable.com/docs/getting-started/) 完成 Reqable 在PC和移动端的安装以及证书安装
1. 将移动端 Reqable 连接PC端，并开启抓包
1. 在PC端的 Reqable 右键点击重写按键，选择 管理规则，导入项目文件夹下的 [/mitm/reqable-rewrites.config](/mitm/reqable-rewrites.config) ，重定向到的 `https://cuagain.one` 可替换为其他CUAgain服务端
1. 打开Hololy，应用发起的 `/asset/Provisioning/hIz5WkFuV6qXgTtQ.json` 请求成功后(约5秒)即可断开Reqable


---
## 快速部署
1. 下载releases中对应架构的文件到服务器
2. 解压压缩包
3. 编辑config.yaml进一步配置
4. 授予下载的二进制文件写入权限
5. 运行二进制文件

---
## 配置说明
### CUAgain config.yaml配置
```yaml
cuagain:
  port: 8080  # CUAgain 应用程序监听端口
  password: 114514  # 角色资源管理密码 必须修改为自定义的强密码
  assets-proxy: true  # 启用角色资源代理
  assets-cache: true  # 启用本地资源缓存
  login-auth: false  # 是否在使用前需要登录认证
  login-password: "Suiseimywife"  # 登录认证密码
  login-timeout: 2880  # 每个IP地址的认证超时时间（分钟）
  get-ip-method: 0  # IP检测方法（0：使用客户端IP或自定义请求头）
  enable-global-holostar-movement: true  # 允许Hololive成员使用Holostar成员的动作
  remove-angle-limit: true  # 移除由角色裙子底部导致的视角限制
  rsa-public-key-path: "rsa/rsa_public_key.pem"  # RSA公钥文件路径
  rsa-private-key-path: "rsa/rsa_private_key.pem"  # RSA私钥文件路径

hololy:
  version-bypass: "2.4.8"  # 绕过Hololy的版本号限制（注意：官方服务停止后无效）
```
---
### `provision.json`配置
- `provision.json` 位于 `/json/provision.json` 

- ``````json
  {
    "provisioningType": 3,
    "api": "https://cuagain.one",	//CUAgain服务端
    "assetbundle": "https://raw.githubusercontent.com/suisei-pettan/hololy-assets/refs/heads/main",	//资源文件位置，当服务端的 assets-proxy 为true时可填入 https://cuagain.one/asset
    "hololiveaccount": "https://account.hololive.net"
  }
  ``````

  

---
### 自定义角色列表
- 请求 `/api/characters`获得，可修改 `./json/characters.json`自定义
---
### 自定义News列表
- 请求 `/api/news`获得，可修改 `./json/news.json`自定义
---

## TODO
- 服务端WebUI管理面板