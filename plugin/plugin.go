package plugin

import (
	"github.com/team4yf/yf-fpm-server-go/fpm"
)

type {{name}}Config struct {
	Foo string
}

func init() {
	fpm.RegisterByPlugin(&fpm.Plugin{
		Name: "fpm-plugin-{{name}}",
		V:    "0.0.1",
		Handler: func(fpmApp *fpm.Fpm) {
			config := {{name}}Config{}
			if fpmApp.HasConfig("{{name}}") {
				if err := fpmApp.FetchConfig("{{name}}", &config); err != nil {
					panic(err)
				}
			}

			fpmApp.Logger.Debugf("Startup : %s, config: %v", "{{name}}", config)

			fpmApp.AddBizModule("{{name}}", &fpm.BizModule{
				"foo": func(param *fpm.BizParam) (data interface{}, err error) {
					req := make(map[string]interface{})
					if err = param.Convert(&req); err != nil {
						return
					}
					fpmApp.Logger.Debugf("req: %#v", req)
					return
				},
			})

		},
	})
}
