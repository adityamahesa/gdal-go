# Linux build/run test for gdal-go using the official OSGeo GDAL image (GDAL 3.12).
#
#   docker build -t gdal-go-example .
#   docker run --rm gdal-go-example
#
FROM ghcr.io/osgeo/gdal:ubuntu-small-3.12.0

# The base image already ships the GDAL 3.12 headers (/usr/include/gdal.h) and
# the gdal.pc pkg-config file, so we must NOT install libgdal-dev (that would
# overwrite them with Ubuntu's older GDAL). We only need Go + the C toolchain
# (cgo requires gcc and pkg-config).
ARG GO_VERSION=1.23.4
RUN apt-get update && apt-get install -y --no-install-recommends \
        build-essential \
        ca-certificates \
        curl \
        pkg-config \
    && rm -rf /var/lib/apt/lists/*

RUN arch="$(dpkg --print-architecture)" \
    && curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${arch}.tar.gz" -o /tmp/go.tgz \
    && tar -C /usr/local -xzf /tmp/go.tgz \
    && rm /tmp/go.tgz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV CGO_ENABLED=1

WORKDIR /src
COPY . .

# Verify GDAL is discoverable via pkg-config, then build the example.
RUN pkg-config --modversion gdal \
    && go build -o /usr/local/bin/basic-example ./example/basic

CMD ["basic-example"]
