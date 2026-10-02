---
page_title: "active_forward_proxy_policies"
subcategory: "Security"
description: "Ordered List of Forward Proxy Policies active."
xcsh_docs: {"aliases": ["active forward proxy policies"], "body_bytes": 2294, "body_sha256": "sha256:e6ed35189c6fa76e53672ac3197e651c7ac66829739d4fc5616aa90ce686e906", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies:forward_proxy_policies"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies", "parent_id": "xcsh-docs:resources:network_firewall:reference", "path": "documentation/resources/network_firewall/properties/active_forward_proxy_policies/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0102101002112321-0202233211212032-1112101221133102-3312333102301221-0010220230112111-0223201103230311-0103311201122130-3010200022203100", "registry_path": "docs/guides/resources--network_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "active_forward_proxy_policies:RequiredObjectAttributes:forward_proxy_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["active_forward_proxy_policies"], "schema_version": 1, "sections": [{"aliases": ["forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies:forward_proxy_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-active_forward_proxy_policies--forward_proxy_policies--name", "enforcement": "provider-schema", "group": "active_forward_proxy_policies.forward_proxy_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:network_firewall:properties:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "schema_path": ["active_forward_proxy_policies", "forward_proxy_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/properties/active_forward_proxy_policies/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ordered List of Forward Proxy Policies active.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# active_forward_proxy_policies

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- active_forward_proxy_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_forward\_proxy\_policies, disable\_forward\_proxy\_policy; Default:
disable\_forward\_proxy\_policy\] Ordered List of Forward Proxy Policies active.

Upstream description:

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_forward_proxy_policies/#section)
- [disable_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/disable_forward_proxy_policy/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_forward_proxy_policies/forward_proxy_policies/): complete subsection reference.

## Next pages

- [active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/active_forward_proxy_policies/forward_proxy_policies/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/properties/)
- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
