# auth_requestを使ったサンプル

[EN](./docs/README-en.md)

## アプリケーションの起動

`docker-compose up -d`で開発環境とnginxを起動します。

nginxは `8888` ポートを公開しています。

開発環境へ devcontainer 経由で入り、以下のコマンドを実施する

`go run main.go app1`

ログインやメインのアプリケーションサーバー

メインアプリケーションでは `app2` のHTMLをiframeで表示している

`go run main.go app2`

セッション情報を表示するアプリケーションサーバー

## URLs

- `app1`
  - `/login` ログイン
  - `/app` アプリケーション
  - `/auth` 認証エンドポイント
- `app2`
  - `/app` アプリケーション

## ログイン情報

ログインユーザー: `user@example.com`  
パスワード: `password`

## 未ログイン状態での検証

`http://localhost:8888/app`

へアクセスする

`auth_request` を利用して、`/auth` に対してセッションの確認を行う。

未ログインなので `/auth` は `Location`ヘッダーにリダイレクト先を設定

nginx の設定で Locationヘッダーの内容をCookieへ設定し、`unauthorized.html` を描画する

`unauthorized.html` は Cookieからリダイレクト先を取得し、親のlocation.hrefへリダイレクト先を設定

`/login`へリダイレクトされると成功
