---
page_title: "tls_parameters.common_params.tls_certificates.private_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["tls parameters common params tls certificates private key"], "body_bytes": 2689, "body_sha256": "sha256:f2c7fc59d1bcf67d5d0f2e0883f54ae7109c264a7205a4adb5340974a020a15c", "capabilities": ["load-balancing.tls", "networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:tls_certificates", "path": "documentation/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0323030313030010-2103100132200310-0233232101002303-0003302312120300-1123101022232022-0232332232222113-1000213320213003-1301123113301303", "registry_path": "docs/guides/data-sources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/)
- tls_parameters.common_params.tls_certificates.private_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/tls_parameters/common_params/tls_certificates/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
