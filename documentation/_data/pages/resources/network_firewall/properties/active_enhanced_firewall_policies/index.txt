---
page_title: "active_enhanced_firewall_policies"
subcategory: "Security"
description: "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion."
xcsh_docs: {"aliases": ["active enhanced firewall policies"], "body_bytes": 2775, "body_sha256": "sha256:1317608d61289640b4d95f1eb75903808908b8dfbb346f71ac35ec3c4ace337d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies:enhanced_firewall_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "documentation/resources/network_firewall/properties/active_enhanced_firewall_policies/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2313013321102121-1013201323302331-2222111323201110-2011320113312130-1311100031233101-2112200023113230-2233122231303002-3003213211332212", "registry_path": "docs/guides/resources--network_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_enhanced_firewall_policies:RequiredObjectAttributes:enhanced_firewall_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_enhanced_firewall_policies"], "schema_version": 1, "sections": [{"aliases": ["enhanced firewall policies"], "anchor": "section", "description": "Ordered List of Enhanced Firewall Policies active.", "document_id": "xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies:enhanced_firewall_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_enhanced_firewall_policies--enhanced_firewall_policies--name", "enforcement": "provider-schema", "group": "active_enhanced_firewall_policies.enhanced_firewall_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "schema_path": ["active_enhanced_firewall_policies", "enhanced_firewall_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_enhanced_firewall_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_enhanced_firewall_policies

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- active_enhanced_firewall_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_enhanced\_firewall\_policies, active\_network\_policies, disable\_network\_policy;
Default: disable\_network\_policy\] List of Enhanced Firewall Policies These policies use
session-based rules and provide all OPTIONS available under firewall policies with an additional
option for service insertion.

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

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_enhanced_firewall_policies/#section)
- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_network_policies/#section)
- [disable_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/disable_network_policy/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_enhanced_firewall_policies/enhanced_firewall_policies/): complete subsection reference.

## Next pages

- [active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_enhanced_firewall_policies/enhanced_firewall_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
