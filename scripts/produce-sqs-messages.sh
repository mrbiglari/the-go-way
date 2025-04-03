#!/bin/bash
set -e

TOPIC_ARN="arn:aws:sns:us-east-1:000000000000:topic1"
ENDPOINT="http://localhost:4566"

# Function to publish message to SNS
publish_message() {
  local message_id=$1

  echo "Sending message $message_id to topic1..."

  aws --endpoint-url=$ENDPOINT sns publish \
    --topic-arn $TOPIC_ARN \
    --message "Test message $message_id - $(date)"

  echo "Message $message_id sent."
  # sleep 1
}

echo "Sending messages to SNS topic: $TOPIC_ARN"

# Send 5 messages
for i in {1..5}
do
  publish_message $i
done

echo "All messages sent."
