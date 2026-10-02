---
page_title: "create_cloud_hosted.testing"
subcategory: ""
description: "Testing"
xcsh_docs: {"aliases": ["create cloud hosted testing"], "body_bytes": 2250, "body_sha256": "sha256:c2a1ba3819ceebc35846d7f6b6b3a8917194d991cfd5a6528a7557c7d269d111", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "parent_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "path": "documentation/resources/bot_infrastructure/properties/create_cloud_hosted/testing/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3033321133303122-0331120203303201-2002100100022211-3211122031301311-3321333331333303-0210211002333012-3221003021312121-1220022220231213", "registry_path": "docs/guides/resources--bot_infrastructure--reference--group-001.md", "relationships": [{"anchor": "schema-create_cloud_hosted--testing--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.testing:RequiredObjectAttributes:region_1", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted", "testing"], "schema_version": 1, "sections": [{"aliases": ["region 1"], "anchor": "schema-create_cloud_hosted--testing--region_1", "description": "This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "testing", "region_1"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/create_cloud_hosted/testing/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Testing", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted.testing

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/)
- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/)
- create_cloud_hosted.testing

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Testing

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("region_1")}
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
testing {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-create_cloud_hosted--testing--region_1"></a>

### region_1 property

Type: `"string"`. Optional.

Active-Passive Infrastructure configuration where traffic is routed to a single region.

Upstream description:

This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.

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
