# 寻雅后端

## 环境变量

以下环境变量皆无默认值，设置是必要的：

- `PGSQL_HOST`: PGSQL 服务器地址;
- `PGSQL_PORT`: PGSQL 服务器端口;
- `PGSQL_USER`: PGSQL 用户名;
- `PGSQL_PASSWORD`: PGSQL 密码;
- `PGSQL_DBNAME`: PGSQL 数据库名;
- `USER_TABLE_NAME`: 用户表名;
- `CONTENT_TABLE_NAME`: 内容表名;
- `IMG_PATH`: 图片存储路径;


## 安装需求

- Go 1.20

    因软件包可能还并未发布到 github 上，因此可能需要在当前目录添加 go.work，文件，内容为：
    ```txt
    go 1.20
    
    use .
    ```

- PostgresSQL

    需要有两个数据库，分别字段如下：
    
    `USER_TABLE`:
    |id|title|desc|user_id| 
    |:---:|:---:|:----:|:---:|
    |主键，`serial`|`varchar`|`varchar`|`integer`|

    `CONTENT_TABLE`:
    |id|username|password|
    |:---:|:---:|:---:|
    |主键，`serial8`|`varchar`|`varchar`|