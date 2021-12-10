# Messaging Service
Api service to serve messaging feature. This service running in [Here](http://101.50.0.19:8005/)
## Feature
- User can send message to another user
- Users can list all messages in a conversation between them and another user.
- Users can reply to a conversation they are involved with.
- User can list all their conversations (if user A has been chatting with user C & D, the list for A will shows A-C & A-D)
- Each conversation is accompanied by unread count.
- Each conversation is accompanied by its last message
## Service Spesifications
- Golang
- MySql
## Images
```md
docker pull messaging:1.0
```
## Sonar Cloud
Check [Source code analysis](https://sonarcloud.io/project/overview?id=mhakimamransyah)

