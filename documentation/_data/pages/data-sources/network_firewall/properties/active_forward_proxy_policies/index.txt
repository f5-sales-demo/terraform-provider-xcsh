---
page_title: "active_forward_proxy_policies"
subcategory: "Security"
description: "Ordered List of Forward Proxy Policies active."
xcsh_docs: {"aliases": ["active forward proxy policies"], "body_bytes": 1487, "body_sha256": "sha256:3f98a44f238c2526385e9490beff999dfcaeb289e1ceb67ec0a7c79a7503bdab", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_firewall:properties:active_forward_proxy_policies:forward_proxy_policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_firewall:properties:active_forward_proxy_policies", "parent_id": "xcsh-docs:data-sources:network_firewall:reference", "path": "documentation/data-sources/network_firewall/properties/active_forward_proxy_policies/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2233311113113112-1012303201230000-3131021231333032-0301110132312132-0000211131030031-0310103000022311-3312020012101120-3122100311021210", "registry_path": "docs/guides/data-sources--network_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["active_forward_proxy_policies"], "schema_version": 1, "sections": [{"aliases": ["active forward proxy policies forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:network_firewall:properties:active_forward_proxy_policies:forward_proxy_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["active_forward_proxy_policies", "forward_proxy_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/properties/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Ordered List of Forward Proxy Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["network_firewallCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_forward_proxy_policies

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/)
- active_forward_proxy_policies

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

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

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/active_forward_proxy_policies/#section)
- [disable_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/disable_forward_proxy_policy/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/active_forward_proxy_policies/forward_proxy_policies/): complete subsection reference.
