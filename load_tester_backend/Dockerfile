FROM golang:1.11.3

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

COPY . /app

RUN cd /app && dep ensure && go build -o server .

RUN ./server