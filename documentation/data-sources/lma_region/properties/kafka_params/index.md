---
page_title: "kafka_params"
subcategory: ""
description: "Configuration parameter for kafka params."
xcsh_docs: {"aliases": ["kafka params"], "body_bytes": 699, "body_sha256": "sha256:e67cd719ce3b0d501799af00200c871c7e0f7bc04d63b629b5da47ebbe4a671e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:kafka_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/kafka_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2323001232031102-0021131322010120-1122322332011201-2001210031120222-2212022230313302-3203321020123320-2302322023212331-2231122210021230", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_params"], "schema_version": 1, "sections": [{"aliases": ["kafka params bootstrap servers"], "anchor": "schema-kafka_params--bootstrap_servers", "description": "Servers in a Kafka cluster that a client should use to bootstrap its connection to the cluster.", "document_id": "xcsh-docs:data-sources:lma_region:properties:kafka_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_params", "bootstrap_servers"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/kafka_params/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration parameter for kafka params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
