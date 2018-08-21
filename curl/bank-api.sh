#!/usr/bin/env bash
count=0
while [ $count -le 10 ]
do
  echo "create new client"
  curl -skw "\ntime_total: %{time_total}s\n\n" -X POST 'http://[::1]:8000/client/new/100'

  echo "create new transaction"
  curl -skw "\ntime_total: %{time_total}s\n\n" --header "Content-Type: application/json" \
    --request POST \
    --data '{"from_client_id":1,"to_client_id":2, "amount":1}' \
    http://[::1]:8000/transaction

  echo "get balance"
  curl -skw "\ntime_total: %{time_total}s\n\n" -X GET 'http://[::1]:8000/client/1/balance'
done