---
page_title: "elastic_params"
subcategory: ""
description: "Configuration parameter for elastic params."
xcsh_docs: {"aliases": ["elastic params"], "body_bytes": 626, "body_sha256": "sha256:0079ea1e0ce0f4d02b8d1ecf15da1581d3a0a6b497cc723838925f6be592a513", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:elastic_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/elastic_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:1b9ea59a790d92083efb68ff1abebe57ebbb3f7b716b91cdb7b0416134676174", "provider_type": "data-sources", "registry_anchor": "canonical-3023202022122021-0232013303221013-1122222121321213-0110223203021121-3001122003332133-2302022120202212-3001202320300322-1233131201220202", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["elastic_params"], "schema_version": 1, "sections": [{"aliases": ["elastic params urls"], "anchor": "schema-elastic_params--urls", "description": "Elastic Search URLs. Elastic Search URL.", "document_id": "xcsh-docs:data-sources:lma_region:properties:elastic_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["elastic_params", "urls"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/elastic_params/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Configuration parameter for elastic params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": [], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
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
