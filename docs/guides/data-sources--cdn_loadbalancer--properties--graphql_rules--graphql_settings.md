---
page_title: "graphql_rules.graphql_settings"
subcategory: "Load Balancing"
description: "graphql_rules.graphql_settings for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4508, "body_sha256": "sha256:650a1fadb64c7cee09dc541ddb3af50a2e2e1fa9f5c9362eb27975bb405df5c2", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:disable_introspection", "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings:enable_introspection"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules:graphql_settings", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:graphql_rules", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--graphql_rules--graphql_settings.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["graphql_rules", "graphql_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/graphql_rules/graphql_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "graphql_rules.graphql_settings for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# graphql_rules.graphql_settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [graphql_rules](data-sources--cdn_loadbalancer--properties--graphql_rules.md)
- graphql_rules.graphql_settings

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for graphql settings.

Upstream description:

GraphQL configuration.

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

## Direct properties

- [disable_introspection](data-sources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--disable_introspection.md): complete subsection reference.

- [enable_introspection](data-sources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--enable_introspection.md): complete subsection reference.

<a id="schema-graphql_rules--graphql_settings--max_batched_queries"></a>

### max_batched_queries property

Type: `"number"`. Computed.

Specify maximum number of queries in a single batched request.

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

Type: `"number"`. Computed.

Specify maximum depth for the GraphQL query.

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

Type: `"number"`. Computed.

Specify maximum length in bytes for the GraphQL query.

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

- [graphql_rules.graphql_settings.disable_introspection](data-sources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--disable_introspection.md)
- [graphql_rules.graphql_settings.enable_introspection](data-sources--cdn_loadbalancer--properties--graphql_rules--graphql_settings--enable_introspection.md)
- [graphql_rules](data-sources--cdn_loadbalancer--properties--graphql_rules.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
