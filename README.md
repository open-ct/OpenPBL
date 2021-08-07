# OpenPBL
System of PBL.


### 前端 react

配置文件内容参考

`web/.env`

```dotenv
REACT_APP_BASE_URL='http://localhost:5000/api'

REACT_APP_OSS_REGION='oss-cn-hangzhou'
REACT_APP_OSS_ACCESSKEYID=
REACT_APP_OSS_ACCESSKEYSECRET=
REACT_APP_OSS_BUCKET=

REACT_APP_CASDOOR_ENDPOINT=

REACT_APP_CLIENT_ID=
REACT_APP_APP_NAME=
REACT_APP_CASDOOR_ORGANIZATION='openct'

GENERATE_SOURCEMAP=false
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
dbName = openpbl_db

jwtSecret = CasdoorSecret

casdoorEndpoint =
clientId =
clientSecret =
casdoorOrganization = "openct"
```

```bash
go build main.go
./main
```