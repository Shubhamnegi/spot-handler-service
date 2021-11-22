# spot-handler-service
Spot interruption notice can be used to drain a node 2 min earlier then the actual interuption. This will allow gracefull handling of the services running on the node. 

The can we extended to other tasks like draing from elb and gracefull shutdown to servers aswell.  

For now we will be using kubectl to drain node

Steps:
- Register cloudwatch event for spot interruption
- Mark target as sns for notification
- Subscribe http call to service. Can depend on sqs as this might delay the request. 
- Use notice api to confirm request and process notification
- To enable kubectl on pod, allow access to context and kubectl from node  

