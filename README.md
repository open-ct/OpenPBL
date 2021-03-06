# OpenPBL
System of PBL.


### 前端 react

配置文件内容参考

`web/.env`

```dotenv
REACT_APP_CASDOOR_ENDPOINT=

REACT_APP_CLIENT_ID=
REACT_APP_APP_NAME=
REACT_APP_CASDOOR_ORGANIZATION='openct'

GENERATE_SOURCEMAP=false
SKIP_PREFLIGHT_CHECK=true
```

```bash
cd web
yarn install
yarn build
```


### 后端 beego

新建配置文件

`vim conf/app-dev.conf`

配置文件内容参考

`conf/app.conf`

```
appname = OpenPBL
httpaddr = 0.0.0.0
autorender = false
copyrequestbody = true
EnableDocs = true
SessionOn = true
copyrequestbody = true

httpport = 5000
driverName = mysql
dataSourceName = root:123@tcp(localhost:3306)/
dbName = openpbl

redisEndpoint = 
jwtSecret = CasdoorSecret

casdoorEndpoint =
clientId =
clientSecret =
casdoorOrganization = "openct"
```

开发环境
```bash
bee run -runargs -RUNMODE=dev
```
或
```bash
go build main.go
./main -RUNMODE dev
```

生产环境
```bash
go build main.go
./main -RUNMODE prod
```

运行参数
```RUNMODE```
```
dev: 加载顺序 conf/app-dev.conf  conf/app.conf
```
```
prod: 加载顺序 conf/app-prod.conf  conf/app.conf
```
```
: 加载顺序 conf/app.conf
```