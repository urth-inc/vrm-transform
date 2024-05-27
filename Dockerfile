FROM golang:1.21rc2-bookworm

COPY . /app
WORKDIR /app

RUN apt update -y && \
    apt install -y libvips-dev wget && \
    wget https://github.com/KhronosGroup/KTX-Software/releases/download/v4.2.1/KTX-Software-4.2.1-Linux-x86_64.deb && \
    dpkg -i KTX-Software-4.2.1-Linux-x86_64.deb && \
    rm KTX-Software-4.2.1-Linux-x86_64.deb && \
    go mod download && \
    go build -o main main.go

CMD ["bash"]
