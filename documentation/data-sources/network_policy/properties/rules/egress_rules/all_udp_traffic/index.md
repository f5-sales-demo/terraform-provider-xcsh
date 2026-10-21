---
page_title: "rules.egress_rules.all_udp_traffic"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rules egress rules all udp traffic"], "body_bytes": 1109, "body_sha256": "sha256:a30e2b527a56d4f7b9a8a133eef538e5d5a71ecb36c714085b6160d58f51ce69", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:all_udp_traffic", "parent_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "path": "documentation/data-sources/network_policy/properties/rules/egress_rules/all_udp_traffic/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0022112302302331-1133222210133011-0221030101201323-3300022220202023-2011313121211202-1202031113131011-1332302122212131-0303311132301312", "registry_path": "docs/guides/data-sources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "egress_rules", "all_udp_traffic"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/rules/egress_rules/all_udp_traffic/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["network_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.all_udp_traffic

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/)
- rules.egress_rules.all_udp_traffic

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
