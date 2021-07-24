# OpenPBL
System of PBL.


## 开发

### 后端

新建开发配置文件

`vim conf/app-dev.conf`

配置文件内容参考

`conf/app.conf`

### 前端

新建开发配置文件

`vim openpbl-landing/.env.development`

配置文件内容参考

`openpbl-landing/.env`


## 部署

### 后端 beego

新建部署配置文件

`vim conf/app-prod.conf`

配置文件内容参考

`conf/app.conf`

```
appname = OpenPBL
httpaddr = 127.0.0.1
autorender = false
copyrequestbody = true
EnableDocs = true
SessionOn = true
copyrequestbody = true

httpport = 5000
driverName = mysql
dataSourceName = root:root@tcp(db:3306)/            # docker-compose 部署环境下 localhost 改为 db
dbName = openpbl_db

casdoorEndpoint = http://localhost:8000
clientId =                                          # casdoor 应用 id
clientSecret =                                      # casdoor 应用 secret
jwtSecret =                                         # jwt 加密密钥
casdoorOrganization = "openct"                      # casdoor 应用所属组织
```

### 前端 react

新建部署配置文件

`vim openpbl-landing/.env.production`

配置文件内容参考

`openpbl-landing/.env`

```dotenv
REACT_APP_BASE_URL='http://localhost:5000/api'
REACT_APP_OSS_REGION: 'oss-cn-hangzhou'             # 阿里云 oss region
REACT_APP_OSS_ACCESSKEYID: '123'                    # oss accesskeyId
REACT_APP_OSS_ACCESSKEYSECRET: '123'                # oss accessKeySecret
REACT_APP_OSS_BUCKET: 'bucket'                      # oss bucket
```