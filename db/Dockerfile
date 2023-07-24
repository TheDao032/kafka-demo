FROM debezium/postgres:latest

COPY postgresql.conf /usr/share/postgresql/9.6/
COPY pg_hba.conf /usr/share/postgresql/9.6/

ADD create-table.sql /docker-entrypoint-initdb.d/
