---
page_title: "rules.egress_rules.ip_prefix_set"
subcategory: "Security"
description: "A list of references to ip_prefix_set objects."
xcsh_docs: {"aliases": ["rules egress rules ip prefix set"], "body_bytes": 1748, "body_sha256": "sha256:9bafe12c3dd7e3b80255528230e7ca61fe2c6d9b2ce6b64693e7bc55b605a9fc", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set:ref"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set", "parent_id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules", "path": "documentation/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0132103010030311-1033233100012302-0121033213310011-2312021232331203-1311112012232203-3311122002212323-1333112021330102-3222210311012011", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "egress_rules", "ip_prefix_set"], "schema_version": 1, "sections": [{"aliases": ["rules egress rules ip prefix set ref"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:network_policy:properties:rules:egress_rules:ip_prefix_set:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "egress_rules", "ip_prefix_set", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of references to ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.ip_prefix_set

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/)
- rules.egress_rules.ip_prefix_set

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

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/): complete subsection reference.

## Next pages

- [rules.egress_rules.ip_prefix_set.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/ip_prefix_set/ref/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/rules/egress_rules/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
