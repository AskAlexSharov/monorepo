FROM umputun/baseimage:buildgo-latest as build-backend
#RUN apk add --no-cache gcc musl-dev
WORKDIR /src/app
COPY . .

ARG host_cache
COPY $host_cache $GOCACHE
RUN ls $GOCACHE

RUN GO111MODULE=on go mod vendor
RUN GOOS=linux GO111MODULE=on go build -o /app .

FROM github.com/grpc-ecosystem/go-grpc-middleware:app-latest
COPY --from=build-backend /app /app

RUN chown -R app:app /app

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080
EXPOSE 2339

# docker inspect --format='{{json .State.Health}}' your-container-name
HEALTHCHECK --interval=30s --timeout=3s \
    CMD curl --silent --fail http://localhost:8080/ping || exit 1

CMD ["server"]
ENTRYPOINT ["/entrypoint.sh"]
