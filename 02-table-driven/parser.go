package tabledriven

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrEmptyURL      = errors.New("URL cannot be empty")
	ErrInvalidScheme = errors.New("URL must start with http:// or https://")
	ErrMissingHost   = errors.New("URL must have a host")
	ErrInvalidPort   = errors.New("invalid port number")
)

type URLInfo struct {
	Scheme   string
	Host     string
	Port     int
	Path     string
	Query    map[string]string
	Fragment string
}

func ParseURL(rawURL string) (*URLInfo, error) {
	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return nil, ErrEmptyURL
	}

	// Parse using standard library
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, ErrInvalidScheme
	}
	if u.Host == "" {
		return nil, ErrMissingHost
	}

	host := u.Hostname()
	port := 0
	if portStr := u.Port(); portStr != "" {
		port, err = strconv.Atoi(portStr)
		if err != nil || port < 1 || port > 65535 {
			return nil, ErrInvalidPort
		}
	}

	queryMap := make(map[string]string)
	for key, values := range u.Query() {
		if len(values) > 0 {
			queryMap[key] = values[0]
		}
	}

	return &URLInfo{
		Scheme:   u.Scheme,
		Host:     host,
		Port:     port,
		Path:     u.Path,
		Query:    queryMap,
		Fragment: u.Fragment,
	}, nil
}

func BuildURL(info *URLInfo) string {
	var b strings.Builder

	b.WriteString(info.Scheme)
	b.WriteString("://")
	b.WriteString(info.Host)

	if info.Port > 0 {
		b.WriteString(":")
		b.WriteString(strconv.Itoa(info.Port))
	}

	if info.Path != "" {
		if !strings.HasPrefix(info.Path, "/") {
			b.WriteString("/")
		}
		b.WriteString(info.Path)
	}

	if len(info.Query) > 0 {
		b.WriteString("?")
		first := true
		for key, value := range info.Query {
			if !first {
				b.WriteString("&")
			}
			b.WriteString(url.QueryEscape(key))
			b.WriteString("=")
			b.WriteString(url.QueryEscape(value))
			first = false
		}
	}

	if info.Fragment != "" {
		b.WriteString("#")
		b.WriteString(info.Fragment)
	}

	return b.String()
}

func ExtractDomain(rawURL string) (string, error) {
	info, err := ParseURL(rawURL)
	if err != nil {
		return "", err
	}
	return info.Host, nil
}
