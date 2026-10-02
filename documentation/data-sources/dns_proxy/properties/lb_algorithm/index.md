---
page_title: "lb_algorithm"
subcategory: ""
description: "Load Balancing Algorithm Type."
xcsh_docs: {"aliases": ["lb algorithm"], "body_bytes": 1312, "body_sha256": "sha256:02c0a6f3f261ee0476487f2ded2c65dbbf1b934de6c2cd70de4a8f30b329f13c", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm:round_robin"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm", "parent_id": "xcsh-docs:data-sources:dns_proxy:reference", "path": "documentation/data-sources/dns_proxy/properties/lb_algorithm/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2013211100010010-0023312102121310-3323002033212331-3103023022222212-0210130112023200-0302223111210332-0203211111130203-1233001021233013", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["lb_algorithm"], "schema_version": 1, "sections": [{"aliases": ["round robin"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm:round_robin", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["lb_algorithm", "round_robin"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/lb_algorithm/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Load Balancing Algorithm Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# lb_algorithm

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- lb_algorithm

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for lb algorithm.

Upstream description:

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

## Direct properties

- [round_robin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/lb_algorithm/round_robin/): complete subsection reference.

## Next pages

- [lb_algorithm.round_robin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/lb_algorithm/round_robin/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
