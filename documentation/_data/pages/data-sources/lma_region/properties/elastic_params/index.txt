---
page_title: "elastic_params"
subcategory: ""
description: "Configuration parameter for elastic params."
xcsh_docs: {"aliases": ["elastic params"], "body_bytes": 626, "body_sha256": "sha256:0079ea1e0ce0f4d02b8d1ecf15da1581d3a0a6b497cc723838925f6be592a513", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:elastic_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/elastic_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-3023202022122021-0232013303221013-1122222121321213-0110223203021121-3001122003332133-2302022120202212-3001202320300322-1233131201220202", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["elastic_params"], "schema_version": 1, "sections": [{"aliases": ["elastic params urls"], "anchor": "schema-elastic_params--urls", "description": "Elastic Search URLs. Elastic Search URL.", "document_id": "xcsh-docs:data-sources:lma_region:properties:elastic_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["elastic_params", "urls"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/elastic_params/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configuration parameter for elastic params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# elastic_params

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- elastic_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for elastic params.

## Direct properties

<a id="schema-elastic_params--urls"></a>

### urls property

Type: `["list", "string"]`. Computed.

Elastic Search URLs. Elastic Search URL.
