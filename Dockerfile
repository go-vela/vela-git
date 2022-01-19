# Copyright (c) 2022 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

######################################################################
##    docker build --no-cache --target certs -t vela-git:certs .    ##
######################################################################

FROM alpine as certs

RUN apk add --update --no-cache ca-certificates

#######################################################
##    docker build --no-cache -t vela-git:local .    ##
#######################################################

FROM alpine:latest

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

RUN apk add --update --no-cache git

COPY release/vela-git /bin/vela-git

ENTRYPOINT [ "/bin/vela-git" ]
