---
page_title: "graphql_rules.graphql_settings"
subcategory: "Load Balancing"
description: "GraphQL configuration."
xcsh_docs: {"aliases": ["graphql rules graphql settings"], "body_bytes": 5865, "body_sha256": "sha256:578d18b6f2943efeec61c8b3141bf3288a87be354ae61f66e0eddb9d5d6f82c9", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules", "path": "documentation/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3011311003031201-0212312023210011-2111310323233011-3101023203021002-3030002113133332-2000020322030231-1120133131022312-1120132001013031", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-018.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:ConflictingObjectAttributes:disable_introspection,enable_introspection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:ConflictingObjectAttributes:disable_introspection,enable_introspection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection", "type": "conflicts"}, {"anchor": "schema-graphql_rules--graphql_settings--max_batched_queries", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:RequiredObjectAttributes:max_batched_queries,max_depth,max_total_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "type": "requires"}, {"anchor": "schema-graphql_rules--graphql_settings--max_depth", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:RequiredObjectAttributes:max_batched_queries,max_depth,max_total_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "type": "requires"}, {"anchor": "schema-graphql_rules--graphql_settings--max_total_length", "enforcement": "provider-schema", "group": "graphql_rules.graphql_settings:RequiredObjectAttributes:max_batched_queries,max_depth,max_total_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["graphql_rules", "graphql_settings"], "schema_version": 1, "sections": [{"aliases": ["disable introspection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "disable_introspection"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable introspection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "enable_introspection"], "syntax": "attribute", "type": "object"}, {"aliases": ["max batched queries"], "anchor": "schema-graphql_rules--graphql_settings--max_batched_queries", "description": "Specify maximum number of queries in a single batched request.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "max_batched_queries"], "syntax": "attribute", "type": "number"}, {"aliases": ["max depth"], "anchor": "schema-graphql_rules--graphql_settings--max_depth", "description": "Specify maximum depth for the GraphQL query.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "max_depth"], "syntax": "attribute", "type": "number"}, {"aliases": ["max total length"], "anchor": "schema-graphql_rules--graphql_settings--max_total_length", "description": "Specify maximum length in bytes for the GraphQL query.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:graphql_rules:graphql_settings", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["graphql_rules", "graphql_settings", "max_total_length"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "GraphQL configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules.graphql_settings

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [graphql_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/)
- graphql_rules.graphql_settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for graphql settings.

Upstream description:

GraphQL configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("max_batched_queries",
    "max_depth",
    "max_total_length"),
  validators.ConflictingObjectAttributes("disable_introspection",
    "enable_introspection")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

Terraform syntax:

```terraform
graphql_settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_introspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/): complete subsection reference.

- [enable_introspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/enable_introspection/): complete subsection reference.

<a id="schema-graphql_rules--graphql_settings--max_batched_queries"></a>

### max_batched_queries property

Type: `"number"`. Optional.

Specify maximum number of queries in a single batched request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="schema-graphql_rules--graphql_settings--max_depth"></a>

### max_depth property

Type: `"number"`. Optional.

Specify maximum depth for the GraphQL query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="schema-graphql_rules--graphql_settings--max_total_length"></a>

### max_total_length property

Type: `"number"`. Optional.

Specify maximum length in bytes for the GraphQL query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 16386),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

## Next pages

- [graphql_rules.graphql_settings.disable_introspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/disable_introspection/)
- [graphql_rules.graphql_settings.enable_introspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/graphql_settings/enable_introspection/)
- [graphql_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/graphql_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
