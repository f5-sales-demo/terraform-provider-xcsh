---
page_title: "create_cloud_hosted.production"
subcategory: ""
description: "Production."
xcsh_docs: {"aliases": ["create cloud hosted production"], "body_bytes": 2747, "body_sha256": "sha256:c4ff9fc4eb2ca37fb38a7e1923e25bea84a9189fbb06aeb91517784373e71384", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "parent_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "path": "documentation/resources/bot_infrastructure/properties/create_cloud_hosted/production/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1331230011332113-1210000300323221-1201310220310031-0223300213320130-3133003023121202-1223223010120302-2302001333030011-1023203021313003", "registry_path": "docs/guides/resources--bot_infrastructure--reference--group-001.md", "relationships": [{"anchor": "schema-create_cloud_hosted--production--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}, {"anchor": "schema-create_cloud_hosted--production--region_2", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted", "production"], "schema_version": 1, "sections": [{"aliases": ["create cloud hosted production region 1"], "anchor": "schema-create_cloud_hosted--production--region_1", "description": "This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "production", "region_1"], "syntax": "attribute", "type": "string"}, {"aliases": ["create cloud hosted production region 2"], "anchor": "schema-create_cloud_hosted--production--region_2", "description": "This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "production", "region_2"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/create_cloud_hosted/production/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Production.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted.production

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/)
- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/)
- create_cloud_hosted.production

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Production.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("region_1",
    "region_2")}
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
production {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-create_cloud_hosted--production--region_1"></a>

### region_1 property

Type: `"string"`. Optional.

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

Type: `"string"`. Optional.

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
