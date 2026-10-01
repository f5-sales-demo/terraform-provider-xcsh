---
page_title: "other_settings.logging_options.client_log_options"
subcategory: "Load Balancing"
description: "other_settings.logging_options.client_log_options for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2291, "body_sha256": "sha256:8e799eee1168197f60c078408e258df37547e8788d9d3b12ba8ae0c5fdb83907", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options", "path": "docs/guides/resources--cdn_loadbalancer--properties--other_settings--logging_options--client_log_options.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings", "logging_options", "client_log_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/other_settings/logging_options/client_log_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings.logging_options.client_log_options for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings.logging_options.client_log_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [other_settings](resources--cdn_loadbalancer--properties--other_settings.md)
- [other_settings.logging_options](resources--cdn_loadbalancer--properties--other_settings--logging_options.md)
- other_settings.logging_options.client_log_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Headers to Log. List of headers to Log.

Upstream description:

List of headers to Log.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
client_log_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-other_settings--logging_options--client_log_options--header_list"></a>

### header_list property

Type: `["list", "string"]`. Optional.

Headers. List of headers.

Upstream description:

List of headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [other_settings.logging_options](resources--cdn_loadbalancer--properties--other_settings--logging_options.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
