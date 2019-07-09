FROM golang:alpine AS builder

WORKDIR $GOPATH/src/github.com/williammartin/storywriter
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /storywriter .

FROM scratch
COPY --from=builder /storywriter ./
