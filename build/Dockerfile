# Postar dockerfile
# Author: fishgoddess
# Version: v0.1.1-alpha

# Use alpine as a based image
# Notice that its shell is sh not bash
FROM alpine:3.12.0
LABEL maintainer="fishgoddess"

# The version of postar
ENV POSTAR_VERSION v0.1.1-alpha
ENV POSTAR_DOWNLOAD_URL https://github.com/avino-plan/postar/releases/download/$POSTAR_VERSION/postar-$POSTAR_VERSION-linux.tar.gz

# Deploy postar to work directory
WORKDIR /opt/
RUN set -e; \
    wget $POSTAR_DOWNLOAD_URL; \
    tar -xf postar-$POSTAR_VERSION-linux.tar.gz; \
    rm postar-$POSTAR_VERSION-linux.tar.gz; \
    mv postar-$POSTAR_VERSION-linux postar-$POSTAR_VERSION;

# Add executable permission
WORKDIR /opt/postar-$POSTAR_VERSION
RUN set -e; \
    mv postar-$POSTAR_VERSION postar; \
    chmod +x ./postar

# Expose two port for services
EXPOSE 5779
EXPOSE 5780

# Run postar
CMD ["./postar"]