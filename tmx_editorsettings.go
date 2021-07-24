package tiled

type EditorSettings struct {
	Export *Export `xml:"export"`
}

type Export struct {
	Target string `xml:"target,attr"`
	Format string `xml:"format,attr"`
}
