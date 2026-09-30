---
page_title: "disable_waf"
subcategory: "Load Balancing"
description: "disable_waf for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 815, "body_sha256": "sha256:f11d111539723086a74f123482973c6064b7fab5dbcc9b66089c0e3f1678caf7", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:disable_waf", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:disable_waf", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--disable_waf.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_waf"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/disable_waf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_waf for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_waf

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- disable_waf

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
