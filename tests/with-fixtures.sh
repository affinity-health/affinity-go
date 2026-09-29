#!/bin/sh
set -eu
sdk_tests_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
sdk_fixture_source=${SDK_FIXTURE_DIRECTORY:-$sdk_tests_dir}
python3 "$sdk_fixture_source/fixture-server.py" > /tmp/affinity-sdk-fixtures.log 2>&1 &
sdk_fixture_pid=$!
trap 'kill "$sdk_fixture_pid" 2>/dev/null || true' EXIT INT TERM
python3 - "$sdk_fixture_pid" <<'PY'
import os,sys,time,urllib.request
for attempt in range(50):
    os.kill(int(sys.argv[1]),0)
    try:
        urllib.request.urlopen('http://127.0.0.1:5199/healthz',timeout=1).read()
        break
    except OSError:
        time.sleep(.1)
else:
    raise SystemExit('SDK fixture server did not start')
PY
"$@"
