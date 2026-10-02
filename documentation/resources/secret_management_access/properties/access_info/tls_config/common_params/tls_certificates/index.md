---
page_title: "access_info.tls_config.common_params.tls_certificates"
subcategory: ""
description: "Set of TLS certificates."
xcsh_docs: {"aliases": ["access info tls config common params tls certificates", "cert", "certificate", "existing certificates", "tls certificates"], "body_bytes": 5911, "body_sha256": "sha256:6dcbe20e9dfc6a2023bb57ec728682b392b55f93ef35c9810eb6b2983c0218bd", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:disable_ocsp_stapling", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params", "path": "documentation/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1222230031301121-1003331022322112-1022200023222230-2333120003231030-1231230303022001-3322203111003121-3233121011323013-3222132303102310", "registry_path": "docs/guides/resources--secret_management_access--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "schema-access_info--tls_config--common_params--tls_certificates--certificate_url", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates:RequiredListObjectAttributes:certificate_url", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "certificate url", "existing certificates", "tls certificates"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms:RequiredObjectAttributes:hash_algorithms", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "type": "requires"}], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "custom_hash_algorithms"], "syntax": "block", "type": "object"}, {"aliases": ["cert", "certificate", "description spec", "existing certificates", "tls certificates"], "anchor": "schema-access_info--tls_config--common_params--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:disable_ocsp_stapling", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "access_info.tls_config.common_params.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "private_key"], "syntax": "block", "type": "object"}, {"aliases": ["use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:use_system_defaults", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Set of TLS certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.tls_certificates

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/)
- [access_info.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/)
- [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/)
- access_info.tls_config.common_params.tls_certificates

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

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
tls_certificates {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-access_info--tls_config--common_params--tls_certificates--certificate_url"></a>

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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-access_info--tls_config--common_params--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/use_system_defaults/): complete subsection reference.

## Next pages

- [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/custom_hash_algorithms/)
- [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/disable_ocsp_stapling/)
- [access_info.tls_config.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/private_key/)
- [access_info.tls_config.common_params.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/use_system_defaults/)
- [access_info.tls_config.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/access_info/tls_config/common_params/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
