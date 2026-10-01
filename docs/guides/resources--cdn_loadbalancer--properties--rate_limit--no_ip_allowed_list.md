---
page_title: "rate_limit.no_ip_allowed_list"
subcategory: "Load Balancing"
description: "rate_limit.no_ip_allowed_list for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1084, "body_sha256": "sha256:b45612c5a996efd121ed6efeacb3c07004239be6459537510b26b317cf374ef4", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit:no_ip_allowed_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:rate_limit", "path": "docs/guides/resources--cdn_loadbalancer--properties--rate_limit--no_ip_allowed_list.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rate_limit", "no_ip_allowed_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/rate_limit/no_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rate_limit.no_ip_allowed_list for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.no_ip_allowed_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
- rate_limit.no_ip_allowed_list

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_ip_allowed_list = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rate_limit](resources--cdn_loadbalancer--properties--rate_limit.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
