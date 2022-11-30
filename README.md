# todo アプリ

- golang、gin、docker を使用した todo アプリを制作
- 目的としては、golang と gin の概要を把握する

### version

- golang; 1.19.2 darwin/amd64
- framework: gin
- docker: 20.10.17

### Alpine Linuxを使用したgolang,ginを動かす為には、下記のパッケージが必須
```sh
RUN apk update && apk add \
    alpine-sdk \
    build-base
```

# 参考サイト
- [Go+Gin+Air+docker-composeでホットリロード可能なローカル環境構築する](https://nothing-behind.com/2022/08/17/goginairdocker-compose%E3%81%A7%E3%83%9B%E3%83%83%E3%83%88%E3%83%AA%E3%83%AD%E3%83%BC%E3%83%89%E5%8F%AF%E8%83%BD%E3%81%AA%E3%83%AD%E3%83%BC%E3%82%AB%E3%83%AB%E7%92%B0%E5%A2%83%E6%A7%8B%E7%AF%89/)
- [Gin Web Framework](https://gin-gonic.com/ja/)
- [gin の github](https://github.com/gin-gonic)
- [Gin(Golang)における HTML テンプレート記述方法](https://qiita.com/lanevok/items/dbf591a3916070fcba0d)
- [cosmtrek/air](https://github.com/cosmtrek/air)
- [Alpineのapkコマンドとaptの違いまとめ](https://kleinblog.net/alpine-apk-cmd.html#:~:text=apk%E3%81%A8%E3%81%AF,%E3%81%A7%E3%81%8D%E3%82%8B%E3%80%81%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E7%AE%A1%E7%90%86%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%81%A7%E3%81%99%E3%80%82)
