---
page_title: "other_settings.logging_options.origin_log_options"
subcategory: "Load Balancing"
description: "List of headers to Log."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "other settings logging options origin log options", "upstream servers"], "body_bytes": 2151, "body_sha256": "sha256:e00cbfbea55d67659db83217b91f3ab3aec150c1c1bacba4f06a296a5fcd3d95", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options", "path": "documentation/resources/cdn_loadbalancer/properties/other_settings/logging_options/origin_log_options/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1101322013323301-2223110001000120-2120123233101121-3222003111233313-0333012210002223-1032223320210211-1130233103030323-1200113031130213", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["other_settings", "logging_options", "origin_log_options"], "schema_version": 1, "sections": [{"aliases": ["other settings logging options origin log options header list"], "anchor": "schema-other_settings--logging_options--origin_log_options--header_list", "description": "List of headers.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["other_settings", "logging_options", "origin_log_options", "header_list"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/other_settings/logging_options/origin_log_options/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "List of headers to Log.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings.logging_options.origin_log_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [other_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/)
- [other_settings.logging_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/logging_options/)
- other_settings.logging_options.origin_log_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin log options.

Additional upstream details:

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
origin_log_options {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-other_settings--logging_options--origin_log_options--header_list"></a>

### header_list property

Type: `["list", "string"]`. Optional.

Headers. List of headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
