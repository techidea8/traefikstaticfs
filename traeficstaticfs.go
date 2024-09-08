package traefikstaticfs

import (
	"context"
	"net/http"
)

/*
alias:
	- mapper: /mnt/:e:/data
	  miss: next
	- rule: /mnt/:e:/data
	  miss: next/404
tryfile:
	root: /mnt/data
	try: index.html /mnt/e/
	miss: next/404

*/

type TryfileRule struct {
	Root string `json:"root,omitempty" toml:"root,omitempty" yaml:"root,omitempty"`
	Try  string `json:"try,omitempty" toml:"try,omitempty" yaml:"try,omitempty"`
	Miss string `json:"miss,omitempty" toml:"miss,omitempty" yaml:"miss,omitempty"`
}
type AliasRule struct {
	Mapper string `json:"mapper,omitempty" toml:"mapper,omitempty" yaml:"mapper,omitempty"`
	Miss   string `json:"miss,omitempty" toml:"miss,omitempty" yaml:"miss,omitempty"`
}

// Config the plugin configuration.
type Config struct {
	Alias   []AliasRule `json:"alias,omitempty" toml:"alias,omitempty" yaml:"alias,omitempty"`
	Tryfile TryfileRule `json:"tryfile,omitempty" toml:"tryfile,omitempty" yaml:"tryfile,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Alias: make([]AliasRule, 0),
		Tryfile: TryfileRule{
			Root: "",
		},
	}
}

// RealIPOverWriter is a plugin that blocks incoming requests depending on their source IP.
type StaticFs struct {
	config Config
	next   http.Handler
	name   string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	println("init" + name + " success")
	return &StaticFs{
		config: *config,
		next:   next,
		name:   name,
	}, nil
}

// 服务
func (r *StaticFs) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	println("(r *StaticFs) ServeHTTP=>", req.RequestURI)
	r.next.ServeHTTP(rw, req)
}
