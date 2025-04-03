#!/bin/sh
set -e

ENDPOINT="http://localhost:4566"

echo "Creating SNS topic 'topic1'..."
TOPIC_ARN=$(aws --endpoint-url=$ENDPOINT sns create-topic --name topic1 --output json | jq -r '.TopicArn')
echo "SNS Topic ARN: $TOPIC_ARN"

echo "Creating queue 'queue1'..."
QUEUE_URL=$(aws --endpoint-url=$ENDPOINT sqs create-queue --queue-name queue1 --output json | jq -r '.QueueUrl')
QUEUE_ARN=$(aws --endpoint-url=$ENDPOINT sqs get-queue-attributes --queue-url $QUEUE_URL --attribute-names QueueArn --output json | jq -r '.Attributes.QueueArn')
echo "Queue URL: $QUEUE_URL"
echo "Queue ARN: $QUEUE_ARN"

echo "Subscribing queue1 to SNS topic..."
aws --endpoint-url=$ENDPOINT sns subscribe \
  --topic-arn $TOPIC_ARN \
  --protocol sqs \
  --notification-endpoint $QUEUE_ARN

echo "Setting queue permissions..."
aws --endpoint-url=$ENDPOINT sqs set-queue-attributes \
  --queue-url $QUEUE_URL \
  --attributes '{"Policy":"{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":\"*\",\"Action\":\"sqs:SendMessage\",\"Resource\":\"*\"}]}"}'

echo "Listing SNS topics:"
aws --endpoint-url=$ENDPOINT sns list-topics

echo "Listing SQS queues:"
aws --endpoint-url=$ENDPOINT sqs list-queues

echo "Listing SNS subscriptions:"
aws --endpoint-url=$ENDPOINT sns list-subscriptions

echo "Setup complete!"
