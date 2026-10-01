---
page_title: "active_forward_proxy_policies"
subcategory: "Security"
description: "active_forward_proxy_policies for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1884, "body_sha256": "sha256:f9d33fd058caeec33a6c7f92811d05325a838a2c41657d5a1d4c21b0627e0a79", "canonical_id": "xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies:forward_proxy_policies"], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "docs/guides/resources--network_firewall--properties--active_forward_proxy_policies.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_forward_proxy_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_forward_proxy_policies for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_forward_proxy_policies

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md)
- [Property reference](resources--network_firewall--reference.md)
- active_forward_proxy_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
```

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

- [active_forward_proxy_policies](resources--network_firewall--properties--active_forward_proxy_policies.md#section)
- [disable_forward_proxy_policy](resources--network_firewall--properties--disable_forward_proxy_policy.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_policies](resources--network_firewall--properties--active_forward_proxy_policies--forward_proxy_policies.md): complete subsection reference.

## Next pages

- [active_forward_proxy_policies.forward_proxy_policies](resources--network_firewall--properties--active_forward_proxy_policies--forward_proxy_policies.md)
- [Property reference](resources--network_firewall--reference.md)
- [xcsh_network_firewall](../resources/network_firewall.md)
