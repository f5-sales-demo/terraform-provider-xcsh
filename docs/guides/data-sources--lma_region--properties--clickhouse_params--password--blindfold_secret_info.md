---
page_title: "clickhouse_params.password.blindfold_secret_info"
subcategory: ""
description: "clickhouse_params.password.blindfold_secret_info for xcsh_lma_region."
xcsh_docs: {"aliases": [], "body_bytes": 1739, "body_sha256": "sha256:aa5579e620bf2fbd5837ef26b1fdb1039a5fe5cc442f03c3ca18e126775dfa41", "canonical_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:lma_region:properties:clickhouse_params:password", "path": "docs/guides/data-sources--lma_region--properties--clickhouse_params--password--blindfold_secret_info.md", "provider_name": "lma_region", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["clickhouse_params", "password", "blindfold_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/clickhouse_params/password/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "clickhouse_params.password.blindfold_secret_info for xcsh_lma_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# clickhouse_params.password.blindfold_secret_info

Breadcrumbs:

- [xcsh_lma_region](../data-sources/lma_region.md)
- [Property reference](data-sources--lma_region--reference.md)
- [clickhouse_params](data-sources--lma_region--properties--clickhouse_params.md)
- [clickhouse_params.password](data-sources--lma_region--properties--clickhouse_params--password.md)
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

- [clickhouse_params.password](data-sources--lma_region--properties--clickhouse_params--password.md)
- [xcsh_lma_region](../data-sources/lma_region.md)
