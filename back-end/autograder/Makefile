SHELL := /bin/bash

all:
	(cd exec; apt-get update && apt-get install -y python3-pip; apt-get install python3.9-venv; python3.9 -m venv env; source env/bin/activate; pip3 install --upgrade pip; pip3 install -r requirements.txt)
