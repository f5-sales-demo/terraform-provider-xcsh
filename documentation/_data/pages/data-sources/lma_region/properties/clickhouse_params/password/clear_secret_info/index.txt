---
page_title: "clickhouse_params.password.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["clickhouse params password clear secret info"], "body_bytes": 1800, "body_sha256": "sha256:5d547896f64340d890876f8cd73c79146fea4ec561aeddd632b4f4b06384f4b3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:clear_secret_info", "parent_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password", "path": "documentation/data-sources/lma_region/properties/clickhouse_params/password/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0023101121323003-0332330003020231-3320101220112333-0001220011110132-0201201120122312-3110102303221011-3032022220033223-1123323332203012", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["clickhouse_params", "password", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["clickhouse params password clear secret info provider ref"], "anchor": "schema-clickhouse_params--password--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "password", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["clickhouse params password clear secret info url"], "anchor": "schema-clickhouse_params--password--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:clear_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "password", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/clickhouse_params/password/clear_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# clickhouse_params.password.clear_secret_info

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [clickhouse_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/)
- [clickhouse_params.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/)
- clickhouse_params.password.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

## Direct properties

<a id="schema-clickhouse_params--password--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-clickhouse_params--password--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

## Next pages

- [clickhouse_params.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
