# 感谢使用！

简体中文 | [English](https://github.com/ryon-wen/own-utils/blob/main/README.md)

## 安装

使用`go get`安装 SDK：

```shell
$ go get -u github.com/ryon-wen/own-utils@latest
```

## Docker 运行示例
###### 注意：`latest` 标签是镜像的版本

#### MySQL: (`xxx`是你的密码)

```shell
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=xxx -d mysql:latest
```

#### Redis:

```shell
$ docker run --name redis -p 6379:6379 -d redis:latest
```

#### Nacos:
```shell
$ docker run --name nacos -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -p 9848:9848 -p 9849:9849 -d nacos/nacos-server:latest
```
对于`macOS (Apple Silicon)`请使用下方命令
```shell
$ docker run --name nacos -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -p 9848:9848 -p 9849:9849 -d nacos/nacos-server-m1:2.0.3
```

#### MinIO:(`xxx`是映射到你本地的文件路径)
```shell
$ docker run -p 9000:9000 -p 9001:9001 --name minio -d --restart=always -e "MINIO_ACCESS_KEY=minioadmin" -e "MINIO_SECRET_KEY=minioadmin" -v xxx:/data -v xxx:/root/.minio minio/minio server /data --console-address ":9001" -address ":9000"
```

#### ElasticSearch:
1. 安装`elasticsearch:7.17.6`并运行
```shell
$ docker run --name elasticsearch -d -e ES_JAVA_OPTS="-Xms512m -Xmx512m" -e "discovery.type=single-node" -p 9200:9200 -p 9300:9300 elasticsearch:7.17.6
```
2. 安装`kibana:7.17.6`并运行
```shell
$ docker run -d --name kibana -p 5601:5601 --link elasticsearch -e "ELASTICSEARCH_URL=http://127.0.0.1:9200" kibana:7.17.6
```
3. 进入`kibana`容器
```shell
$ docker exec -u 0 -it kibana /bin/bash
```
4. 在容器中安装`vim`编辑器
```shell
$ apt-get update
$ apt-get install vim
```
5. 编辑容器中的`kibana.yml`配置文件
```shell
$ vi /opt/kibana/config/kibana.yml
```
6. 按下`insert`(mac请按`i`)键后，将下方命令拷贝到配置文件的最后一行
```shell
$ i18n.locale: "zh-CN"
```
7. 按下`ESC`键结束编辑，然后输入`:wq`再按下`Enter`键保存更改，最后输入`exit`退出容器

8. 下载 [elasticsearch-analysis-ik-7.17.6.zip](https://share.feijipan.com/s/3SUofVC7) 到本地

9. 将下载好的文件拷贝到`elasticsearch`容器中
   (`xxx`是你的文件路径)
```shell
$ docker cp xxx elasticsearch:/usr/share/elasticsearch/plugins
```

10. 进入`elasticsearch`容器
```shell
$ docker exec -it elasticsearch bash
```

11. 进入`plugins`目录
```shell
$ cd plugins/
```

12. 新建名为`ik`的文件夹
```shell
$ mkdir ik
```

13. 将`elasticsearch-analysis-ik-7.17.6.zip`文件复制到`ik`目录中
```shell
$ mv elasticsearch-analysis-ik-7.17.6.zip ik/
```

14. 进入`ik`目录
```shell
$ cd ik/
```

15. 解压压缩文件
```shell
$ unzip elasticsearch-analysis-ik-7.17.6.zip
```

16. 删除压缩文件
```shell
$ rm -rf elasticsearch-analysis-ik-7.17.6.zip
```

17. 输入`exit`退出容器，随后先重启`elasticsearch`容器，再重启`kibana`容器
```shell
$ exit
$ docker restart elasticsearch
$ docker restart kibana
```