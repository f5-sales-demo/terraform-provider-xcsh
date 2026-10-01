---
page_title: "active_enhanced_firewall_policies"
subcategory: "Security"
description: "active_enhanced_firewall_policies for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 2314, "body_sha256": "sha256:8e2db13de6ea51520d01d8a9fec604ddb295687466ffcbdb5bcc68c157eacbda", "canonical_id": "xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "docs/guides/resources--network_firewall--properties--active_enhanced_firewall_policies.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_enhanced_firewall_policies for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_network_firewall](../resources/network_firewall.md)
- [Property reference](resources--network_firewall--reference.md)
- active_enhanced_firewall_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_enhanced\_firewall\_policies, active\_network\_policies, disable\_network\_policy;
Default: disable\_network\_policy\] List of Enhanced Firewall Policies These policies use
session-based rules and provide all OPTIONS available under firewall policies with an additional
option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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

- [active_enhanced_firewall_policies](resources--network_firewall--properties--active_enhanced_firewall_policies.md#section)
- [active_network_policies](resources--network_firewall--properties--active_network_policies.md#section)
- [disable_network_policy](resources--network_firewall--properties--disable_network_policy.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enhanced_firewall_policies](resources--network_firewall--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md): complete subsection reference.

## Next pages

- [active_enhanced_firewall_policies.enhanced_firewall_policies](resources--network_firewall--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md)
- [Property reference](resources--network_firewall--reference.md)
- [xcsh_network_firewall](../resources/network_firewall.md)
