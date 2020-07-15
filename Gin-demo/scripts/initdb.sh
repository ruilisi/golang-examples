#!/bin/sh
export PGPASSWORD=postgres 
psql -U postgres -c "create database test;"
psql -U postgres -d test -c '
CREATE TABLE "user"(
    "id" varchar(45) PRIMARY KEY,
    "email" varchar(45) NOT NULL,
    "mobile" varchar(45) NOT NULL,
    "password" varchar(45) NOT NULL
);'
