# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

FROM alpine:latest

RUN apk add --update --no-cache ca-certificates git

COPY release/vela-git /bin/vela-git

ENTRYPOINT [ "/bin/vela-git" ]
