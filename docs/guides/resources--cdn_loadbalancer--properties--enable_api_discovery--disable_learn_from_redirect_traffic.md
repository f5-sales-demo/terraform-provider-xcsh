---
page_title: "enable_api_discovery.disable_learn_from_redirect_traffic"
subcategory: "Load Balancing"
description: "enable_api_discovery.disable_learn_from_redirect_traffic for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1168, "body_sha256": "sha256:be4085af6412e72649eeb589c82ab47c7285a5cfc82e80d72fa9736a97e13037", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery", "path": "docs/guides/resources--cdn_loadbalancer--properties--enable_api_discovery--disable_learn_from_redirect_traffic.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "disable_learn_from_redirect_traffic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/disable_learn_from_redirect_traffic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.disable_learn_from_redirect_traffic for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.disable_learn_from_redirect_traffic

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [enable_api_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery.md)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

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
disable_learn_from_redirect_traffic = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [enable_api_discovery](resources--cdn_loadbalancer--properties--enable_api_discovery.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
