package pause

import (
	"errors"
	"spotify/internal"
	"spotify/internal/status"

	"github.com/brianstrauch/spotify"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause",
		Short: "pause music",
		RunE: func(cmd *cobra.Command, _ []string) error {
			api, err := internal.Authenticate()
			if err != nil {
				return err
			}

			status, err := Pause(api)
			if err != nil {
				return err
			}

			cmd.Print(status)
			return nil
		},
	}

	return cmd
}

func Pause(api internal.APIInterface) (string, error) {
	playback, err := api.GetPlayback()
	if err != nil {
		return "", err
	}

	if playback == nil {
		return "", errors.New(internal.ErrNoActiveDevice)
	}

	if err := api.Pause(); err != nil {
		return "", err
	}

	playback, err = internal.WaitForUpdatedPlayback(api, func(playback *spotify.Playback) bool {
		return !playback.IsPlaying
	})
	if err != nil {
		return "", err
	}

	return status.Show(playback), nil
}
