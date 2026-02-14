package tabledriven

import (
	"errors"
	"testing"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *URLInfo
		wantErr error
	}{
		{
			name:  "simple http URL",
			input: "http://example.com",
			want: &URLInfo{
				Scheme: "http",
				Host:   "example.com",
				Port:   0,
				Path:   "",
				Query:  map[string]string{},
			},
			wantErr: nil,
		},
		{
			name:  "https with path",
			input: "https://example.com/path/to/resource",
			want: &URLInfo{
				Scheme: "https",
				Host:   "example.com",
				Port:   0,
				Path:   "/path/to/resource",
				Query:  map[string]string{},
			},
			wantErr: nil,
		},
		{
			name:  "with port",
			input: "http://example.com:8080",
			want: &URLInfo{
				Scheme: "http",
				Host:   "example.com",
				Port:   8080,
				Path:   "",
				Query:  map[string]string{},
			},
			wantErr: nil,
		},
		{
			name:  "with query parameters",
			input: "https://example.com/search?q=golang&page=1",
			want: &URLInfo{
				Scheme: "https",
				Host:   "example.com",
				Port:   0,
				Path:   "/search",
				Query: map[string]string{
					"q":    "golang",
					"page": "1",
				},
			},
			wantErr: nil,
		},
		{
			name:  "with fragment",
			input: "https://example.com/page#section",
			want: &URLInfo{
				Scheme:   "https",
				Host:     "example.com",
				Port:     0,
				Path:     "/page",
				Query:    map[string]string{},
				Fragment: "section",
			},
			wantErr: nil,
		},
		{
			name:    "empty URL",
			input:   "",
			want:    nil,
			wantErr: ErrEmptyURL,
		},
		{
			name:    "invalid scheme",
			input:   "ftp://example.com",
			want:    nil,
			wantErr: ErrInvalidScheme,
		},
		{
			name:    "missing host",
			input:   "http://",
			want:    nil,
			wantErr: ErrMissingHost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.input)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("ParseURL(%q) error = nil; want %v", tt.input, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ParseURL(%q) error = %v; want %v", tt.input, err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseURL(%q) unexpected error: %v", tt.input, err)
			}

			if got.Scheme != tt.want.Scheme {
				t.Errorf("Scheme = %q; want %q", got.Scheme, tt.want.Scheme)
			}
			if got.Host != tt.want.Host {
				t.Errorf("Host = %q; want %q", got.Host, tt.want.Host)
			}
			if got.Port != tt.want.Port {
				t.Errorf("Port = %d; want %d", got.Port, tt.want.Port)
			}
			if got.Path != tt.want.Path {
				t.Errorf("Path = %q; want %q", got.Path, tt.want.Path)
			}
			if got.Fragment != tt.want.Fragment {
				t.Errorf("Fragment = %q; want %q", got.Fragment, tt.want.Fragment)
			}

			if len(got.Query) != len(tt.want.Query) {
				t.Errorf("Query length = %d; want %d", len(got.Query), len(tt.want.Query))
			}
			for key, wantVal := range tt.want.Query {
				if gotVal, ok := got.Query[key]; !ok || gotVal != wantVal {
					t.Errorf("Query[%q] = %q; want %q", key, gotVal, wantVal)
				}
			}
		})
	}
}
