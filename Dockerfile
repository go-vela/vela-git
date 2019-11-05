# Copyright (c) 2019 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

FROM alpine:latest

RUN apk add --update --no-cache ca-certificates git

COPY release/git-plugin /bin/git-plugin

ENTRYPOINT [ "/bin/git-plugin" ]
