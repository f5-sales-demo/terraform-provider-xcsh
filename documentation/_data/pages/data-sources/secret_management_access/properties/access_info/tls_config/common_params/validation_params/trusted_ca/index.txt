---
page_title: "access_info.tls_config.common_params.validation_params.trusted_ca"
subcategory: ""
description: "Reference to Root CA Certificate."
xcsh_docs: {"aliases": ["access info tls config common params validation params trusted ca"], "body_bytes": 1764, "body_sha256": "sha256:135ce29837240787ecc67bec6bfe9b2fb803c26a969bcc14f39df7f0fd70090f", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params", "path": "documentation/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3323022232231121-0202001001312303-2231100110230321-2130122233101103-3231100313122112-1003222223102130-1121012032301232-0213203023200030", "registry_path": "docs/guides/data-sources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca"], "schema_version": 1, "sections": [{"aliases": ["access info tls config common params validation params trusted ca trusted ca list"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:validation_params:trusted_ca:trusted_ca_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reference to Root CA Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/)
- [access_info.tls_config.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/)
- access_info.tls_config.common_params.validation_params.trusted_ca

<a id="section"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

- [trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/properties/access_info/tls_config/common_params/validation_params/trusted_ca/trusted_ca_list/): complete subsection reference.
