---
page_title: "origin_pool.no_tls"
subcategory: "Load Balancing"
description: "origin_pool.no_tls for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 882, "body_sha256": "sha256:318eb26c15b535a357a2db097d4ddc3022fca31b2844dd1bd25b3d34f5b2282c", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:no_tls", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool", "path": "docs/guides/resources--cdn_loadbalancer--properties--origin_pool--no_tls.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_pool", "no_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/no_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_pool.no_tls for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# origin_pool.no_tls

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [origin_pool](resources--cdn_loadbalancer--properties--origin_pool.md)
- origin_pool.no_tls

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
no_tls = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [origin_pool](resources--cdn_loadbalancer--properties--origin_pool.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
