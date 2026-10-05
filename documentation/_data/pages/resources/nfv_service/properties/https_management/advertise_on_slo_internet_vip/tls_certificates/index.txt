---
page_title: "https_management.advertise_on_slo_internet_vip.tls_certificates"
subcategory: ""
description: "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "https management advertise on slo internet vip tls certificates", "tls certificates"], "body_bytes": 6699, "body_sha256": "sha256:02329079755bef1b29339a335fce5f08b3c0c05b5eee5742fd5d71a898a89597", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:custom_hash_algorithms", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:disable_ocsp_stapling", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip", "path": "documentation/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3000022222203033-3231020103202213-3023210232220211-3031323131333002-2330312032131112-2210223330301032-0222333023223300-3031111021111331", "registry_path": "docs/guides/resources--nfv_service--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "schema-https_management--advertise_on_slo_internet_vip--tls_certificates--certificate_url", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates:RequiredListObjectAttributes:certificate_url", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "https management advertise on slo internet vip tls certificates certificate url", "tls certificates"], "anchor": "schema-https_management--advertise_on_slo_internet_vip--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["https management advertise on slo internet vip tls certificates custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:custom_hash_algorithms", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-https_management--advertise_on_slo_internet_vip--tls_certificates--custom_hash_algorithms--hash_algorithms", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms:RequiredObjectAttributes:hash_algorithms", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:custom_hash_algorithms", "type": "requires"}], "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "custom_hash_algorithms"], "syntax": "block", "type": "object"}, {"aliases": ["https management advertise on slo internet vip tls certificates description spec"], "anchor": "schema-https_management--advertise_on_slo_internet_vip--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["https management advertise on slo internet vip tls certificates disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:disable_ocsp_stapling", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo internet vip tls certificates private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https_management.advertise_on_slo_internet_vip.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:private_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "private_key"], "syntax": "block", "type": "object"}, {"aliases": ["https management advertise on slo internet vip tls certificates use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates:use_system_defaults", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_internet_vip", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_internet_vip.tls_certificates

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/)
- [https_management.advertise_on_slo_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/)
- https_management.advertise_on_slo_internet_vip.tls_certificates

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
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

<a id="schema-https_management--advertise_on_slo_internet_vip--tls_certificates--certificate_url"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-https_management--advertise_on_slo_internet_vip--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/use_system_defaults/): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/custom_hash_algorithms/)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/disable_ocsp_stapling/)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/private_key/)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/tls_certificates/use_system_defaults/)
- [https_management.advertise_on_slo_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
