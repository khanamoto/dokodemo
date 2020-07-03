#!/bin/bash
set -xe

mysqladmin -uroot create dokodemo
mysqladmin -uroot create dokodemo_test

mysql -uroot dokodemo < /app/db/schema.sql
mysql -uroot dokodemo_test < /app/db/schema.sql