---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules.tls_list"
subcategory: ""
description: "DomainListType."
xcsh_docs: {"aliases": ["forward proxy pbr forward proxy pbr rules tls list"], "body_bytes": 1925, "body_sha256": "sha256:1e35dc7df314aab4a0927e89179b7393ec64c5d1b08b69b65f8b929ae03d8846", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "path": "documentation/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2302000223132332-0120121212222231-3102303013320021-1233103202231113-2200123132102030-2110101311030132-0331111101023210-1311133032213030", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "tls_list"], "schema_version": 1, "sections": [{"aliases": ["forward proxy pbr forward proxy pbr rules tls list tls list"], "anchor": "section", "description": "Domains in SNI for TLS connections.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--exact_value", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list:ConflictingListObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list", "type": "conflicts"}, {"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--exact_value", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list", "type": "conflicts"}, {"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--regex_value", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list:ConflictingListObjectAttributes:exact_value,regex_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list", "type": "conflicts"}, {"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--regex_value", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list:ConflictingListObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list", "type": "conflicts"}, {"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--suffix_value", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list:ConflictingListObjectAttributes:exact_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list", "type": "conflicts"}, {"anchor": "schema-forward_proxy_pbr--forward_proxy_pbr_rules--tls_list--tls_list--suffix_value", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list:ConflictingListObjectAttributes:regex_value,suffix_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list:tls_list", "type": "conflicts"}], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "tls_list", "tls_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "DomainListType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr.forward_proxy_pbr_rules.tls_list

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/)
- [forward_proxy_pbr.forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/)
- forward_proxy_pbr.forward_proxy_pbr_rules.tls_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DomainListType.

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

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/tls_list/): complete subsection reference.

## Next pages

- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/tls_list/)
- [forward_proxy_pbr.forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
