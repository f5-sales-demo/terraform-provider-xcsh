---
page_title: "custom_network_config.active_enhanced_firewall_policies"
subcategory: ""
description: "custom_network_config.active_enhanced_firewall_policies for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1667, "body_sha256": "sha256:cb38eeb52bf57e3ba2a10c4b43f947ab387b2d24088d974b6dab620a05309992", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config", "path": "docs/guides/data-sources--voltstack_site--properties--custom_network_config--active_enhanced_firewall_policies.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.active_enhanced_firewall_policies for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- custom_network_config.active_enhanced_firewall_policies

<a id="section"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

## Direct properties

- [enhanced_firewall_policies](data-sources--voltstack_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md): complete subsection reference.

## Next pages

- [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--voltstack_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md)
- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
