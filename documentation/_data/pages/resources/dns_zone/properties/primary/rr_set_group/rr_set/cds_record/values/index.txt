---
page_title: "primary.rr_set_group.rr_set.cds_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary rr set group rr set cds record values"], "body_bytes": 7103, "body_sha256": "sha256:61937d8cca39a34ae50cb4917ea41e72bbc241361ed5cf28f05935280f9b0d2a", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha1_digest", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha256_digest", "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha384_digest"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record", "path": "documentation/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values:ConflictingListObjectAttributes:sha1_digest,sha256_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha1_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values:ConflictingListObjectAttributes:sha1_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha1_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values:ConflictingListObjectAttributes:sha1_digest,sha256_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha256_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values:ConflictingListObjectAttributes:sha256_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha256_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values:ConflictingListObjectAttributes:sha1_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha384_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values:ConflictingListObjectAttributes:sha256_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha384_digest", "type": "conflicts"}, {"anchor": "schema-primary--rr_set_group--rr_set--cds_record--values--key_tag", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values:RequiredListObjectAttributes:key_tag", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "cds_record", "values"], "schema_version": 1, "sections": [{"aliases": ["ds key algorithm"], "anchor": "schema-primary--rr_set_group--rr_set--cds_record--values--ds_key_algorithm", "description": "DS key value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1: RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 - ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448: ED448.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "cds_record", "values", "ds_key_algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["key tag"], "anchor": "schema-primary--rr_set_group--rr_set--cds_record--values--key_tag", "description": "A short numeric value which can help quickly identify the referenced DNSKEY-record.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "cds_record", "values", "key_tag"], "syntax": "attribute", "type": "number"}, {"aliases": ["sha1 digest"], "anchor": "section", "description": "Configuration parameter for sha1 digest.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha1_digest", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--cds_record--values--sha1_digest--digest", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values.sha1_digest:RequiredObjectAttributes:digest", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha1_digest", "type": "requires"}], "schema_path": ["primary", "rr_set_group", "rr_set", "cds_record", "values", "sha1_digest"], "syntax": "block", "type": "object"}, {"aliases": ["sha256 digest"], "anchor": "section", "description": "Configuration parameter for sha256 digest.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha256_digest", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--cds_record--values--sha256_digest--digest", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values.sha256_digest:RequiredObjectAttributes:digest", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha256_digest", "type": "requires"}], "schema_path": ["primary", "rr_set_group", "rr_set", "cds_record", "values", "sha256_digest"], "syntax": "block", "type": "object"}, {"aliases": ["sha384 digest"], "anchor": "section", "description": "Configuration parameter for sha384 digest.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha384_digest", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--cds_record--values--sha384_digest--digest", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.cds_record.values.sha384_digest:RequiredObjectAttributes:digest", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:cds_record:values:sha384_digest", "type": "requires"}], "schema_path": ["primary", "rr_set_group", "rr_set", "cds_record", "values", "sha384_digest"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.cds_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [primary.rr_set_group.rr_set.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/)
- primary.rr_set_group.rr_set.cds_record.values

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
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

<a id="schema-primary--rr_set_group--rr_set--cds_record--values--ds_key_algorithm"></a>

### ds_key_algorithm property

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--cds_record--values--key_tag"></a>

### key_tag property

Type: `"number"`. Optional.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha1_digest/): complete subsection reference.

- [sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha256_digest/): complete subsection reference.

- [sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha384_digest/): complete subsection reference.

## Next pages

- [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha1_digest/)
- [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha256_digest/)
- [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha384_digest/)
- [primary.rr_set_group.rr_set.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
