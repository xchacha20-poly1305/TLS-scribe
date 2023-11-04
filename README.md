# TLS-scribe

很多时候，我们需要获取 TLS 网站的证书或者它的证书指纹，但是用浏览器看不够优雅，用 openssl 又记不住命令，所以我们需要一个工具让这个过程简单化。

「scribe」 意为抄写员，让他帮您抄写目标网站的证书吧！

# Build

```shell
./scripts/build.sh
```

得到 `build/scribe`。

# Usage

## Program

TODO.

## Command

```
scribe <Target> <Flags...>

Target:
   A url such as "https://github.com?fmt=sha256"
                 "quic://www.google.com"

Flags:
   -s, --sni      Server name.
```

链接后面可跟随查询字符串，可选：`pem`（默认）, `sha256`。

# TODO

* [ ] 更多文档、示例程序。

* [x] QUIC 支持。

# Credits

* [spf13/cobra](https://github.com/spf13/cobra)


* [quic-go/quic-go](https://github.com/quic-go/quic-go)

* ~~[refraction-networking/utls](https://github.com/refraction-networking/utls)~~

  为了防止获取证书被识别为恶意行为，我们使用了 uTLS 以期待能骗过对方的 waf XD。

  由于效果不明显，并且引入 uTLS 会导致二进制体积增加，所以移除。