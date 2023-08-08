# Copyright (c) 2023 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

######################################################################
##    docker build --no-cache --target certs -t vela-git:certs .    ##
######################################################################

FROM alpine@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a as certs

RUN apk add --update --no-cache ca-certificates

#######################################################
##    docker build --no-cache -t vela-git:local .    ##
#######################################################

FROM alpine:latest@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

RUN apk add --update --no-cache git

COPY release/vela-git /bin/vela-git

ENTRYPOINT [ "/bin/vela-git" ]
