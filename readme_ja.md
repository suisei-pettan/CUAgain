# CUAgain —— 高度にカスタマイズ可能なHololyオープンソースサーバー

---
Readme: [简体中文](/readme_zh.md)  [English](/readme.md)  日本語
## 使用説明
1. Reqableを使用
2. mitmproxyスクリプトを使用（開発中）

### [Reqable](https://reqable.com)の使用
1. モバイルデバイスにHololyをインストール
2. [Reqableドキュメント](https://reqable.com/docs/getting-started/)に従い、PCおよびモバイルでのReqableのインストールと証明書設定を完了
3. モバイルのReqableをPCのReqableに接続し、パケットキャプチャを開始
4. PCのReqableで書き換えボタンを右クリックし、ルール管理を選択。プロジェクトフォルダの `/mitm/reqable-rewrites.config` をインポート。`https://cuagain.one` へのリダイレクトは、他のCUAgainサーバーに置き換え可能
5. Hololyを開く。アプリが `/asset/Provisioning/hIz5WkFuV6qXgTtQ.json` リクエストを正常に開始した後（約5秒）、Reqableを切断できます

---
## クイックデプロイメント
1. リリースから対応するアーキテクチャのファイルをサーバーにダウンロード
2. パッケージを解凍
3. config.yamlを編集してさらに設定
4. ダウンロードしたバイナリファイルに書き込み権限を付与
5. バイナリファイルを実行

---
## 設定説明
### CUAgain config.yaml 設定
```yaml
cuagain:
  port: 8080  # CUAgainアプリケーションのリスニングポート
  password: 114514  # ロールリソース管理パスワード（カスタムの強力なパスワードに変更する必要があります）
  assets-proxy: true  # ロールリソースプロキシを有効化
  assets-cache: true  # ローカルリソースキャッシュを有効化
  login-auth: false  # 使用前にログイン認証が必要かどうか
  login-password: "Suiseimywife"  # ログイン認証パスワード
  login-timeout: 2880  # 各IPアドレスの認証タイムアウト（分）
  get-ip-method: 0  # IP検出方法（0：クライアントIPまたはカスタムリクエストヘッダーを使用）
  enable-global-holostar-movement: true  # HololiveメンバーにHolostarメンバーのアクションを許可
  remove-angle-limit: true  # キャラクターのスカート下部による視角制限を削除
  rsa-public-key-path: "rsa/rsa_public_key.pem"  # RSA公開鍵ファイルのパス
  rsa-private-key-path: "rsa/rsa_private_key.pem"  # RSA秘密鍵ファイルのパス

hololy:
  version-bypass: "2.4.8"  # Hololyバージョン番号の制限をバイパス（注意：公式サービス停止後は無効）
```
---
### `provision.json` 設定
- `provision.json` は `/json/provision.json` に配置

- ``````json
  {
    "provisioningType": 3,
    "api": "https://cuagain.one",	// CUAgainサーバー
    "assetbundle": "https://raw.githubusercontent.com/suisei-pettan/hololy-assets/refs/heads/main",	// リソースファイルの場所、サーバーのassets-proxyがtrueの場合はhttps://cuagain.one/assetを入力可能
    "hololiveaccount": "https://account.hololive.net"
  }
  ``````

---
### カスタムキャラクターリスト
- `/api/characters` をリクエストして取得、`./json/characters.json` を修正してカスタマイズ可能
---
### カスタムニュースリスト
- `/api/news` をリクエストして取得、`./json/news.json` を修正してカスタマイズ可能
---

## TODO
- サーバーWebUI管理パネル