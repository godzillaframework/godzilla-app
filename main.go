/* making this file exec */
package main

import (
	"log"

	"github.com/godzillaframework/godzilla"
)

/* main func */
func main() {
	/* new godzilla app */
	gz := godzilla.New()

	gz.Get("/index", func(ctx godzilla.Context) {
		ctx.SendString("Hey EveryOne")
	})

	/* params */
	gz.Get("/param/:param", func(ctx godzilla.Context) {
		ctx.SendString(ctx.Param("param"))
	})

	/* godzilla framework supports hosting static files */
	gz.Static("/static", "./imgs")

	gz.Static("/main", "./hello.html")

	/* middlewares */

	/* logger */
	logMiddleware := func(ctx godzilla.Context) {
		log.Printf("log!!")

		ctx.Next()
	}

	gz.Use(logMiddleware)

	/* unauthorized */
	unAuthorizedMiddleware := func(ctx godzilla.Context) {
		ctx.Status(godzilla.StatusUnauthorized).SendString("Hey get out of here!!")
	}

	gz.Get("/unauthorized", unAuthorizedMiddleware, func(ctx godzilla.Context) {
		ctx.SendString("UNAUTHORIZED PAGE ACCESSED!!!")
	})

	gz.Start(":8080")
}
