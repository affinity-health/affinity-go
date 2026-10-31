# Adds Fern's names and pagination to the Affinity contract. Operation IDs are <resource>.<method>
# (orders.create); every `list` pages forward with startingAfter / nextCursor / data.
.paths |= with_entries(.value |= with_entries(
  if (.value | type) == "object" and (.value.operationId? // null) != null then
    (.value.operationId | split(".")) as [$group, $method]
    | .value["x-fern-sdk-group-name"] = [$group]
    | .value["x-fern-sdk-method-name"] = $method
    | if $method == "list" then
        .value["x-fern-pagination"] = {cursor: "$request.startingAfter", next_cursor: "$response.nextCursor", results: "$response.data"}
      else . end
  else . end))
