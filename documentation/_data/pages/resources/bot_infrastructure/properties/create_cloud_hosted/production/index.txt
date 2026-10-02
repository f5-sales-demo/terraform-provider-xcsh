---
page_title: "create_cloud_hosted.production"
subcategory: ""
description: "Production."
xcsh_docs: {"aliases": ["create cloud hosted production"], "body_bytes": 3238, "body_sha256": "sha256:aca2127c533817be619a32e7ed6ce1645d702afcdcd40ef7434b47d348d0669f", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "parent_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "path": "documentation/resources/bot_infrastructure/properties/create_cloud_hosted/production/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1331230011332113-1210000300323221-1201310220310031-0223300213320130-3133003023121202-1223223010120302-2302001333030011-1023203021313003", "registry_path": "docs/guides/resources--bot_infrastructure--reference--group-001.md", "relationships": [{"anchor": "schema-create_cloud_hosted--production--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}, {"anchor": "schema-create_cloud_hosted--production--region_2", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted", "production"], "schema_version": 1, "sections": [{"aliases": ["region 1"], "anchor": "schema-create_cloud_hosted--production--region_1", "description": "This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "production", "region_1"], "syntax": "attribute", "type": "string"}, {"aliases": ["region 2"], "anchor": "schema-create_cloud_hosted--production--region_2", "description": "This is an Active-Active Infrastructure configuration where traffic is routed equally between the two regions.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "production", "region_2"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/create_cloud_hosted/production/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Production.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Active-Active Infrastructure configuration where traffic is routed equally between the two regions.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/)
- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
