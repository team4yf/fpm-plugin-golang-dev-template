# fpm-go-plugin-{{name}}

### Install

`$ go get -u github.com/team4yf/fpm-go-plugin-{{name}}`

```golang

import _ "github.com/team4yf/fpm-go-plugin-{{name}}/plugin"

```

### Config

`conf/config.local.yaml`

```yaml
{{name}}:
    foo: bar
```

### Usage

```golang

fpmApp.Execute("{{name}}.demo", &fpm.BizParam{
    "body":    "ok",
})

```