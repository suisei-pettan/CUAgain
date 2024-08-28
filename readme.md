## CUAgain

[中文文档](https://github.com/suisei-pettan/CUAgain/tree/main/readmeCN.md)

### Features
- Restore removed models
- Bypass version update checks
- Circumvent regional access restrictions to official servers

### Note
This service only restores models for characters that were once freely available and have officially graduated.

---

### How to Use
Before the application loads, intercept and rewrite the response for `/asset/Provisioning/*` to:
```json
{
  "provisioningType": 3,
  "api": "https://cuagain.one:8443",
  "assetbundle": "https://cuagain.one:8443/asset",
  // Other fields remain unchanged
}
```

---

## Screenshots

![Rendering-AChan](https://github.com/suisei-pettan/CUAgain/blob/main/img/Rendering-AChan.jpg?raw=true)

---