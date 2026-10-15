// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/egress/pkg/types"
)

func TestValidateUrl(t *testing.T) {
	var twitchUpdated = regexp.MustCompile("rtmps://(.*).contribute.live-video.net/app/streamkey")
	var twitchRedacted = regexp.MustCompile(`rtmps://(.*).contribute.live-video.net/app/\{str\.\.\.key}`)

	o := &StreamConfig{}

	for _, test := range []struct {
		url      string
		twitch   bool
		parsed   string
		redacted string
	}{
		{
			url:      "mux://streamkey",
			parsed:   "rtmps://global-live.mux.com:443/app/streamkey",
			redacted: "rtmps://global-live.mux.com:443/app/{str...key}",
		},
		{
			url:    "twitch://streamkey",
			twitch: true,
		},
		{
			url:    "rtmp://fake.contribute.live-video.net/app/streamkey",
			twitch: true,
		},
		{
			url:      "rtmp://localhost:1935/live/streamkey",
			parsed:   "rtmp://localhost:1935/live/streamkey",
			redacted: "rtmp://localhost:1935/live/{str...key}",
		},
		{
			url:      "rtmps://localhost:1935/live/streamkey",
			parsed:   "rtmps://localhost:1935/live/streamkey",
			redacted: "rtmps://localhost:1935/live/{str...key}",
		},
	} {
		parsed, redacted, streamID, err := o.ValidateUrl(test.url, types.OutputTypeRTMP)
		require.NoError(t, err)
		require.NotEmpty(t, streamID)

		if test.twitch {
			require.NotEmpty(t, twitchUpdated.FindString(parsed), parsed)
			require.NotEmpty(t, twitchRedacted.FindString(redacted), redacted)
		} else {
			require.Equal(t, test.parsed, parsed)
			require.Equal(t, test.redacted, redacted)
		}
	}
}

func TestGetUrl(t *testing.T) {
	o := &StreamConfig{}
	require.NoError(t, o.updateTwitchTemplate())

	parsedTwitchUrl := strings.ReplaceAll(o.twitchTemplate, "{stream_key}", "streamkey")
	urls := []string{
		"rtmps://global-live.mux.com:443/app/streamkey",
		parsedTwitchUrl,
		parsedTwitchUrl,
		"rtmp://localhost:1935/live/streamkey",
	}

	for _, url := range []string{urls[0], urls[1], urls[3]} {
		_, err := o.AddStream(url, types.OutputTypeRTMP)
		require.NoError(t, err)
	}

	for i, rawUrl := range []string{
		"mux://streamkey",
		"twitch://streamkey",
		"rtmp://any.contribute.live-video.net/app/streamkey",
		"rtmp://localhost:1935/live/streamkey",
	} {
		stream, err := o.GetStream(rawUrl)
		require.NoError(t, err)
		require.Equal(t, urls[i], stream.ParsedUrl)
	}
}

// BEGIN OPENVIDU BLOCK
func TestStreamTag(t *testing.T) {
	o := &StreamConfig{}

	// the tag is reported after the redacted url, and kept out of the url streamed to
	parsed, redacted, streamID, err := o.ValidateUrl("rtmp://localhost:1935/live/streamkey#ov-stream-tag=room-1--kX9pL2mQ7a", types.OutputTypeRTMP)
	require.NoError(t, err)
	require.Equal(t, "rtmp://localhost:1935/live/streamkey", parsed)
	require.Equal(t, "rtmp://localhost:1935/live/{str...key}#ov-stream-tag=room-1--kX9pL2mQ7a", redacted)
	require.Equal(t, "streamkey", streamID)

	// keys with the same ends, or the same url streamed to again, are told apart by their tags
	_, first, _, err := o.ValidateUrl("rtmp://localhost:1935/live/abcSAMExyz#ov-stream-tag=a", types.OutputTypeRTMP)
	require.NoError(t, err)
	_, second, _, err := o.ValidateUrl("rtmp://localhost:1935/live/abcDIFFxyz#ov-stream-tag=b", types.OutputTypeRTMP)
	require.NoError(t, err)
	require.NotEqual(t, first, second)

	// the tag follows the rewrite of a url, and the url keeps working without one
	parsed, redacted, _, err = o.ValidateUrl("mux://streamkey#ov-stream-tag=mux_1", types.OutputTypeRTMP)
	require.NoError(t, err)
	require.Equal(t, "rtmps://global-live.mux.com:443/app/streamkey", parsed)
	require.Equal(t, "rtmps://global-live.mux.com:443/app/{str...key}#ov-stream-tag=mux_1", redacted)

	// a tagged url finds the stream it added, with or without its tag
	_, err = o.AddStream("rtmp://localhost:1935/live/taggedkey#ov-stream-tag=t1", types.OutputTypeRTMP)
	require.NoError(t, err)
	for _, rawUrl := range []string{"rtmp://localhost:1935/live/taggedkey#ov-stream-tag=t1", "rtmp://localhost:1935/live/taggedkey"} {
		stream, err := o.GetStream(rawUrl)
		require.NoError(t, err)
		require.Equal(t, "rtmp://localhost:1935/live/taggedkey", stream.ParsedUrl)
		require.Equal(t, "rtmp://localhost:1935/live/{tag...key}#ov-stream-tag=t1", stream.StreamInfo.Url)
	}

	// a malformed tag is refused without showing the stream key
	for _, rawUrl := range []string{
		"rtmp://localhost:1935/live/secretkey#ov-stream-tag=",
		"rtmp://localhost:1935/live/secretkey#ov-stream-tag=a b",
		"rtmp://localhost:1935/live/secretkey#ov-stream-tag=" + strings.Repeat("a", 129),
	} {
		_, _, _, err = o.ValidateUrl(rawUrl, types.OutputTypeRTMP)
		require.ErrorContains(t, err, "stream tags must be")
		require.NotContains(t, err.Error(), "secretkey")
	}
}

// END OPENVIDU BLOCK
