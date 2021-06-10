package unsave

import (
	"spotify/internal"
	"testing"

	"github.com/brianstrauch/spotify"
	"github.com/stretchr/testify/require"
)

func TestUnsaveCommand(t *testing.T) {
	api := new(spotify.MockAPI)

	var id string

	playback := &spotify.Playback{
		Item: spotify.Item{
			ID: id,
		},
	}

	api.On("GetPlayback").Return(playback, nil)
	api.On("RemoveSavedTracks", []string{id}).Return(nil)

	err := unsave(api)
	require.NoError(t, err)
}

func TestSavePodcast(t *testing.T) {
	api := new(spotify.MockAPI)

	playback := &spotify.Playback{
		Item: spotify.Item{
			Type: "episode",
		},
	}

	api.On("GetPlayback").Return(playback, nil)

	err := unsave(api)
	require.Equal(t, internal.SavePodcastErr, err.Error())
}
