---
page_title: "limits.disabled"
subcategory: "Security"
description: "limits.disabled for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 929, "body_sha256": "sha256:60f22263ec6732995bb4900ceda9654eb838d545427cc6597e6d879184acbdeb", "canonical_id": "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits:disabled", "parent_id": "xcsh-docs:resources:rate_limiter:properties:limits", "path": "docs/guides/resources--rate_limiter--properties--limits--disabled.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["limits", "disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "limits.disabled for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.disabled

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Property reference](resources--rate_limiter--reference.md)
- [limits](resources--rate_limiter--properties--limits.md)
- limits.disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [limits](resources--rate_limiter--properties--limits.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
