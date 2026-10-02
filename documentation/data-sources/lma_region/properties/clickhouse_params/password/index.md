---
page_title: "clickhouse_params.password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["clickhouse params password"], "body_bytes": 1700, "body_sha256": "sha256:a2f8a0de53bdb6dc201a5ad981cc3915ff44d315e9559742000166b3530e57bf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password", "parent_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params", "path": "documentation/data-sources/lma_region/properties/clickhouse_params/password/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2001211213223022-0121320233312311-2223220010102012-3223001333301232-2132322021323001-0013202031111022-1330311302113220-2022303333231132", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["clickhouse_params", "password"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["clickhouse_params", "password", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["clickhouse_params", "password", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/clickhouse_params/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# clickhouse_params.password

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [clickhouse_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/)
- clickhouse_params.password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/clear_secret_info/): complete subsection reference.

## Next pages

- [clickhouse_params.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/)
- [clickhouse_params.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/clear_secret_info/)
- [clickhouse_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
