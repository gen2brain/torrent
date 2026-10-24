//go:build nowebtorrent
// +build nowebtorrent

package torrent

import (
	"net/url"
)

type webtorrentConfig struct{}

type clientWebtorrent struct{}

func (cl *Client) initWebtorrent() {}

func (t *Torrent) startWebsocketAnnouncer(u url.URL, shortInfohash [20]byte) torrentTrackerAnnouncer {
	return nil
}
