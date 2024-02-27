#!/bin/sh

# Load env from file .env
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
    echo "Environment variables loaded from .env file."
else
    echo "No .env file found."
fi

# Example usage
echo "APP_VERSION: $APP_VERSION"
echo "MYSQL_URI: $MYSQL_URI"

# Start project
./api
