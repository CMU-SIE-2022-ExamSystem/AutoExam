SHELL := /bin/bash

all:
	(cd exec; python3.8 -m venv env; source env/bin/activate; pip3 install --upgrade pip; pip3 install -r requirements.txt; chmod +x driver.sh)
