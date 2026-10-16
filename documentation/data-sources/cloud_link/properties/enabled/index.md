---
page_title: "enabled"
subcategory: ""
description: "CloudLink ADN Network Config."
xcsh_docs: {"aliases": ["enabled"], "body_bytes": 1708, "body_sha256": "sha256:d6aecf043e3b7d9773353b7e65141267a25b06c5e63b6379f8f5ffb8edfeb765", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:properties:enabled", "parent_id": "xcsh-docs:data-sources:cloud_link:reference", "path": "documentation/data-sources/cloud_link/properties/enabled/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2133211113302331-1111103313201201-2300313320232302-1001210232102002-3322223320212110-1202033223333021-1002113012122212-3121000132101120", "registry_path": "docs/guides/data-sources--cloud_link--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enabled"], "schema_version": 1, "sections": [{"aliases": ["enabled cloudlink network name"], "anchor": "schema-enabled--cloudlink_network_name", "description": "Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN network. To provision a Private ADN network, please contact F5 Distributed Cloud support.", "document_id": "xcsh-docs:data-sources:cloud_link:properties:enabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled", "cloudlink_network_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/properties/enabled/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "CloudLink ADN Network Config.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enabled

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- enabled

<a id="section"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

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

## Direct properties

<a id="schema-enabled--cloudlink_network_name"></a>

### cloudlink_network_name property

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```
