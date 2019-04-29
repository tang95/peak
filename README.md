# peak
``` bash
├── Gopkg.lock  依赖管理
├── Gopkg.toml 依赖管理
├── README.md 自述文件
├── config 配置文件及处理
├── controllers 视图
├── forms 表单验证
├── main.go 入口文件
├── models 模型
├── results 视图对象
├── routers 路由
├── services 业务逻辑
├── vendor 依赖
└── web 前端
```

``` 
# 后端开发：
dep ensure -v
go run main.go

# 前端开发：
cd web
npm install 
npm run dev

# 打包（先打包后端文件）：
cd web
npm run build

cd ..
go build

# 运行
./peak -C config/config.yaml
```


### 配置文件

```yaml
mongo:
  host: localhost
  port: 27017
  user: peak
  pass: peak
  db: peak

server:
  mode: release
  port: 8080
  host: 127.0.0.1
```