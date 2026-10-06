---
page_title: "tls_tcp.tls_parameters.tls_certificates"
subcategory: "Load Balancing"
description: "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls tcp tls parameters tls certificates"], "body_bytes": 4895, "body_sha256": "sha256:aff27c84edc8ae1208023c5dd063f36532aaf2ed5da3551084d41a6d4a0b55ff", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:custom_hash_algorithms", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:disable_ocsp_stapling", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:use_system_defaults"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters", "path": "documentation/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0002320022201130-0310033133223112-1210131011131021-0233000132012002-0030201320130223-3020131322232320-1330011330112100-0231301213311120", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "schema-tls_tcp--tls_parameters--tls_certificates--certificate_url", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates:RequiredListObjectAttributes:certificate_url", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls tcp tls parameters tls certificates certificate url"], "anchor": "schema-tls_tcp--tls_parameters--tls_certificates--certificate_url", "description": "TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "certificate_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls tcp tls parameters tls certificates custom hash algorithms"], "anchor": "section", "description": "Specifies the hash algorithms to be used.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:custom_hash_algorithms", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tls_tcp--tls_parameters--tls_certificates--custom_hash_algorithms--hash_algorithms", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms:RequiredObjectAttributes:hash_algorithms", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:custom_hash_algorithms", "type": "requires"}], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "custom_hash_algorithms"], "syntax": "block", "type": "object"}, {"aliases": ["tls tcp tls parameters tls certificates description spec"], "anchor": "schema-tls_tcp--tls_parameters--tls_certificates--description_spec", "description": "Description. Description for the certificate.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls tcp tls parameters tls certificates disable ocsp stapling"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:disable_ocsp_stapling", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "disable_ocsp_stapling"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls tcp tls parameters tls certificates private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "private_key"], "syntax": "block", "type": "object"}, {"aliases": ["tls tcp tls parameters tls certificates use system defaults"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:use_system_defaults", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "use_system_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters.tls_certificates

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/)
- [tls_tcp.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/)
- tls_tcp.tls_parameters.tls_certificates

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="schema-tls_tcp--tls_parameters--tls_certificates--certificate_url"></a>

### certificate_url property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/custom_hash_algorithms/): complete subsection reference.

<a id="schema-tls_tcp--tls_parameters--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/disable_ocsp_stapling/): complete subsection reference.

- [private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/): complete subsection reference.

- [use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/use_system_defaults/): complete subsection reference.
