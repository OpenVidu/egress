#!/usr/bin/env bash
# Copyright 2023 LiveKit, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

# BEGIN OPENVIDU BLOCK
# Backup EG_* folders before cleaning tmp
mkdir -p /home/egress/backup_storage || true
if [ -d /home/egress/tmp ]; then
    find /home/egress/tmp -maxdepth 1 -type d -name 'EG_*' -exec cp -r {} /home/egress/backup_storage/ \; 2>/dev/null || true
fi

# Read chrome into the page cache in the background, at the lowest CPU and I/O priority. Its first
# launch then reads ~10 MB from disk instead of ~200 MB, which on a slow or busy disk takes ~20 s and
# delays the first web egress. It waits 5 s so that it does not slow down the egress's own startup,
# and the subshell leaves it to tini, which reaps it.
if [ -d /opt/google/chrome ]; then
    (sleep 5 && nice -n 19 ionice -c3 find /opt/google/chrome -type f -exec cat {} + > /dev/null 2>&1 &)
fi
# END OPENVIDU BLOCK

# Clean out tmp
rm -rf /home/egress/tmp/*

# Adjust max open file
ulimit -n 65536 || true

# Start pulseaudio
rm -rf /var/run/pulse /var/lib/pulse /home/egress/.config/pulse /home/egress/.cache/xdgr/pulse
pulseaudio -D --verbose --exit-idle-time=-1 --disallow-exit > /dev/null 2>&1

# Run egress service
exec egress
