---
page_title: "disable_api_discovery"
subcategory: "Load Balancing"
description: "disable_api_discovery for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1222, "body_sha256": "sha256:7e65379a43ed843a6630920cb0dc15f1b46db7a20765e0b6c3c45f81b983301e", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:disable_api_discovery", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:disable_api_discovery", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--disable_api_discovery.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_api_discovery"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/disable_api_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_api_discovery for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# disable_api_discovery

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- disable_api_discovery

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_api_discovery](resources--cdn_loadbalancer--properties--disable_api_discovery.md#section)
- [enable_api_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
