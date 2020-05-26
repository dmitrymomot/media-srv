FROM alpine:latest
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*
WORKDIR /
COPY ./.env /.env
COPY ./repository/migrations/* /migrations/
COPY ./media-srv /
ENV APP_PORT=8080
EXPOSE $APP_PORT
ENTRYPOINT ["/media-srv"]
