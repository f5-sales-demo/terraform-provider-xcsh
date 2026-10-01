---
page_title: "policy_based_challenge.always_enable_js_challenge"
subcategory: "Load Balancing"
description: "policy_based_challenge.always_enable_js_challenge for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1144, "body_sha256": "sha256:a5cef118626d24269239b1f4d857635b4fd4f643ab0247785c50d82bee33bc8a", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:always_enable_js_challenge", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:always_enable_js_challenge", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge", "path": "docs/guides/resources--cdn_loadbalancer--properties--policy_based_challenge--always_enable_js_challenge.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "always_enable_js_challenge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/always_enable_js_challenge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.always_enable_js_challenge for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.always_enable_js_challenge

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- policy_based_challenge.always_enable_js_challenge

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

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
always_enable_js_challenge = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [policy_based_challenge](resources--cdn_loadbalancer--properties--policy_based_challenge.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
