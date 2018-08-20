#!/usr/bin/env bash

curl -X GET 'http://[::1]:8000/'

curl -X POST 'http://[::1]:8000/client/new/100'

curl -X GET 'http://[::1]:8000/client/1/balance'