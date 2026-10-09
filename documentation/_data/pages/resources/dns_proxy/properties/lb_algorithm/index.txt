---
page_title: "lb_algorithm"
subcategory: ""
description: "Load Balancing Algorithm Type."
xcsh_docs: {"aliases": ["lb algorithm"], "body_bytes": 1054, "body_sha256": "sha256:1d621fb07b119ed8fbf8f69ef20d6cc906b120ee44ed695012ee5ae8a19ae63c", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:lb_algorithm:round_robin"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm", "parent_id": "xcsh-docs:resources:dns_proxy:reference", "path": "documentation/resources/dns_proxy/properties/lb_algorithm/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0123302030312300-1300310310103200-2010132201320120-3002100002202030-2300121210113102-3033032012230111-1301131132112112-3013012210331233", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["lb_algorithm"], "schema_version": 1, "sections": [{"aliases": ["lb algorithm round robin"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_proxy:properties:lb_algorithm:round_robin", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["lb_algorithm", "round_robin"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/lb_algorithm/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Load Balancing Algorithm Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# lb_algorithm

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- lb_algorithm

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for lb algorithm.

Additional upstream details:

Load Balancing Algorithm Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lb_algorithm_choice": "[\"round_robin\"]"
}
```

Terraform syntax:

```terraform
lb_algorithm {
  # Configure direct properties listed below.
}
```

## Direct properties

- [round_robin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/lb_algorithm/round_robin/): complete subsection reference.
