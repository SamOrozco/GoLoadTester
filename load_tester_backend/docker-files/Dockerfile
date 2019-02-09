FROM base_tester:latest
# PROJECT LOCATION
ARG base=/tester
# BINARY NAME
ARG binary=load_tester
# COPY PROJECTS
COPY . $base
# WORKDIR
WORKDIR $base/load_tester_backend
# BUILD BINARY
RUN go build -o $binary .
#
CMD ["/tester/load_tester_backend/load_tester"]

