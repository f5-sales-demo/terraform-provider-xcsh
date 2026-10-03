---
page_title: "use_tls.use_mtls.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["use tls use mtls tls certificates private key"], "body_bytes": 2717, "body_sha256": "sha256:5ffb9ea904913eac81421fa7ffd1c48c266ba8343170505f00cd86873f23e8db", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "path": "documentation/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3113313001013311-3301020000020230-0113331133101031-2333002303322011-0321223332133002-0313202012121001-3102010030122110-2303221132303312", "registry_path": "docs/guides/resources--origin_pool--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_tls", "use_mtls", "tls_certificates", "private_key"], "schema_version": 1, "sections": [{"aliases": ["use tls use mtls tls certificates private key blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "private_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["use tls use mtls tls certificates private key clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates.private_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "type": "requires"}], "schema_path": ["use_tls", "use_mtls", "tls_certificates", "private_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls.tls_certificates.private_key

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/)
- [use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/)
- [use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/)
- use_tls.use_mtls.tls_certificates.private_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/)
- [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/)
- [use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
