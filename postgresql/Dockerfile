FROM postgres

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

ENV POSTGRES_PASSWORD=1234

EXPOSE 5432

COPY initpostgre.sql /docker-entrypoint-initdb.d/
