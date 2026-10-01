---
page_title: "forward_proxy_pbr.forward_proxy_pbr_rules"
subcategory: ""
description: "forward_proxy_pbr.forward_proxy_pbr_rules for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 5663, "body_sha256": "sha256:0bd4a33e0b7f9d5e4cba002731b3dcbebb721d387940ecd843c2fbc955d72505", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:metadata", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list"], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "parent_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "path": "documentation/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "forward_proxy_pbr.forward_proxy_pbr_rules for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Network(L3/L4) routing policy rules.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [forward_proxy_pbr.forward_proxy_pbr_rules.all_destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_destinations/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.all_sources](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/all_sources/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.forwarding_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/forwarding_class_list/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/http_list/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/ip_prefix_set/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/label_selector/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/metadata/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/prefix_list/)
- [forward_proxy_pbr.forward_proxy_pbr_rules.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/tls_list/)
- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
