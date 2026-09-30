---
page_title: "custom_network_config.active_enhanced_firewall_policies"
subcategory: ""
description: "custom_network_config.active_enhanced_firewall_policies for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1577, "body_sha256": "sha256:03aa9be4e08bc2370947d767335ced74b5112f8305282edec91b2e77ef63e5d2", "canonical_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config", "path": "docs/guides/data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.active_enhanced_firewall_policies for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
- [Property reference](data-sources--securemesh_site--reference.md)
- [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md)
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

- [enhanced_firewall_policies](data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md): complete subsection reference.

## Next pages

- [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--securemesh_site--properties--custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies.md)
- [custom_network_config](data-sources--securemesh_site--properties--custom_network_config.md)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
