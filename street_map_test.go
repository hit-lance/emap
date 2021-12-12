package tinymap

import (
	"testing"
)

func TestStreetMap(t *testing.T) {

	t.Run("street map find closest", func(t *testing.T) {
		fn := "./berkeley.osm.xml"
		sns := SimpleNodeSet{}
		sm := NewStreetMapFrom(fn, &sns)

		got := sm.Closest(37.875613, -122.26009)
		want := int64(1281866063)
		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})

}
