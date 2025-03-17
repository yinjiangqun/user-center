# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-user-center"
LABEL REPO="https://github.com/yinjiangqun/user-center"

ENV PROJPATH=/go/src/github.com/yinjiangqun/user-center

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/yinjiangqun/user-center
WORKDIR /go/src/github.com/yinjiangqun/user-center

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/yinjiangqun/user-center"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/user-center/bin

WORKDIR /opt/user-center/bin

COPY --from=build-stage /go/src/github.com/yinjiangqun/user-center/bin/user-center /opt/user-center/bin/
RUN chmod +x /opt/user-center/bin/user-center

# Create appuser
RUN adduser -D -g '' user-center
USER user-center

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/user-center/bin/user-center"]
