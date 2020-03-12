#!/bin/bash

set -e

until mysqladmin ping -h $DB_HOST -P 3306 --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 3
done

echo 'mysqld is connect !'
/server
