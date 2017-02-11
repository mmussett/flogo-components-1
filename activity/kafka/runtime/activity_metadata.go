package kafka

var jsonMetadata = `{
  "name": "kafka",
  "version": "0.0.1",
  "title": "Send to Kafka",
  "description": "Send message to Kafka",
  "homepage": "https://github.com/TIBCOSoftware/jvanderl/activity/kafka",
  "inputs":[
    {
      "name": "server",
      "type": "string"
    },
    {
      "name": "configid",
      "type": "string"
    },
    {
      "name": "topic",
      "type": "string"
    },
    {
      "name": "partition",
      "type": "int"
    },
    {
      "name": "message",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}`
