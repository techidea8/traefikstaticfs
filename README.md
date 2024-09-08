# traefikstaticfs
静态资源文件服务插件
#
## Configuration

### Static

```yaml
experimental:
  localPlugins:
    traefikstaticfs:
      modulename: github.com/techidea8/traefikstaticfs
```


```yaml
experimental:
  plugins:
    traefikstaticfs:
      modulename: github.com/techidea8/traefikstaticfs
      version: v0.0.1
```

### Dynamic configuration

```yaml
http:
  routers:
    my-router:
      middlewares:
        - traefikstaticfs
  middlewares:
    traefikstaticfs:
      plugin:
        traefikstaticfs:
          alias:
            - mapper: /mnt:e:/data/storage
              miss: next
            - mapper: /:e:/data/www/site
              miss: next
          tryfile:
            root: /var/www/site1
            try:  $uri $uri/ /index.html
            miss: next
          
```

#
## Configuration documentation

Supported configurations per body

| Setting           | Allowed values      | Required    | Description |
| :--               | :--                 | :--         | :--         |
| alias         | []AliasRule            | No          | 不需要鉴权的白名单 |
| tryfile         | Tryfile            | No          | 鉴权参数对应的字段名 |

AliasRule defined as follow
mapper: maper rule  like  /paterna=>/path/to/a
miss: if miss then  next /405 ,can only use 404 or next

Tryfile defined as follow
root: fs dir
try: tryfile as nginx rule try_files
miss : if miss then do next or response 404 , 404/next

#