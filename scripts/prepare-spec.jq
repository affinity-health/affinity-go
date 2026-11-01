# Adds Fern's names and pagination to the Affinity contract. Operation IDs are <resource>.<method>
# (orders.create); every `list` pages forward with startingAfter / nextCursor / data
# (API 2026-11-01 lists page forward only, which Fern's pager fits).
.paths |= with_entries(.value |= with_entries(
  if (.value | type) == "object" and (.value.operationId? // null) != null then
    (.value.operationId | split(".")) as [$group, $method]
    | .value["x-fern-sdk-group-name"] = [$group]
    | .value["x-fern-sdk-method-name"] = $method
    | if $method == "list" then
        if any(.value.parameters[]?; .name == "startingAfter") | not then
          error("\(.value.operationId) has no startingAfter cursor") else . end
        | if any(.value.parameters[]?; .name == "endingBefore") then
            error("\(.value.operationId) pages backward; Fern's pager pages forward only") else . end
        | .value["x-fern-pagination"] = {cursor: "$request.startingAfter", next_cursor: "$response.nextCursor", results: "$response.data"}
      else . end
  else . end))
