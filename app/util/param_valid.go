package util

import (
	"log"
)

func (c *Context) Validate(p interface{}) error {
	// validata := validator.New()
	// if err := validata.Struct(p); err != nil {
	// 	log.Println("param validate err:", err)
	// 	c.Response(err.Error(), nil)
	// 	return err
	// }
	if err := c.Ctx.ShouldBind(p); err != nil {
		log.Println("param validate err:", err)
		c.Response(400, err.Error(), nil)
		return err
	}
	return nil
}
