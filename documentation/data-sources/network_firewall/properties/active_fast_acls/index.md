---
page_title: "active_fast_acls"
subcategory: "Security"
description: "active_fast_acls for xcsh_network_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1785, "body_sha256": "sha256:9abbb465b12429a0eb7416f3d395b498ff942f3fe2185336a53afd4cc285f392", "child_ids": ["xcsh-docs:data-sources:network_firewall:properties:active_fast_acls:fast_acls"], "collection_id": "xcsh-docs:data-sources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_firewall:properties:active_fast_acls", "parent_id": "xcsh-docs:data-sources:network_firewall:reference", "path": "documentation/data-sources/network_firewall/properties/active_fast_acls/index.md", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["active_fast_acls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_firewall/properties/active_fast_acls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "active_fast_acls for xcsh_network_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
