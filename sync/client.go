package qbittorrent_sync

import (
	"github.com/5rahim/go-qbit/model"
	"github.com/5rahim/go-qbit/util"
	"github.com/rs/zerolog"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	BaseUrl string
	Client  *http.Client
	Logger  *zerolog.Logger
}

func (c Client) GetMainData(rid int) (*qbittorrent_model.SyncMainData, error) {
	params := url.Values{}
	params.Add("rid", strconv.Itoa(rid))
	endpoint := c.BaseUrl + "/maindata?" + params.Encode()
	var res qbittorrent_model.SyncMainData
	if err := qbittorrent_util.GetInto(c.Client, &res, endpoint, nil); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c Client) GetTorrentPeersData(hash string, rid int) (*qbittorrent_model.SyncPeersData, error) {
	params := url.Values{}
	params.Add("hash", hash)
	params.Add("rid", strconv.Itoa(rid))
	endpoint := c.BaseUrl + "/torrentPeers?" + params.Encode()
	var res qbittorrent_model.SyncPeersData
	if err := qbittorrent_util.GetInto(c.Client, &res, endpoint, nil); err != nil {
		return nil, err
	}
	return &res, nil
}
