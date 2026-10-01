---
page_title: "create_cloud_hosted.testing"
subcategory: ""
description: "create_cloud_hosted.testing for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 1993, "body_sha256": "sha256:130a13938db8e16a227236c13ecf884c28808db0fc5e7e414452407f3223c92e", "canonical_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "child_ids": [], "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "parent_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "path": "docs/guides/resources--bot_infrastructure--properties--create_cloud_hosted--testing.md", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["create_cloud_hosted", "testing"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/create_cloud_hosted/testing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "create_cloud_hosted.testing for xcsh_bot_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted.testing

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md)
- [Property reference](resources--bot_infrastructure--reference.md)
- [create_cloud_hosted](resources--bot_infrastructure--properties--create_cloud_hosted.md)
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

- [create_cloud_hosted](resources--bot_infrastructure--properties--create_cloud_hosted.md)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md)
