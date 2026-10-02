---
page_title: "datadog_receiver.use_tls.mtls_enable.key_url"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["datadog receiver use tls mtls enable key url"], "body_bytes": 2584, "body_sha256": "sha256:a78458249e5c24262a424d4cce25b385607f0b84530662726dfa462a3491f32b", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_enable:key_url:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_enable:key_url", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_enable", "path": "documentation/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0300002020110133-0233110100022230-3212002123010013-3200231212311312-2033203000022311-3302001213032330-1032112333301132-1020231313201010", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["datadog_receiver", "use_tls", "mtls_enable", "key_url"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_enable:key_url:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["datadog_receiver", "use_tls", "mtls_enable", "key_url", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_enable:key_url:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["datadog_receiver", "use_tls", "mtls_enable", "key_url", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# datadog_receiver.use_tls.mtls_enable.key_url

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/)
- [datadog_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/)
- [datadog_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/)
- datadog_receiver.use_tls.mtls_enable.key_url

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/clear_secret_info/): complete subsection reference.

## Next pages

- [datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/blindfold_secret_info/)
- [datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/key_url/clear_secret_info/)
- [datadog_receiver.use_tls.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/use_tls/mtls_enable/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
