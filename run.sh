#
# This file is part of LTW-server project (https://github.com/intellivoid/LTW-server).
# Copyright (c) 2021 WotoTeam.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, version 2.
#
# This program is distributed in the hope that it will be useful, but
# WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
# General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.
#


buildApp() 
{
	# clear the screen (the terminal)
	clear

	echo -e "bulding it, please wait a bit..."

	go build -o ltw-server
}

runApp()
{
	# clear the screen (the terminal)
	clear

	echo -e "we are done building it,\n->now running the server...\n-------------------"

	./ltw-server
}

operations=0

if [ -z "$1" ] || [ "$1" == "true" ] || [ "$1" == "1" ];
then
	buildApp;
	operations=$((i+1))
fi;

if [ -z "$2" ] || [ "$2" == "true" ] || [ "$2" == "1" ];
then
	runApp;
	operations=$((i+1))
fi;

if [ $operations == 0 ]
then
	echo "You have done nothing!"
fi;