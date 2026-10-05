---
page_title: "client_side_defense"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Client-Side Defense Policy."
xcsh_docs: {"aliases": ["client side defense"], "body_bytes": 2025, "body_sha256": "sha256:3e29711faf9d9f2f7812c9f995b09ba3c3e352072251be6da05478f5b0ef4055", "capabilities": ["cdn", "security.client-side-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/client_side_defense/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy"], "anchor": "section", "description": "This defines various configuration OPTIONS for Client-Side Defense policy.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:disable_js_insert", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insert_all_pages_except", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages_except,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:disable_js_insert,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "client_side_defense.policy:ConflictingObjectAttributes:js_insert_all_pages_except,js_insertion_rules", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:client_side_defense:policy:js_insertion_rules", "type": "conflicts"}], "schema_path": ["client_side_defense", "policy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/client_side_defense/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This defines various configuration OPTIONS for Client-Side Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- client_side_defense

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

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

- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/#section)
- [disable_client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/disable_client_side_defense/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/): complete subsection reference.

## Next pages

- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/client_side_defense/policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
