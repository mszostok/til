# Dockerfile frontend experimental syntaxes

The problem that I wanted to solve was a long docker build time for images. I had a Go application and each time I've changed the source code and wanted to rebuild the image it took so much time. The issue was that the layers were wrongly defined and there was no support for Go caching mechanisms.     

To reduce the time need for a build I've added:
- cache for `go mod download`
- cache for `go build`
- remove `COPY` and mount volume with sources instead
- reduce data copied to docker build context

I've also used the [Dockerfile frontend experimental syntaxes][1] as was describe nicely in [this][2] blog post series. 
 
Final Dockerfile:
 
```dockerfile
# syntax = docker/dockerfile:1-experimental

FROM golang:1.15.2-alpine as builder

ARG COMPONENT
ARG SOURCE_PATH="./cmd/$COMPONENT/main.go"

WORKDIR /src

# Copy the go modules files before copying and building source, so that we don't need to re-download them
# and source code changes don't invalidate our dependency layer.
COPY go.mod go.sum ./
# Use frontend syntax to cache dependencies.
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Replace `COPY . .` with `--mount=target=.` to speed up as we do not need them to persist in the final image.
# Use frontend syntax to cache go build.
# Mount the cached dependencies from the previous layer, not doing it, cause that Go build will download them once again.
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-s -w" -o /bin/$COMPONENT $SOURCE_PATH

FROM scratch
ARG COMPONENT

COPY --from=builder /bin/$COMPONENT /app

LABEL source=git@github.com:mszostok/til.git
LABEL app=$COMPONENT

CMD ["/app"]
```

Now to build an image, run:

```bash
# enable the BuildKit builder in the Docker CLI.
export DOCKER_BUILDKIT=1
# build image
docker build --build-arg COMPONENT=agent -t mszostok/agent:0.0.1 .
```

[1]: https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/experimental.md
[2]: https://www.docker.com/blog/tag/go-env-series/