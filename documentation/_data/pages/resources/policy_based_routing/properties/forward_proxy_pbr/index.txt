---
page_title: "forward_proxy_pbr"
subcategory: ""
description: "Network(L3/L4) routing policy rule."
xcsh_docs: {"aliases": ["forward proxy pbr"], "body_bytes": 1962, "body_sha256": "sha256:fa9023f244a6f0ec5429645dffe62aec05466937a8127924a65b6a2003b56e4c", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr", "parent_id": "xcsh-docs:resources:policy_based_routing:reference", "path": "documentation/resources/policy_based_routing/properties/forward_proxy_pbr/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2220332212101100-0001123331002020-3213301113012312-0301203031100010-1323210312231213-2222133130213322-1323010233200233-2200000202030220", "registry_path": "docs/guides/resources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["forward_proxy_pbr"], "schema_version": 1, "sections": [{"aliases": ["forward proxy pbr rules"], "anchor": "section", "description": "Network(L3/L4) routing policy rules.", "document_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,http_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_destinations", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:all_sources", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,http_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:http_list,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:http_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_sources,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:all_destinations,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:ConflictingListObjectAttributes:http_list,tls_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:tls_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "forward_proxy_pbr.forward_proxy_pbr_rules:RequiredListObjectAttributes:forwarding_class_list", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules:forwarding_class_list", "type": "requires"}], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/properties/forward_proxy_pbr/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Network(L3/L4) routing policy rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- forward_proxy_pbr

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: forward\_proxy\_pbr, network\_pbr\] Configuration parameter for forward proxy pbr.

Upstream description:

Network(L3/L4) routing policy rule.

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

- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/#section)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/network_pbr/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
forward_proxy_pbr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/): complete subsection reference.

## Next pages

- [forward_proxy_pbr.forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
