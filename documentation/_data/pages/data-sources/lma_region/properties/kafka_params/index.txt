---
page_title: "kafka_params"
subcategory: ""
description: "Configuration parameter for kafka params."
xcsh_docs: {"aliases": ["kafka params"], "body_bytes": 933, "body_sha256": "sha256:2606c61a7b24e0f2f3ff2e0532ea3dc1f0db7fdd40cb40ac58291dd78702612e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:kafka_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/kafka_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2323001232031102-0021131322010120-1122322332011201-2001210031120222-2212022230313302-3203321020123320-2302322023212331-2231122210021230", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_params"], "schema_version": 1, "sections": [{"aliases": ["kafka params bootstrap servers"], "anchor": "schema-kafka_params--bootstrap_servers", "description": "Servers in a Kafka cluster that a client should use to bootstrap its connection to the cluster.", "document_id": "xcsh-docs:data-sources:lma_region:properties:kafka_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_params", "bootstrap_servers"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/kafka_params/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for kafka params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_params

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- kafka_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for kafka params.

## Direct properties

<a id="schema-kafka_params--bootstrap_servers"></a>

### bootstrap_servers property

Type: `["list", "string"]`. Computed.

Servers in a Kafka cluster that a client should use to bootstrap its connection to the cluster.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
