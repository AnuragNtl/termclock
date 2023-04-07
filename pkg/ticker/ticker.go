package ticker

import (
    "fmt"
    ui "github.com/gizak/termui/v3"
    "github.com/gizak/termui/widgets"

)

func Show() {
    if err := ui.Init(); err != nil {
        fmt.Println("Init error")
    }
    defer ui.Close()
    p := widgets.NewParagraph()
    p.Text = "sample"
    p.SetRect(0, 0, 25, 5)
    ui.Render(p)
}


