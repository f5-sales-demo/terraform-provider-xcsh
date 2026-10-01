---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-38d0c9ae18bc7012d75bb569480465d4ceb9551b5ca26122f6fbdb585b6c276a"></a>

## primary.default_rr_set_group.ds_record — primary.default_rr_set_group.ds_record / 667005a502d0 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.ds_record

<a id="canonical-9a8239d3f6bd83c459c22ef486f4c378221782004dfb8e3a5d39360c90d26599"></a>

Type: `"object"`. single nested block, Optional.

DNS DS Record. DNS DS Record.

Upstream description:

DNS DS Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-619e792335b3b2519992c5e67b93e22ecdd55830f8e7a37a85db6c293682d2b9"></a>

## Direct properties — primary.default_rr_set_group.ds_record / 667005a502d0 / 3

<a id="canonical-030fe7896e1262c76289a9a3f9091d258e31213fbe2bc6c41082f253915b7de1"></a>

<a id="canonical-14321e90a235bd5bf6413b0f3dd130cb89da049699efc55990723554bf0677a6"></a>

## name property — primary.default_rr_set_group.ds_record / 667005a502d0 / 4

Type: `"string"`. Optional.

DS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6): complete subsection reference.

<a id="canonical-20e4404e4a1765dac21b8e6e5e295b49c354a788531a6d790759ec2b93b34e4c"></a>

## Next pages — primary.default_rr_set_group.ds_record / 667005a502d0 / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e8f2055ca9f02e69188b1b24b714c0499ced6ffbb37b8072f5985126070f4d1"></a>

## primary.default_rr_set_group.ds_record.values — primary.default_rr_set_group.ds_record.values / b702cb952865 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-001.md#canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673)
- primary.default_rr_set_group.ds_record.values

<a id="canonical-e3c78ae8cda23c10de36087f9932beab06ecfb1e8428a8efc086365130c3d376"></a>

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

<a id="canonical-f1dfec83bb2e9810f7e90b99295c73715bf9d814bddd63d6a9adeeb9762ce4af"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values / b702cb952865 / 3

<a id="canonical-dc814eb94f7e4540d97aa0f674efe083d85be8ab38b631f425c83ff8769aa0b1"></a>

<a id="canonical-c31969ddc42c24d4c580204f8bf6b0ea99d719ac06f66e3b5f5d6d7845cd9b77"></a>

## ds_key_algorithm property — primary.default_rr_set_group.ds_record.values / b702cb952865 / 4

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

<a id="canonical-b0d259435bbce603cb39c15c71202071116b260097462ea7460239baa77947f9"></a>

<a id="canonical-983ded9c35d3616521af7eb5ce366f0c8c471431e828c0ec53bdd4b24bdd8e04"></a>

## key_tag property — primary.default_rr_set_group.ds_record.values / b702cb952865 / 5

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

- [sha1_digest](resources--dns_zone--reference--group-002.md#canonical-fc4ca6e45ce8d6e7edca86a5e1650d419c1b985e4adfc2c363abab49fb9b1101): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-002.md#canonical-e7a64cdb1b99d0d75a784ac37799fa528296793f06296974bfd6ffa071e9a69a): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-002.md#canonical-a4678c50e712d5752b79d4ce4c2e6ea26a82977fd00ae0c2e496d74ddb1f520d): complete subsection reference.

<a id="canonical-c4fac5560e27c6e6ce195236cea3d5d5996941434c5978f9c9ee15cccb915208"></a>

## Next pages — primary.default_rr_set_group.ds_record.values / b702cb952865 / 6

- [primary.default_rr_set_group.ds_record.values.sha1_digest](resources--dns_zone--reference--group-002.md#canonical-fc4ca6e45ce8d6e7edca86a5e1650d419c1b985e4adfc2c363abab49fb9b1101)
- [primary.default_rr_set_group.ds_record.values.sha256_digest](resources--dns_zone--reference--group-002.md#canonical-e7a64cdb1b99d0d75a784ac37799fa528296793f06296974bfd6ffa071e9a69a)
- [primary.default_rr_set_group.ds_record.values.sha384_digest](resources--dns_zone--reference--group-002.md#canonical-a4678c50e712d5752b79d4ce4c2e6ea26a82977fd00ae0c2e496d74ddb1f520d)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-001.md#canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-fc4ca6e45ce8d6e7edca86a5e1650d419c1b985e4adfc2c363abab49fb9b1101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-371fcfa5f45f6369a462b425ab77c606cf3b1c9b7c8ed276efd18f7401f6a976"></a>

## primary.default_rr_set_group.ds_record.values.sha1_digest — primary.default_rr_set_group.ds_record.values.sha1_digest / 0d307ef2e61b / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-001.md#canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673)
- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6)
- primary.default_rr_set_group.ds_record.values.sha1_digest

<a id="canonical-660c74bcb25854ab8d14ffb718bbddd445fa7ab9a7df83405671a32301c46bc6"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-25e13a6c5efdf1f6bfb8b6e8fb511e4ea8ef61fab6a550c2e149466edf8dba9e"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values.sha1_digest / 0d307ef2e61b / 3

<a id="canonical-5fd0911a73ebed263dbf35bf53e7ee69434fe7f6f94ae2500451bd8f41e8c79e"></a>

<a id="canonical-8c2b0455f11534541bca70d2fd98f7ccde6d42683c4c8f6b9a3cdb0b3d1d313c"></a>

## digest property — primary.default_rr_set_group.ds_record.values.sha1_digest / 0d307ef2e61b / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-a9347670a5e21fc8506d8e3d3c479973983cb6119ad093ed216fde769d393ac1"></a>

## Next pages — primary.default_rr_set_group.ds_record.values.sha1_digest / 0d307ef2e61b / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-e7a64cdb1b99d0d75a784ac37799fa528296793f06296974bfd6ffa071e9a69a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8604fdd36a26e4e655e7333fbaed64af439c35f0e09e92be7297e510a5cacbb"></a>

## primary.default_rr_set_group.ds_record.values.sha256_digest — primary.default_rr_set_group.ds_record.values.sha256_digest / 92793ca6f30c / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-001.md#canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673)
- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6)
- primary.default_rr_set_group.ds_record.values.sha256_digest

<a id="canonical-b97fdf25b052ddf367d6b53680ee1f754afc5353965f8c0153ec098fbda20e18"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-b9e388261534a264667cb9fb482ef684e27d38ab3fa348819383aeeac52d5aa0"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values.sha256_digest / 92793ca6f30c / 3

<a id="canonical-cd0e5a6bf2aad3fdd226e1283345c5b36db455557fffbbbde1aa6aa85d85a065"></a>

<a id="canonical-28dc6cd0be37057e98cc4c87f86078a4b48d26504303269abb7cef566921b57b"></a>

## digest property — primary.default_rr_set_group.ds_record.values.sha256_digest / 92793ca6f30c / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-ec8a0b0cb9e1e1e9f0bbc83fd1f1c70ad17908b74d895755b98d3e94e6b07460"></a>

## Next pages — primary.default_rr_set_group.ds_record.values.sha256_digest / 92793ca6f30c / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-a4678c50e712d5752b79d4ce4c2e6ea26a82977fd00ae0c2e496d74ddb1f520d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e73b666186b0a4ba244a1c8150df9256234af529079653305b0c75d66fe2fcb"></a>

## primary.default_rr_set_group.ds_record.values.sha384_digest — primary.default_rr_set_group.ds_record.values.sha384_digest / 7ffdebfabfa9 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.ds_record](resources--dns_zone--reference--group-001.md#canonical-a95a01f6aa08e943ea1267e9b696669ccf3c5fcf7acfd21343f4a70bf68b5673)
- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6)
- primary.default_rr_set_group.ds_record.values.sha384_digest

<a id="canonical-8c6af6f4266a18cc3cc3cbd7461359a94e19fc65ba1f80337cb31d8e9a4a193f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-adf9ecf82e46ad050dac8ed41e03fb0d5c50021fcf75293558fb873ec999ec1e"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values.sha384_digest / 7ffdebfabfa9 / 3

<a id="canonical-51cbd828a95301b45ae40bdff967a9b714ee07b6b28d1a09bb390bef2b57ef2c"></a>

<a id="canonical-bdf3d11be98e27e3e6a41aa7cb4c7954513a551248e303dabbb4504c8c5a1bf0"></a>

## digest property — primary.default_rr_set_group.ds_record.values.sha384_digest / 7ffdebfabfa9 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-812fa410b88f6f8286f18218805a6f90d8352c9f9393e236abba4d65bf51076e"></a>

## Next pages — primary.default_rr_set_group.ds_record.values.sha384_digest / 7ffdebfabfa9 / 5

- [primary.default_rr_set_group.ds_record.values](resources--dns_zone--reference--group-002.md#canonical-6b1b4cb1858f054ad726ff01099d162634515675cfddc739a2727069b7a43bb6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-d0061623363fa2c7be32ad499afa216264a6374fea4ded410c8eb10ecae277a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3b5aa25df1abc513aef42c21cbdc05b22263cfa262ab7cae98f036f86e89a7d"></a>

## primary.default_rr_set_group.eui48_record — primary.default_rr_set_group.eui48_record / 5590eeaf26ff / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.eui48_record

<a id="canonical-6d9f2388659ebdaf2fcfe6b0f2f271cbc4d5c63056a473a10138589f2d8f28e2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui48 record.

Upstream description:

DNS EUI48 Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
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
eui48_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-4635e9b9c64728a8c796630053f4656ee2310c05010aad9a8eb2000bf6783d32"></a>

## Direct properties — primary.default_rr_set_group.eui48_record / 5590eeaf26ff / 3

<a id="canonical-4643cc862286990aa94e26b0dc51730fc91a3e1f2fba83d2f444c5c112766b56"></a>

<a id="canonical-8678d7a8955b97def6ab020b9945b03b14ec144b23d810725c150b0e639596df"></a>

## name property — primary.default_rr_set_group.eui48_record / 5590eeaf26ff / 4

Type: `"string"`. Optional.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-e15a020edb3937514f16667c7688377b386cdf56083660d8f0f3316bd52ba8b7"></a>

<a id="canonical-564006c602fef49f7f6d556e34b223a6dee494621a80829f2bbce7e210a95c7b"></a>

## value property — primary.default_rr_set_group.eui48_record / 5590eeaf26ff / 5

Type: `"string"`. Optional.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Upstream description:

A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-165f1610a488effbee7297ed8746204c200d2097d7b82a0c246b1aebcf46f1a3"></a>

## Next pages — primary.default_rr_set_group.eui48_record / 5590eeaf26ff / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-34b84b028eff581da400c6b35883d2c8d1506ce9b2ab4580ae71ad7b2311a5f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae1e8ad8cfcc457d85e6e8ba91e3c57fd0656e1758090c7e8a3f234eb13952da"></a>

## primary.default_rr_set_group.eui64_record — primary.default_rr_set_group.eui64_record / f3933532edac / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.eui64_record

<a id="canonical-ebe180c5d3e1c34ef5f2f1e977ad310c579a32590a1ac31657d9e3ab7aad149e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
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
eui64_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-e217c348595299b8ed86ebf8d06a89e1bf999ade0d7f060a60e3e9ef38cd52b9"></a>

## Direct properties — primary.default_rr_set_group.eui64_record / f3933532edac / 3

<a id="canonical-546dbc00927c0775ab0758ce4678f60b57eacd7dfe3be49a724ffa24f519b500"></a>

<a id="canonical-d2127e63fa28ab8cc4c136d406387c516abc364a1256dd548defd80d8867503d"></a>

## name property — primary.default_rr_set_group.eui64_record / f3933532edac / 4

Type: `"string"`. Optional.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-cc2ae5c35a1070f40a050b53c9a858221c3fa14ebdaf5580ed47cfc867f5bde7"></a>

<a id="canonical-827649d21aead94fba69a4f4324a81a6954a624fcf0a1144f0b37712e5926cf0"></a>

## value property — primary.default_rr_set_group.eui64_record / f3933532edac / 5

Type: `"string"`. Optional.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(23, 23),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-98957127f1aba693a3d8c7f8f60daad620dfee17a5413655c32bbbc88d75a43a"></a>

## Next pages — primary.default_rr_set_group.eui64_record / f3933532edac / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-802c5376839db15e30f65fef43748962271c0826de4f85bc9c570e4b2968f34c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bb89d7077e69199de10c1f6223ed5565b9a9de1501c4c49004af06a74826e61"></a>

## primary.default_rr_set_group.lb_record — primary.default_rr_set_group.lb_record / 3ecaf8034b8c / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.lb_record

<a id="canonical-1b095e419ce2862479d03cc3908cd350afa5d9c3ce508a054705bae9513efbb4"></a>

Type: `"object"`. single nested block, Optional.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

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
lb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-82057e63db7eb51234a17401fa86bb867829e46604f88de713274f6939700012"></a>

## Direct properties — primary.default_rr_set_group.lb_record / 3ecaf8034b8c / 3

<a id="canonical-6218b9dc934bebc42c63c59e66ac55428a771e4443c2adbc1312743a610593c1"></a>

<a id="canonical-1ba03a548cfee78b479fd12a0588ce161e6ad1c82f755221f95b827198cf8b13"></a>

## name property — primary.default_rr_set_group.lb_record / 3ecaf8034b8c / 4

Type: `"string"`. Optional.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](resources--dns_zone--reference--group-002.md#canonical-5c13e70a293d680c6d1cca797d5de9705eae33693fd8c6fb9b9c61d50d1610b7): complete subsection reference.

<a id="canonical-b09f61fc2209870a28f091d48c5093fea8124de883630ddeb7317d1cae5c31c4"></a>

## Next pages — primary.default_rr_set_group.lb_record / 3ecaf8034b8c / 5

- [primary.default_rr_set_group.lb_record.value](resources--dns_zone--reference--group-002.md#canonical-5c13e70a293d680c6d1cca797d5de9705eae33693fd8c6fb9b9c61d50d1610b7)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-5c13e70a293d680c6d1cca797d5de9705eae33693fd8c6fb9b9c61d50d1610b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d0704c0d47927ee0f17b2e95c7f3a25903e57120f99cce54b0f7ebe78b0f277"></a>

## primary.default_rr_set_group.lb_record.value — primary.default_rr_set_group.lb_record.value / 4931cfb8db32 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-802c5376839db15e30f65fef43748962271c0826de4f85bc9c570e4b2968f34c)
- primary.default_rr_set_group.lb_record.value

<a id="canonical-fb4077bba3708dedad1cfaaa331b2bbe4fc84f66f8665441370ae38095453c8e"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ef1c948c8566fbcd24589290c9c80ce519409ee8824ee52f52aba2612aeefa6"></a>

## Direct properties — primary.default_rr_set_group.lb_record.value / 4931cfb8db32 / 3

<a id="canonical-24aab09cac04bc1c11f181fd830609a1bcc4788e48f7b677bc5448f4bc934372"></a>

<a id="canonical-73356acad232239d22f801cec61cecba505ba09b6ce87be5e4ac69ba09ed4fb0"></a>

## name property — primary.default_rr_set_group.lb_record.value / 4931cfb8db32 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-9d8bafa6c2451cf317d8ae5b0aed0555ea76ed23f60f75c5a55f3c53742499c7"></a>

<a id="canonical-21207f0c53931530d9afd0660215f67f09d8005376355e36acf1e0dbec29e7a0"></a>

## namespace property — primary.default_rr_set_group.lb_record.value / 4931cfb8db32 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-b5ba9b0e6a3fe065e76a365155db93486c227ea8239cc7ae88a35059ddea1adb"></a>

<a id="canonical-342efd192d4b1acecdd016ae3b59d5760b41b478089401263aca7087e0416897"></a>

## tenant property — primary.default_rr_set_group.lb_record.value / 4931cfb8db32 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1f245609e297bad78991eefc6d2cd5684ce4f05a3acd1d12b9f33476d5e7cdca"></a>

## Next pages — primary.default_rr_set_group.lb_record.value / 4931cfb8db32 / 7

- [primary.default_rr_set_group.lb_record](resources--dns_zone--reference--group-002.md#canonical-802c5376839db15e30f65fef43748962271c0826de4f85bc9c570e4b2968f34c)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-01f7a3b77dbf0a6d8939bf5a77c55698576dea8907d7f0a4c95d5604d913baff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78071167d8d6fcb7619ac77e974daee20af4ad34980be58c1db9bc8167dcf0cd"></a>

## primary.default_rr_set_group.loc_record — primary.default_rr_set_group.loc_record / 922781663f82 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.loc_record

<a id="canonical-9e557309f9eac6a4bc2b884e0a577755d66b9ec3f95e1616bda7f60f2e20c5c4"></a>

Type: `"object"`. single nested block, Optional.

DNS LOC Record. DNS LOC Record.

Upstream description:

DNS LOC Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
loc_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-f47abb848415b0a261dc2b1095d70e50bd11bc6117a4d7aca7b1bace1be4be50"></a>

## Direct properties — primary.default_rr_set_group.loc_record / 922781663f82 / 3

<a id="canonical-066fefe0825e0c60ff626d11568add5af771d513cef10beccdfd45d09af062e2"></a>

<a id="canonical-c34ad24361f95b87da10c3fcfb967ef7294311ba584883cce92e6ca5294103b6"></a>

## name property — primary.default_rr_set_group.loc_record / 922781663f82 / 4

Type: `"string"`. Optional.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-96655741f3b917377f5222fb915ad64653dc57fa4f701baf2e2fc1696b3e056c): complete subsection reference.

<a id="canonical-745cc28617e78c54bfa9bcc67286d26fa6799690d5ea5e5efc5b772b9f9ba315"></a>

## Next pages — primary.default_rr_set_group.loc_record / 922781663f82 / 5

- [primary.default_rr_set_group.loc_record.values](resources--dns_zone--reference--group-002.md#canonical-96655741f3b917377f5222fb915ad64653dc57fa4f701baf2e2fc1696b3e056c)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-96655741f3b917377f5222fb915ad64653dc57fa4f701baf2e2fc1696b3e056c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4da7804f0ae1167858248892dc90a6d28d81b89b3c368bc2e0d171cdebf7bbf9"></a>

## primary.default_rr_set_group.loc_record.values — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-01f7a3b77dbf0a6d8939bf5a77c55698576dea8907d7f0a4c95d5604d913baff)
- primary.default_rr_set_group.loc_record.values

<a id="canonical-00cf7713e51c975ba65921feb73b1efd972ffcc99a03db844fdfb3302d0c4472"></a>

Type: `"object"`. list nested block, Optional.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("altitude",
    "latitude_degree",
    "longitude_degree")}
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

<a id="canonical-6cc0a6b9d2c594673533953af4650f67eec9da3afeab6812b628942639819740"></a>

## Direct properties — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 3

<a id="canonical-9d8964e9a5d11bf5c75f9739eb9db4816b06e1939542ccd485978cbe0c03e86d"></a>

<a id="canonical-7d5271b8ded6582223a096ec77a721e0c4f12069e18de3dc057e554216ccdbeb"></a>

## altitude property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 4

Type: `"number"`. Optional.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2bd350ed9c5acc6a60094649cb92facf875aa91ed44165fe0a0e358c077a8579"></a>

<a id="canonical-f4484771e441b925b0ec86ef08fec6a752bfe9876733fc1de5cce9f534487ed9"></a>

## horizontal_precision property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 5

Type: `"number"`. Optional.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-d1fec37de3bb14c81be96593c7de296d82476ee78f059744f77607267ce6abc5"></a>

<a id="canonical-cd996abd400daa652f1e7fff90f5eab3cc0c7b9e442c74d2fdb3ce79b4f527e8"></a>

## latitude_degree property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 6

Type: `"number"`. Optional.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 90),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-c4eb1990e977902e83d715742d4680814c0d2948318d2b38e92e7bc14292cb47"></a>

<a id="canonical-321fcacec1c77213309a8738fdb77b2e92e3fa0d0980591e74145f47204b940e"></a>

## latitude_hemisphere property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 7

Type: `"string"`. Optional.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("N",
    "S"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c3804645da21022ed8c33cd070f2434b516b93c5f874f17457c2fa055c1c3d5c"></a>

<a id="canonical-b29d63ca53d4e7f3d0629a6a6d6d06ab37e4637ff97e25ecc755b09e4d410e14"></a>

## latitude_minute property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 8

Type: `"number"`. Optional.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-5ed597640e5f0e50ac085c907670d70a0b7462c1c5b5c57174915c8456410adb"></a>

<a id="canonical-6dc664681a9c8b62288aca72b5a4e571ade5e3d723c1c5c99d10fda4da9e79af"></a>

## latitude_second property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 9

Type: `"number"`. Optional.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-81c800cb145f1528cad20e877ad40147a3efd7ea7e68438f4aca96eab4a6f6ed"></a>

<a id="canonical-080f7877f20a596c12c972a72183a0767184e3bbdcb594eb9602a47d66c65892"></a>

## location_diameter property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 10

Type: `"number"`. Optional.

Diameter of a sphere enclosing the described entity, in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-1b382697513f05b9d2fd51a7d563df7c255f443961c550a9978656ceadd365e3"></a>

<a id="canonical-30626c4b6d03444e44292a5e9978e1f38e40f99ff9ecff0e1b5908f6b3a75e66"></a>

## longitude_degree property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 11

Type: `"number"`. Optional.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-4f8378def2c49a8086e385a4e1a2d13107ed84e36e2fd23db11657738700a90e"></a>

<a id="canonical-ca3b7de7b18e10618e21d6856a8b5ef9e900fff766357591e64820fcec420661"></a>

## longitude_hemisphere property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 12

Type: `"string"`. Optional.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("E",
    "W"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d1d1e8f16eef15134ba677473cc59b624faab18b29c13977e09dce3c21ecbead"></a>

<a id="canonical-5348e150a329a818b733f000b3a0f81746cdead0a0248486d1ca18eac042fd59"></a>

## longitude_minute property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 13

Type: `"number"`. Optional.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-8bb3542dfa5035da4ff7fb753e2ddb24af7c34758f46aad7d41e30a23a6216dc"></a>

<a id="canonical-8579dde487d3193da50ef2b3a6fc194f8fa954cdab440c8dc6c1aa8030786912"></a>

## longitude_second property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 14

Type: `"number"`. Optional.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-b3925685de3c303112d42201206a99c349a5ea127ac676b7a8a73434681d5e70"></a>

<a id="canonical-1cb6d959b48ff4d0f69ea68fd0f8252cac9d3a66327f7a8f471b28559302618b"></a>

## vertical_precision property — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 15

Type: `"number"`. Optional.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-10eeae09a5e8bf740abf92b316306ba25b04f85ad5d90a4fc4961ace417e4359"></a>

## Next pages — primary.default_rr_set_group.loc_record.values / 68323aeca0d7 / 16

- [primary.default_rr_set_group.loc_record](resources--dns_zone--reference--group-002.md#canonical-01f7a3b77dbf0a6d8939bf5a77c55698576dea8907d7f0a4c95d5604d913baff)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-42513c672c75487deb183579303e80a4c071e22205160f1118a0cea4e78708f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c3845c152c6d007701e4e17860958a4e22a4b0690e67b7793c39fed9cabd3fc"></a>

## primary.default_rr_set_group.mx_record — primary.default_rr_set_group.mx_record / e9fe4697ea83 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.mx_record

<a id="canonical-cd2866444f0039a87f9dfb81b18fcac562f91cd82c82cc457509c35d17996c09"></a>

Type: `"object"`. single nested block, Optional.

DNSMXResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
mx_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed3c647ffa46438ddd9f461552b3dc5ff9b4cdb93102347d5ca264be844567a9"></a>

## Direct properties — primary.default_rr_set_group.mx_record / e9fe4697ea83 / 3

<a id="canonical-790a25b4b60b9c945cd20ffbe5f39f564dc45cd3537cba2e3289e0e22cf06107"></a>

<a id="canonical-83476db6b6b5776d104be07373af009da763314fdaf5b67adfd12bd3f0ff8d7a"></a>

## name property — primary.default_rr_set_group.mx_record / e9fe4697ea83 / 4

Type: `"string"`. Optional.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-9e9a26b88ce507431f0f17d228a56d19dab7e281ceba20b81756c432f76428b5): complete subsection reference.

<a id="canonical-78af59e3184aff841c3b7066ce6b1bd76f2c8835fda5fdeb2fcc591f714e8b86"></a>

## Next pages — primary.default_rr_set_group.mx_record / e9fe4697ea83 / 5

- [primary.default_rr_set_group.mx_record.values](resources--dns_zone--reference--group-002.md#canonical-9e9a26b88ce507431f0f17d228a56d19dab7e281ceba20b81756c432f76428b5)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-9e9a26b88ce507431f0f17d228a56d19dab7e281ceba20b81756c432f76428b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7502d89f7828a7c291b7250f80321ddcdecfd88940fffc7f5418b22d07ed2331"></a>

## primary.default_rr_set_group.mx_record.values — primary.default_rr_set_group.mx_record.values / 5471e04ec642 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-42513c672c75487deb183579303e80a4c071e22205160f1118a0cea4e78708f6)
- primary.default_rr_set_group.mx_record.values

<a id="canonical-033347cf71da161052c41945518d91e0ee663af7e0873affe6c518b5944bef2c"></a>

Type: `"object"`. list nested block, Optional.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-e95b47ecbc014af684a814a543f1ba43a5a8b408a3713ecbadf70409be5634bb"></a>

## Direct properties — primary.default_rr_set_group.mx_record.values / 5471e04ec642 / 3

<a id="canonical-e6afeab19592737972142b63a3a2bd58af8c0921a902a1d10761271288007cb1"></a>

<a id="canonical-9bc01d4039171e9a27ef597173aa310a15b8b334a98f70ab55d9c1bfecd3999e"></a>

## domain property — primary.default_rr_set_group.mx_record.values / 5471e04ec642 / 4

Type: `"string"`. Optional.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-31254d157ac754a16d019b2e7f1120b1a262fdc2f39a2da6ccec5a6e8ec51870"></a>

<a id="canonical-7f6bc15acb477e8efa85d01cdb251613d62f3d7b0f41fd979452a10893974fa1"></a>

## priority property — primary.default_rr_set_group.mx_record.values / 5471e04ec642 / 5

Type: `"number"`. Optional.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-c75008adc45f431e7f86b3f0c3e090eb86e98f73ecd93b57d17e9c361f75343c"></a>

## Next pages — primary.default_rr_set_group.mx_record.values / 5471e04ec642 / 6

- [primary.default_rr_set_group.mx_record](resources--dns_zone--reference--group-002.md#canonical-42513c672c75487deb183579303e80a4c071e22205160f1118a0cea4e78708f6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-29861e75cf9362a30496f5183138a125bae01dd5822188db7f03bc58fccf8ff2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1dcd138d0d10d95f55ea777ccb624079bc929578f773f86f2855ea00d6ef8df"></a>

## primary.default_rr_set_group.naptr_record — primary.default_rr_set_group.naptr_record / e3a57d421adf / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.naptr_record

<a id="canonical-f284891cca16b8890fcd1a2102116fda0d3c6ba9f73a04eed2e4986551723574"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
naptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-90e64b3588cea1715f53f10f1741b91849ab1d5f20940def1e4e8904298794c7"></a>

## Direct properties — primary.default_rr_set_group.naptr_record / e3a57d421adf / 3

<a id="canonical-f382a8f5facda2b934541d6e75bd0a8280922745fadc550bdf19ab589833efaa"></a>

<a id="canonical-18d57df81a7aaa53cfd56ea40cd6efcd36608e513eb0b3314a51e0aac985dd81"></a>

## name property — primary.default_rr_set_group.naptr_record / e3a57d421adf / 4

Type: `"string"`. Optional.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-5531bf2fc60f71247484f901622e8a8c9a4045fbae7034e90d95a73f08cbba75): complete subsection reference.

<a id="canonical-1ba99320acada1e9f3c4f3dd6e748c54866a567fc43082ff20604876280e222e"></a>

## Next pages — primary.default_rr_set_group.naptr_record / e3a57d421adf / 5

- [primary.default_rr_set_group.naptr_record.values](resources--dns_zone--reference--group-002.md#canonical-5531bf2fc60f71247484f901622e8a8c9a4045fbae7034e90d95a73f08cbba75)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-5531bf2fc60f71247484f901622e8a8c9a4045fbae7034e90d95a73f08cbba75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84bff3fe34b34c42730ae6e27b03c5616e00b187e582a8565135ce489ffe41b6"></a>

## primary.default_rr_set_group.naptr_record.values — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-29861e75cf9362a30496f5183138a125bae01dd5822188db7f03bc58fccf8ff2)
- primary.default_rr_set_group.naptr_record.values

<a id="canonical-067e57912a5a5bc322daaf5f5520e66975cff88dd64671618c8399f6c84007b5"></a>

Type: `"object"`. list nested block, Optional.

NAPTR Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("flags",
    "order",
    "preference")}
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

<a id="canonical-a2789d7855e44ec2094d7cba1de3c0a5a2a9ade9bc782f87d7a620c05b0c610a"></a>

## Direct properties — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 3

<a id="canonical-4f26f91ae0b817661cf20551820ff6c0f434050dece4fbac1c94629bf30b64db"></a>

<a id="canonical-63c858fbd230998d3384c8bdd2c1caa31e7282d375c3611e1e1cbcf9b6a0cb20"></a>

## flags property — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 4

Type: `"string"`. Optional.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-03eb2b6a420a1dc4e1b25b286b85ae38f5f208230066ff0635b8e47ddbdc7fae"></a>

<a id="canonical-fab075c6cf60b2ad82a307e1779219f85ed7a8462341235dbd7858ea7e484c38"></a>

## order property — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 5

Type: `"number"`. Optional.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-23bd86196bcddf126bbb4782f1a574d0f6bcfcb50523ef9105eefc5ee9a50db3"></a>

<a id="canonical-e2d8fac05a2e84873186e0784bf003626aa1f1d3cfa56bf3a1a876291b618348"></a>

## preference property — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 6

Type: `"number"`. Optional.

Preference when records have the same order. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-5f44a64cea55cbbbc3dab71cc05755781225a286db4b4f49bdd2490a507961d8"></a>

<a id="canonical-a0c9f830a60a72e3dc3de47fe492546fa9481a18ec34a8f85d9b4483c31a0fc5"></a>

## regexp property — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 7

Type: `"string"`. Optional.

Regular expression to construct the next domain name to lookup.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-730a8606cfd6b899a434e53828353c715462c3753499595e112bddf269cd429d"></a>

<a id="canonical-15c54a7b611b0b4f4c216b8235db95201c570080a98d74e1d0976acf7260275b"></a>

## replacement property — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 8

Type: `"string"`. Optional.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-10e80ffc47cbf4dea28ab4b6515b6367f426a526683ea820de32b46018716c69"></a>

<a id="canonical-a2f5a9ee904d2a5b7dbd587f75e146299332d8506dfb05d240f751a30f87c879"></a>

## service property — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 9

Type: `"string"`. Optional.

Specifies the service(s) available down this rewrite path.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-2ce0fab442f4ab3dd20d2a86b4f63dc0f558d9a7efb6698f2d4d75ff4c801bbc"></a>

## Next pages — primary.default_rr_set_group.naptr_record.values / c1cf03f21cf6 / 10

- [primary.default_rr_set_group.naptr_record](resources--dns_zone--reference--group-002.md#canonical-29861e75cf9362a30496f5183138a125bae01dd5822188db7f03bc58fccf8ff2)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-781801294746217c2a3523dd3f86ab70b44e13a8ea217f252c073dad72ab2732"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e7cf55e4d140e91af276983142513cf3abf7d8c34b456c7c39c19d9457818f9"></a>

## primary.default_rr_set_group.ns_record — primary.default_rr_set_group.ns_record / fda55f0e4970 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.ns_record

<a id="canonical-dfed0301fbe283864bf9bd2b9fbeb55a94f74a45eb7554e38a91103e063c7d4d"></a>

Type: `"object"`. single nested block, Optional.

DNSNSResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ns_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-552dd1d5cdac92cd4462f01ff88ad1954807330802fb544eb2aead8821aab776"></a>

## Direct properties — primary.default_rr_set_group.ns_record / fda55f0e4970 / 3

<a id="canonical-8f0d5f6e90e707e21635ecce39b204e2ff11ad6bf8b4cdf8b0249d15959962d4"></a>

<a id="canonical-bb9eb8958c72fb8ad294b59e07686b0c71d03476df7257f330c502890889aff8"></a>

## name property — primary.default_rr_set_group.ns_record / fda55f0e4970 / 4

Type: `"string"`. Optional.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-d7ff7ae2fa88e9c462bc0fa0ea4fa5d2aa7f3c6f388e4c9ded55cc42b66ca912"></a>

<a id="canonical-55e375f3edcaed08edf902c203fc25fe4ef61f830ea98d84587c7f29f1535511"></a>

## values property — primary.default_rr_set_group.ns_record / fda55f0e4970 / 5

Type: `["list", "string"]`. Optional.

Name Servers. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-474b350d9af11cab07a18157d71e4c3e935a82e5ea212c51153ed87b062eec05"></a>

## Next pages — primary.default_rr_set_group.ns_record / fda55f0e4970 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-cc42794f9714583aafa8ea4d6820d1615e7962dce5b458be9befc0a7a0b921ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bed46ad49b643c0aab064cc0bd126ced2d06b73ded25435119404b87fc7c8d3d"></a>

## primary.default_rr_set_group.ptr_record — primary.default_rr_set_group.ptr_record / 4deff00fe505 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.ptr_record

<a id="canonical-807291065fb9cfa38b6c64db67fee14fc9efde9e28e255ffb56390a12222486d"></a>

Type: `"object"`. single nested block, Optional.

DNSPTRResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
ptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0cc3509a0eca6bc62ff67de77ecceb733433244a260e4e4d52e9368322d15c8d"></a>

## Direct properties — primary.default_rr_set_group.ptr_record / 4deff00fe505 / 3

<a id="canonical-52403f8cc3e81e7dae112077eb51259578ee73b95e9acd42dcb700941e3f3cc0"></a>

<a id="canonical-8e336c63745fd0d4837d594ecb01b280f52179b81d9ecc67fc6fd9a4d0b536c6"></a>

## name property — primary.default_rr_set_group.ptr_record / 4deff00fe505 / 4

Type: `"string"`. Optional.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-26ecdee7f42f5bdcc7493a1c8ca348e4e9343f20ec41278d1460d3daca6ca1ec"></a>

<a id="canonical-c63ceabc35f796cd812649561e1cb3095ce49e4949830b4d21b06550668624e5"></a>

## values property — primary.default_rr_set_group.ptr_record / 4deff00fe505 / 5

Type: `["list", "string"]`. Optional.

Domain Name. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a2933deee3fa92654340ce5086e4a0f1473e2a0114425b03a551f82965d4a3f6"></a>

## Next pages — primary.default_rr_set_group.ptr_record / 4deff00fe505 / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-1706c61fb364648d60d5a0aecef6e3fb6011b6a966e1b65b079e1ef8f7de562f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f8e8f5b87100d5652efdf4e998dc81d840a36f557057a09c4ceeaebcc9321a4"></a>

## primary.default_rr_set_group.srv_record — primary.default_rr_set_group.srv_record / eb2ed2e72b17 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.srv_record

<a id="canonical-ee78b04d1479134cb6e3b100325c9637a0cee5180efb3f38a0b43a06c255cf04"></a>

Type: `"object"`. single nested block, Optional.

DNSSRVResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "values")}
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
srv_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-a112a9f5edae8c237eb0d6d6b39be2366368caf9c44d5a389da1a78243b2a797"></a>

## Direct properties — primary.default_rr_set_group.srv_record / eb2ed2e72b17 / 3

<a id="canonical-892b2e142ed598d71de4cabef9c1bc4aa2987574c21aca20c4652d0265254250"></a>

<a id="canonical-4ac99a8b65803e73649b2b8d5f7c79ba046b5f984168b6f5371eda0187669dab"></a>

## name property — primary.default_rr_set_group.srv_record / eb2ed2e72b17 / 4

Type: `"string"`. Optional.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-2882ce77717042791597b7c9afc89ebe354541336a84ea5040696b4f93b5ff2e): complete subsection reference.

<a id="canonical-fac30062f2f1c657751667d50cf2ede6fcd0dc4ba96b17711f53b8afeb670f03"></a>

## Next pages — primary.default_rr_set_group.srv_record / eb2ed2e72b17 / 5

- [primary.default_rr_set_group.srv_record.values](resources--dns_zone--reference--group-002.md#canonical-2882ce77717042791597b7c9afc89ebe354541336a84ea5040696b4f93b5ff2e)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-2882ce77717042791597b7c9afc89ebe354541336a84ea5040696b4f93b5ff2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9647371ff78054cace741cc69ea5e43b7e188c2cc74c54f00b78b3971bd7fab"></a>

## primary.default_rr_set_group.srv_record.values — primary.default_rr_set_group.srv_record.values / dd04e9086bf3 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-1706c61fb364648d60d5a0aecef6e3fb6011b6a966e1b65b079e1ef8f7de562f)
- primary.default_rr_set_group.srv_record.values

<a id="canonical-6f51465f25fa658a28feb5082a46701b6be7e3cc66862ffd2a90e0f31fa35128"></a>

Type: `"object"`. list nested block, Optional.

SRV Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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

<a id="canonical-11f5fc7f4c5409e932d0f15808667257dd3a01c54832d431076a68425225d475"></a>

## Direct properties — primary.default_rr_set_group.srv_record.values / dd04e9086bf3 / 3

<a id="canonical-c8926f76973e7c43fe3242c26af502f5e73f43e5c84663c8c2c4066f0bc99937"></a>

<a id="canonical-7ccc9a7afe34ab041e72fa5bb07ca2fc647bc207ae19e372bbe62774e8bfb307"></a>

## port property — primary.default_rr_set_group.srv_record.values / dd04e9086bf3 / 4

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-266a72825273cb0a974f4d62ca9c3e55cd68dae4b23bd68e0faf8312c93fa042"></a>

<a id="canonical-a6081ba26407c890e4e992c9b762a5e63b31322dc65d5d083b9bdd8bdc2a2e3e"></a>

## priority property — primary.default_rr_set_group.srv_record.values / dd04e9086bf3 / 5

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-ebe90cf90d41a1802d631a3bd1ec8968f92cdf8c80bc3c4a46c6e9b92b7624ac"></a>

<a id="canonical-91369aa13541153be9c245e2afb4e5fb782dbc07a875c147038063a7e7c81d6b"></a>

## target property — primary.default_rr_set_group.srv_record.values / dd04e9086bf3 / 6

Type: `"string"`. Optional.

Hostname of the machine providing the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-2be083b3b7e09102505865e3aa594af7285fdadd7c581c9752ae30dfb3791c52"></a>

<a id="canonical-8022d1e8343b81cb05137118e3a1377967d1cac4693856431fb34b7d1ea4bab2"></a>

## weight property — primary.default_rr_set_group.srv_record.values / dd04e9086bf3 / 7

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-b02a4aa59479c2d07b9e3efcf363eb27a61078517743ee7b51d7554a4abd3a6f"></a>

## Next pages — primary.default_rr_set_group.srv_record.values / dd04e9086bf3 / 8

- [primary.default_rr_set_group.srv_record](resources--dns_zone--reference--group-002.md#canonical-1706c61fb364648d60d5a0aecef6e3fb6011b6a966e1b65b079e1ef8f7de562f)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-027f21fbf6635b09d7990db940a88858d0819a3562710a5cfec958515218b6e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-462b2bda56ab14cc1f7db0b21feccc39e783504b90ac911a50d7eaf5b0f0cedc"></a>

## primary.default_rr_set_group.sshfp_record — primary.default_rr_set_group.sshfp_record / f8ca938b5f99 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.sshfp_record

<a id="canonical-0fd30bf7c1875249db91f9c537b8225215894cf72c7cd23b94b26cf0f8dc7271"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sshfp record.

Upstream description:

DNS SSHFP Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
sshfp_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-bed37c82a085648c407c86a7ac43702ef2db5a74f0b85d18b192f7b1fa488037"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record / f8ca938b5f99 / 3

<a id="canonical-76ec0625dadda3c1b142ed0604077f5953ee57cb55917294178508097b910665"></a>

<a id="canonical-82fe4699eccd9888c31052e9e5ff66b7fdaa5bb2e18f2deec47ebdb4d6f64bf9"></a>

## name property — primary.default_rr_set_group.sshfp_record / f8ca938b5f99 / 4

Type: `"string"`. Optional.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-51639c33b8690b40750a67cd194f7f4bd36bbc3e2256448f31f2227006118766): complete subsection reference.

<a id="canonical-8f8c0e066fb9616bf5d6415460a23e1eafd3d4114d3003d3f7ed40d6bc63da58"></a>

## Next pages — primary.default_rr_set_group.sshfp_record / f8ca938b5f99 / 5

- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-51639c33b8690b40750a67cd194f7f4bd36bbc3e2256448f31f2227006118766)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-51639c33b8690b40750a67cd194f7f4bd36bbc3e2256448f31f2227006118766"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2aaa4578061d3b4c56055670802faa4e121a47afa01c2387e2bfd5e0eabbd8e7"></a>

## primary.default_rr_set_group.sshfp_record.values — primary.default_rr_set_group.sshfp_record.values / d6167bfd4547 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-027f21fbf6635b09d7990db940a88858d0819a3562710a5cfec958515218b6e9)
- primary.default_rr_set_group.sshfp_record.values

<a id="canonical-3b834ce5df8abc2cffb5fb163a98eb0558f04369d3132dd2f4ac9f222ae7b0ff"></a>

Type: `"object"`. list nested block, Optional.

SSHFP Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-6d39809f0a8b7ea0ed6e0c903dd6c454aa3137d8a12226709202a196bfba59ea"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record.values / d6167bfd4547 / 3

<a id="canonical-10db005a50e8c44fc7290f9d2f9684a9de7bb779c702ab31490bb808784570e3"></a>

<a id="canonical-cb7f518ee10ac55e9a68ddb2c1f80c819412df4a64a7125b3288c4774111cd75"></a>

## algorithm property — primary.default_rr_set_group.sshfp_record.values / d6167bfd4547 / 4

Type: `"string"`. Optional.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Upstream description:

SSHFP algorithm value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM

&#8203;- RSA: RSA

&#8203;- DSA: DSA

&#8203;- ECDSA: ECDSA

&#8203;- Ed25519: Ed25519

&#8203;- Ed448: Ed448.

Provider validators and defaults (from schema source):

```go
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

- [sha1_fingerprint](resources--dns_zone--reference--group-002.md#canonical-8bb4f7579489b74863561cbe779815f1ff5f8f498864141b28d846f480c36614): complete subsection reference.

- [sha256_fingerprint](resources--dns_zone--reference--group-002.md#canonical-d32be527d037b986290c80ee2cce53310f0cc5c20e5c3bbabb064fd3495ea9d5): complete subsection reference.

<a id="canonical-5ea555965eb798ceffb2423a103ab439be419af9c0cf27c996d63d703fac3a32"></a>

## Next pages — primary.default_rr_set_group.sshfp_record.values / d6167bfd4547 / 5

- [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-002.md#canonical-8bb4f7579489b74863561cbe779815f1ff5f8f498864141b28d846f480c36614)
- [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-002.md#canonical-d32be527d037b986290c80ee2cce53310f0cc5c20e5c3bbabb064fd3495ea9d5)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-027f21fbf6635b09d7990db940a88858d0819a3562710a5cfec958515218b6e9)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-8bb4f7579489b74863561cbe779815f1ff5f8f498864141b28d846f480c36614"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-794624b27350dd41d411e35daa3948baf0fa9e307edc06daebe473f079879e55"></a>

## primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 51c72f8dc294 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-027f21fbf6635b09d7990db940a88858d0819a3562710a5cfec958515218b6e9)
- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-51639c33b8690b40750a67cd194f7f4bd36bbc3e2256448f31f2227006118766)
- primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint

<a id="canonical-f0b12d0e07d64327d5d8522a9a6c2396a9c87f31e247be32f21a726ab895dfc4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 fingerprint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
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
sha1_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-4503330c39dfe6f90565c512b072d8cc95e2833efbc798109036be20147adb2a"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 51c72f8dc294 / 3

<a id="canonical-ecf32208677aa77b02118ecdef17b087a18313a019ea116a77a5d6e4819a5e5a"></a>

<a id="canonical-e6b06f2dfd7b8eb457a9902a429225f0a0a88e5e0d2863db0ff53c66c04f6fea"></a>

## fingerprint property — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 51c72f8dc294 / 4

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-f397491468abadcf23f5fcb99f26aa9b92536622c8dd719191e0cbf2c773fa97"></a>

## Next pages — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 51c72f8dc294 / 5

- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-51639c33b8690b40750a67cd194f7f4bd36bbc3e2256448f31f2227006118766)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-d32be527d037b986290c80ee2cce53310f0cc5c20e5c3bbabb064fd3495ea9d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b331bd66e48d1607b68fe6ba3f36b4c500a990a0f93b0343c0e277f907e42ea7"></a>

## primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 6d86a2569516 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.sshfp_record](resources--dns_zone--reference--group-002.md#canonical-027f21fbf6635b09d7990db940a88858d0819a3562710a5cfec958515218b6e9)
- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-51639c33b8690b40750a67cd194f7f4bd36bbc3e2256448f31f2227006118766)
- primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint

<a id="canonical-1bcaf9f2f5ba78cd87b099f5c0680865084e889559b268b6a13a071bf3e13dd2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 fingerprint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
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
sha256_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-339dcec2db8a8fce7ef70ccf038b00364ea27b5893ae7361cac286a119114b95"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 6d86a2569516 / 3

<a id="canonical-7286c2536ebe17bb4c1bf9e05fe034acfb305b536c2b548e8001a8abf25b40a0"></a>

<a id="canonical-24e6179d8e9e6f7281f8f80784e76976278fe4ba0d0a872eaa5c5e537ebbe8c5"></a>

## fingerprint property — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 6d86a2569516 / 4

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-fa2d1d3b93180a926ca4f0af40aaad69b40adf67488a43ec99691e216efbe9e8"></a>

## Next pages — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 6d86a2569516 / 5

- [primary.default_rr_set_group.sshfp_record.values](resources--dns_zone--reference--group-002.md#canonical-51639c33b8690b40750a67cd194f7f4bd36bbc3e2256448f31f2227006118766)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-549151928d884e1b5ee19bbe822cdb0e8a9cebaf669d1ae353e4aa003c5bb1da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf5dbb4d3b82b354c97c3066e550d4e86ac753545bba21e28096f8ebc8c95293"></a>

## primary.default_rr_set_group.tlsa_record — primary.default_rr_set_group.tlsa_record / d0d92c6a72f9 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.tlsa_record

<a id="canonical-7852f0cbf82c87e7752c2701c7bafc40b7acdf7562f056eac5af057e4a7ac413"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tlsa record.

Upstream description:

DNS TLSA Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
tlsa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-fbdcf17778dc46c1ab133eebd7dd5c786fdab021f4e3c0db0bbd2b7ab78bf08c"></a>

## Direct properties — primary.default_rr_set_group.tlsa_record / d0d92c6a72f9 / 3

<a id="canonical-6bde7fa9b9b7f10792fe47688d3e14ad174747f466c91137bf3f334d72f57b27"></a>

<a id="canonical-0da94ccf05ca4b74e98cb8600602ef1ff9181bddc905c0117020d0ff3c90ff7c"></a>

## name property — primary.default_rr_set_group.tlsa_record / d0d92c6a72f9 / 4

Type: `"string"`. Optional.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-53be77ddb81c8e401179c710726470411115d2bf67f8bc61f101b5f022b5e04d): complete subsection reference.

<a id="canonical-62f639780e676ca343348daf3eebfdf0bea76fd6700ccca3e3c953bcfd9e8923"></a>

## Next pages — primary.default_rr_set_group.tlsa_record / d0d92c6a72f9 / 5

- [primary.default_rr_set_group.tlsa_record.values](resources--dns_zone--reference--group-002.md#canonical-53be77ddb81c8e401179c710726470411115d2bf67f8bc61f101b5f022b5e04d)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-53be77ddb81c8e401179c710726470411115d2bf67f8bc61f101b5f022b5e04d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf45f337a922e8bc148655cbca0960b420f609016c7f8cd45696bf1741b961ee"></a>

## primary.default_rr_set_group.tlsa_record.values — primary.default_rr_set_group.tlsa_record.values / 1473ab6ed43d / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-549151928d884e1b5ee19bbe822cdb0e8a9cebaf669d1ae353e4aa003c5bb1da)
- primary.default_rr_set_group.tlsa_record.values

<a id="canonical-8faf6eb248629989492dcb2ed2ea43e55f91dcc0a2b420a190096a3797d74fef"></a>

Type: `"object"`. list nested block, Optional.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_association_data")}
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

<a id="canonical-f32cb9cfc680e907bccee07b3b58db9a4aa136b9365bbf210a99f6871fefe213"></a>

## Direct properties — primary.default_rr_set_group.tlsa_record.values / 1473ab6ed43d / 3

<a id="canonical-a2d596fe8bb4dbf1f31b02b1083e7d66a4cbe490b59971c8ebec0989774398af"></a>

<a id="canonical-2403999fc7229663114c66a1e54725fb59406e7f08be822efb1f4d263fc03e69"></a>

## certificate_association_data property — primary.default_rr_set_group.tlsa_record.values / 1473ab6ed43d / 4

Type: `"string"`. Optional.

The actual data to be matched given the settings of the other fields.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1728f6bff7d5d58ecc2b8670bfcb9b0a56ed09dd6eab5ab4aa947db80e606d94"></a>

<a id="canonical-50ed4f852724bc88694379e62a41b0fb86173404bd9d3f18e946d31529167787"></a>

## certificate_usage property — primary.default_rr_set_group.tlsa_record.values / 1473ab6ed43d / 5

Type: `"string"`. Optional.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4b892046d6b20111b985e40f055edf027b29245f6ed5eea383caa2f0d638accb"></a>

<a id="canonical-2a77ed30ed994b39d44a210a1968c6af6d4f53ebfbef90ce2f1904b6f7cf1169"></a>

## matching_type property — primary.default_rr_set_group.tlsa_record.values / 1473ab6ed43d / 6

Type: `"string"`. Optional.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NoHash",
    "SHA256",
    "SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-63aa0f3a70d812b3106bb95d9a8fcfd330d198468302dd8fba83f07f70217d75"></a>

<a id="canonical-e6317aa64c3645a93e0f5bb5d57b4e8d45c7cbac8b855de7798be393e245aaef"></a>

## selector property — primary.default_rr_set_group.tlsa_record.values / 1473ab6ed43d / 7

Type: `"string"`. Optional.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("FullCertificate",
    "UseSubjectPublicKey"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c29225fcd6268519fea15af6a483800e68c54247b62861e25428544ccf1ca83a"></a>

## Next pages — primary.default_rr_set_group.tlsa_record.values / 1473ab6ed43d / 8

- [primary.default_rr_set_group.tlsa_record](resources--dns_zone--reference--group-002.md#canonical-549151928d884e1b5ee19bbe822cdb0e8a9cebaf669d1ae353e4aa003c5bb1da)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-fdb7631ec4d613093c3880c355db38c0eec42adcd38295c86269c3657197c858"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0c133571a482789533ae658462f5d3be129ec34cb309781a6ac41e16a0df0eb"></a>

## primary.default_rr_set_group.txt_record — primary.default_rr_set_group.txt_record / 23e54837a4fa / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- primary.default_rr_set_group.txt_record

<a id="canonical-51eb4b36eacea5e3ba339251f95e878b86ed4f201376c314f06494824a96b4cd"></a>

Type: `"object"`. single nested block, Optional.

DNSTXTResourceRecord.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
txt_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2aa111812454be35772ba64a7bebc48b12fe13235a04c5a4b388eb51ee61ba28"></a>

## Direct properties — primary.default_rr_set_group.txt_record / 23e54837a4fa / 3

<a id="canonical-4968725855434b4643f7b2034a2ccdc7e31b857736974d8e805a1665f0b99801"></a>

<a id="canonical-160618e407bc467693dce35ead470498ab740f6ae58e2fd372f26273ad4902e9"></a>

## name property — primary.default_rr_set_group.txt_record / 23e54837a4fa / 4

Type: `"string"`. Optional.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-e3e41e3d2e77427ff8f6a06044316afd824e365605175f9c4406490c46d5b60b"></a>

<a id="canonical-165391258341da33dd77393c29774598b7d391f3e6c07d62172e5223f495a46e"></a>

## values property — primary.default_rr_set_group.txt_record / 23e54837a4fa / 5

Type: `["list", "string"]`. Optional.

Text. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d1cdc66678f57e8fca4825fb98110416ec2f91e1e23a554cfc574f72c6f3b507"></a>

## Next pages — primary.default_rr_set_group.txt_record / 23e54837a4fa / 6

- [primary.default_rr_set_group](resources--dns_zone--reference--group-001.md#canonical-986c5e57e32d476396c311d1b129867ed4d39420ea5228cbcade13276249814d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-011cd0aebba1ce421435bec972cff46926c8a27bfd644b58dd735e03f67530f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5c92485e400cd01988ad17a38e006909044d9f3fd72118bef1b390567a01e01"></a>

## primary.default_soa_parameters — primary.default_soa_parameters / 75ed79c505a5 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- primary.default_soa_parameters

<a id="canonical-e095b6512a3a9b738e0020e504d7db4fcef2c0cd5f503558aa285f3fd21dd86a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default soa parameters.

Upstream description:

This can be used for messages where no values are needed.

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
default_soa_parameters = {}
```

<a id="canonical-e83aec62ec9fee655d81805ce90910ab498250fe29e0863526e99acd67ea1e26"></a>

## Direct properties — primary.default_soa_parameters / 75ed79c505a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bfc4e908160d14fc0affe1a2d46816c2488be63cd5c2a2f20c585cb5df092b02"></a>

## Next pages — primary.default_soa_parameters / 75ed79c505a5 / 4

- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-a0163bba2e298b8beeb8b130e64bd81043610e52aeb892e347bb2aab203f53d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71e5ce8fa8b9b9f36d4e9c5fd095e55ce5fbf253b9cfd51a5255fd04ab055429"></a>

## primary.dnssec_mode — primary.dnssec_mode / 6dc2e28cdb33 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- primary.dnssec_mode

<a id="canonical-8420098636cb935a0db948198999de0e4ce6f8e65a9b01c432e93c44b17b0f10"></a>

Type: `"object"`. single nested block, Optional.

DNSSEC Mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
dnssec_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-536cc03d8b038cc0574256f92d7d199abac1dd2caed547ba5f7738f3245e53d0"></a>

## Direct properties — primary.dnssec_mode / 6dc2e28cdb33 / 3

- [disable_spec](resources--dns_zone--reference--group-002.md#canonical-94545c0514aa0880bd044c1f914852c5e45471c281f403ab5ad7d7234c9393b3): complete subsection reference.

- [enable](resources--dns_zone--reference--group-002.md#canonical-74dc6f9a4c09fcd94348713bfe5bbc6bd4d8c6a33625f84e50ee205e9d731cf1): complete subsection reference.

<a id="canonical-1a08023d720a6a889229722909a931fc9f50c5d1dbfefc7312f55817e672d536"></a>

## Next pages — primary.dnssec_mode / 6dc2e28cdb33 / 4

- [primary.dnssec_mode.disable_spec](resources--dns_zone--reference--group-002.md#canonical-94545c0514aa0880bd044c1f914852c5e45471c281f403ab5ad7d7234c9393b3)
- [primary.dnssec_mode.enable](resources--dns_zone--reference--group-002.md#canonical-74dc6f9a4c09fcd94348713bfe5bbc6bd4d8c6a33625f84e50ee205e9d731cf1)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-94545c0514aa0880bd044c1f914852c5e45471c281f403ab5ad7d7234c9393b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0ce0bb0a0e0a3cabdb14d7c8b571ca9c690f4544f2f0871546d49333bc20e05"></a>

## primary.dnssec_mode.disable_spec — primary.dnssec_mode.disable_spec / 2ff2f4c6756c / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-a0163bba2e298b8beeb8b130e64bd81043610e52aeb892e347bb2aab203f53d2)
- primary.dnssec_mode.disable_spec

<a id="canonical-12161a39eb4a6edf5d9c31b9b0a896fb95df47794f43279921d14fc27cc5a329"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-41b4a81a8fdf6ecc1feaec3996b8aa11899939535dc02374a879bc3e8ae0e63e"></a>

## Direct properties — primary.dnssec_mode.disable_spec / 2ff2f4c6756c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b355d55af79ba753026374604055eb6370045f04f2387fa4aaa4854dd0578e73"></a>

## Next pages — primary.dnssec_mode.disable_spec / 2ff2f4c6756c / 4

- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-a0163bba2e298b8beeb8b130e64bd81043610e52aeb892e347bb2aab203f53d2)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-74dc6f9a4c09fcd94348713bfe5bbc6bd4d8c6a33625f84e50ee205e9d731cf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c5e61292730d82f27fe1b3291d9b33871029c26c97cd211c7b0b33ebf696b26"></a>

## primary.dnssec_mode.enable — primary.dnssec_mode.enable / 23c1dd4ace00 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-a0163bba2e298b8beeb8b130e64bd81043610e52aeb892e347bb2aab203f53d2)
- primary.dnssec_mode.enable

<a id="canonical-a315e64ebb2dddba4926df681d252f8fe0f797565c3281d1b52ca5c320a8d4ab"></a>

Type: `["object", {}]`. Optional.

Enable. DNSSEC enable.

Upstream description:

DNSSEC enable.

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
enable = {}
```

<a id="canonical-0996fa5e9e0af05cbe22471a3f37d73389daab6c24cda9318992e582b14c7c9b"></a>

## Direct properties — primary.dnssec_mode.enable / 23c1dd4ace00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a0b77229b2b640460d4fbdc2e61d3043a0711555bd6f8e1548ad804143d6992"></a>

## Next pages — primary.dnssec_mode.enable / 23c1dd4ace00 / 4

- [primary.dnssec_mode](resources--dns_zone--reference--group-002.md#canonical-a0163bba2e298b8beeb8b130e64bd81043610e52aeb892e347bb2aab203f53d2)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c2e1d0f392f5953a23daee9e821db6451390749e8516fc04a99c991934f3f47"></a>

## primary.rr_set_group — primary.rr_set_group / 2c5df548d2e0 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- primary.rr_set_group

<a id="canonical-ba051d9ab361e297286390078ab285b50cfc51e3aed2e56572cfc4dfa9833b5b"></a>

Type: `"object"`. list nested block, Optional.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rr_set_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc9b078ed2a0eebe8044f848a300fada6a2feb02dd0af19b76754c7989ca62c7"></a>

## Direct properties — primary.rr_set_group / 2c5df548d2e0 / 3

- [metadata](resources--dns_zone--reference--group-002.md#canonical-0c9b8611b937a4a22ecba4980c4d193fbb248603f02e62a88389fb61719901f3): complete subsection reference.

- [rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237): complete subsection reference.

<a id="canonical-e5c4f16e7f6aff592e63f3406361c82f020ff752003592422892d695418b083c"></a>

## Next pages — primary.rr_set_group / 2c5df548d2e0 / 4

- [primary.rr_set_group.metadata](resources--dns_zone--reference--group-002.md#canonical-0c9b8611b937a4a22ecba4980c4d193fbb248603f02e62a88389fb61719901f3)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-0c9b8611b937a4a22ecba4980c4d193fbb248603f02e62a88389fb61719901f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b15d58ac58d0671b1d19ed0d59c5a5d22f17d6f034045f5a79b0037cff000fdc"></a>

## primary.rr_set_group.metadata — primary.rr_set_group.metadata / aa694ef342cf / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- primary.rr_set_group.metadata

<a id="canonical-4213c2db1f14c049218d6f76e32da858c4416482a5f6ccf3d0849845318a1570"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-562ebe591ac2db8b1fa55ce39f637a5c95b2fde73be2d3f802df34545c0a62c3"></a>

## Direct properties — primary.rr_set_group.metadata / aa694ef342cf / 3

<a id="canonical-7a10cd34072cadd9797218e748cc495e7fb91ec9c1b8f0a0c9f76ce220a317e3"></a>

<a id="canonical-73156c61c1d254887295a220aafaf2a6a6db35734a00489dc4063df5d3f4cdb9"></a>

## description_spec property — primary.rr_set_group.metadata / aa694ef342cf / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-8790e588869b7922f402c751d2bd03243fd6e79ec9dd7bc4142bfbc2b09c8a06"></a>

<a id="canonical-a17013ccf5fcfb4fc21e395f16cb32bf743099f5499f9eb8609d28c0d7ed8e16"></a>

## name property — primary.rr_set_group.metadata / aa694ef342cf / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-c4ede846d7d452220610f23f7a0b9e0a12420b08ab1718016b096a1e289988ea"></a>

## Next pages — primary.rr_set_group.metadata / aa694ef342cf / 6

- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-143f7279fcc33e231fd42f4340d6170cde2a6fee23e96edb222fd7ceef2aa9bd"></a>

## primary.rr_set_group.rr_set — primary.rr_set_group.rr_set / 551b29e38fe2 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- primary.rr_set_group.rr_set

<a id="canonical-79799cf7b30e119a5eac90eac36a5a208ee438e2a64272523bfd74e1854c9651"></a>

Type: `"object"`. list nested block, Optional.

Resource Record Sets. Collection of DNS resource record sets.

Upstream description:

Collection of DNS resource record sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ttl"),
  validators.ConflictingListObjectAttributes("a_record",
    "aaaa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("a_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "afsdb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("aaaa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "alias_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("afsdb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "caa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("alias_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("caa_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cert_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "cname_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cert_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ds_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("cname_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui48_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ds_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "eui64_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui48_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "lb_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("eui64_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "loc_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("lb_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "mx_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("loc_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "naptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("mx_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ns_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("naptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "ptr_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ns_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "srv_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("ptr_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "sshfp_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("srv_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "tlsa_record"),
  validators.ConflictingListObjectAttributes("sshfp_record",
    "txt_record"),
  validators.ConflictingListObjectAttributes("tlsa_record",
    "txt_record")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

Terraform syntax:

```terraform
rr_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-f0ee51cc85e020ac42fd97277a5eb2c78f2e65a9bfdc2a94613de41240e70084"></a>

## Direct properties — primary.rr_set_group.rr_set / 551b29e38fe2 / 3

- [a_record](resources--dns_zone--reference--group-002.md#canonical-58a25a45eceba3c56f3a3a39525416be7e65f8714fea8818b72efeb9c4b265c3): complete subsection reference.

- [aaaa_record](resources--dns_zone--reference--group-002.md#canonical-7f2d2684d4faec41988c68e8c54a0b00ebfd3dd47db15ab7632cf0391627bef3): complete subsection reference.

- [afsdb_record](resources--dns_zone--reference--group-002.md#canonical-4dc4c9b00f3a1a922fb17c7214ad1ad6ee3bd7b4091ea45539d44fc2f59cc87c): complete subsection reference.

- [alias_record](resources--dns_zone--reference--group-002.md#canonical-2fa41e86d64f3ba676454454d9f701d605082ade6d0d2aad66859b3ffd2021b8): complete subsection reference.

- [caa_record](resources--dns_zone--reference--group-002.md#canonical-cf848ffc0e401aea9da9672384ad97dae88717076541f41d5528f704f564b350): complete subsection reference.

- [cds_record](resources--dns_zone--reference--group-002.md#canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f): complete subsection reference.

- [cert_record](resources--dns_zone--reference--group-003.md#canonical-9a7828a9b291db43cc63061b9ff7ab475c2d5736cdd55e78a2c55d8331c83211): complete subsection reference.

- [cname_record](resources--dns_zone--reference--group-003.md#canonical-5371974e821ae367da729184bd9adb0dd4371d3d39f0c278e60572e945875efa): complete subsection reference.

<a id="canonical-9ff8e9a41b9eca99db5ee61a44fa5bd8d0e3d7183c700a0e974fdd8c7ab4dba5"></a>

<a id="canonical-05d517129b90234f8d9c1acd4a5ae085691200a5b577d3a8faef2901f63b3478"></a>

## description_spec property — primary.rr_set_group.rr_set / 551b29e38fe2 / 4

Type: `"string"`. Optional.

Comment. Human-readable description text

- [ds_record](resources--dns_zone--reference--group-003.md#canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb): complete subsection reference.

- [eui48_record](resources--dns_zone--reference--group-003.md#canonical-cc21bd523c0b066121389cde3dd32c5982bbc85de11047c259a20d1d194410e1): complete subsection reference.

- [eui64_record](resources--dns_zone--reference--group-003.md#canonical-5bb96ff6312f0f9a10bef3014b08ea4c07368ac1b1d153fd187cfd6316e078cf): complete subsection reference.

- [lb_record](resources--dns_zone--reference--group-003.md#canonical-8501fd8f15ca3bf884866766a26bb1357a6f17826763d164f5c395e1e3b77ee3): complete subsection reference.

- [loc_record](resources--dns_zone--reference--group-003.md#canonical-c3aa28a94f83cfd171303d2f7d99d864f23b1e65162c4a67600df832064ab922): complete subsection reference.

- [mx_record](resources--dns_zone--reference--group-003.md#canonical-8307a17271b7fe3063cb6a7c8fcca75ed9e1c50dbe43413f9c1574ce4ae344a2): complete subsection reference.

- [naptr_record](resources--dns_zone--reference--group-003.md#canonical-0d06865cae15e71aebaf490992eaf16baf306801ee5df6c7ac205445e30c9aeb): complete subsection reference.

- [ns_record](resources--dns_zone--reference--group-003.md#canonical-d64af960376e35b05c1d2b17852ee628721ac928e27617d40527087752d7e49e): complete subsection reference.

- [ptr_record](resources--dns_zone--reference--group-003.md#canonical-ee0343545539ce1c7c1282eb8b827c02dbeb555c99152d4b631c303af0039914): complete subsection reference.

- [srv_record](resources--dns_zone--reference--group-003.md#canonical-fd47b913cf7776b21ce65e86af95d67b63abfca6bec77f6fdb114ffd8b4e4959): complete subsection reference.

- [sshfp_record](resources--dns_zone--reference--group-003.md#canonical-c9cddb199c143e2778d9e2884ca288256832b00a0f2f4d886a453c679ff0a375): complete subsection reference.

- [tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0db24ebdf992bc441effe1df0838893fc99bd28af63535fadf4301d50c3119d4): complete subsection reference.

<a id="canonical-a6e4f2f33b06dd83e43b40c32106049d1d4ee1159ad909cb7307028520494070"></a>

<a id="canonical-5415ffa74a33b6ff9ad1c3c98aa760d2abe2e77e8aa88860447f31009af35a2f"></a>

## ttl property — primary.rr_set_group.rr_set / 551b29e38fe2 / 5

Type: `"number"`. Optional.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](resources--dns_zone--reference--group-003.md#canonical-59755a283fb35414e8ffd5fcfc23876c4f29dc873fbb94d6d71dbb702cbaf497): complete subsection reference.

<a id="canonical-260b56198a2da317d1d9fed4d9408f4a162863a205c8e38f590de82aa560a29d"></a>

## Next pages — primary.rr_set_group.rr_set / 551b29e38fe2 / 6

- [primary.rr_set_group.rr_set.a_record](resources--dns_zone--reference--group-002.md#canonical-58a25a45eceba3c56f3a3a39525416be7e65f8714fea8818b72efeb9c4b265c3)
- [primary.rr_set_group.rr_set.aaaa_record](resources--dns_zone--reference--group-002.md#canonical-7f2d2684d4faec41988c68e8c54a0b00ebfd3dd47db15ab7632cf0391627bef3)
- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-4dc4c9b00f3a1a922fb17c7214ad1ad6ee3bd7b4091ea45539d44fc2f59cc87c)
- [primary.rr_set_group.rr_set.alias_record](resources--dns_zone--reference--group-002.md#canonical-2fa41e86d64f3ba676454454d9f701d605082ade6d0d2aad66859b3ffd2021b8)
- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-cf848ffc0e401aea9da9672384ad97dae88717076541f41d5528f704f564b350)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f)
- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-9a7828a9b291db43cc63061b9ff7ab475c2d5736cdd55e78a2c55d8331c83211)
- [primary.rr_set_group.rr_set.cname_record](resources--dns_zone--reference--group-003.md#canonical-5371974e821ae367da729184bd9adb0dd4371d3d39f0c278e60572e945875efa)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb)
- [primary.rr_set_group.rr_set.eui48_record](resources--dns_zone--reference--group-003.md#canonical-cc21bd523c0b066121389cde3dd32c5982bbc85de11047c259a20d1d194410e1)
- [primary.rr_set_group.rr_set.eui64_record](resources--dns_zone--reference--group-003.md#canonical-5bb96ff6312f0f9a10bef3014b08ea4c07368ac1b1d153fd187cfd6316e078cf)
- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-8501fd8f15ca3bf884866766a26bb1357a6f17826763d164f5c395e1e3b77ee3)
- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-c3aa28a94f83cfd171303d2f7d99d864f23b1e65162c4a67600df832064ab922)
- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-8307a17271b7fe3063cb6a7c8fcca75ed9e1c50dbe43413f9c1574ce4ae344a2)
- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0d06865cae15e71aebaf490992eaf16baf306801ee5df6c7ac205445e30c9aeb)
- [primary.rr_set_group.rr_set.ns_record](resources--dns_zone--reference--group-003.md#canonical-d64af960376e35b05c1d2b17852ee628721ac928e27617d40527087752d7e49e)
- [primary.rr_set_group.rr_set.ptr_record](resources--dns_zone--reference--group-003.md#canonical-ee0343545539ce1c7c1282eb8b827c02dbeb555c99152d4b631c303af0039914)
- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-fd47b913cf7776b21ce65e86af95d67b63abfca6bec77f6fdb114ffd8b4e4959)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-c9cddb199c143e2778d9e2884ca288256832b00a0f2f4d886a453c679ff0a375)
- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0db24ebdf992bc441effe1df0838893fc99bd28af63535fadf4301d50c3119d4)
- [primary.rr_set_group.rr_set.txt_record](resources--dns_zone--reference--group-003.md#canonical-59755a283fb35414e8ffd5fcfc23876c4f29dc873fbb94d6d71dbb702cbaf497)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-58a25a45eceba3c56f3a3a39525416be7e65f8714fea8818b72efeb9c4b265c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce411b1b5791d25eca11bbab910b889a5b5af8aa4df4a405cfc5f5b5d1db2243"></a>

## primary.rr_set_group.rr_set.a_record — primary.rr_set_group.rr_set.a_record / 3da20cb1ef51 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.a_record

<a id="canonical-8dd140157ae9c87cfbc41526dfd3c7c7241cf49da3b2080ca71e87a77149390d"></a>

Type: `"object"`. single nested block, Optional.

DNSAResourceRecord. A Records

Upstream description:

A Records

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
a_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-adb3e3d02773cdb26039138acc29a775196e479ab34be63fdcf8081828448463"></a>

## Direct properties — primary.rr_set_group.rr_set.a_record / 3da20cb1ef51 / 3

<a id="canonical-282ce1c269b09e98675174a6a6eb7532098f46073faadfe7f1a355d5c813c44d"></a>

<a id="canonical-4242cf1493d2c28efe1eb4c619196eb25900bcdfebcea7c68bd3425e0e2ebb33"></a>

## name property — primary.rr_set_group.rr_set.a_record / 3da20cb1ef51 / 4

Type: `"string"`. Optional.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-30282b52b230c3a3ccbf3bf92b3bdd8fbb302049f4828470ca53b1fb3fade7da"></a>

<a id="canonical-1a601db16de5c4de0b30ff44b04d0d3c85a4c5cb02bbf8d7ad9080e69a12f291"></a>

## values property — primary.rr_set_group.rr_set.a_record / 3da20cb1ef51 / 5

Type: `["list", "string"]`. Optional.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9d28ddbb7f806890153f1bc32e06a83819af1be37584d62488183df65eb42919"></a>

## Next pages — primary.rr_set_group.rr_set.a_record / 3da20cb1ef51 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-7f2d2684d4faec41988c68e8c54a0b00ebfd3dd47db15ab7632cf0391627bef3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24e8f33c67094b3d0b09e0734e5c4b4046069f558e03cc203e0dd1c0416d4e7c"></a>

## primary.rr_set_group.rr_set.aaaa_record — primary.rr_set_group.rr_set.aaaa_record / 52c84eff3812 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.aaaa_record

<a id="canonical-7b1128bd0d6eeaa2259c980bc6ac16205e543406b20055f7986986483f340f03"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
aaaa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-b68fb3d5d1cefaf9e58b70fe1781dcdcc0fcc8dc54ddda039d0e1973e55e53cb"></a>

## Direct properties — primary.rr_set_group.rr_set.aaaa_record / 52c84eff3812 / 3

<a id="canonical-ebd80a66c368b51a8308bc3f2517f58006c3f136aa521b7d2382975c1e66e36b"></a>

<a id="canonical-81b19810877a6135685495ca797e3423f9270692cf299c1b0969c3ffcdd33208"></a>

## name property — primary.rr_set_group.rr_set.aaaa_record / 52c84eff3812 / 4

Type: `"string"`. Optional.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-c8a71a044e9497b75ba1411c1a5dd599ad5ae07cd3dce415ae54e1714abde1e0"></a>

<a id="canonical-337d6d521ef83da382b5d0e5cbee5ff59b14e39ffd0ab39f5a9fa5e0e2abada7"></a>

## values property — primary.rr_set_group.rr_set.aaaa_record / 52c84eff3812 / 5

Type: `["list", "string"]`. Optional.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-46c84009cba1b2507f378625e02fab9af5ec294f67e984ebe556775fb418993a"></a>

## Next pages — primary.rr_set_group.rr_set.aaaa_record / 52c84eff3812 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-4dc4c9b00f3a1a922fb17c7214ad1ad6ee3bd7b4091ea45539d44fc2f59cc87c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b7f6655b3f11dd3ea7fa5909a74f77bb2d18a9a17704aa5774c9c83e4100377"></a>

## primary.rr_set_group.rr_set.afsdb_record — primary.rr_set_group.rr_set.afsdb_record / 1e04ac2e3f73 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.afsdb_record

<a id="canonical-66d3aaebb6048a48acab610fb8028d67037db69cb4cf1443d5c192aba7cd57ca"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
afsdb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-d0d41127c1a8521e8a9b248692a3ef6e33c74135b80c899efb581daad8abe3f5"></a>

## Direct properties — primary.rr_set_group.rr_set.afsdb_record / 1e04ac2e3f73 / 3

<a id="canonical-67ca590d9fe208c0154ace0469fb06192b0cde6a337baf030ca0c8d356acba0d"></a>

<a id="canonical-9d9d8e0e82446acc962c8335d8c7fb59b1f0fbbfd1c98f54d87bc849bad7370d"></a>

## name property — primary.rr_set_group.rr_set.afsdb_record / 1e04ac2e3f73 / 4

Type: `"string"`. Optional.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-f1ac78e7d0b8b8db3bd4de41035942f8b3d08b53ba7d2ec45c36a69b14cac775): complete subsection reference.

<a id="canonical-a6a9e53652e9b3ca1527ab65f0133d616a9a62005420405af5c2ccbeb6e70e39"></a>

## Next pages — primary.rr_set_group.rr_set.afsdb_record / 1e04ac2e3f73 / 5

- [primary.rr_set_group.rr_set.afsdb_record.values](resources--dns_zone--reference--group-002.md#canonical-f1ac78e7d0b8b8db3bd4de41035942f8b3d08b53ba7d2ec45c36a69b14cac775)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-f1ac78e7d0b8b8db3bd4de41035942f8b3d08b53ba7d2ec45c36a69b14cac775"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38af4f733448bd023b722539efc36a3f143ad40ee14f5212f59f96aba47ce9fd"></a>

## primary.rr_set_group.rr_set.afsdb_record.values — primary.rr_set_group.rr_set.afsdb_record.values / 65c05115d488 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-4dc4c9b00f3a1a922fb17c7214ad1ad6ee3bd7b4091ea45539d44fc2f59cc87c)
- primary.rr_set_group.rr_set.afsdb_record.values

<a id="canonical-b282af3970514edb73bc073ae1a9c82d03fbdfd9e1aad7092e5c089c2e7b6363"></a>

Type: `"object"`. list nested block, Optional.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("hostname")}
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

<a id="canonical-a01c99253ac4821d037af98459944934897261c1dcd51fee6a222b751c40b812"></a>

## Direct properties — primary.rr_set_group.rr_set.afsdb_record.values / 65c05115d488 / 3

<a id="canonical-920a10fc824691abc18d9d82e59591fcd5d9ab0a64cdb272d52ff1deb7a723fd"></a>

<a id="canonical-b2be9a7b9e7f007b0004f17e480bcb79e5f1a4e438c7bb3db6be4dd2fd0ba452"></a>

## hostname property — primary.rr_set_group.rr_set.afsdb_record.values / 65c05115d488 / 4

Type: `"string"`. Optional.

Server name of the AFS cell database server or the DCE name server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-8055fd07af621914a80b330494014cebc42d79cbdb404d0861b0843bcec5a839"></a>

<a id="canonical-c9e1b5a428de20a398af9617df67c8d818db1a41fb2457d65dfabb3aace2033f"></a>

## subtype property — primary.rr_set_group.rr_set.afsdb_record.values / 65c05115d488 / 5

Type: `"string"`. Optional.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-419367b48ec7666979edfc67cf6a91aec7c3173c653ef6e9832af6711bb745d6"></a>

## Next pages — primary.rr_set_group.rr_set.afsdb_record.values / 65c05115d488 / 6

- [primary.rr_set_group.rr_set.afsdb_record](resources--dns_zone--reference--group-002.md#canonical-4dc4c9b00f3a1a922fb17c7214ad1ad6ee3bd7b4091ea45539d44fc2f59cc87c)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-2fa41e86d64f3ba676454454d9f701d605082ade6d0d2aad66859b3ffd2021b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3d215b4d4f843bc9fb6231defbdfc9da44afc3503df324abfe20391fa90343c"></a>

## primary.rr_set_group.rr_set.alias_record — primary.rr_set_group.rr_set.alias_record / 6305cee3052e / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.alias_record

<a id="canonical-cb37358245a37f0df01bf2235362769b3e5106bb39fa7bc443c7e279fab5f890"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for alias record.

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
alias_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-ddffe48414ebdf8e3a6021f061d89a73243fa8b41f61dbff032f772530690f89"></a>

## Direct properties — primary.rr_set_group.rr_set.alias_record / 6305cee3052e / 3

<a id="canonical-e34f3ac0a97f1936774fc5c34034c8691a66add26920cd07c9145b8fed955666"></a>

<a id="canonical-0450201bb2922d58667b8bed6ee0cdb06e7f965927b9a4801fce87f1e7b533e5"></a>

## value property — primary.rr_set_group.rr_set.alias_record / 6305cee3052e / 4

Type: `"string"`. Optional.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-597b4d2493ca9a62b6963074a05fc20f08ef37a3a5c932995d55aff2dc290b7e"></a>

## Next pages — primary.rr_set_group.rr_set.alias_record / 6305cee3052e / 5

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-cf848ffc0e401aea9da9672384ad97dae88717076541f41d5528f704f564b350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b8ff887453b8cf121b82be26ea49c42b99b4b655ed3d0faa2f63d8ab0a7ea18"></a>

## primary.rr_set_group.rr_set.caa_record — primary.rr_set_group.rr_set.caa_record / be639c368bc3 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.caa_record

<a id="canonical-e32abe2d4e40424bd27ae885e935e982be381024e47243cd2cb22f4c1eba16d3"></a>

Type: `"object"`. single nested block, Optional.

DNSCAAResourceRecord.

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
caa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-83654eb57f202d685b464a9f664ad1c1c9e0c453f0940589114ec71be125044c"></a>

## Direct properties — primary.rr_set_group.rr_set.caa_record / be639c368bc3 / 3

<a id="canonical-22bb4b2d10fc4a7bb51e28a46ca20aee1cce67765dab55d84f6d80eec48f6d63"></a>

<a id="canonical-17dc091fc0642a8b8580c5b9c5e3c016de862e11da7e5eb53b0b776bb1c46247"></a>

## name property — primary.rr_set_group.rr_set.caa_record / be639c368bc3 / 4

Type: `"string"`. Optional.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-efc54e2c998a1793c2d311daa08e55150686803412dcc3e05d92fb9e1484b291): complete subsection reference.

<a id="canonical-79627e1965311345f21e825c8b34f984b36192167212d46657b3b2906b3275f9"></a>

## Next pages — primary.rr_set_group.rr_set.caa_record / be639c368bc3 / 5

- [primary.rr_set_group.rr_set.caa_record.values](resources--dns_zone--reference--group-002.md#canonical-efc54e2c998a1793c2d311daa08e55150686803412dcc3e05d92fb9e1484b291)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-efc54e2c998a1793c2d311daa08e55150686803412dcc3e05d92fb9e1484b291"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-618808b8390a764ba5ff9433b704a7db89c4f56991437cc1ca255a8a22899688"></a>

## primary.rr_set_group.rr_set.caa_record.values — primary.rr_set_group.rr_set.caa_record.values / d41c22497546 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-cf848ffc0e401aea9da9672384ad97dae88717076541f41d5528f704f564b350)
- primary.rr_set_group.rr_set.caa_record.values

<a id="canonical-97234d1a0474aa1584b1f4ca274d45d94201e2ee95e3a77416092e3752cae84c"></a>

Type: `"object"`. list nested block, Optional.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-6a954d84b03ba835b2f92116169d4ec6e9958e49d5413f74c9dafc9a0d0ca236"></a>

## Direct properties — primary.rr_set_group.rr_set.caa_record.values / d41c22497546 / 3

<a id="canonical-311e1baeed18139c9290b06bd70eef3f6c34409adb7e2fd72628a58a93420435"></a>

<a id="canonical-3dca3fd8abc8d1a63601d9e54164e037c51207d7f9663d49006a04e046071d1f"></a>

## flags property — primary.rr_set_group.rr_set.caa_record.values / d41c22497546 / 4

Type: `"number"`. Optional.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-3cb28684816b038f222aaee02ff5c4fa5de30d157ee44cbd234f11f8ed558b05"></a>

<a id="canonical-d0c51e0cb9a64ed2224faba548983c0c7b4ae34ceb7be9c5b11b1eb56806f5fb"></a>

## tag property — primary.rr_set_group.rr_set.caa_record.values / d41c22497546 / 5

Type: `"string"`. Optional.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("issue",
    "issuewild",
    "iodef"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-00748efaf6475feaf4091b6ea7bcbd67c29b5af03464b53b71b2619c25cdc4de"></a>

<a id="canonical-175bb0bbbb5e7983c54aa8881812355fe29fe76df66bbef01febd851c10ab352"></a>

## value property — primary.rr_set_group.rr_set.caa_record.values / d41c22497546 / 6

Type: `"string"`. Optional.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-9af7494e1cc9c23a2e215e2887068616169d3a25284f6e236448b84a03517006"></a>

## Next pages — primary.rr_set_group.rr_set.caa_record.values / d41c22497546 / 7

- [primary.rr_set_group.rr_set.caa_record](resources--dns_zone--reference--group-002.md#canonical-cf848ffc0e401aea9da9672384ad97dae88717076541f41d5528f704f564b350)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c804903abc64ad2530b3e07982c099539bd458cb7b2fb8a60b65b12ca066b79"></a>

## primary.rr_set_group.rr_set.cds_record — primary.rr_set_group.rr_set.cds_record / 06ec78e30f9d / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.cds_record

<a id="canonical-c6a39ab792433adef8073b7657a226afef3e28b5f3c9ce3adc071918ec516c1b"></a>

Type: `"object"`. single nested block, Optional.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
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
cds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-5464ba1d164d356f5350c50e979135ef2e70f77ab4e16dce0fdd870ef703aa8d"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record / 06ec78e30f9d / 3

<a id="canonical-622273816904063720be4a445807bd6e2d9bb772682524be554e16290b4ce85e"></a>

<a id="canonical-1ae525c75b987d79e9cb53ecf3cc6ead29e06686adc7500b9ca70be26d1bc275"></a>

## name property — primary.rr_set_group.rr_set.cds_record / 06ec78e30f9d / 4

Type: `"string"`. Optional.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf): complete subsection reference.

<a id="canonical-896751709df68a985a9cb4a5bf94c8f9ab0eb00e2e660f08dc23ac70fcbfc091"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record / 06ec78e30f9d / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01c26a3b7fb48cf62625f10e0ad6fa9082a8cf92f31cd135ebb745e8995df599"></a>

## primary.rr_set_group.rr_set.cds_record.values — primary.rr_set_group.rr_set.cds_record.values / fc12c41b6beb / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f)
- primary.rr_set_group.rr_set.cds_record.values

<a id="canonical-b9d03af879fbf303ba3c9f6b7ce82300e99118f460c9c5d724debecf8b2e6b22"></a>

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

<a id="canonical-08d99c080ccc5eb7ef6fc6d9c1d3c3b9099b9ada4ac41575c982e56df66f8c3c"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values / fc12c41b6beb / 3

<a id="canonical-e9dcbc8572f5e29292dee0e84557230c9eb84220e6769e0511d8907785d1d17d"></a>

<a id="canonical-56d27e9cc5533b834194e327cb9ab03c7548cf8d4534fad6ad34e08623be3c11"></a>

## ds_key_algorithm property — primary.rr_set_group.rr_set.cds_record.values / fc12c41b6beb / 4

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

<a id="canonical-ed19a1f02148abe7ec3420024abfb01e5eacf1fb644b4fc3215d8afc52573950"></a>

<a id="canonical-1938bfdab65f0f328a61b81700189dd260b58689211e2ad6320c7a1744035fe1"></a>

## key_tag property — primary.rr_set_group.rr_set.cds_record.values / fc12c41b6beb / 5

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

- [sha1_digest](resources--dns_zone--reference--group-002.md#canonical-d06211a6832a8d30e1b53980970c381c202cec54eef28aff8f8d601113844273): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-002.md#canonical-b2d38f2437fc2f4096c31baefa36d10591127ea8d5e0586bf03b4218f9edb04d): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-003.md#canonical-8ea2a071490c6012110bcc3b368f2088038165a5c93a058283dbf7e0d7f60529): complete subsection reference.

<a id="canonical-dec072cec77bd945ccfcfffc08b32cb2690add1b23baa935711d645e08379ef9"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values / fc12c41b6beb / 6

- [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](resources--dns_zone--reference--group-002.md#canonical-d06211a6832a8d30e1b53980970c381c202cec54eef28aff8f8d601113844273)
- [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](resources--dns_zone--reference--group-002.md#canonical-b2d38f2437fc2f4096c31baefa36d10591127ea8d5e0586bf03b4218f9edb04d)
- [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-8ea2a071490c6012110bcc3b368f2088038165a5c93a058283dbf7e0d7f60529)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-d06211a6832a8d30e1b53980970c381c202cec54eef28aff8f8d601113844273"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d3c47c87a92b30a9c227320a27fbdbaeec13789e5884d40aee9bb53b0bf30b5"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha1_digest — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / fb6f3068b0eb / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf)
- primary.rr_set_group.rr_set.cds_record.values.sha1_digest

<a id="canonical-cf7936a2b8e89093a68cc243f70840ff38113d84fe24a5601ca4edb96c5f210d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-7c9278254243fe7b93e7b41efe10844b0190625c2c5e5e5a982f4c7eba16b5ea"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / fb6f3068b0eb / 3

<a id="canonical-9e2065437348b11c448b7d54908708e3a167c5f528123c8aab79cde5f6d4cea7"></a>

<a id="canonical-3d5a85d0d1401a8746c7fa3ccab5664d2205a46fbce9b1f9fcb73c044f27768c"></a>

## digest property — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / fb6f3068b0eb / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-2e7c8e286f13c5323574522b6d6e340faee6e9bef0fb380a6b146e9d1dc27ffe"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / fb6f3068b0eb / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-b2d38f2437fc2f4096c31baefa36d10591127ea8d5e0586bf03b4218f9edb04d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62f4633063d54ed4a7d23bd652c7138bd2df451646a862c0529dc4baaac5a7b2"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha256_digest — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 3ea35dc88288 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf)
- primary.rr_set_group.rr_set.cds_record.values.sha256_digest

<a id="canonical-04bdbaef8648dc1ef6eca5d4cdd5249991d7b8e29d3a7bcf5dd68126a9a82042"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
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
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-5a58406610509da3781f4764831d900b9183beeb28257ce728eeaf0b34a63e72"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 3ea35dc88288 / 3

<a id="canonical-29b2299625740f56b9b27980dad2bdb4248b88f2adb4846eec8375dec02618a7"></a>
