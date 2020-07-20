# Postar dockerfile
# Author: fishgoddess
# Version: v0.1.0-alpha

# Use alpine as a based image
# Notice that its shell is sh not bash
FROM alpine:3.12.0
LABEL maintainer="fishgoddess"

# The version of postar
ENV POSTAR_VERSION v0.1.0-alpha

# Deploy postar to work directory
WORKDIR /opt/postar/
COPY ./postar-$POSTAR_VERSION-linux-amd64 ./
COPY ./logit.conf ./
COPY ./postar.ini ./

# Expose two port for services
EXPOSE 5779
EXPOSE 5780

# Run postar
CMD ["./postar-$POSTAR_VERSION"]