## CUAgain

### 功能
- 恢复被移除的模型
- 绕过版本更新检测
- 绕过部分地区对官方服务器的访问限制

### 注意事项
此服务仅对曾经免费公开且正常毕业的角色的模型提供恢复。

---

### 使用方法
在应用程序加载前抓包重写`/asset/Provisioning/*`的返回值为
```json
{
  "provisioningType": 3,
  "api": "https://cuagain.one:8443",
  "assetbundle": "https://cuagain.one:8443/asset",
  //剩余项不变
}
```

---

## 效果图

![Rendering-AChan](https://github.com/suisei-pettan/CUAgain/blob/main/img/Rendering-AChan.jpg?raw=true)

---