package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var (
	dropTarget string = "Drop here"
)

func loop() {
	g.SingleWindow("Drag and Drop").Layout(
		g.Row(
			g.Button("Drag me: 9"),
			g.Custom(func() {
				if imgui.BeginDragDropSource() {
					imgui.SetDragDropPayload("DND_DEMO", 9)
					g.Label("9").Build()
					imgui.EndDragDropSource()
				}
			}),
			g.Button("Drag me: 10"),
			g.Custom(func() {
				if imgui.BeginDragDropSource() {
					imgui.SetDragDropPayload("DND_DEMO", 10)
					g.Label("10").Build()
					imgui.EndDragDropSource()
				}
			}),
		),
		g.InputTextMultiline("##DropTarget", &dropTarget).Size(-1, -1).Flags(g.InputTextFlags_ReadOnly),
		g.Custom(func() {
			if imgui.BeginDragDropTarget() {
				payload := imgui.AcceptDragDropPayload("DND_DEMO")
				if payload != 0 {
					dropTarget = fmt.Sprintf("Dropped value: %d", payload.Data())
				}
				imgui.EndDragDropTarget()
			}
		}),
	)
}

func main() {
	wnd := g.NewMasterWindow("Drag and Drop", 600, 400, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
