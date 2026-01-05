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

set -euxo pipefail

# BEGIN OPENVIDU BLOCK
# Backup EG_* folders before cleaning tmp
mkdir -p /home/egress/backup_storage || true
if [ -d /home/egress/tmp ]; then
    find /home/egress/tmp -maxdepth 1 -type d -name 'EG_*' -exec cp -r {} /home/egress/backup_storage/ \; 2>/dev/null || true
fi
# END OPENVIDU BLOCK

# Clean out tmp
rm -rf /home/egress/tmp/*

# Start pulseaudio
rm -rf /var/run/pulse /var/lib/pulse /home/egress/.config/pulse /home/egress/.cache/xdgr/pulse
pulseaudio -D --verbose --exit-idle-time=-1 --disallow-exit

# Run egress service
exec /tini -- egress
