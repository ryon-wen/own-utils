# Thanks for using this project !

English | [简体中文](https://github.com/ryon-wen/own-utils/blob/main/README_CN.md)

## Installation

Use `go get` to install SDK：

```shell
$ go get -u github.com/ryon-wen/own-utils@latest
```

## Docker Run (example)
###### Attention: `latest` tag is the image's version

#### MySQL: (`xxx` is your password)

```shell
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=xxx -d mysql:latest
```

#### Redis:

```shell
$ docker run --name redis -p 6379:6379 -d redis:latest
```

#### Nacos: 
```unsupport macOS (Apple Silicon)```
```shell
$ docker run --name nacos -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -p 9848:9848 -p 9849:9849 -d nacos/nacos-server:latest
```
```support macOS (Apple Silicon)```
```shell
$ docker run --name nacos -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -p 9848:9848 -p 9849:9849 -d nacos/nacos-server-m1:2.0.3
```

#### MinIO:(`xxx` is your local path)
```shell
$ docker run -p 9000:9000 -p 9001:9001 --name minio -d --restart=always -e "MINIO_ACCESS_KEY=minioadmin" -e "MINIO_SECRET_KEY=minioadmin" -v xxx:/data -v xxx:/root/.minio minio/minio server /data --console-address ":9001" -address ":9000"
```

#### ElasticSearch:
1. Install `elasticsearch:7.17.6` and run
```shell
$ docker run --name elasticsearch -d -e ES_JAVA_OPTS="-Xms512m -Xmx512m" -e "discovery.type=single-node" -p 9200:9200 -p 9300:9300 elasticsearch:7.17.6
```
2. Install `kibana:7.17.6` and run
```shell
$ docker run -d --name kibana -p 5601:5601 --link elasticsearch -e "ELASTICSEARCH_URL=http://127.0.0.1:9200" kibana:7.17.6
```
3. Entering the `kibana` container
```shell
$ docker exec -u 0 -it kibana /bin/bash
```
4. Install `vim` into the container
```shell
$ apt-get update
$ apt-get install vim
```
5. Edit `kibana.yml` in the container
```shell
$ vi /opt/kibana/config/kibana.yml
```
6. Press `insert` key then copy to the last row
```shell
$ i18n.locale: "zh-CN"
```
7. Press `ESC` to exit edit then input `:wq` and Press `Enter` to save the change, at last input `exit` to exit the container

8. Download [elasticsearch-analysis-ik-7.17.6.zip](https://share.feijipan.com/s/3SUofVC7)

9. Copy the downloaded file into the elasticsearch container
   (`xxx` is your file path)
```shell
$ docker cp xxx elasticsearch:/usr/share/elasticsearch/plugins
```

10. Entering the `elasticsearch` container
```shell
$ docker exec -it elasticsearch bash
```

11. Entering the `plugins` directory
```shell
$ cd plugins/
```

12. Create new folder named `ik`
```shell
$ mkdir ik
```

13. Move the `elasticsearch-analysis-ik-7.17.6.zip` file to the `ik` directory
```shell
$ mv elasticsearch-analysis-ik-7.17.6.zip ik/
```

14. Entering the `ik` directory
```shell
$ cd ik/
```

15. Unzip the zip file
```shell
$ unzip elasticsearch-analysis-ik-7.17.6.zip
```

16. Delete the zip file
```shell
$ rm -rf elasticsearch-analysis-ik-7.17.6.zip
```

17. Exit the container and restart `elasticsearch` and `kibana`
```shell
$ exit
$ docker restart elasticsearch
$ docker restart kibana
```