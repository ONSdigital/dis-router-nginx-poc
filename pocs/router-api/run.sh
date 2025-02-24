#!/bin/bash

json-server "$DATABASE_FILE" -p "$PORT" & ( cd webhook || exit; npm run dev )
