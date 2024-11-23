package qbittorrent

import (
	"fmt"
	"github.com/5rahim/go-qbit/application"
	"github.com/5rahim/go-qbit/log"
	"github.com/5rahim/go-qbit/rss"
	"github.com/5rahim/go-qbit/search"
	"github.com/5rahim/go-qbit/sync"
	"github.com/5rahim/go-qbit/torrent"
	"github.com/5rahim/go-qbit/transfer"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	baseURL     string
	logger      *zerolog.Logger
	client      *http.Client
	Username    string
	Password    string
	Port        int
	Host        string
	BinaryPath  string
	Application qbittorrent_application.Client
	Log         qbittorrent_log.Client
	RSS         qbittorrent_rss.Client
	Search      qbittorrent_search.Client
	Sync        qbittorrent_sync.Client
	Torrent     qbittorrent_torrent.Client
	Transfer    qbittorrent_transfer.Client
}

type NewClientOptions struct {
	HTTPS      bool
	Username   string
	Password   string
	Port       int
	Host       string
	BinaryPath string
}

func NewClient(opts *NewClientOptions) *Client {
	protocol := "http"
	if opts.HTTPS {
		protocol = "https"
	}
	baseURL := fmt.Sprintf("%s://%s:%d/api/v2", protocol, opts.Host, opts.Port)
	client := &http.Client{}
	return &Client{
		baseURL:    baseURL,
		client:     client,
		Username:   opts.Username,
		Password:   opts.Password,
		Port:       opts.Port,
		BinaryPath: opts.BinaryPath,
		Host:       opts.Host,
		Application: qbittorrent_application.Client{
			BaseUrl: baseURL + "/app",
			Client:  client,
		},
		Log: qbittorrent_log.Client{
			BaseUrl: baseURL + "/log",
			Client:  client,
		},
		RSS: qbittorrent_rss.Client{
			BaseUrl: baseURL + "/rss",
			Client:  client,
		},
		Search: qbittorrent_search.Client{
			BaseUrl: baseURL + "/search",
			Client:  client,
		},
		Sync: qbittorrent_sync.Client{
			BaseUrl: baseURL + "/sync",
			Client:  client,
		},
		Torrent: qbittorrent_torrent.Client{
			BaseUrl: baseURL + "/torrents",
			Client:  client,
		},
		Transfer: qbittorrent_transfer.Client{
			BaseUrl: baseURL + "/transfer",
			Client:  client,
		},
	}
}

func (c *Client) Login() error {
	endpoint := c.baseURL + "/auth/login"
	data := url.Values{}
	data.Add("username", c.Username)
	data.Add("password", c.Password)
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {

		}
	}()
	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status %s", resp.Status)
	}
	if len(resp.Cookies()) < 1 {
		return fmt.Errorf("no cookies in login response")
	}
	apiURL, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}
	jar.SetCookies(apiURL, []*http.Cookie{resp.Cookies()[0]})
	c.client.Jar = jar
	return nil
}

func (c *Client) Logout() error {
	endpoint := c.baseURL + "/auth/logout"
	request, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status %s", resp.Status)
	}
	return nil
}
