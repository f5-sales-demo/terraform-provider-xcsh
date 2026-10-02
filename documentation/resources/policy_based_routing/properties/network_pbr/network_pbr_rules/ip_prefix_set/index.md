---
page_title: "network_pbr.network_pbr_rules.ip_prefix_set"
subcategory: ""
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["network pbr network pbr rules ip prefix set"], "body_bytes": 1919, "body_sha256": "sha256:31f7e444bc81d0ca0a9206194112a23459b065d13b5c0e643e17b48ab976a769", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules", "path": "documentation/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0220301111020322-0213132333011101-3030220100103223-3333332013320203-0210321310010013-2021213310222113-2311132013120122-0323322231211312", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["network_pbr", "network_pbr_rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:network_pbr:network_pbr_rules:ip_prefix_set:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["network_pbr", "network_pbr_rules", "ip_prefix_set", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.network_pbr_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/)
- [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/)
- network_pbr.network_pbr_rules.ip_prefix_set

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/): complete subsection reference.

## Next pages

- [network_pbr.network_pbr_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/ip_prefix_set/ref/)
- [network_pbr.network_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/network_pbr_rules/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
