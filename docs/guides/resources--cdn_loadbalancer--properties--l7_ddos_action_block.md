---
page_title: "l7_ddos_action_block"
subcategory: "Load Balancing"
description: "l7_ddos_action_block for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1471, "body_sha256": "sha256:3523960f27785b28b55c4dbfd59f802dd3e77c40fbbebaf21b67da3ff735d609", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:l7_ddos_action_block", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:l7_ddos_action_block", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "docs/guides/resources--cdn_loadbalancer--properties--l7_ddos_action_block.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["l7_ddos_action_block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/l7_ddos_action_block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "l7_ddos_action_block for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# l7_ddos_action_block

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- l7_ddos_action_block

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

- [l7_ddos_action_block](resources--cdn_loadbalancer--properties--l7_ddos_action_block.md#section)
- [l7_ddos_action_default](resources--cdn_loadbalancer--properties--l7_ddos_action_default.md#section)
- [l7_ddos_action_js_challenge](resources--cdn_loadbalancer--properties--l7_ddos_action_js_challenge.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
l7_ddos_action_block = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
