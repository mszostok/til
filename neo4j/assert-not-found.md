## Assert not found nodes

TL;DR; The key here is this fact:

> MATCH - return nothing if it finds nothing
> OPTIONAL MATCH - always return at least null

So without `OPTIONAL` if node doesn't exist, the `MATCH` returns **no records** and Cyper query is not continued, which makes sens as we don't have "items" on which we can operate later.

In my example, I had:
```cypher
  //...trimmed...

  // Find node with `TypeInstance` label and a given `id`
  WITH *
  MATCH (backendTI:TypeInstance {id: typeInstance.backend.id})
  CREATE (ti)-[:USES]->(backendTI)

  //...trimmed...
```

For me, it was problematic as this script exists with non error, but also the relation is not created if node `TypeInstance` doesn't exit with a given `id`.
I could do an assertion on the number of executed statements, but it's more guessing that really knowing what was the issue.
Fortunately, we can change that behavior if we will use the `OPTIONAL MATCH`:

```cypher
  //...trimmed...

  // Find node with `TypeInstance` label and a given `id`
  WITH *
  OPTIONAL MATCH (backendTI:TypeInstance {id: typeInstance.backend.id})

  // Check if a given node was found
  CALL apoc.util.validate(backendTI IS NULL, apoc.convert.toJson({code: 404}), null)

  //...trimmed...
```

_source: https://community.neo4j.com/t/check-if-node-exists-and-use-the-result-as-a-condition/32069_
