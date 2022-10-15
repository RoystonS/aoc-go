#!/usr/bin/env bash

# Based on https://github.com/lucianoq/adventofcode/commit/6469f7bf0bdcb31d50658e63477f503d429ae66b

. .env
YEAR="$1"
DAY="$2"

#create dir
mkdir -p "$YEAR/$DAY"

# download input files
curl --cookie "session=${session}" "https://adventofcode.com/$YEAR/day/$DAY/input" >"$YEAR/$DAY/input" 2>/dev/null

# download assignment
curl --cookie "session=${session}" "https://adventofcode.com/$YEAR/day/$DAY" 2>/dev/null | pup 'article.day-desc' >"$YEAR/$DAY/tmp.html"
lynx -dump "$YEAR/$DAY/tmp.html" -width 80 >"$YEAR/$DAY/assignment"
rm -f "$YEAR/$DAY/tmp.html"

# add skeleton files if they don't already exist
for f in day-skeleton/*;
do
    base=$(basename "${f}")
    if [ ! -f "${YEAR}/${DAY}/${base}" ];
    then
      echo "Creating skeleton file ${YEAR}/${DAY}/${base}"
      cp -a "day-skeleton/${base}" "${YEAR}/${DAY}/"
    fi
done
