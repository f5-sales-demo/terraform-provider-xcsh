---
page_title: "clickhouse_params"
subcategory: ""
description: "Configuration parameter for clickhouse params."
xcsh_docs: {"aliases": ["clickhouse params"], "body_bytes": 1428, "body_sha256": "sha256:3d8763e7ee657a2fd312d823c876f5b1eaab319439a2b0c35057da69dbdb7570", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/clickhouse_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3202202033210222-0021013013103132-3120111201112300-2011201302010333-3130120331133230-3313021303011003-0023112010023002-3221302133332001", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["clickhouse_params"], "schema_version": 1, "sections": [{"aliases": ["clickhouse params host"], "anchor": "schema-clickhouse_params--host", "description": "Clickhouse Host. Clickhouse Host.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "host"], "syntax": "attribute", "type": "string"}, {"aliases": ["clickhouse params password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["clickhouse_params", "password"], "syntax": "attribute", "type": "object"}, {"aliases": ["clickhouse params port"], "anchor": "schema-clickhouse_params--port", "description": "Clickhouse Port. Clickhouse Port.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["clickhouse params user"], "anchor": "schema-clickhouse_params--user", "description": "Clickhouse User. Clickhouse User.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "user"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/clickhouse_params/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration parameter for clickhouse params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [clickhouse_params.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
