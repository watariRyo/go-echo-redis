# go-echo-redis
## 概要
そのままgoでechoとredisを使ってみた、というだけ
ディレクトリ構成はアーキごちゃごちゃ変えたので適当

登録・ログイン・ログアウト
認証後のエンドポイント１つ実装したのみ

## サード
go.mod見れば自明だが以下使用

### DBまわり
ORM：gorm

### Redis
Redis:go-redis
Session:gorilla

### ログインとか認証機能
JWT:golang-jwt

### DI
DI:wire

### config周り
config:viper

### docker
dbとredisだけ構成
環境作ってそのうえで動かしたりとかはしていない

### frontendとの連携
現在ログインのみ（しかもリクエスト投げても何もしない）
