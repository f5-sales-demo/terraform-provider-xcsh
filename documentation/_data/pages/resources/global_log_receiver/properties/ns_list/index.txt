---
page_title: "ns_list"
subcategory: ""
description: "Namespace List."
xcsh_docs: {"aliases": ["ns list"], "body_bytes": 1996, "body_sha256": "sha256:e27a0328f9a57cf7621fb07f2d301855c1690666d58afd53b8c7621b3c3025ea", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:ns_list", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/ns_list/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1323013202223030-2213130333303312-1001111102133131-1001220102221312-3002003012031130-3113012231221100-0102230132130332-0202133321133203", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-004.md", "relationships": [{"anchor": "schema-ns_list--namespaces", "enforcement": "provider-schema", "group": "ns_list:RequiredObjectAttributes:namespaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:ns_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ns_list"], "schema_version": 1, "sections": [{"aliases": ["ns list namespaces"], "anchor": "schema-ns_list--namespaces", "description": "List of namespaces to stream logs for.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:ns_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ns_list", "namespaces"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/ns_list/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Namespace List.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ns_list

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- ns_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Namespace List. Namespace List.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("namespaces")}
```

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
ns_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ns_list--namespaces"></a>

### namespaces property

Type: `["list", "string"]`. Optional.

Namespaces. List of namespaces to stream logs for.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```
