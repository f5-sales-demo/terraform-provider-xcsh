---
page_title: "access_info.tls_config.cert_params.tls_validation_params.trusted_ca"
subcategory: ""
description: "Reference to Root CA Certificate."
xcsh_docs: {"aliases": ["access info tls config cert params tls validation params trusted ca", "cert", "certificate", "existing certificates", "tls certificates"], "body_bytes": 2478, "body_sha256": "sha256:90cb659bd9663f4e4bc0e16bd7a00b5701eb49c82d3e87536298a7be9efd3707", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca:trusted_ca_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "path": "documentation/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0021003012231102-1113101021031010-1123303330032223-1230302120220031-3233030320222203-0030320213202202-3132332220213000-3023131312113012", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "trusted ca list"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca:trusted_ca_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca", "trusted_ca_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Reference to Root CA Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.cert_params.tls_validation_params.trusted_ca

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/)
- [access_info.tls_config.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca

<a id="section"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/): complete subsection reference.

## Next pages

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/)
- [access_info.tls_config.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
