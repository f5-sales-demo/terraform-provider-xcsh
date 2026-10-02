---
page_title: "active_fast_acls"
subcategory: "Security"
description: "List of Fast ACL(s)."
xcsh_docs: {"aliases": ["active fast acls"], "body_bytes": 1785, "body_sha256": "sha256:9abbb465b12429a0eb7416f3d395b498ff942f3fe2185336a53afd4cc285f392", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_firewall:properties:active_fast_acls:fast_acls"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_firewall:properties:active_fast_acls", "parent_id": "xcsh-docs:data-sources:network_firewall:reference", "path": "documentation/data-sources/network_firewall/properties/active_fast_acls/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102", "registry_path": "docs/guides/data-sources--network_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["active_fast_acls"], "schema_version": 1, "sections": [{"aliases": ["fast acls"], "anchor": "section", "description": "Ordered List of Fast ACL(s) active for this network firewall.", "document_id": "xcsh-docs:data-sources:network_firewall:properties:active_fast_acls:fast_acls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["active_fast_acls", "fast_acls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/properties/active_fast_acls/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Fast ACL(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_fast_acls

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/)
- active_fast_acls

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Upstream description:

List of Fast ACL(s).

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

- [active_fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/active_fast_acls/#section)
- [disable_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/disable_fast_acl/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/active_fast_acls/fast_acls/): complete subsection reference.

## Next pages

- [active_fast_acls.fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/active_fast_acls/fast_acls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/properties/)
- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_firewall/)
