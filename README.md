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

* [x] QUIC 支持。

# Credits

* [spf13/cobra](https://github.com/spf13/cobra)

* [SagerNet/quic-go](https://github.com/SagerNet/quic-go)
