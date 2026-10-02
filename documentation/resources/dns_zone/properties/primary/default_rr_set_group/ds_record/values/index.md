---
page_title: "primary.default_rr_set_group.ds_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary default rr set group ds record values"], "body_bytes": 6969, "body_sha256": "sha256:d75182ef3e4172a1745f156d46d917d292cdf6303e81bb98b1016535fa8229ce", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha256_digest", "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha384_digest"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1223012310302301-2011203300111022-3113021233330001-0021213101120212-0310110111121311-3033313130130321-2202130213001221-2313221003232312", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha256_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha256_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha256_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha256_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha256_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha384_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha256_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha384_digest", "type": "conflicts"}, {"anchor": "schema-primary--default_rr_set_group--ds_record--values--key_tag", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:RequiredListObjectAttributes:key_tag", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "ds_record", "values"], "schema_version": 1, "sections": [{"aliases": ["ds key algorithm"], "anchor": "schema-primary--default_rr_set_group--ds_record--values--ds_key_algorithm", "description": "DS key value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1: RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 - ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448: ED448.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "ds_record", "values", "ds_key_algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["key tag"], "anchor": "schema-primary--default_rr_set_group--ds_record--values--key_tag", "description": "A short numeric value which can help quickly identify the referenced DNSKEY-record.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "ds_record", "values", "key_tag"], "syntax": "attribute", "type": "number"}, {"aliases": ["sha1 digest"], "anchor": "section", "description": "Configuration parameter for sha1 digest.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--default_rr_set_group--ds_record--values--sha1_digest--digest", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values.sha1_digest:RequiredObjectAttributes:digest", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "type": "requires"}], "schema_path": ["primary", "default_rr_set_group", "ds_record", "values", "sha1_digest"], "syntax": "block", "type": "object"}, {"aliases": ["sha256 digest"], "anchor": "section", "description": "Configuration parameter for sha256 digest.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha256_digest", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--default_rr_set_group--ds_record--values--sha256_digest--digest", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values.sha256_digest:RequiredObjectAttributes:digest", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha256_digest", "type": "requires"}], "schema_path": ["primary", "default_rr_set_group", "ds_record", "values", "sha256_digest"], "syntax": "block", "type": "object"}, {"aliases": ["sha384 digest"], "anchor": "section", "description": "Configuration parameter for sha384 digest.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha384_digest", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-primary--default_rr_set_group--ds_record--values--sha384_digest--digest", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values.sha384_digest:RequiredObjectAttributes:digest", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha384_digest", "type": "requires"}], "schema_path": ["primary", "default_rr_set_group", "ds_record", "values", "sha384_digest"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.ds_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/)
- primary.default_rr_set_group.ds_record.values

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

<a id="schema-primary--default_rr_set_group--ds_record--values--ds_key_algorithm"></a>

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

<a id="schema-primary--default_rr_set_group--ds_record--values--key_tag"></a>

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

- [sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha1_digest/): complete subsection reference.

- [sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha256_digest/): complete subsection reference.

- [sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha384_digest/): complete subsection reference.

## Next pages

- [primary.default_rr_set_group.ds_record.values.sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha1_digest/)
- [primary.default_rr_set_group.ds_record.values.sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha256_digest/)
- [primary.default_rr_set_group.ds_record.values.sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha384_digest/)
- [primary.default_rr_set_group.ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
