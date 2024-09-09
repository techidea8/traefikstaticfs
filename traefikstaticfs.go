package traefikstaticfs

import (
	"context"
	"net/http"
	"os"
	"strings"
)

// Config the plugin configuration.
//
/*alias:
- path: /static
  dir: /data/storage
*/
type Config struct {
	Alias []string `json:"alias,omitempty" toml:"alias,omitempty" yaml:"alias,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Alias: []string{},
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
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

// 静态资源服务
func (r *StaticFs) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	os.Stdout.WriteString("(r *StaticFs) ServeHTTP" + uri)
	for _, rule := range r.config.Alias {
		// 如果有这个开头的路径
		// /mnt/123/1.png   =>  /data/storage/123/1.png
		ruleIndex := strings.Index(rule, ":")

		if strings.HasPrefix(req.RequestURI, rule[:ruleIndex]) {
			dstpath := strings.ReplaceAll(uri, rule[:ruleIndex], rule[ruleIndex+1:])
			exist, err := PathExists(dstpath)
			// 如果存在
			if exist {
				http.ServeFile(rw, req, dstpath)
				return
			} else {
				if err != nil {
					rw.Write([]byte(err.Error()))
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
	}
	r.next.ServeHTTP(rw, req)
}
