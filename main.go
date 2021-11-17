/* making this file exec */
package main

import "github.com/godzillaframework/godzilla"

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

	gz.Start(":8080")
}
