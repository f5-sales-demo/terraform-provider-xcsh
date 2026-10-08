---
page_title: "primary.default_rr_set_group.sshfp_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary default rr set group sshfp record values"], "body_bytes": 3949, "body_sha256": "sha256:6b8c671f8f2abb61dd675e678b47aaf8247ae1c817621c50ce02b6febc1c48fe", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha1_fingerprint", "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha256_fingerprint"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1101120321300303-2320122100231000-1311002212133031-0121103313331023-3103122323300332-0202111210102033-0301330202021300-0012010120131212", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.sshfp_record.values:ConflictingListObjectAttributes:sha1_fingerprint,sha256_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha1_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.sshfp_record.values:ConflictingListObjectAttributes:sha1_fingerprint,sha256_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha256_fingerprint", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "sshfp_record", "values"], "schema_version": 1, "sections": [{"aliases": ["primary default rr set group sshfp record values algorithm"], "anchor": "schema-primary--default_rr_set_group--sshfp_record--values--algorithm", "description": "SSHFP algorithm value must be compatible with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA - ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DSA", "ECDSA", "Ed25519", "Ed448", "RSA", "UNSPECIFIEDALGORITHM"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "sshfp_record", "values", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary default rr set group sshfp record values sha1 fingerprint"], "anchor": "section", "description": "Configuration parameter for sha1 fingerprint.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha1_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--default_rr_set_group--sshfp_record--values--sha1_fingerprint--fingerprint", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint:RequiredObjectAttributes:fingerprint", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha1_fingerprint", "type": "requires"}], "schema_path": ["primary", "default_rr_set_group", "sshfp_record", "values", "sha1_fingerprint"], "syntax": "block", "type": "object"}, {"aliases": ["primary default rr set group sshfp record values sha256 fingerprint"], "anchor": "section", "description": "Configuration parameter for sha256 fingerprint.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha256_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--default_rr_set_group--sshfp_record--values--sha256_fingerprint--fingerprint", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint:RequiredObjectAttributes:fingerprint", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:sshfp_record:values:sha256_fingerprint", "type": "requires"}], "schema_path": ["primary", "default_rr_set_group", "sshfp_record", "values", "sha256_fingerprint"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.sshfp_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/)
- primary.default_rr_set_group.sshfp_record.values

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

SSHFP Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("sha1_fingerprint",
    "sha256_fingerprint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--default_rr_set_group--sshfp_record--values--algorithm"></a>

### algorithm property

Type: `"string"`. Optional.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DSA","ECDSA","Ed25519","Ed448","RSA","UNSPECIFIEDALGORITHM"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/sha1_fingerprint/): complete subsection reference.

- [sha256_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/sha256_fingerprint/): complete subsection reference.
