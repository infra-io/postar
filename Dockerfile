# Copyright 2021 Ye Zi Jie.  All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.
#
# Postar dockerfile
# Author: fishgoddess

# Use alpine as a based image
# Notice that its shell is sh not bash
FROM alpine:3.12.0
LABEL maintainer="fishgoddess"

# The version of postar
ENV POSTAR_VERSION v0.2.0-alpha
ENV POSTAR_DOWNLOAD_URL https://github.com/avino-plan/postar/releases/download/$POSTAR_VERSION/postar-$POSTAR_VERSION.tar.gz

# Download postar
WORKDIR /opt/
RUN set -e; \
    wget $POSTAR_DOWNLOAD_URL -P postar-$POSTAR_VERSION

# Deploy and add executable permission
WORKDIR /opt/postar-$POSTAR_VERSION
RUN set -e; \
    tar -xf postar-$POSTAR_VERSION.tar.gz; \
    rm postar-$POSTAR_VERSION.tar.gz; \
    chmod +x bin/postar-$POSTAR_VERSION-linux

# Expose ports for services
EXPOSE 5897

# Run postar
CMD ["bin/postar-$POSTAR_VERSION-linux"]