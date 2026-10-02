---
page_title: "create_cloud_hosted.testing"
subcategory: ""
description: "Testing"
xcsh_docs: {"aliases": ["create cloud hosted testing"], "body_bytes": 2002, "body_sha256": "sha256:042f03c14ad327bce591837ca5ed58335023d7cd4bbdad844bd328e7565475f6", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:testing", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted", "path": "documentation/data-sources/bot_infrastructure/properties/create_cloud_hosted/testing/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0213311311211033-0120002121021033-1011203122121200-0221322330031300-3112200022130021-1000120033110023-0220210100330020-0320231303320302", "registry_path": "docs/guides/data-sources--bot_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted", "testing"], "schema_version": 1, "sections": [{"aliases": ["region 1"], "anchor": "schema-create_cloud_hosted--testing--region_1", "description": "This is an Active-Passive Infrastructure configuration where traffic is routed to a single region.", "document_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:testing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "testing", "region_1"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/properties/create_cloud_hosted/testing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Testing", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted.testing

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/)
- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/)
- create_cloud_hosted.testing

<a id="section"></a>

Type: `"single"`. Computed.

Testing

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

<a id="schema-create_cloud_hosted--testing--region_1"></a>

### region_1 property

Type: `"string"`. Computed.

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

- [create_cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/)
- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
