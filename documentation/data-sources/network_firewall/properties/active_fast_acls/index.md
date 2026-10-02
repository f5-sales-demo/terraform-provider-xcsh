---
page_title: "active_fast_acls"
subcategory: "Security"
description: "List of Fast ACL(s)."
xcsh_docs: {"aliases": ["active fast acls"], "body_bytes": 1785, "body_sha256": "sha256:9abbb465b12429a0eb7416f3d395b498ff942f3fe2185336a53afd4cc285f392", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_firewall:properties:active_fast_acls:fast_acls"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_firewall:properties:active_fast_acls", "parent_id": "xcsh-docs:data-sources:network_firewall:reference", "path": "documentation/data-sources/network_firewall/properties/active_fast_acls/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3110031110222221-1202323102011100-3331003003330103-0303323021130010-3102212020312231-0021130321211312-1222212111102211-3101302030000102", "registry_path": "docs/guides/data-sources--network_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["active_fast_acls"], "schema_version": 1, "sections": [{"aliases": ["fast acls"], "anchor": "section", "description": "Ordered List of Fast ACL(s) active for this network firewall.", "document_id": "xcsh-docs:data-sources:network_firewall:properties:active_fast_acls:fast_acls", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["active_fast_acls", "fast_acls"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/properties/active_fast_acls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Fast ACL(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
