FROM umputun/baseimage:buildgo-latest as build-backend
#RUN apk add --no-cache gcc musl-dev
WORKDIR /src/app
COPY . .

RUN go test ./... && mkdir -p target && /script/coverage.sh

#RUN gometalinter --disable-all --deadline=300s --vendor --enable=vet --enable=vetshadow --enable=golint \
#    --enable=staticcheck --enable=ineffassign --enable=goconst --enable=errcheck --enable=unconvert \
#    --enable=deadcode  --enable=gosimple --enable=gas --exclude=test --exclude=mock --exclude=vendor ./...

RUN \
#    version=$(/script/git-rev.sh) && \
    echo "version $version" && \
#    GO111MODULE=on go build -o app -ldflags "-X main.revision=${version} -s -w" .
    GO111MODULE=on go build -o app -ldflags "-s -w" .

FROM umputun/baseimage:app-latest
COPY --from=build-backend /src/app/app /app

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
