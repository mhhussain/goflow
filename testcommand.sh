curl -X POST -H 'Content-type: application/json' -d '{"caller":"test-service","request":{"verb":"GET","serviceLink":"http://buckets-dev.xby2-rnd.com/buckets/drop-four/testfile0009110000000","headers":{},"body":{}}}' http://localhost:110/outbox

# {
#     caller: 'test-service',
#     request: {
#         verb: "GET",
#         serviceLink: "buckets",
#         headers: {},
#         body: {}
#     }
# }