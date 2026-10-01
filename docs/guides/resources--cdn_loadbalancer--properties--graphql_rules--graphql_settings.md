---
page_title: "graphql_rules.graphql_settings"
subcategory: "Load Balancing"
description: "graphql_rules.graphql_settings for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5401, "body_sha256": "sha256:3cc6c7a1d184a23b7b8e2531070784cce75780fde4bfc952d68287a4eaad1e07", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:graphql_rules", "path": "docs/guides/resources--cdn_loadbalancer--properties--graphql_rules--graphql_settings.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["graphql_rules", "graphql_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "graphql_rules.graphql_settings for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# graphql_rules.graphql_settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [graphql_rules](resources--cdn_loadbalancer--properties--graphql_rules.md)
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

- [disable_introspection](resources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--disable_introspection.md): complete subsection reference.

- [enable_introspection](resources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--enable_introspection.md): complete subsection reference.

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

- [graphql_rules.graphql_settings.disable_introspection](resources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--disable_introspection.md)
- [graphql_rules.graphql_settings.enable_introspection](resources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--enable_introspection.md)
- [graphql_rules](resources--cdn_loadbalancer--properties--graphql_rules.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
