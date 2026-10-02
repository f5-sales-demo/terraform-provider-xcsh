---
page_title: "origin_pool.use_tls.use_mtls.tls_certificates"
subcategory: "Load Balancing"
description: "MTLS Client Certificate."
xcsh_docs: {"aliases": ["backend servers", "cert", "certificate", "existing certificates", "origin pool use tls use mtls tls certificates", "origin servers", "tls certificates", "upstream servers"], "body_bytes": 6304, "body_sha256": "sha256:3136736cd933d974373e458a2985573461f004b8d009ded0735ad14ba0b0ae51", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key", "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls", "path": "documentation/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--certificate_url", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates:RequiredListObjectAttributes:certificate_url", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--custom_hash_algorithms--hash_algorithms", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms:RequiredObjectAttributes:hash_algorithms", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "type": "requires"}], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "custom_hash_algorithms"], "syntax": "block", "type": "object"}, {"aliases": ["cert", "certificate", "description spec", "existing certificates", "tls certificates"], "anchor": "schema-origin_pool--use_tls--use_mtls--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pool.use_tls.use_mtls.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:private_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "private_key"], "syntax": "block", "type": "object"}, {"aliases": ["use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:origin_pool:use_tls:use_mtls:tls_certificates:use_system_defaults", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pool", "use_tls", "use_mtls", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "MTLS Client Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pool.use_tls.use_mtls.tls_certificates

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/)
- [origin_pool.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/)
- [origin_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/)
- origin_pool.use_tls.use_mtls.tls_certificates

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--certificate_url"></a>

### certificate_url property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-origin_pool--use_tls--use_mtls--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/use_system_defaults/): complete subsection reference.

## Next pages

- [origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/custom_hash_algorithms/)
- [origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/private_key/)
- [origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/tls_certificates/use_system_defaults/)
- [origin_pool.use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/origin_pool/use_tls/use_mtls/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
