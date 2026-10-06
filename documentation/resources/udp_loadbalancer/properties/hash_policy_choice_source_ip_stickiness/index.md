---
page_title: "hash_policy_choice_source_ip_stickiness"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["hash policy choice source ip stickiness"], "body_bytes": 923, "body_sha256": "sha256:a28100ca1a98780b8c114549bf1ec88e54ce277234ad7eb4dca6c2fcc176e3d6", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:properties:hash_policy_choice_source_ip_stickiness", "parent_id": "xcsh-docs:resources:udp_loadbalancer:reference", "path": "documentation/resources/udp_loadbalancer/properties/hash_policy_choice_source_ip_stickiness/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1232202330301332-2233232122032300-3311123230120100-0221111330223031-0212333133112002-2320202110313302-2212102131223121-2130011100002212", "registry_path": "docs/guides/resources--udp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["hash_policy_choice_source_ip_stickiness"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/properties/hash_policy_choice_source_ip_stickiness/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hash_policy_choice_source_ip_stickiness

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/)
- hash_policy_choice_source_ip_stickiness

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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
hash_policy_choice_source_ip_stickiness = {}
```

This is an empty object or choice marker. It has no direct properties.
