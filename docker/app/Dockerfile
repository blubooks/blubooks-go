# Build environment
# -----------------

FROM golang:1.19-alpine
# as build-env
ARG LOCAL_USER_ID
ENV LOCAL_USER_ID=$LOCAL_USER_ID
WORKDIR /app_blubooks
COPY . .

RUN apk update && apk add --no-cache git bash su-exec
#RUN apk update && apk add --no-cache gcc musl-dev git bash su-exec vips-dev 

COPY ./docker/app/bin/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

RUN go install github.com/codegangsta/gin@latest


#RUN go get github.com/codegangsta/gin
#CMD bash -c "go mod vendor && gin -i --appPort 4070 run main.go"
#CMD bash -c "go mod init blubooks && go mod tidy && go mod vendor &&  gin -i --appPort 4070 run main.go"
#CMD bash -c "gin  --buildArgs -buildvcs=false -i --appPort 4070 run main.go"
