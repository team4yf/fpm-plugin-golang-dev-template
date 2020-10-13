# fpm-go-plugin-{{name}}

### Install

`$ go get -u github.com/team4yf/fpm-go-plugin-{{name}}`

```golang

import _ "github.com/team4yf/fpm-go-plugin-{{name}}/plugin"

```

### Config
```yaml
xxx:
    xxx: xxx
```

### Usage

```golang

fpmApp.Execute("x.y", &fpm.BizParam{
    "body":    "ok",
})

```