FROM golang:latest
RUN go get github.com/HAL-RO-Developer/caseTeamB_server

EXPOSE 8080

ENTRYPOINT ["caseTeamB_server"]