---
page_title: "active_fast_acls"
subcategory: "Security"
description: "List of Fast ACL(s)."
xcsh_docs: {"aliases": ["active fast acls"], "body_bytes": 2034, "body_sha256": "sha256:303c13b02362b2be29428b83d0255e43a3d5f86b08701cae2d412591d7507e29", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "documentation/resources/network_firewall/properties/active_fast_acls/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3112210130022303-2030213233130033-2121231033222300-2033311300122320-0021030220003103-2012033313212000-3303200320201022-1230312021203230", "registry_path": "docs/guides/resources--network_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_fast_acls:RequiredObjectAttributes:fast_acls", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_fast_acls"], "schema_version": 1, "sections": [{"aliases": ["fast acls"], "anchor": "section", "description": "Ordered List of Fast ACL(s) active for this network firewall.", "document_id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_fast_acls--fast_acls--name", "enforcement": "provider-schema", "group": "active_fast_acls.fast_acls:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_fast_acls:fast_acls", "type": "requires"}], "schema_path": ["active_fast_acls", "fast_acls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_fast_acls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Fast ACL(s).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_fast_acls

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- active_fast_acls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_fast\_acls, disable\_fast\_acl; Default: disable\_fast\_acl\] Configuration
parameter for active fast acls.

Upstream description:

List of Fast ACL(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fast_acls")}
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

- [active_fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/#section)
- [disable_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/disable_fast_acl/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_fast_acls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/fast_acls/): complete subsection reference.

## Next pages

- [active_fast_acls.fast_acls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_fast_acls/fast_acls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
