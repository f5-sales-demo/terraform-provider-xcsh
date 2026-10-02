---
page_title: "enable_forward_proxy.tls_intercept.custom_certificate.private_key"
subcategory: "Networking"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["cert", "certificate", "enable forward proxy tls intercept custom certificate private key", "existing certificates", "tls certificates"], "body_bytes": 2822, "body_sha256": "sha256:8ce6d3f0bd121b1828db0135c9b6c17bfde130b9bb4a9dfc91c82f33d61162cf", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:blindfold_secret_info", "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key", "parent_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate", "path": "documentation/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0210121113331122-2223233132220213-3020031300302212-0323332120211011-0111011220000200-2021220120312010-2000112133312310-1130021333331212", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:network_connector:properties:enable_forward_proxy:tls_intercept:custom_certificate:private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_forward_proxy", "tls_intercept", "custom_certificate", "private_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_forward_proxy.tls_intercept.custom_certificate.private_key

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [enable_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/)
- [enable_forward_proxy.tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/)
- [enable_forward_proxy.tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/blindfold_secret_info/)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/private_key/clear_secret_info/)
- [enable_forward_proxy.tls_intercept.custom_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/enable_forward_proxy/tls_intercept/custom_certificate/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
