# Adds Fern's names and pagination to the Affinity contract. Operation IDs are <resource>.<method>
# (orders.create); every `list` pages forward with startingAfter / nextCursor / data.
def pascal: [splits("[^a-zA-Z0-9]+") | select(length > 0) | (.[:1] | ascii_upcase) + .[1:]] | join("");
.paths |= with_entries(.value |= with_entries(
  if (.value | type) == "object" and (.value.operationId? // null) != null then
    (.value.operationId | split(".")) as [$group, $method]
    | .value["x-fern-sdk-group-name"] = [$group]
    | .value["x-fern-sdk-method-name"] = $method
    | if $method == "list" then
        if any(.value.parameters[]?; .name == "startingAfter") | not then
          error("\(.value.operationId) has no startingAfter cursor") else . end
        # Fern's pages only go forward: after a page fetched with endingBefore, the next page
        # would add startingAfter and keep endingBefore, which the API refuses.
        | .value.parameters |= map(select(.name != "endingBefore"))
        | .value["x-fern-pagination"] = {cursor: "$request.startingAfter", next_cursor: "$response.nextCursor", results: "$response.data"}
      else . end
  else . end))
# A sort value and its reverse (`created`, `-created`) give Fern the same name, and Fern keeps only
# one, so callers could not send the reverse. Name each `-` value `<Value>Descending`.
| .paths |= walk(
    if type == "object" and (.enum | type) == "array" and any(.enum[]; type == "string" and startswith("-")) then
      .["x-fern-enum"] = ([.enum[] | select(type == "string" and startswith("-")) | {key: ., value: {name: "\(.[1:] | pascal)Descending"}}] | from_entries)
    else . end)
