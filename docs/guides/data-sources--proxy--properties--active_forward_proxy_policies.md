---
page_title: "active_forward_proxy_policies"
subcategory: ""
description: "active_forward_proxy_policies for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 1380, "body_sha256": "sha256:e2c2edf9efc478e5aec9136b77f9b46a465c9df45f0a44736e9fd968d109094e", "canonical_id": "xcsh-docs:data-sources:proxy:properties:active_forward_proxy_policies", "child_ids": ["xcsh-docs:data-sources:proxy:properties:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:active_forward_proxy_policies", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "docs/guides/data-sources--proxy--properties--active_forward_proxy_policies.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_forward_proxy_policies for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# active_forward_proxy_policies

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- active_forward_proxy_policies

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: active\_forward\_proxy\_policies, no\_forward\_proxy\_policy; Default:
no\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

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

- [active_forward_proxy_policies](data-sources--proxy--properties--active_forward_proxy_policies.md#section)
- [no_forward_proxy_policy](data-sources--proxy--properties--no_forward_proxy_policy.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [forward_proxy_policies](data-sources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [active_forward_proxy_policies.forward_proxy_policies](data-sources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md)
- [Property reference](data-sources--proxy--reference.md)
- [xcsh_proxy](../data-sources/proxy.md)
