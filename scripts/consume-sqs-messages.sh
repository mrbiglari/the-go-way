#!/bin/bash
set -e

QUEUE_URL="http://localhost:4566/000000000000/queue1"
ENDPOINT="http://localhost:4566"

echo "Consuming messages from SQS queue: $QUEUE_URL"

while true; do
  MESSAGES=$(aws --endpoint-url=$ENDPOINT sqs receive-message \
    --queue-url $QUEUE_URL \
    --max-number-of-messages 10 \
    --wait-time-seconds 2)

  if [[ -z "$MESSAGES" || "$MESSAGES" == "{}" ]]; then
    echo "No more messages."
    break
  fi

  echo "$MESSAGES" | jq -c '.Messages[]?' | while read -r msg; do
    BODY=$(echo "$msg" | jq -r '.Body')
    RECEIPT_HANDLE=$(echo "$msg" | jq -r '.ReceiptHandle')

    echo "Message received: $BODY"

    aws --endpoint-url=$ENDPOINT sqs delete-message \
      --queue-url $QUEUE_URL \
      --receipt-handle "$RECEIPT_HANDLE"

    echo "Message deleted."
  done
done

echo "Done consuming messages."