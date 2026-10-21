---
page_title: "access_info.tls_config.cert_params.tls_validation_params.trusted_ca"
subcategory: ""
description: "Reference to Root CA Certificate."
xcsh_docs: {"aliases": ["access info tls config cert params tls validation params trusted ca"], "body_bytes": 1868, "body_sha256": "sha256:2c88648b31ae9fdcde78dad089ea2d910c47a9568a7a6bdb0234bc5f52990faf", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca:trusted_ca_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "path": "documentation/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2031030213320222-0001001333133110-0130332232123201-0022203002221000-2200312220313122-0103013012013030-1313302100301021-0233110310001321", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca"], "schema_version": 1, "sections": [{"aliases": ["access info tls config cert params tls validation params trusted ca trusted ca list"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca:trusted_ca_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca", "trusted_ca_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reference to Root CA Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.cert_params.tls_validation_params.trusted_ca

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/cert_params/)
- [access_info.tls_config.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

## Direct properties

- [trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/): complete subsection reference.
