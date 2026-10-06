---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules"
subcategory: ""
description: "Network(L3/L4) routing policy rules."
xcsh_docs: {"aliases": ["forward proxy pbr forward proxy pbr rules"], "body_bytes": 3362, "body_sha256": "sha256:0e1fba9402c1b8fd4821a7bf1d7e5fb013d55fd4e5b9ad085fab233d45fc75ef", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "parent_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "path": "documentation/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2033231232023322-1123220021132001-1113001322220122-3322102000322103-3113211023303333-3022331000233322-1230123031120013-0020212321011122", "registry_path": "docs/guides/data-sources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules"], "schema_version": 1, "sections": [{"aliases": ["forward proxy pbr forward proxy pbr rules all destinations"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "all_destinations"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules all sources"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "all_sources"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules forwarding class list"], "anchor": "section", "description": "Ordered list of forwarding Class to be used if no rule match.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "forwarding_class_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules http list"], "anchor": "section", "description": "URLListType.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "http_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules ip prefix set"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "prefix_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy pbr forward proxy pbr rules tls list"], "anchor": "section", "description": "DomainListType.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules", "tls_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Network(L3/L4) routing policy rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr.forward_proxy_pbr_rules

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/)
- forward_proxy_pbr.forward_proxy_pbr_rules

<a id="section"></a>

Type: `"list"`. Computed.

L3/L4 routing rules. Network(L3/L4) routing policy rules.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_destinations/): complete subsection reference.

- [all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_sources/): complete subsection reference.

- [forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/forwarding_class_list/): complete subsection reference.

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/ip_prefix_set/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/label_selector/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/metadata/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/prefix_list/): complete subsection reference.

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/): complete subsection reference.
