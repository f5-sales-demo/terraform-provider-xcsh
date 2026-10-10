---
page_title: "tgw_security.active_network_policies"
subcategory: ""
description: "List of firewall policy views."
xcsh_docs: {"aliases": ["tgw security active network policies"], "body_bytes": 1108, "body_sha256": "sha256:10158bb4cbd49d9376a12c3c47fbcdd2a6fbbc85b3edd3eb8c13242f84f07da6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_network_policies:network_policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_network_policies", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security", "path": "documentation/data-sources/aws_tgw_site/properties/tgw_security/active_network_policies/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tgw_security", "active_network_policies"], "schema_version": 1, "sections": [{"aliases": ["tgw security active network policies network policies"], "anchor": "section", "description": "Ordered List of Firewall Policies active for this network firewall.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:tgw_security:active_network_policies:network_policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["tgw_security", "active_network_policies", "network_policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/tgw_security/active_network_policies/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of firewall policy views.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.active_network_policies

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [tgw_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/)
- tgw_security.active_network_policies

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Additional upstream details:

List of firewall policy views.

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

- [network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/tgw_security/active_network_policies/network_policies/): complete subsection reference.
