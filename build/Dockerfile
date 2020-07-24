# Postar dockerfile
# Author: fishgoddess
# Version: v0.1.0-alpha

# Use alpine as a based image
# Notice that its shell is sh not bash
FROM alpine:3.12.0
LABEL maintainer="fishgoddess"

# The version of postar
ENV POSTAR_VERSION v0.1.0-alpha
ENV POSTAR_PACKAGE_NAME postar-$POSTAR_VERSION-linux-amd64
ENV POSTAR_DOWNLOAD_URL https://github.com/avino-plan/postar/releases/download/$POSTAR_VERSION/$POSTAR_PACKAGE_NAME.tar.gz

# Deploy postar to work directory
WORKDIR /opt/
RUN set -e; \
    wget $POSTAR_DOWNLOAD_URL; \
    tar -xf $POSTAR_PACKAGE_NAME.tar.gz; \
    rm $POSTAR_PACKAGE_NAME.tar.gz; \
WORKDIR /opt/postar-$POSTAR_VERSION-linux

# Expose two port for services
EXPOSE 5779
EXPOSE 5780

# Run postar
CMD ["./$POSTAR_PACKAGE_NAME"]