---
page_title: "active_enhanced_firewall_policies"
subcategory: ""
description: "active_enhanced_firewall_policies for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1800, "body_sha256": "sha256:36cfe0e74938242976a2b34d7fb6ad6b8ec9184ba1208e805ec8d72290dc1d49", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:active_enhanced_firewall_policies", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:active_enhanced_firewall_policies:enhanced_firewall_policies"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "docs/guides/data-sources--securemesh_site_v2--properties--active_enhanced_firewall_policies.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["active_enhanced_firewall_policies"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_enhanced_firewall_policies for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- active_enhanced_firewall_policies

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: active\_enhanced\_firewall\_policies, no\_network\_policy; Default: no\_network\_policy\]
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

OneOf alternatives in this subsection:

- [active_enhanced_firewall_policies](data-sources--securemesh_site_v2--properties--active_enhanced_firewall_policies.md#section)
- [no_network_policy](data-sources--securemesh_site_v2--properties--no_network_policy.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [enhanced_firewall_policies](data-sources--securemesh_site_v2--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md): complete subsection reference.

## Next pages

- [active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--securemesh_site_v2--properties--active_enhanced_firewall_policies--enhanced_firewall_policies.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
