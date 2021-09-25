# OpenPBL
System of PBL.

---
### 前端 react

```bash
cd web
yarn install
yarn build
```

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

#### 配置文件加载顺序

开发环境`yarn start`
```bash
.env.development -> .env
```

生产环境`yarn build`
```bash
.env.production -> .env
```

---
### 后端 beego

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
bee run -runargs -RUNMODE=prod
```
或
```bash
go build main.go
./main -RUNMODE prod
```

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
applicationName = app-openpbl
```


#### 配置文件加载顺序
开发环境下`dev`
```bash
app-dev.conf -> app.conf
```
生产环境下`prod`
```bash
app-prod.conf -> app.conf
```