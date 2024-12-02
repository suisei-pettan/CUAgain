# CUAgain —— A Highly Customizable Hololy Open-Source Server

---
Readme: [简体中文](/readme_zh.md) | English | [日本語](/readme_ja.md)

## Features
1. Works independently from the official server
2. Unlock all characters and costumes
3. Remove some view angle restrictions (excluding AR camera)
4. Allow Hololive members to use Holostar member movements

---
## Demo
`https://cuagain.one`

---
## Usage Instructions
1. Using Reqable
2. Using mitmproxy script (Work in Progress)

### Using [Reqable](https://reqable.com)
1. Install Hololy on your mobile device
2. Follow the [Reqable documentation](https://reqable.com/docs/getting-started/) to complete Reqable installation and certificate setup on PC and mobile
3. Connect mobile Reqable to PC Reqable and enable packet capture
4. Right-click the rewrite button in PC Reqable, select Manage Rules, import the [/mitm/reqable-rewrites.config](/mitm/reqable-rewrites.config) from the project folder. The redirection to `https://cuagain.one` can be replaced with another CUAgain server
5. Open Hololy. After the `/asset/Provisioning/hIz5WkFuV6qXgTtQ.json` request is successful (about 5 seconds), you can disconnect Reqable

---
## Quick Deployment
1. Install Go runtime environment
2. Execute `git clone https://github.com/suisei-pettan/CUAgain.git` in the directory where you want to deploy CUAgain on your server
3. Enter the project directory and complete server configuration according to the configuration instructions below
4. Set environment variable `CGO_ENABLED=1`
5. Run `go run main.go`

---
## Configuration Instructions
### CUAgain config.yaml Configuration
```yaml
cuagain:
  port: 8080  # CUAgain application listening port
  password: 114514  # Character resource management password - must be changed to a custom strong password
  assets-proxy: true  # Enable character resource proxy
  assets-cache: true  # Enable local resource caching
  login-auth: false  # Whether login authentication is required before use
  login-password: "Suiseimywife"  # Login authentication password
  login-timeout: 2880  # Authentication timeout for each IP address (minutes)
  get-ip-method: 0  # IP detection method (0: Use client IP or custom request header)
  enable-global-holostar-movement: true  # Allow Hololive members to use Holostar member movements
  remove-angle-limit: true  # Remove view angle restrictions caused by character skirt bottom
  rsa-public-key-path: "rsa/rsa_public_key.pem"  # RSA public key file path
  rsa-private-key-path: "rsa/rsa_private_key.pem"  # RSA private key file path

hololy:
  version-bypass: "2.4.8"  # Bypass Hololy version number limit (Note: Invalid after official service stops)
```
---
### `provision.json` Configuration
- `provision.json` is located at `/json/provision.json`

- ``````json
  {
    "provisioningType": 3,
    "api": "https://cuagain.one",    // CUAgain server
    "assetbundle": "https://raw.githubusercontent.com/suisei-pettan/hololy-assets/refs/heads/main",    // Asset file location, can be filled with https://cuagain.one/asset when the server's assets-proxy is true
    "hololiveaccount": "https://account.hololive.net"
  }
  ``````

---
### Custom Character List
- Request `/api/characters` to obtain, can be customized by modifying `./json/characters.json`
---
### Custom News List
- Request `/api/news` to obtain, can be customized by modifying `./json/news.json`
---

## TODO
- Server-side WebUI management panel

---
## Precautions
- Please do not abuse this project to avoid causing trouble for Hololive members or the official team
- The demo's character and costume unlock function is limited, only restoring the models from December 2, 2024, and all graduated members that previously existed in Hololy. Paid costumes are not provided in the demo
- Some hidden models may produce errors due to abnormalities in the official models themselves, which CUAgain cannot resolve