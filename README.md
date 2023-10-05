# Go Modular-Monolith Template

Go言語で行うモジュラモノリスのテンプレートリポジトリです．

本テンプレートでは簡単なToDoアプリを作成しています．


## Quick Start
```sh
cp .env.example .env
docker compose up -d --build
```

## 使用技術
- Go
  - API: [Echo](https://echo.labstack.com/guide/)
  - ORM: [GORM](https://gorm.io/ja_JP/docs/index.html)
- DB
  - PostgreSQL

## ディレクトリ構成
```
.
├── cmd           # エントリーポイント
│   └── main.go
├── modules       # モジュール
│   ├── bff
│   ├── todo
│   └── user
├── pkg           # 汎用パッケージ
│   ├── config
│   ├── database
│   ├── logger
│   └── middleware
```

## 注意事項
DockerComposeを用意していますが，あくまで開発用であり本番環境での使用は想定していません．

また，DBのマイグレーションの簡単化のため，ORMによるマイグレーションを行っています．<br>
実際のプロダクにョンにおいては，マイグレーションツールを用いてマイグレーションを行うことを推奨します．

より実際のプロダクションに近い環境は[こちら](https://github.com/Kurichi/go-template-pro)(※まだ作っていません)を参照してください．
