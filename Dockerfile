FROM golang:latest
RUN go get github.com/HAL-RO-Developer/caseTeamB_server

ADD config.yml.template ./config.yml
EXPOSE 8000

ENTRYPOINT ["caseTeamB_server"]