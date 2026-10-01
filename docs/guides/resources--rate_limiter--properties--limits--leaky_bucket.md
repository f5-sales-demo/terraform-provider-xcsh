---
page_title: "limits.leaky_bucket"
subcategory: "Security"
description: "limits.leaky_bucket for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 899, "body_sha256": "sha256:a082541d073940bc405cb8a1334c47a818a2351929e59840dd71f01f75e87846", "canonical_id": "xcsh-docs:resources:rate_limiter:properties:limits:leaky_bucket", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:properties:limits:leaky_bucket", "parent_id": "xcsh-docs:resources:rate_limiter:properties:limits", "path": "docs/guides/resources--rate_limiter--properties--limits--leaky_bucket.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["limits", "leaky_bucket"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/properties/limits/leaky_bucket/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "limits.leaky_bucket for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# limits.leaky_bucket

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Property reference](resources--rate_limiter--reference.md)
- [limits](resources--rate_limiter--properties--limits.md)
- limits.leaky_bucket

<a id="section"></a>

Type: `["object", {}]`. Optional.

Leaky-Bucket is the default rate limiter algorithm for F5.

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
leaky_bucket = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [limits](resources--rate_limiter--properties--limits.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
