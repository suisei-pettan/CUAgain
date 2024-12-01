# CUAgain —— A Highly Customizable Hololy Open Source Server

---

Readme: [简体中文](/readme_zh.md)  English  [日本語](/readme_ja.md)
## Usage Instructions
1. Using Reqable
2. Using mitmproxy script (WIP)

### Using [Reqable](https://reqable.com)
1. Install Hololy on your mobile device
2. Follow [Reqable documentation](https://reqable.com/docs/getting-started/) to complete Reqable installation and certificate setup on PC and mobile
3. Connect mobile Reqable to PC Reqable and start packet capture
4. Right-click the rewrite button on PC Reqable, select Manage Rules, import [/mitm/reqable-rewrites.config](/mitm/reqable-rewrites.config) from the project folder. The redirection to `https://cuagain.one` can be replaced with other CUAgain servers
5. Open Hololy. After the app successfully initiates the `/asset/Provisioning/hIz5WkFuV6qXgTtQ.json` request (about 5 seconds), you can disconnect Reqable

---
## Quick Deployment

1. Install the Go runtime environment

2. In the directory on your server where you want to deploy CUAgain, execute: `git clone https://github.com/suisei-pettan/CUAgain.git`

3. Enter the project directory and complete server configuration according to the configuration instructions below

---
## Configuration Instructions
### CUAgain config.yaml Configuration
```yaml
cuagain:
  port: 8080  # CUAgain application listening port
  password: 114514  # Role resource management password (must be changed to a custom strong password)
  assets-proxy: true  # Enable role resource proxy
  assets-cache: true  # Enable local resource caching
  login-auth: false  # Whether login authentication is required before use
  login-password: "Suiseimywife"  # Login authentication password
  login-timeout: 2880  # Authentication timeout for each IP address (minutes)
  get-ip-method: 0  # IP detection method (0: use client IP or custom request header)
  enable-global-holostar-movement: true  # Allow Hololive members to use Holostar member actions
  remove-angle-limit: true  # Remove view angle limit caused by character skirt bottom
  rsa-public-key-path: "rsa/rsa_public_key.pem"  # RSA public key file path
  rsa-private-key-path: "rsa/rsa_private_key.pem"  # RSA private key file path

hololy:
  version-bypass: "2.4.8"  # Bypass Hololy version number limit (note: invalid after official service stops)
```
---
### `provision.json` Configuration
- `provision.json` is located at `/json/provision.json`

- ``````json
  {
    "provisioningType": 3,
    "api": "https://cuagain.one",	// CUAgain server
    "assetbundle": "https://raw.githubusercontent.com/suisei-pettan/hololy-assets/refs/heads/main",	// Resource file location, can be filled with https://cuagain.one/asset when server's assets-proxy is true
    "hololiveaccount": "https://account.hololive.net"
  }
  ``````

---
### Custom Character List
- Obtain by requesting `/api/characters`, can be customized by modifying `./json/characters.json`
---
### Custom News List
- Obtain by requesting `/api/news`, can be customized by modifying `./json/news.json`
---

## TODO
- Server WebUI management panel