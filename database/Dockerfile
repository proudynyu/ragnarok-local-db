FROM alpine:latest

RUN apk --no-cache add sqlite

RUN mkdir /database

WORKDIR /database

COPY database.db /database

EXPOSE 7777

CMD ["sqlite3", "database.db"]
