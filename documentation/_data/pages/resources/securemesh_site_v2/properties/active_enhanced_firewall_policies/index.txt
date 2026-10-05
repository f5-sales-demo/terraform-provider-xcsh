---
page_title: "active_enhanced_firewall_policies"
subcategory: ""
description: "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion."
xcsh_docs: {"aliases": ["active enhanced firewall policies"], "body_bytes": 2592, "body_sha256": "sha256:689cf6f9809f33510eac8dad26ec8fff38773192b5e24fe24f8e079a46a2e2de", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:active_enhanced_firewall_policies:enhanced_firewall_policies"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/active_enhanced_firewall_policies/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3310233122203010-0332320002300320-2331123231131321-0313313323121201-1111200133001032-0002303100331000-3011120302021001-1133033311121312", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_enhanced_firewall_policies:RequiredObjectAttributes:enhanced_firewall_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_enhanced_firewall_policies"], "schema_version": 1, "sections": [{"aliases": ["active enhanced firewall policies enhanced firewall policies"], "anchor": "section", "description": "Ordered List of Enhanced Firewall Policies active.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:active_enhanced_firewall_policies:enhanced_firewall_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_enhanced_firewall_policies--enhanced_firewall_policies--name", "enforcement": "provider-schema", "group": "active_enhanced_firewall_policies.enhanced_firewall_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "schema_path": ["active_enhanced_firewall_policies", "enhanced_firewall_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- active_enhanced_firewall_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_enhanced\_firewall\_policies, no\_network\_policy; Default: no\_network\_policy\]
List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/active_enhanced_firewall_policies/#section)
- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/no_network_policy/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/active_enhanced_firewall_policies/enhanced_firewall_policies/): complete subsection reference.

## Next pages

- [active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/active_enhanced_firewall_policies/enhanced_firewall_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
