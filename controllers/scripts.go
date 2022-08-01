package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	client "gitlab.hho-inc.com/devops/oops-api/utils"
)

// @summary 获取脚本
// @description 获取脚本
// @tags 脚本
// @produce plain
// @accept plain
// @param script path string true "脚本名"
// @success 200 string string
// @failure 404 string string
// @router /scripts/{script} [get]
func ExecShellScripts(c *gin.Context) {
	shell := c.Param("script")
    consul := client.NewConsul("/devops/scripts")
	data, _ := consul.GetKV(shell)
	if len(data) == 0 {
		c.String(404, fmt.Sprintf("%s脚本未找到", shell))
	}
	c.String(200, string(data))
}
