---
page_title: "virtual_server.connection_rate_limit_mode.per_virtual_server_source_address"
subcategory: ""
description: "Source Address Mask."
xcsh_docs: {"aliases": ["virtual server connection rate limit mode per virtual server source address"], "body_bytes": 2332, "body_sha256": "sha256:0ffd684f31017e6f55d9117bd7947dbf604703ad5835a854bfd49c8970b61fd5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "parent_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "path": "documentation/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-2120210301030312-0201031221120200-1030032320011313-3212303321331232-1011211111003121-3001112002001300-0312113022222301-3110233212220201", "registry_path": "docs/guides/resources--application_profiles--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_source_address"], "schema_version": 1, "sections": [{"aliases": ["virtual server connection rate limit mode per virtual server source address source mask"], "anchor": "schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_address--source_mask", "description": "Configuration parameter for source mask", "document_id": "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode:per_virtual_server_source_address", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode", "per_virtual_server_source_address", "source_mask"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Source Address Mask.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["application_profilesCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/)
- [virtual_server.connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Source Address Mask.

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
per_virtual_server_source_address {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_address--source_mask"></a>

### source_mask property

Type: `"number"`. Optional.

Configuration parameter for source mask.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```
