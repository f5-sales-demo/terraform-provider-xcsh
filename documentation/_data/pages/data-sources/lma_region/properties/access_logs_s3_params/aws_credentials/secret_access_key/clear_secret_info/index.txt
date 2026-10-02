---
page_title: "access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["access logs s3 params aws credentials secret access key clear secret info", "authentication", "credential setup", "credentials"], "body_bytes": 2215, "body_sha256": "sha256:f985c20bb5fe6392cb97e5bef48f128fba991cfc9b82ca5aba2584b393b432cd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key:clear_secret_info", "parent_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key", "path": "documentation/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0213033322332300-2311301012301311-0201211333331012-3203332101000301-3120133220020001-2320331001112223-0002000011101230-3230101322320033", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_logs_s3_params", "aws_credentials", "secret_access_key", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["provider ref"], "anchor": "schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "secret_access_key", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["url"], "anchor": "schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key:clear_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "secret_access_key", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [access_logs_s3_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/)
- [access_logs_s3_params.aws_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/)
- [access_logs_s3_params.aws_credentials.secret_access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/)
- access_logs_s3_params.aws_credentials.secret_access_key.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

## Direct properties

<a id="schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-access_logs_s3_params--aws_credentials--secret_access_key--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

## Next pages

- [access_logs_s3_params.aws_credentials.secret_access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
