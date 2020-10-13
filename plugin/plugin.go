package plugin

import (
	"github.com/team4yf/yf-fpm-server-go/fpm"
)

type xxConfig struct {
	Xx string
}

func init() {
	fpm.RegisterByPlugin(&fpm.Plugin{
		Name: "fpm-plugin-xx",
		V:    "0.0.1",
		Handler: func(fpmApp *fpm.Fpm) {
			config := xxConfig{}
			if fpmApp.HasConfig("xx") {
				if err := fpmApp.FetchConfig("xx", &config); err != nil {
					panic(err)
				}
			}

			fpmApp.Logger.Debugf("Startup : %s, config: %v", "xx", config)

			fpmApp.AddBizModule("xx", &fpm.BizModule{
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
