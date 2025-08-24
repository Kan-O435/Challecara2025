# Challecara2025

# challecara2025-back

このリポジトリは、challecara2025のバックエンドアプリケーションです。
Go言語とGinフレームワークで開発されており、Dockerコンテナとして動作します。

---

## 🚀 環境構築手順

以下の手順に従って、開発環境をセットアップしてください。

### 前提条件

- [Docker Desktop](https://www.docker.com/products/docker-desktop/) がインストール済みであること。
- [Git](https://git-scm.com/book/ja/v2/%E5%85%A5%E9%96%80-Git%E3%81%AE%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB) がインストール済みであること。

### 1. リポジトリのクローン

まず、このリポジトリをローカルにクローンします。

```bash
git clone [https://github.com/](https://github.com/)<ユーザー名>/challecara2025-back.git
cd challecara2025-back
```

### 2. Dockerイメージのビルド
```
docker build -t challecara2025-back .
```

### 3. コンテナの起動
```
docker run -p 8080:8080　challecara2025-back
```
アプリケーションの起動に成功すると以下のログがでるはず
```
[GIN-debug] Listening and serving HTTP on :8080
```

### 4. 動作確認
```
curl http://localhost:8080
```
以下のJSONレスポンスが返れば成功
```
curl http://localhost:8080
```