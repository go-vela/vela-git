# SPDX-License-Identifier: Apache-2.0

######################################################################
##    docker build --no-cache --target certs -t vela-git:certs .    ##
######################################################################

FROM alpine:3.23.3@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659 as certs

RUN apk add --update --no-cache ca-certificates

#######################################################
##    docker build --no-cache -t vela-git:local .    ##
#######################################################

FROM alpine:3.23.3@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

RUN apk add --update --no-cache git git-lfs github-cli

COPY release/vela-git /bin/vela-git

ENTRYPOINT [ "/bin/vela-git" ]
