FROM golang


WORKDIR $PWD/src/talk-with-kiritan
ENV CGO_CFLAGS="-I/path/to/include" \
    CGO_LDFLAGS="-L/path/to/lib -lmecab -lstdc++"\
    GO111MODULE=on

RUN apt update && \
    apt upgrade -y && \
    apt install -y mecab libmecab-dev mecab-ipadic-utf8 ffmpeg

ADD . ./

EXPOSE 8080

CMD go run main.go