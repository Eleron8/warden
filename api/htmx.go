package api

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
)

func getHtmxHead() *elem.Element {
	head := elem.Head(nil,
		elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@1.9.6"}),
	)
	return head
}

func greetingsHtmx() string {
	bodyStyle := styles.Props{
		styles.BackgroundColor: "#f4f4f4",
		styles.FontFamily:      "Arial, sans-serif",
		styles.Height:          "100vh",
		styles.Display:         "flex",
		styles.FlexDirection:   "column",
		styles.AlignItems:      "center",
		styles.JustifyContent:  "center",
	}

	body := elem.Body(attrs.Props{attrs.Style: bodyStyle.ToInline()},
		elem.H1(nil, elem.Text("Warden App")),
	)
	pagecontent := elem.Html(nil, getHtmxHead(), body)
	return pagecontent.Render()
}
