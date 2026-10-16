---
page_title: "clickhouse_params"
subcategory: ""
description: "Configuration parameter for clickhouse params."
xcsh_docs: {"aliases": ["clickhouse params"], "body_bytes": 1043, "body_sha256": "sha256:37c7a9b8a3c06dc0ae43d52e58a7132d57372e2ee9a7937d383877dbf5ce3e29", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/clickhouse_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3202202033210222-0021013013103132-3120111201112300-2011201302010333-3130120331133230-3313021303011003-0023112010023002-3221302133332001", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["clickhouse_params"], "schema_version": 1, "sections": [{"aliases": ["clickhouse params host"], "anchor": "schema-clickhouse_params--host", "description": "Clickhouse Host. Clickhouse Host.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "host"], "syntax": "attribute", "type": "string"}, {"aliases": ["clickhouse params password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["clickhouse_params", "password"], "syntax": "attribute", "type": "object"}, {"aliases": ["clickhouse params port"], "anchor": "schema-clickhouse_params--port", "description": "Clickhouse Port. Clickhouse Port.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["clickhouse params user"], "anchor": "schema-clickhouse_params--user", "description": "Clickhouse User. Clickhouse User.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "user"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/clickhouse_params/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Configuration parameter for clickhouse params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# clickhouse_params

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- clickhouse_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for clickhouse params.

## Direct properties

<a id="schema-clickhouse_params--host"></a>

### host property

Type: `"string"`. Computed.

Clickhouse Host. Clickhouse Host.

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/): complete subsection reference.

<a id="schema-clickhouse_params--port"></a>

### port property

Type: `"number"`. Computed.

Clickhouse Port. Clickhouse Port.

<a id="schema-clickhouse_params--user"></a>

### user property

Type: `"string"`. Computed.

Clickhouse User. Clickhouse User.
