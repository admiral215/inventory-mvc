package helpers

import beego "github.com/beego/beego/v2/server/web"

func GetFlash(c *beego.Controller) {
	flash := beego.ReadFromRequest(c)

	if flash != nil {
		if success, ok := flash.Data["success"]; ok {
			c.Data["flash"] = map[string]string{
				"success": success,
			}
		}

		if err, ok := flash.Data["error"]; ok {
			c.Data["flash"] = map[string]string{
				"error": err,
			}
		}
	}
}
