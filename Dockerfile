FROM alpine:3.12.0
MAINTAINER  fishgoddess

# 部署 postar 到工作目录
WORKDIR /postar/
COPY ./postar-v0.0.1-alpha ./
EXPOSE 5779

# 这个 entrypoint 文件的写法很有讲究，比如开头的 set -e 和结尾的 exec "$@"
# COPY ./docker-entrypoint.sh ./
# RUN chmod 755 ./docker-entrypoint.sh
# ENTRYPOINT ["./docker-entrypoint.sh"]

CMD ["./postar-v0.0.1-alpha"]