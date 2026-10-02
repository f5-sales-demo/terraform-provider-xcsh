---
page_title: "clickhouse_params.password.blindfold_secret_info"
subcategory: ""
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["clickhouse params password blindfold secret info"], "body_bytes": 2045, "body_sha256": "sha256:75e12a4d3d2faa7ffccb6ea386f7f7d82eee1366b18b3d0cf8276c4584da0a9a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password", "path": "documentation/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3113103313221122-0322102011203302-3020130021133022-3233103110232132-1101123220311002-3221200120211020-2221113103313002-0020211102103313", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["clickhouse_params", "password", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "decryption provider", "origin servers", "upstream servers"], "anchor": "schema-clickhouse_params--password--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "password", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["location"], "anchor": "schema-clickhouse_params--password--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "password", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["store provider"], "anchor": "schema-clickhouse_params--password--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["clickhouse_params", "password", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# clickhouse_params.password.blindfold_secret_info

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [clickhouse_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/)
- [clickhouse_params.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/)
- clickhouse_params.password.blindfold_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

## Direct properties

<a id="schema-clickhouse_params--password--blindfold_secret_info--decryption_provider"></a>

### decryption_provider property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

<a id="schema-clickhouse_params--password--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

<a id="schema-clickhouse_params--password--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

## Next pages

- [clickhouse_params.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/clickhouse_params/password/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
