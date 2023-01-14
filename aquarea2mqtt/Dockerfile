FROM golang as builder

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/fbloemhof/hassio-addons/tree/main/aquarea2mqtt
ADD . /go/src/github.com/fbloemhof/hassio-addons/tree/main/aquarea2mqtt
RUN go get github.com/fbloemhof/hassio-addons/tree/main/aquarea2mqtt
RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /go/bin/aquarea2mqtt

FROM alpine
COPY --from=builder /go/bin/aquarea2mqtt /aquarea/aquarea2mqtt
COPY --from=builder /go/src/github.com/fbloemhof/hassio-addons/tree/main/aquarea2mqtt/translation.json /aquarea/translation.json
WORKDIR /aquarea
ENTRYPOINT ./aquarea2mqtt

# Labels
LABEL \
    io.hass.name="${BUILD_NAME}" \
    io.hass.description="${BUILD_DESCRIPTION}" \
    io.hass.arch="${BUILD_ARCH}" \
    io.hass.type="addon" \
    io.hass.version=${BUILD_VERSION}