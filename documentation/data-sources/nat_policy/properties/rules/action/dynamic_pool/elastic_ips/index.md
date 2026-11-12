---
page_title: "rules.action.dynamic_pool.elastic_ips"
subcategory: ""
description: "List of references to Cloud Elastic IP Object."
xcsh_docs: {"aliases": ["rules action dynamic pool elastic ips"], "body_bytes": 1092, "body_sha256": "sha256:970d417f180147b47fb33a7c8242cc959282a133b29b845feb2ad2f1643e5984", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool", "path": "documentation/data-sources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-2231231022310311-3123220132313302-2202122002213221-2331113010211202-0332331131231212-1113231210301003-0303332133130231-3232122213311201", "registry_path": "docs/guides/data-sources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips"], "schema_version": 1, "sections": [{"aliases": ["rules action dynamic pool elastic ips refs"], "anchor": "section", "description": "Reference to one or more cloud elastic IP objects.", "document_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic_pool:elastic_ips:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rules", "action", "dynamic_pool", "elastic_ips", "refs"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "List of references to Cloud Elastic IP Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["nat_policyCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic_pool.elastic_ips

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/)
- [rules.action.dynamic_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic_pool/)
- rules.action.dynamic_pool.elastic_ips

<a id="section"></a>

Type: `"single"`. Computed.

List of references to Cloud Elastic IP Object.

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic_pool/elastic_ips/refs/): complete subsection reference.
