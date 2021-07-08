# Bot.go Project
# Copyright (C) 2021 Sayan Biswas, ALiwoto
# This file is subject to the terms and conditions defined in
# file 'LICENSE', which is part of the source code.

#!/bin/sh

clear

echo -e "building it, please wait a bit..."

go build -o ltw-server

clear

echo -e "done building it, now running...\n-------------------"

./ltw-server