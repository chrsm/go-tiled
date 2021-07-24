package tiled_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chrsm/go-tiled"
)

func TestLoadEditorSettings(t *testing.T) {
	const mapPath = "assets/editorsettings.tmx"

	m, err := tiled.LoadFromFile(mapPath)
	if err != nil {
		t.Fatal(err)
	}

	r := require.New(t)
	r.NotNil(m.EditorSettings, "expected EditorSettings to be populated")
	r.NotNil(m.EditorSettings.Export, "expected Export to be populated")
	r.Equal("dst/lua/file.lua", m.EditorSettings.Export.Target)
	r.Equal("lua", m.EditorSettings.Export.Format)
}
