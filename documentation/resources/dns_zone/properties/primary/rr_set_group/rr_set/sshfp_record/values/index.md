---
page_title: "primary.rr_set_group.rr_set.sshfp_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary rr set group rr set sshfp record values"], "body_bytes": 4075, "body_sha256": "sha256:3d251b3a96501d88be350ffa77e24da2870b496bb43d36afe8c28c65850882d1", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha1_fingerprint", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record", "path": "documentation/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322", "registry_path": "docs/guides/resources--dns_zone--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record.values:ConflictingListObjectAttributes:sha1_fingerprint,sha256_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha1_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record.values:ConflictingListObjectAttributes:sha1_fingerprint,sha256_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "values"], "schema_version": 1, "sections": [{"aliases": ["primary rr set group rr set sshfp record values algorithm"], "anchor": "schema-primary--rr_set_group--rr_set--sshfp_record--values--algorithm", "description": "SSHFP algorithm value must be compatible with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA - ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["DSA", "ECDSA", "Ed25519", "Ed448", "RSA", "UNSPECIFIEDALGORITHM"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "values", "algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set sshfp record values sha1 fingerprint"], "anchor": "section", "description": "Configuration parameter for sha1 fingerprint.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha1_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--sshfp_record--values--sha1_fingerprint--fingerprint", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint:RequiredObjectAttributes:fingerprint", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha1_fingerprint", "type": "requires"}], "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "values", "sha1_fingerprint"], "syntax": "block", "type": "object"}, {"aliases": ["primary rr set group rr set sshfp record values sha256 fingerprint"], "anchor": "section", "description": "Configuration parameter for sha256 fingerprint.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint--fingerprint", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint:RequiredObjectAttributes:fingerprint", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint", "type": "requires"}], "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "values", "sha256_fingerprint"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.sshfp_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [primary.rr_set_group.rr_set.sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/)
- primary.rr_set_group.rr_set.sshfp_record.values

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="schema-primary--rr_set_group--rr_set--sshfp_record--values--algorithm"></a>

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

- [sha1_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha1_fingerprint/): complete subsection reference.

- [sha256_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha256_fingerprint/): complete subsection reference.
