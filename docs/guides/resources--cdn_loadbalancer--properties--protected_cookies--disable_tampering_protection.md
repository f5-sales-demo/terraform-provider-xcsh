---
page_title: "protected_cookies.disable_tampering_protection"
subcategory: "Load Balancing"
description: "protected_cookies.disable_tampering_protection for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1122, "body_sha256": "sha256:9ffa6bf06ed083ca65df695b7576aea1f1d4d7d2c57f278214da29d6253a04ce", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:protected_cookies:disable_tampering_protection", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:protected_cookies:disable_tampering_protection", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:protected_cookies", "path": "docs/guides/resources--cdn_loadbalancer--properties--protected_cookies--disable_tampering_protection.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protected_cookies", "disable_tampering_protection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/protected_cookies/disable_tampering_protection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protected_cookies.disable_tampering_protection for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protected_cookies.disable_tampering_protection

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [protected_cookies](resources--cdn_loadbalancer--properties--protected_cookies.md)
- protected_cookies.disable_tampering_protection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [protected_cookies](resources--cdn_loadbalancer--properties--protected_cookies.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
