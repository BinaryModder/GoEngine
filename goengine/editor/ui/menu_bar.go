package ui

import "github.com/AllenDang/giu"

func MenuBar() giu.Widget {

	return giu.Row(

		giu.Button("File"),

		giu.Button("Edit"),

		giu.Button("Assets"),

		giu.Button("GameObject"),

		giu.Button("Window"),

		giu.Button("Help"),
	)
}
