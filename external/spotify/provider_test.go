package spotify

import "testing"

// TestSpotifyTrackPageSizeRespectsAPILimit asserts spotifyTrackPageSize stays
// within the Spotify Web API's silent 50-item cap; see the constant's comment
// in provider.go for why exceeding it silently drops tracks.
func TestSpotifyTrackPageSizeRespectsAPILimit(t *testing.T) {
	tests := []struct {
		name string
		got  int
		max  int
	}{
		{"spotifyTrackPageSize within /v1/playlists/{id}/items cap", spotifyTrackPageSize, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got > tt.max {
				t.Fatalf("page size = %d, want <= %d (Spotify Web API cap)", tt.got, tt.max)
			}
		})
	}
}

