# CUAgain —— 高度にカスタマイズ可能なHololyオープンソースサーバー

---
Readme: [简体中文](/readme_zh.md) | [English](/readme.md) | 日本語

## 特徴
1. 公式サーバーとは独立して動作
2. すべてのキャラクターと衣装をアンロック
3. 一部の視点制限を解除（AR カメラを除く）
4. HololiveメンバーがHolostarメンバーの動きを使用することを許可

---
## デモ
`https://cuagain.one`

---
## 使用方法
1. Reqableを使用
2. mitmproxyスクリプトを使用（作業中）

### [Reqable](https://reqable.com)の使用
1. モバイルデバイスにHololyをインストール
2. [Reqableドキュメント](https://reqable.com/docs/getting-started/)に従って、PCおよびモバイルでのReqableのインストールと証明書のセットアップを完了
3. モバイルReqableをPCのReqableに接続し、パケットキャプチャを有効化
4. PC上のReqableで右クリックし、書き換えボタンから「ルール管理」を選択。プロジェクトフォルダーの[/mitm/reqable-rewrites.config](/mitm/reqable-rewrites.config)をインポート。`https://cuagain.one`は別のCUAgainサーバーに置き換え可能
5. Hololyを開く。`/asset/Provisioning/hIz5WkFuV6qXgTtQ.json`のリクエストが成功したら（約5秒）、Reqableを切断できます

---
## クイックデプロイ
1. Go実行環境をインストール
2. サーバーのCUAgainをデプロイするディレクトリで `git clone https://github.com/suisei-pettan/CUAgain.git` を実行
3. プロジェクトディレクトリに入り、以下の設定説明に従ってサーバー設定を完了
4. 環境変数 `CGO_ENABLED=1` を設定
5. `go run main.go` を実行

---
## 設定説明
### CUAgain config.yaml 設定
```yaml
cuagain:
  port: 8080  # CUAgainアプリケーションのリスニングポート
  password: 114514  # キャラクターリソース管理パスワード - カスタムの強力なパスワードに変更する必要があります
  assets-proxy: true  # キャラクターリソースプロキシを有効化
  assets-cache: true  # ローカルリソースキャッシュを有効化
  login-auth: false  # 使用前にログイン認証が必要かどうか
  login-password: "Suiseimywife"  # ログイン認証パスワード
  login-timeout: 2880  # 各IPアドレスの認証タイムアウト（分）
  get-ip-method: 0  # IPDetection方法（0：クライアントIPまたはカスタムリクエストヘッダーを使用）
  enable-global-holostar-movement: true  # HololiveメンバーがHolostarメンバーの動きを使用することを許可
  remove-angle-limit: true  # キャラクターのスカート下部による視点制限を解除
  rsa-public-key-path: "rsa/rsa_public_key.pem"  # RSA公開鍵ファイルパス
  rsa-private-key-path: "rsa/rsa_private_key.pem"  # RSA秘密鍵ファイルパス

hololy:
  version-bypass: "2.4.8"  # Hololyのバージョン番号制限をバイパス（注意：公式サービス停止後は無効）
```
---
### `provision.json`の設定
- `provision.json` は `/json/provision.json` に配置

- ``````json
  {
    "provisioningType": 3,
    "api": "https://cuagain.one",    // CUAgainサーバー
    "assetbundle": "https://raw.githubusercontent.com/suisei-pettan/hololy-assets/refs/heads/main",    // アセットファイルの場所、サーバーのassets-proxyがtrueの場合はhttps://cuagain.one/assetを入力可能
    "hololiveaccount": "https://account.hololive.net"
  }
  ``````

---
### カスタムキャラクターリスト
- `/api/characters` にリクエストして取得、`./json/characters.json` を修正してカスタマイズ可能
---
### カスタムニュースリスト
- `/api/news` にリクエストして取得、`./json/news.json` を修正してカスタマイズ可能
---

## TODO
- サーバーサイドWebUI管理パネル

---
## 注意事項
- このプロジェクトを悪用せず、Hololiveメンバーや公式チームに迷惑をかけないようにしてください
- デモのキャラクターと衣装のアンロック機能は制限されており、2024年12月2日のモデルと、以前Hololyに存在した卒業メンバーのみを復元します。有料の衣装はデモでは提供されません
- 一部の隠しモデルは、公式モデル自体の異常により、エラーが発生する可能性があります。これはCUAgainでは解決できません