curl -X POST -H 'Content-type: application/json' -d '{"caller":"test-service","request":{"verb":"GET","serviceLink":"httpbuckets","headers":{},"body":{}}}' http://localhost:110/outbox

# {
#     caller: 'test-service',
#     request: {
#         verb: "GET",
#         serviceLink: "buckets",
#         headers: {},
#         body: {}
#     }
# }