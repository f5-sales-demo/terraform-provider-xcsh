---
page_title: "create_cloud_hosted.production"
subcategory: ""
description: "Production."
xcsh_docs: {"aliases": ["create cloud hosted production"], "body_bytes": 2444, "body_sha256": "sha256:4199a69f49969262b2d37eae474ffa6ff391a80127bfc4746d42c6da5bec0854", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:production", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted", "path": "documentation/data-sources/bot_infrastructure/properties/create_cloud_hosted/production/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-3103113012102212-1110021012012323-2311131123000022-0031100313022110-2213303012213310-2101002231133011-2020002333012001-0333003212110131", "registry_path": "docs/guides/data-sources--bot_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted", "production"], "schema_version": 1, "sections": [{"aliases": ["create cloud hosted production region 1"], "anchor": "schema-create_cloud_hosted--production--region_1", "description": "This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions.", "document_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:production", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "production", "region_1"], "syntax": "attribute", "type": "string"}, {"aliases": ["create cloud hosted production region 2"], "anchor": "schema-create_cloud_hosted--production--region_2", "description": "This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions.", "document_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:production", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "production", "region_2"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/properties/create_cloud_hosted/production/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Production.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted.production

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/)
- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/)
- create_cloud_hosted.production

<a id="section"></a>

Type: `"single"`. Computed.

Production.

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

<a id="schema-create_cloud_hosted--production--region_1"></a>

### region_1 property

Type: `"string"`. Computed.

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-create_cloud_hosted--production--region_2"></a>

### region_2 property

Type: `"string"`. Computed.

This is an Active-Active Infrastructure configuration where traffic is routed equally between the
two regions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-29T12:14:42+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```
