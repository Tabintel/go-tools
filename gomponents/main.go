package main

import (
	"fmt"
	"net/http"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func main() {
	http.Handle("/", createHandler("Welcome!", simpleComponent("Hello, this is the main page!")))
	http.Handle("/contact", createHandler("Contact", simpleComponent("Contact us!")))
	http.Handle("/about", createHandler("About", simpleComponent("About this site!")))

	// Print a message indicating that the server is running
	fmt.Println("Server is running on http://localhost:8080")

	_ = http.ListenAndServe("localhost:8080", nil)
}

func createHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = Page(title, r.URL.Path, body).Render(w)
	}
}

func simpleComponent(content string) g.Node {
	return Div(
		H1(g.Text(content)),
		P(g.Text("This is a simple component.")),
	)
}

func Page(title, path string, body g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Body: []g.Node{
			Navbar(path),
			Container(body),
		},
	})
}

func Navbar(currentPath string) g.Node {
	return Nav(Class("navbar"),
		Container(
			NavbarLink("/", "Home", currentPath == "/"),
			NavbarLink("/contact", "Contact", currentPath == "/contact"),
			NavbarLink("/about", "About", currentPath == "/about"),
		),
	)
}

func NavbarLink(path, text string, active bool) g.Node {
	return A(Href(path), g.Text(text),
		c.Classes{
			"active": active,
		},
	)
}

func Container(children ...g.Node) g.Node {
	return Div(Class("container"), g.Group(children))
}
