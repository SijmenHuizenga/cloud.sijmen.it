#!/usr/bin/env bash
set -e
set -o pipefail

[ -e fdisk.layout ] && rm install-os
[ -e fdisk.layout ] && rm setup-os
[ -e fdisk.layout ] && rm fdisk.layout

wget -nv https://sijmen.it/wp-content/uploads/tmp/install-os
wget -nv https://sijmen.it/wp-content/uploads/tmp/setup-os
wget -nv https://sijmen.it/wp-content/uploads/tmp/fdisk.layout

chmod +x install-os
chmod +x setup-os

echo "Now run ./install-os to install the os"