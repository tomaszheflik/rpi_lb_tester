FROM golang
MAINTAINER Tomasz Heflim <tomasz.heflik@sap.com>
ADD . /go/src/github.com/tomaszheflik/rpi_lb_tester/
RUN cd  /go/src/github.com/tomaszheflik/rpi_lb_tester/ && go get
RUN go install github.com/tomaszheflik/rpi_lb_tester/
ENTRYPOINT /go/bin/rpi_lb_tester
EXPOSE 9090
