FROM golang:1.11.3
# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
#
ARG go_src=$GOPATH/src
ARG binary=server
# creating a backend dir and copying to the dir so our imports
# will work correctly
ARG back_end=$GOPATH/src/load_tester/load_tester_backend
COPY . $back_end
#build our golang binary
RUN cd $back_end && dep ensure && go build -o $binary .

ARG command=$back_end/$binary
RUN chmod +x $command
RUN echo $command
CMD $command

