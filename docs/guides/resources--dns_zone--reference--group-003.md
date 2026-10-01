---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-37c3e27cb6c72fbe60fad7d2e8080285c8c2793165530b5cb59932d50c9ef0d4"></a>

## digest property — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 3ea35dc88288 / 4

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

<a id="canonical-adc19e7a8fb49f424658cde363310bc8ca7b7d5a5495201b89c50304818b6667"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 3ea35dc88288 / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-8ea2a071490c6012110bcc3b368f2088038165a5c93a058283dbf7e0d7f60529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72a5021cfc53795e3a1c1d0d11f65689fc39ac2f7ea993b7c50c08d91476a8b2"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha384_digest — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / a79c4fbf5909 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-26760ef230e47b0f99c69ba32f7c15d3e1dc00895dd3af7847b727104fffa31f)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf)
- primary.rr_set_group.rr_set.cds_record.values.sha384_digest

<a id="canonical-d8f281ec335da889a1ba25fb8a5c29073c8afeda2f6b315664b68e8e11d185a1"></a>

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

<a id="canonical-1d30db0346056b62e7f2e3ed71d0acd0655c1ec5032bb3f4e167d77ee888eeae"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / a79c4fbf5909 / 3

<a id="canonical-5c5eb52018291efb8da653ca186de1798913273a3d50e12af0a2a9b536d8c5b5"></a>

<a id="canonical-f5c0e7dd355e682d4b96329236375e21f54b16ed6362adb5a9938aaadaf7f0cd"></a>

## digest property — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / a79c4fbf5909 / 4

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

<a id="canonical-762f7e22fa9e4144a9e3d98383adca6b311e85f72cd61373ba584d766b045d05"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / a79c4fbf5909 / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-002.md#canonical-e4f820c5e68b358c2fbc884c279e1b596ccbb45359fca83492e7c7fbcff2c1cf)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-9a7828a9b291db43cc63061b9ff7ab475c2d5736cdd55e78a2c55d8331c83211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-825ba25cd1ae77c4fc0f5279bdfd43c68996984dfa58238e488e5adbaaee310b"></a>

## primary.rr_set_group.rr_set.cert_record — primary.rr_set_group.rr_set.cert_record / c614ff760d52 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.cert_record

<a id="canonical-c6c789327ff609e0c4d0f7246b08495c0a44715d2cdfe6693e61181f6d9fabbc"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

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
cert_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-44127f7f60309ca8f9383faa480938d60cf5a553eadeaa97ceb375ab86196fc0"></a>

## Direct properties — primary.rr_set_group.rr_set.cert_record / c614ff760d52 / 3

<a id="canonical-883b8f4d056f3ba5391dc53a07a971eb151a6af5553b9fbe2af8dcbd2cfc7c48"></a>

<a id="canonical-5715aeb811ab3826be506942c07a7c064f184218f5b3e25756f48f5f7772a5d7"></a>

## name property — primary.rr_set_group.rr_set.cert_record / c614ff760d52 / 4

Type: `"string"`. Optional.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

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

- [values](resources--dns_zone--reference--group-003.md#canonical-549b2fb17c75bf875ab1a44bc00071d9a7245b51507ee3e95bebdf06c6bbcbfc): complete subsection reference.

<a id="canonical-3a4c49418be66d33763c3015d62b37545d75c5e462a5bf8b3b43d7a24411f6b9"></a>

## Next pages — primary.rr_set_group.rr_set.cert_record / c614ff760d52 / 5

- [primary.rr_set_group.rr_set.cert_record.values](resources--dns_zone--reference--group-003.md#canonical-549b2fb17c75bf875ab1a44bc00071d9a7245b51507ee3e95bebdf06c6bbcbfc)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-549b2fb17c75bf875ab1a44bc00071d9a7245b51507ee3e95bebdf06c6bbcbfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34a08be83f2f8189d65545ebb211c9a12ef6f692a6eda26078071a8c7655e5c4"></a>

## primary.rr_set_group.rr_set.cert_record.values — primary.rr_set_group.rr_set.cert_record.values / a7c993ff1d08 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-9a7828a9b291db43cc63061b9ff7ab475c2d5736cdd55e78a2c55d8331c83211)
- primary.rr_set_group.rr_set.cert_record.values

<a id="canonical-b90e051c6f5245cd6b931c4799adac5d7b6ccdc0cbba85220f4e66e643c9a028"></a>

Type: `"object"`. list nested block, Optional.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cert_key_tag",
    "certificate")}
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

<a id="canonical-64e8fd4468fbf17d8f145ca9b94e9299144250b799688a7663578a301fdef18e"></a>

## Direct properties — primary.rr_set_group.rr_set.cert_record.values / a7c993ff1d08 / 3

<a id="canonical-6519a134216376074eea5e54df670d57e4aed2ac75e75ba6b1994e4656a07d0b"></a>

<a id="canonical-eb1ce0d02ac6b610497700814b2ffc9339382109081dace17753b82c953e915a"></a>

## algorithm property — primary.rr_set_group.rr_set.cert_record.values / a7c993ff1d08 / 4

Type: `"string"`. Optional.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Upstream description:

CERT algorithm value must be compatible with the specified algorithm.

&#8203;- RESERVEDALGORITHM: RESERVEDALGORITHM

&#8203;- RSAMD5: RSAMD5

&#8203;- DH: DH

&#8203;- DSASHA1: DSASHA1

&#8203;- ECC: ECC

&#8203;- RSASHA1ALGORITHM: RSA-SHA1

&#8203;- INDIRECT: INDIRECT

&#8203;- PRIVATEDNS: PRIVATEDNS

&#8203;- PRIVATEOID: PRIVATEOID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7313060d7ba1accc9e97a6e90ade63803a87d21d58d889389ad56d9966ecacba"></a>

<a id="canonical-de874f8bac2a5626d0c164ad950cda2ae9fdc9913670a401f905da877b1e003e"></a>

## cert_key_tag property — primary.rr_set_group.rr_set.cert_record.values / a7c993ff1d08 / 5

Type: `"number"`. Optional.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

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

<a id="canonical-9dcf7e2cd9163ae1d6ba7e638993b0aadf203c81cf43f7feecc02080a49563cc"></a>

<a id="canonical-9afbf771ef4e2be89deff87c04b876e99f134137e0784c63f14bd985b1ae81a2"></a>

## cert_type property — primary.rr_set_group.rr_set.cert_record.values / a7c993ff1d08 / 6

Type: `"string"`. Optional.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Upstream description:

CERT type value must be compatible with the specified types.

&#8203;- INVALIDCERTTYPE: INVALIDCERTTYPE

&#8203;- PKIX: PKIX

&#8203;- SPKI: SPKI

&#8203;- PGP: PGP

&#8203;- IPKIX: IPKIX

&#8203;- ISPKI: ISPKI

&#8203;- IPGP: IPGP

&#8203;- ACPKIX: ACPKIX

&#8203;- IACPKIX: IACPKIX

&#8203;- URI\_: URI

&#8203;- OID: OID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4c30b9997468b3c96afca5388c5bc71e89e301ebed5db1c7297874023c83f3b9"></a>

<a id="canonical-505e9c530b64840b3c530b0f663cbf52f77ccdeb1ac982a2f3c2ed065cd401f0"></a>

## certificate property — primary.rr_set_group.rr_set.cert_record.values / a7c993ff1d08 / 7

Type: `"string"`. Optional.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

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
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-fc71b66bee4b37e2cc396074b97fc26f17f39ba5e6f6c9d00216a993674b12d1"></a>

## Next pages — primary.rr_set_group.rr_set.cert_record.values / a7c993ff1d08 / 8

- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-9a7828a9b291db43cc63061b9ff7ab475c2d5736cdd55e78a2c55d8331c83211)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-5371974e821ae367da729184bd9adb0dd4371d3d39f0c278e60572e945875efa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39b35d5554430d0da82402b0a5897ffa8ea2d0947cd10727df36d01d677061e1"></a>

## primary.rr_set_group.rr_set.cname_record — primary.rr_set_group.rr_set.cname_record / e9b47ef20e2c / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.cname_record

<a id="canonical-34d5905aa5802da9f46e479b550c7491281f0405fb25fae463fa6f458d9884db"></a>

Type: `"object"`. single nested block, Optional.

DNSCNAMEResourceRecord.

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
cname_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-cf63337e394cb6490233ef26b96d65b4c77acad911b8f8cb36b710f440e59e7e"></a>

## Direct properties — primary.rr_set_group.rr_set.cname_record / e9b47ef20e2c / 3

<a id="canonical-685796b1c132cc30372659ebf71a3aa8ef0af128f0675536969172b1ca5f9f5f"></a>

<a id="canonical-fb3083064fadf2e861c5ff3b5133df7f94b7247a6963f0254731f8951dd3721c"></a>

## name property — primary.rr_set_group.rr_set.cname_record / e9b47ef20e2c / 4

Type: `"string"`. Optional.

CName Record name, please provide only the specific subdomain or record name without the base
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
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-984dd91ec5eba33dc59729ba2e7cf1b6084260aab91f13f3eddaefd6c2598991"></a>

<a id="canonical-5b4370a3f7a1ad6e2871ce949d1b12b54520c8f897cee286884607b56febe982"></a>

## value property — primary.rr_set_group.rr_set.cname_record / e9b47ef20e2c / 5

Type: `"string"`. Optional.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

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

<a id="canonical-6a2f36d7bb96fa0c008ce190366918e5f3f92238de831e102e8b460f4b8d0132"></a>

## Next pages — primary.rr_set_group.rr_set.cname_record / e9b47ef20e2c / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5aa043444fd6979e088ff9c3030c745b23877b29d7bda5a8cdb492b071e3383"></a>

## primary.rr_set_group.rr_set.ds_record — primary.rr_set_group.rr_set.ds_record / 05783b6383e0 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.ds_record

<a id="canonical-76e22f22caa83f6f4831e7a0452d5f93fc822ec65b28dd68a95f137eb2650252"></a>

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

<a id="canonical-dab131bfe075eaa7af34cff64a9ff3115a2e20d5f4dd39525cda6e4b62cf8caf"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record / 05783b6383e0 / 3

<a id="canonical-0ee745cb5fe1cf9e99c4061c4bdd7ef53c8436108b0edcf5d103ab3f68f9424d"></a>

<a id="canonical-5708fc2c23c24ccfe4a23eaa258122fe986672f20d3c12e00409698694a4cd4d"></a>

## name property — primary.rr_set_group.rr_set.ds_record / 05783b6383e0 / 4

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

- [values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d): complete subsection reference.

<a id="canonical-102e4a8290c3f8bf5adbdc14be52676b0ce996ef3a0a0d340d2b0c183796f3e0"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record / 05783b6383e0 / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8919f3dab405945da19f17cf7605963e2c9133fbec2b71e802306808a2bdd78"></a>

## primary.rr_set_group.rr_set.ds_record.values — primary.rr_set_group.rr_set.ds_record.values / 3d5b1d934a5f / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb)
- primary.rr_set_group.rr_set.ds_record.values

<a id="canonical-908e4fcb9d02f4efd483ec03224e499c35f27be5c3c5a8022415fec8defa57a1"></a>

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

<a id="canonical-21db5869777305826b7bb99fe84482d722fc148d45b2506757e872dd79cc304c"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values / 3d5b1d934a5f / 3

<a id="canonical-bba743dd2af769047a7ab5600837bd376326100d105f97c633cd015d7cb21d30"></a>

<a id="canonical-fdfb8440bd44582ff421b0f2e9c77800f4b37e48f078d55610abea6dbf9c104c"></a>

## ds_key_algorithm property — primary.rr_set_group.rr_set.ds_record.values / 3d5b1d934a5f / 4

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

<a id="canonical-4a1f693cd57aab471d3cae5363ebcd752fe3982993bbcbe45df6e30d73b8f52d"></a>

<a id="canonical-1176394d5d537057158474e02d2bc93a0d29bf6377a9bfc093bef01c154c5323"></a>

## key_tag property — primary.rr_set_group.rr_set.ds_record.values / 3d5b1d934a5f / 5

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

- [sha1_digest](resources--dns_zone--reference--group-003.md#canonical-699e527c48b44654670cfd2a8244f5abb986c3fb1398d7cb0c0bda71153f0086): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-003.md#canonical-528e65660c23579c09abe674c518abe53ad9afe2be6276410484fcbc506367a4): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-003.md#canonical-756cc12a07ca6fe892a45ca63fb0b0eac0e27e7cd76f4d84fd70b4e1abb7bdb8): complete subsection reference.

<a id="canonical-02ac0806f99f74dfe7dc40fe1a68bbeb1587d6c41ed6fcc9d32679365bdfef4b"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values / 3d5b1d934a5f / 6

- [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](resources--dns_zone--reference--group-003.md#canonical-699e527c48b44654670cfd2a8244f5abb986c3fb1398d7cb0c0bda71153f0086)
- [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](resources--dns_zone--reference--group-003.md#canonical-528e65660c23579c09abe674c518abe53ad9afe2be6276410484fcbc506367a4)
- [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-756cc12a07ca6fe892a45ca63fb0b0eac0e27e7cd76f4d84fd70b4e1abb7bdb8)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-699e527c48b44654670cfd2a8244f5abb986c3fb1398d7cb0c0bda71153f0086"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8246381dc6cd22e9e88265a7d6cf3498d8287c7b9990ee4c1770e476de29681"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha1_digest — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 692ddd9a8031 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d)
- primary.rr_set_group.rr_set.ds_record.values.sha1_digest

<a id="canonical-057dde4bdc8f160e145a875295645849a5f7a61c8236827c96b69e882be9bee0"></a>

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

<a id="canonical-9c90a8d94aeb3b996c8bba4ff1005619305eae64ac7b8c1915d7f1ae1ee188a3"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 692ddd9a8031 / 3

<a id="canonical-ae773b8f3fd8a2b65550424391a9838cfc6d0b1a9caa4bf60a4c754e22b4794f"></a>

<a id="canonical-34880bd6ead19b421f19af613be8de5187f151ccb083a76ca2f1ef59c371160b"></a>

## digest property — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 692ddd9a8031 / 4

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

<a id="canonical-3e5a7dfa96c59ece202ab05eb17e3fc6b13f64be7e05b7d4e9b31a222a0f493f"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values.sha1_digest / 692ddd9a8031 / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-528e65660c23579c09abe674c518abe53ad9afe2be6276410484fcbc506367a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ca78182759e0882fd921a8b1e389a280e956c0da47a0dac1c4c7b20ff7cdb85"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha256_digest — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 417e94e9f81a / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d)
- primary.rr_set_group.rr_set.ds_record.values.sha256_digest

<a id="canonical-ebc154f09d5c3662308baae866ca53275f66009db0414bbd28684b2ea7a262d8"></a>

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

<a id="canonical-335b688fcfaf825b6abe886c6c7dbe3d9764b6b3130009dceb89d9a480b0670f"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 417e94e9f81a / 3

<a id="canonical-52ba0723e72fde65372e463a99933bb3c85201def6d246fded085fec6731a6c0"></a>

<a id="canonical-cf8a2313d354e779f7b4d3282f89019a4cce5c3c645483ed601b7761bc0d6f0d"></a>

## digest property — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 417e94e9f81a / 4

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

<a id="canonical-2a5c54050e14808623191858789745827660666613adafbed6e9e1a81298f63b"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values.sha256_digest / 417e94e9f81a / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-756cc12a07ca6fe892a45ca63fb0b0eac0e27e7cd76f4d84fd70b4e1abb7bdb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fadbeacff5b63403813f6e0818256091e9548b5c87550335a1aeec3b25864c4"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha384_digest — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / 272e5642b4e2 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-1f54e5cfc137c90adf139755cc3bdf2b4c4ecc516a9c94ed57fef32370a386eb)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d)
- primary.rr_set_group.rr_set.ds_record.values.sha384_digest

<a id="canonical-88e4f14212003d0ca2e0023f047c43081cdc4f24f50ed0d126428644a87988b0"></a>

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

<a id="canonical-91b037b610c935a04f0c0427c4a9319da7f402de44778ba840c82f86edc8da37"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / 272e5642b4e2 / 3

<a id="canonical-78de4e9dad0a8cf23ead45a1ae6c6e466937a222244fec1babfa6c315d09a80c"></a>

<a id="canonical-335f22b677f02479fad217ab915f45a92854f59e6d0646918b688ee96b68bbf3"></a>

## digest property — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / 272e5642b4e2 / 4

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

<a id="canonical-ff48fd1f53ac208f3ea19e9dbcc680da2c70c421a3d85c2ac120d7662eebc5fa"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values.sha384_digest / 272e5642b4e2 / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-c24384a7f07aefedd9169ace52c2b3d805d690fcc9f133d100e0cb073418e13d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-cc21bd523c0b066121389cde3dd32c5982bbc85de11047c259a20d1d194410e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31a7198fb45e4cc37d3ef1b89da0895728b3a381c46784cd2b7fde3cd746ca83"></a>

## primary.rr_set_group.rr_set.eui48_record — primary.rr_set_group.rr_set.eui48_record / 0be1729edfe6 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.eui48_record

<a id="canonical-216f2a93b61d77861a018d38382993d0c72bba0cfdae7a7fb91d0bf9f2f1466e"></a>

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

<a id="canonical-b4e9bb535ef0c2c929e8acdc25313f6d10bb413f5fae1936eec9c9ca264c8329"></a>

## Direct properties — primary.rr_set_group.rr_set.eui48_record / 0be1729edfe6 / 3

<a id="canonical-826be9317cc42134d49761bcbb9da7be7f3687e66838719d1f4134827d25f2c8"></a>

<a id="canonical-ff88eeba5ec246c4d7a4a1742d114681b5ea2273ecca904f4c8e46ce0d624dca"></a>

## name property — primary.rr_set_group.rr_set.eui48_record / 0be1729edfe6 / 4

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

<a id="canonical-022deb453359246e49bf1093f8fe79c766548ddd753e11d794810a5f999e8273"></a>

<a id="canonical-e3c7ac00d86e9d494ae57da358ab5046c784b5055e527b83e92005c9c8b74b5d"></a>

## value property — primary.rr_set_group.rr_set.eui48_record / 0be1729edfe6 / 5

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

<a id="canonical-ce1aa3e632d7041427c618bfef3476a5de9f16e8388612c3de9d496a9cfe3ac6"></a>

## Next pages — primary.rr_set_group.rr_set.eui48_record / 0be1729edfe6 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-5bb96ff6312f0f9a10bef3014b08ea4c07368ac1b1d153fd187cfd6316e078cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d975a45066eb3f4440778985ea9e448ce48a609fdd0456e0f1023d8eb6db311"></a>

## primary.rr_set_group.rr_set.eui64_record — primary.rr_set_group.rr_set.eui64_record / cc7cf0d438c4 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.eui64_record

<a id="canonical-8795d399362d50041a23c675da96a20f0c13bab3158eda96a55b1d98648cf36f"></a>

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

<a id="canonical-fa9f87a48b686596fd89a1fc13174f3761fecb0875b90cea038bd55fefa21bd4"></a>

## Direct properties — primary.rr_set_group.rr_set.eui64_record / cc7cf0d438c4 / 3

<a id="canonical-c31ee61a959964b682f4300d2b6ec8294a8a6bc2d164db34ec2b6b62a9193bce"></a>

<a id="canonical-b68005b99520d8ee5f5f783f4f9abaf4cb7f5cd3ff4a96c323552a5f2b803df4"></a>

## name property — primary.rr_set_group.rr_set.eui64_record / cc7cf0d438c4 / 4

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

<a id="canonical-526fb9d2d6a2f79daeeeb0f21113b86ddc1b5595d0ab579ad7ec4f8169e67a26"></a>

<a id="canonical-d9877c85a95238e7df1c1f6cfc4bf59301ceac1ff33938fed8170a51d7d98526"></a>

## value property — primary.rr_set_group.rr_set.eui64_record / cc7cf0d438c4 / 5

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

<a id="canonical-4037882ae49a501d930153bc72d3b1c7d64309ec2e021008a345762cd9c40026"></a>

## Next pages — primary.rr_set_group.rr_set.eui64_record / cc7cf0d438c4 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-8501fd8f15ca3bf884866766a26bb1357a6f17826763d164f5c395e1e3b77ee3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a06bfc7dd18ea572eff7f123388872f0278d4c6478f8f8d46881b983f44f38b0"></a>

## primary.rr_set_group.rr_set.lb_record — primary.rr_set_group.rr_set.lb_record / d0799ce78f2b / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.lb_record

<a id="canonical-0963412d2570059ec0fd1bb7f81afd24b014a937af631f868bf8c11c278a1caf"></a>

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

<a id="canonical-d658c017da68afd6db1665f34626b53e96f376dd9df04a902b7dd3b0a9388618"></a>

## Direct properties — primary.rr_set_group.rr_set.lb_record / d0799ce78f2b / 3

<a id="canonical-092301468a776edec18b35e09de595cc76754d3e183ad7ea64d85436f95e90a0"></a>

<a id="canonical-70ed2d23dcd8c6aacc53628c6db778d1de8657e87b26c34407e7f1b79b3b1a3e"></a>

## name property — primary.rr_set_group.rr_set.lb_record / d0799ce78f2b / 4

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

- [value](resources--dns_zone--reference--group-003.md#canonical-abd5eeacf8eb6c1e32b100a2d49124fe11e7eceb6dccdeb3c80afdff13deec62): complete subsection reference.

<a id="canonical-c298eb149ec920eb5062cfa1111c0fede06d87601ade3f3d2460c66a0f65a9dd"></a>

## Next pages — primary.rr_set_group.rr_set.lb_record / d0799ce78f2b / 5

- [primary.rr_set_group.rr_set.lb_record.value](resources--dns_zone--reference--group-003.md#canonical-abd5eeacf8eb6c1e32b100a2d49124fe11e7eceb6dccdeb3c80afdff13deec62)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-abd5eeacf8eb6c1e32b100a2d49124fe11e7eceb6dccdeb3c80afdff13deec62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92ae210b2929461b641b968762579eb053292fb8eb52f3320319b78ede600278"></a>

## primary.rr_set_group.rr_set.lb_record.value — primary.rr_set_group.rr_set.lb_record.value / a9a91a2d7b12 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-8501fd8f15ca3bf884866766a26bb1357a6f17826763d164f5c395e1e3b77ee3)
- primary.rr_set_group.rr_set.lb_record.value

<a id="canonical-8c09cb5e801be9d767f469e89945bcb793c3dd09d8126a958ab21c4b42fd7c08"></a>

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

<a id="canonical-6cebe39c1a80c6d221887616b954f858129d31ff8c7256f51869ff49c9734104"></a>

## Direct properties — primary.rr_set_group.rr_set.lb_record.value / a9a91a2d7b12 / 3

<a id="canonical-9c17df4933a8e2f72a38ca9d5ff468100e685fb9cba7861f9c64fb2b13b18cf1"></a>

<a id="canonical-84a2e3c24e8ac3abb4bdcc20543a1ecd7576512fdb4bd3d2321f278bc5c87beb"></a>

## name property — primary.rr_set_group.rr_set.lb_record.value / a9a91a2d7b12 / 4

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

<a id="canonical-96912b2a8e8449d3f6d2d5c2dca72ce8b6c171c6cbbe60bd2c9de1bd6c46c9b9"></a>

<a id="canonical-99c35c5e6357ed9c7e07c55449f13250f81734f29e4376dd3f8f131b2d599163"></a>

## namespace property — primary.rr_set_group.rr_set.lb_record.value / a9a91a2d7b12 / 5

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

<a id="canonical-7a766e0d21f09e5d37e5d86f5e7a627a80e821c75cd12d6b4f3d821242128862"></a>

<a id="canonical-b2b9be53095dea1b3d8b6b023246e570869b46fcc978895e830bbc42c04061fc"></a>

## tenant property — primary.rr_set_group.rr_set.lb_record.value / a9a91a2d7b12 / 6

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

<a id="canonical-297cd16c800a20e1023c5d6a6bc2992fb9003cfdf3b102bf4640b47b75f21cb0"></a>

## Next pages — primary.rr_set_group.rr_set.lb_record.value / a9a91a2d7b12 / 7

- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-8501fd8f15ca3bf884866766a26bb1357a6f17826763d164f5c395e1e3b77ee3)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-c3aa28a94f83cfd171303d2f7d99d864f23b1e65162c4a67600df832064ab922"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d227e0dbb75c34ee13f95941c4357a2801083726042c500aa5c3ae4e44514fd0"></a>

## primary.rr_set_group.rr_set.loc_record — primary.rr_set_group.rr_set.loc_record / e9e0967c0b6f / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.loc_record

<a id="canonical-636c023c97807452576c77476421f5eb6bb35edf416e0167139921b53a2eb4ae"></a>

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

<a id="canonical-fcf0c16d8884cf903daa15d68924c3eacd7128a028f1e6164b06c01957a0085d"></a>

## Direct properties — primary.rr_set_group.rr_set.loc_record / e9e0967c0b6f / 3

<a id="canonical-a117fbd3d9fa53fdadc0489b3e9b7fa1970fc1f9900fde483e46847dd2efae5d"></a>

<a id="canonical-9c553119e7f78d2e0802bdd23a20e34a8c50adce6f9123a4e94c214fa62f6aec"></a>

## name property — primary.rr_set_group.rr_set.loc_record / e9e0967c0b6f / 4

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

- [values](resources--dns_zone--reference--group-003.md#canonical-7b1f25a3ce98b4c745e8c9165ffcd66fc865a609b3e5988cea6e508a46e11576): complete subsection reference.

<a id="canonical-f9660811241a14725c205b73a8657facbc488125f8ab8f5929295865dafb5f7c"></a>

## Next pages — primary.rr_set_group.rr_set.loc_record / e9e0967c0b6f / 5

- [primary.rr_set_group.rr_set.loc_record.values](resources--dns_zone--reference--group-003.md#canonical-7b1f25a3ce98b4c745e8c9165ffcd66fc865a609b3e5988cea6e508a46e11576)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-7b1f25a3ce98b4c745e8c9165ffcd66fc865a609b3e5988cea6e508a46e11576"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0650128744042537954ceefe3bf6b33c963e9bf6a54148ecbc314b3cfbe8fa2c"></a>

## primary.rr_set_group.rr_set.loc_record.values — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-c3aa28a94f83cfd171303d2f7d99d864f23b1e65162c4a67600df832064ab922)
- primary.rr_set_group.rr_set.loc_record.values

<a id="canonical-f7d6025c5d0a342c8dc31082d1464284f59eb8c7bac6c824c8b66447a777e894"></a>

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

<a id="canonical-7877c538e50569aebfd9d5d45d1454cc6640edda1aeb83e15df5c3d7c70794b1"></a>

## Direct properties — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 3

<a id="canonical-075b763ef87bfa9e7d04f05a0f075e873294e905dd2cfc04c0ee53c23732345c"></a>

<a id="canonical-fadc49c7bb55dab30fa05a7be1ce4d4d9969f0907686d7da03af4da9d66b3a07"></a>

## altitude property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 4

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

<a id="canonical-3202ef51df86585adebd7356bbddc8a8f922daa59c3f9b810657a63137218a41"></a>

<a id="canonical-98c7fe469d33136f0caf430b3dce1660b9a1f590255204098e52e851fcb00805"></a>

## horizontal_precision property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 5

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

<a id="canonical-cad3a6f735448a8648b20c0dba50d2b49245d81f1e5fb65df425cd075b80d5e5"></a>

<a id="canonical-c7eac757f0535272abb6b00b2d32f771ee17865ab5a8989af34e254710136d8e"></a>

## latitude_degree property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 6

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

<a id="canonical-f472b64ba9f0475581d835682ebb6785745d0cc57999f63e8cc4f20bdb7c0953"></a>

<a id="canonical-86a82910b41dd271dca8e4de84ad559551933e31d597272224405b277b86bd39"></a>

## latitude_hemisphere property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 7

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

<a id="canonical-809a050700878ffe17f48214928ab27c9e51225f7434920edfde73b7bc261227"></a>

<a id="canonical-d5119c3fcbc7333e07231b9f32b4addec8a7df04aa39817c9907060bae43cde0"></a>

## latitude_minute property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 8

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

<a id="canonical-396f2a2b678dd4b0405a7d44f2e8c2221573dbb4024c1eab56b642d76b6f6bba"></a>

<a id="canonical-f2ec9d871b4093acb0e5673b112c098ee17299f3439d05759136d935fe032cb4"></a>

## latitude_second property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 9

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

<a id="canonical-0071547aaf7adfc1b2628d11f6276a45161ba9af834cc15d0ca9aaf09e5734b1"></a>

<a id="canonical-d55c8acf16cffebc1c081488da02b1156156e272bf45854be5f5c1349287188f"></a>

## location_diameter property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 10

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

<a id="canonical-562f2dae7d0c91e049ca0f91942daec69bd656023e692f1dd9e1b175c8babeda"></a>

<a id="canonical-001fa9f457c0fd21947f62801ad051622e67d6de59dfb760bb7ecb01ed7ed205"></a>

## longitude_degree property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 11

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

<a id="canonical-f59cd940a3507533c5774de792684ab235e4dce3f30d5cb169d9a6a4e7491ec6"></a>

<a id="canonical-8f9009da71624dbc51b0ae08273832539501280e8d3146efd416c93c82a4e9ba"></a>

## longitude_hemisphere property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 12

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

<a id="canonical-ffd5f4d66b6a3eb3519dc0538369951c19867fd6321857f3f20020ed0431e474"></a>

<a id="canonical-ad74009e8da77626337671ac94e94d4bb7333e1e0e6cb746e2bdf2a276deaba2"></a>

## longitude_minute property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 13

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

<a id="canonical-03c35979f23f11c073add48109d06c6ac2a540d9d2414eb1f14c4adb22880e5d"></a>

<a id="canonical-1de9850e45237d9a7f7fbe2ab34b4b907f1c13270360ae274a993bf640a18552"></a>

## longitude_second property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 14

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

<a id="canonical-cbf9fa4da7821153e718b3ff369c7183a0ca0f973869f840c0c3c42a0c961e01"></a>

<a id="canonical-140a5b0c86d37154a76864e303bb5820c19c7d67576a7520d489090affeb0552"></a>

## vertical_precision property — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 15

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

<a id="canonical-4c8fb473d7cfbb4d8564b62df488aae02960213955630d9a47502c3b736a0bb4"></a>

## Next pages — primary.rr_set_group.rr_set.loc_record.values / 9f552d5e4857 / 16

- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-c3aa28a94f83cfd171303d2f7d99d864f23b1e65162c4a67600df832064ab922)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-8307a17271b7fe3063cb6a7c8fcca75ed9e1c50dbe43413f9c1574ce4ae344a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a222a9ae03e7137e580ed002cdd68e10e62239961b84d0f320c1190f559eb9f"></a>

## primary.rr_set_group.rr_set.mx_record — primary.rr_set_group.rr_set.mx_record / 35d4799333a3 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.mx_record

<a id="canonical-9e3370feec0c55291c08db96df2194baabb12c768b33cb3cddc10437e97ab47d"></a>

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

<a id="canonical-6b05bceeb5c92b8f8eda6c7649b43b76fbdcc770522b3a7b4658b2662abb272c"></a>

## Direct properties — primary.rr_set_group.rr_set.mx_record / 35d4799333a3 / 3

<a id="canonical-1236085d5ecf20eae4f5fa14fee88c0c0fd3c8ec84fbb68b0e31ad25560c4874"></a>

<a id="canonical-3d0abf268a92462c91ca83de976b9b0169bfb996a8bba6bb4b5bc1fe5016508a"></a>

## name property — primary.rr_set_group.rr_set.mx_record / 35d4799333a3 / 4

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

- [values](resources--dns_zone--reference--group-003.md#canonical-e3165dc1027e4b1ad215c18d9b7c1a3d7f35fdb052bf761ac8f781d8f2512864): complete subsection reference.

<a id="canonical-9d75ea6eb355cd12142d1f0556503bb4bfc2fbbacf1571bfe0f6a0189f59d879"></a>

## Next pages — primary.rr_set_group.rr_set.mx_record / 35d4799333a3 / 5

- [primary.rr_set_group.rr_set.mx_record.values](resources--dns_zone--reference--group-003.md#canonical-e3165dc1027e4b1ad215c18d9b7c1a3d7f35fdb052bf761ac8f781d8f2512864)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-e3165dc1027e4b1ad215c18d9b7c1a3d7f35fdb052bf761ac8f781d8f2512864"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6801a249966dcd209500727bb461aad4b73ce415ff4ba35875b234025cfa46d"></a>

## primary.rr_set_group.rr_set.mx_record.values — primary.rr_set_group.rr_set.mx_record.values / 4995ec68a6ca / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-8307a17271b7fe3063cb6a7c8fcca75ed9e1c50dbe43413f9c1574ce4ae344a2)
- primary.rr_set_group.rr_set.mx_record.values

<a id="canonical-a2fe245d5e4931c056f19fb8e6ac8a6d8b72d4bde3cd29240fa02a603778ca85"></a>

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

<a id="canonical-816244d01ac22c20c3c9a29a9038e59c6f45bb605cc87e713b5dd18bc3e2efe8"></a>

## Direct properties — primary.rr_set_group.rr_set.mx_record.values / 4995ec68a6ca / 3

<a id="canonical-5f4ba535d1edbeddd1c6356deb9bb9677d512444efd8afbc8cf8fce15227439d"></a>

<a id="canonical-961869bb45ae91217c3b76ff1f6f2178a42ec9f036f3688f85b75a20b92b92c8"></a>

## domain property — primary.rr_set_group.rr_set.mx_record.values / 4995ec68a6ca / 4

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

<a id="canonical-d5dca6c95f9fa58ea364282e59c407fcc07f6bd37f78cdb05232b6fb0a5afff1"></a>

<a id="canonical-d5777042b77025054139d71fd8b95d941d894cae0b3ab36a821ee09e1c0d504d"></a>

## priority property — primary.rr_set_group.rr_set.mx_record.values / 4995ec68a6ca / 5

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

<a id="canonical-2ac6d249b9d371b8e4b1c58db2a5c3f03b2891607066dd7a5f3d42795132cfcf"></a>

## Next pages — primary.rr_set_group.rr_set.mx_record.values / 4995ec68a6ca / 6

- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-8307a17271b7fe3063cb6a7c8fcca75ed9e1c50dbe43413f9c1574ce4ae344a2)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-0d06865cae15e71aebaf490992eaf16baf306801ee5df6c7ac205445e30c9aeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-195770b591554c62ae52721d9c0390c161ab4917e364b5937b1abe60faa26f8a"></a>

## primary.rr_set_group.rr_set.naptr_record — primary.rr_set_group.rr_set.naptr_record / 3d5bb8e8b198 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.naptr_record

<a id="canonical-174a598251a89294d1034269e76dd69eaacf4124db2746df363fd297c318c3c5"></a>

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

<a id="canonical-9a613467f849dcbd7eddac681e8781a93959700b5fe21db89ff6187cf5e2f162"></a>

## Direct properties — primary.rr_set_group.rr_set.naptr_record / 3d5bb8e8b198 / 3

<a id="canonical-601aeb4171436d4d9fc174284c44912a8eaef4383f8537141b4fed46d80a9306"></a>

<a id="canonical-f1a1aec21707c9da69d5d2c5ad5bff9060cdd9f108983b2a15cb0ecea0837487"></a>

## name property — primary.rr_set_group.rr_set.naptr_record / 3d5bb8e8b198 / 4

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

- [values](resources--dns_zone--reference--group-003.md#canonical-10a2a7f36cb8b7f6defc49f30027467d192cdcc3832e2091f80425805fb94d8d): complete subsection reference.

<a id="canonical-44d480b71ed2281d1d03e4d9986f6bc27faa91ff2020665b1a01d390f4410731"></a>

## Next pages — primary.rr_set_group.rr_set.naptr_record / 3d5bb8e8b198 / 5

- [primary.rr_set_group.rr_set.naptr_record.values](resources--dns_zone--reference--group-003.md#canonical-10a2a7f36cb8b7f6defc49f30027467d192cdcc3832e2091f80425805fb94d8d)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-10a2a7f36cb8b7f6defc49f30027467d192cdcc3832e2091f80425805fb94d8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cae35e0d70605107eee433c49ef4559db333845433cdd3f807f313d7f71186c9"></a>

## primary.rr_set_group.rr_set.naptr_record.values — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0d06865cae15e71aebaf490992eaf16baf306801ee5df6c7ac205445e30c9aeb)
- primary.rr_set_group.rr_set.naptr_record.values

<a id="canonical-2a82f0fb3a36ddd0612945fa5c2d664cd3bd1dd977332cbfba1de92d0e5574cf"></a>

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

<a id="canonical-21fd37d872742fbca0dabe3da5ae9e84ddee11e7f8de9621b7996ac2a59b5629"></a>

## Direct properties — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 3

<a id="canonical-80f9ed0b5ccdadc391bf09dd8c08d861d599974f723a3ab1c754ce5b16c30f2b"></a>

<a id="canonical-ddeee128303c8f17256df8e53e654a58f97e242d357dea470b23593f01f9b91b"></a>

## flags property — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 4

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

<a id="canonical-5d1cd4ad75ca4ccd4d3cef0b07f217d67027a1d0a410619adec6381b2dd8194d"></a>

<a id="canonical-c5c33097c4d6b3808a10751f93f8a5d212828de6182600efaba7d89522c9de3b"></a>

## order property — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 5

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

<a id="canonical-8566f9301985ace9f00cb0ea1eba778ab4c6d99463e38d0b973d457663e4e166"></a>

<a id="canonical-6df09696c5440522aa9014c34bf2ee9b1e1b3be772d0540eb7de855965cce68d"></a>

## preference property — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 6

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

<a id="canonical-4a68c86f0c4c7dd168a996f36ca9a7db961cc9a1fe6454253c9bb062f2958b30"></a>

<a id="canonical-f9bbab68db3acef935236cb61ee7f05224fb72dfd9c94ffee921939fafc1805b"></a>

## regexp property — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 7

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

<a id="canonical-62ee7daadb719740fd2b76bfabaac8e34803844ed2a11ce3668f37a165ae45e5"></a>

<a id="canonical-7ad32ad8de953583cebb3272dfd4b3790e7a87b5362fd2caa3350a034b2c8b3d"></a>

## replacement property — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 8

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

<a id="canonical-4c145849b97de94db32f9e2261d596f8b8892da67f55d59b9d7718bb5bf759c0"></a>

<a id="canonical-a0871b9d6510aa724ada280e5c7583b2d1df01d4f98113541641d2e2aad981ed"></a>

## service property — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 9

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

<a id="canonical-d962dc0e1bee75b52d9b4dc2420c44a6d3e332b18e8bbbbfa75a22173fcf63cc"></a>

## Next pages — primary.rr_set_group.rr_set.naptr_record.values / 5b12d1b8b1ba / 10

- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0d06865cae15e71aebaf490992eaf16baf306801ee5df6c7ac205445e30c9aeb)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-d64af960376e35b05c1d2b17852ee628721ac928e27617d40527087752d7e49e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10c5b708f2a3517bb4cabac7f064f7f7ab1ff275aa51a21d2f960a8f45d040cd"></a>

## primary.rr_set_group.rr_set.ns_record — primary.rr_set_group.rr_set.ns_record / 13a55f2223af / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.ns_record

<a id="canonical-6d0da6f39f1c1a95debdc1c1c6d52a6893916066fb023ed7d8cbd02e89967db9"></a>

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

<a id="canonical-2670ecf7a9079fe8b88f7794c03ca61fe6a1d069b369ae719a4f3926620cd2f6"></a>

## Direct properties — primary.rr_set_group.rr_set.ns_record / 13a55f2223af / 3

<a id="canonical-f82119045a704373f1efeea7b1f0404a1992d72a4e6b1739783610c596b52c97"></a>

<a id="canonical-d0094cef407439813befa0f3ccbb76821693a531fe600639b207e3e8e5c94f89"></a>

## name property — primary.rr_set_group.rr_set.ns_record / 13a55f2223af / 4

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

<a id="canonical-871f908069ca417673b4062df2115e047d13e5eab2ff7a4fcf015af266861b32"></a>

<a id="canonical-ca01328b27d2345e4756f0a5ecf450cab634af046e35117e52fd5b18986f8bfb"></a>

## values property — primary.rr_set_group.rr_set.ns_record / 13a55f2223af / 5

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

<a id="canonical-3e03ac8ff45954a4caffb8af2d6e68b2f1a5f6056bce26b6e3ce6991239331f3"></a>

## Next pages — primary.rr_set_group.rr_set.ns_record / 13a55f2223af / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-ee0343545539ce1c7c1282eb8b827c02dbeb555c99152d4b631c303af0039914"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60b9885ccb6a291509767ecab48a8d64995d526d866c4c4db6182cfed5428cfb"></a>

## primary.rr_set_group.rr_set.ptr_record — primary.rr_set_group.rr_set.ptr_record / 21ccee002f87 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.ptr_record

<a id="canonical-c7a4bd4a0895c873dd89a3c37091fac0ea10b7944148db1708a9c7c05ace08e0"></a>

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

<a id="canonical-030c8c2163c3ecb2a5df2652a11c3a9a33a0eaa67e1817d3060b81a628a18c0f"></a>

## Direct properties — primary.rr_set_group.rr_set.ptr_record / 21ccee002f87 / 3

<a id="canonical-d0af55dd139061e43de94a5f083b92b88a8e046ab40fab0127d6c5456219898e"></a>

<a id="canonical-1fbce7517d91cd2c39e50b67ac183fb06ee5a3aaf0c6e50697b35821038ab6f1"></a>

## name property — primary.rr_set_group.rr_set.ptr_record / 21ccee002f87 / 4

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

<a id="canonical-8bdb1ba55e569f47921b0f9936034d730c78e1a5ed1f9f826339692c2cfc69a7"></a>

<a id="canonical-0015fd3dbdadb50ed4ea402fea2029f581df77f88ee54932cdedd6fc6dd89250"></a>

## values property — primary.rr_set_group.rr_set.ptr_record / 21ccee002f87 / 5

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

<a id="canonical-7cf1f1019534db7b1be5300191670d2c8d739d8aba256b7ef10087e44516553e"></a>

## Next pages — primary.rr_set_group.rr_set.ptr_record / 21ccee002f87 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-fd47b913cf7776b21ce65e86af95d67b63abfca6bec77f6fdb114ffd8b4e4959"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84167a9d73d91e149a2724854490536f597a5b04e54d205b8a38dd8b3d93fb70"></a>

## primary.rr_set_group.rr_set.srv_record — primary.rr_set_group.rr_set.srv_record / a78718ec2018 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.srv_record

<a id="canonical-817b0e7c207fc263379fa1d01565484986050a2c402bfa62d6d323d349a6f417"></a>

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

<a id="canonical-04930b1b59f2aae0e284ea35eb49b26db25a6e7621e42ce32d745a31172d0a9a"></a>

## Direct properties — primary.rr_set_group.rr_set.srv_record / a78718ec2018 / 3

<a id="canonical-ca746817db665410290c46f927cdd64edaecc6cc3875b097d1ee8d6dbcd3a7e1"></a>

<a id="canonical-4f6b34522b70a3e73ae6292880584bbee0a8c5b8e65236fcbfb1762df9949fa1"></a>

## name property — primary.rr_set_group.rr_set.srv_record / a78718ec2018 / 4

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

- [values](resources--dns_zone--reference--group-003.md#canonical-77b9e9eafbd9daa872b6cc74a4d6850d0361dc73cea63678a72a1f09eb1046c9): complete subsection reference.

<a id="canonical-32de26645d69d809102d9ccd96f3a60055b053c3e042a56c4e8c97e5f49feb5d"></a>

## Next pages — primary.rr_set_group.rr_set.srv_record / a78718ec2018 / 5

- [primary.rr_set_group.rr_set.srv_record.values](resources--dns_zone--reference--group-003.md#canonical-77b9e9eafbd9daa872b6cc74a4d6850d0361dc73cea63678a72a1f09eb1046c9)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-77b9e9eafbd9daa872b6cc74a4d6850d0361dc73cea63678a72a1f09eb1046c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4fb9924f152107b6dceb74adfbcb978eb717ed059688c1de8f84c58637adb45"></a>

## primary.rr_set_group.rr_set.srv_record.values — primary.rr_set_group.rr_set.srv_record.values / 4fc29e06b0c7 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-fd47b913cf7776b21ce65e86af95d67b63abfca6bec77f6fdb114ffd8b4e4959)
- primary.rr_set_group.rr_set.srv_record.values

<a id="canonical-3c5642a92bf376371be812a3d267061b806a34915c17603f8b93dc55555dce55"></a>

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

<a id="canonical-d081902fa44bd12e912c8b1b26eb2ffde601412cfacfdfa60094679fb895ba81"></a>

## Direct properties — primary.rr_set_group.rr_set.srv_record.values / 4fc29e06b0c7 / 3

<a id="canonical-38819f0f92a3313ebd7c19221a6e3553e9db171c7fab6a29543aaaf1c4acee21"></a>

<a id="canonical-209bf2d966563b6bb6eefb2aa018b1393ce0165a28699b101d26b829ebc5395c"></a>

## port property — primary.rr_set_group.rr_set.srv_record.values / 4fc29e06b0c7 / 4

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

<a id="canonical-3c3eea868da3b9132eaa62ea62bbbe0117b650b1f9f70fc25edd49feb3abbb04"></a>

<a id="canonical-70460d18a964feaf6840f0d31ec0404b103844163420d67707e66c4c02288dd3"></a>

## priority property — primary.rr_set_group.rr_set.srv_record.values / 4fc29e06b0c7 / 5

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

<a id="canonical-2cbc29682d5eeb9817e01110f128d129a79dbfefe2478c6334f3fddd9a612a8b"></a>

<a id="canonical-057f02f9ae9a7645ca0c2d9ebd818faf3b6cb9cec8a2bb2b7a092872a3ed11f9"></a>

## target property — primary.rr_set_group.rr_set.srv_record.values / 4fc29e06b0c7 / 6

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

<a id="canonical-558d091d205638efe260332d23671ac0225c5c09e5e4f54bf1801cf076d59018"></a>

<a id="canonical-f3cca3d4d02ca8482983ecd9f288c19891b136701d088fdbe7b90c3cb7c92c12"></a>

## weight property — primary.rr_set_group.rr_set.srv_record.values / 4fc29e06b0c7 / 7

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

<a id="canonical-1c38a0023a8f2ae44eba76d56ec87bc4c75b9eaffb13ce1b2242b56cb09feed1"></a>

## Next pages — primary.rr_set_group.rr_set.srv_record.values / 4fc29e06b0c7 / 8

- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-fd47b913cf7776b21ce65e86af95d67b63abfca6bec77f6fdb114ffd8b4e4959)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-c9cddb199c143e2778d9e2884ca288256832b00a0f2f4d886a453c679ff0a375"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afa7c2229567216f745d1307574f71ec9153d8767914bd93fe9ac34b0f94f8ff"></a>

## primary.rr_set_group.rr_set.sshfp_record — primary.rr_set_group.rr_set.sshfp_record / ce632e1d504d / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.sshfp_record

<a id="canonical-1967a16573c65b03213f601611ad04a692421a6a6f1b81d789ffbc9d27f766d9"></a>

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

<a id="canonical-da1409ba47b681206a1c594a2805c7439e2232112200635600ec09f5df8fbaac"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record / ce632e1d504d / 3

<a id="canonical-67396cbc8ac3069dcc1d732631262d787d510e87d6af64fa6116f1c14729c7cf"></a>

<a id="canonical-c61e3239c7a1698748beee21010f989516106f6fa3ad9acac53c87e2fb1cc588"></a>

## name property — primary.rr_set_group.rr_set.sshfp_record / ce632e1d504d / 4

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

- [values](resources--dns_zone--reference--group-003.md#canonical-56e4f399c24481bb24e731d26dfed5cc70a726f98b192c4ca2878e6d4a45267a): complete subsection reference.

<a id="canonical-c9c64908cc5c65c97af31a2987ab736263f296043d9900075000d97019220d67"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record / ce632e1d504d / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-56e4f399c24481bb24e731d26dfed5cc70a726f98b192c4ca2878e6d4a45267a)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-56e4f399c24481bb24e731d26dfed5cc70a726f98b192c4ca2878e6d4a45267a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52753c6a7d475bec78e6377829df4b873fea312dfe7dcb43b2ae528389e48a22"></a>

## primary.rr_set_group.rr_set.sshfp_record.values — primary.rr_set_group.rr_set.sshfp_record.values / 639fc4b58085 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-c9cddb199c143e2778d9e2884ca288256832b00a0f2f4d886a453c679ff0a375)
- primary.rr_set_group.rr_set.sshfp_record.values

<a id="canonical-f0428d7ad2c072f342aae4d74424ea1bc946e3741750d5234c23038ab5ffbeed"></a>

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

<a id="canonical-77d2fcc78e2d9396d82bca1bd28e8c632d2a5b85fcb4cd1447444daf8a5b028a"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record.values / 639fc4b58085 / 3

<a id="canonical-82a6a44f04985998dbeb4deffd9dfd784b47f9cef84175038a357b4d59fa2f35"></a>

<a id="canonical-5653b971a4b8f838c6fc237d44dd956f23019e6725c9a9105591cc38abc06b52"></a>

## algorithm property — primary.rr_set_group.rr_set.sshfp_record.values / 639fc4b58085 / 4

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

- [sha1_fingerprint](resources--dns_zone--reference--group-003.md#canonical-bb81f4666f234f3d48f180dc31985e770f3e812ed3da3ed5a3d665aa80edf841): complete subsection reference.

- [sha256_fingerprint](resources--dns_zone--reference--group-003.md#canonical-d30a2eb2ae1275a55206b396d1c4ae1912ae459dcc275d595f8ae776b2509cfb): complete subsection reference.

<a id="canonical-19b30ed7c3563c21781923288962e5766a1b90994c769daaba262f50519cdcf3"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record.values / 639fc4b58085 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-003.md#canonical-bb81f4666f234f3d48f180dc31985e770f3e812ed3da3ed5a3d665aa80edf841)
- [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-003.md#canonical-d30a2eb2ae1275a55206b396d1c4ae1912ae459dcc275d595f8ae776b2509cfb)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-c9cddb199c143e2778d9e2884ca288256832b00a0f2f4d886a453c679ff0a375)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-bb81f4666f234f3d48f180dc31985e770f3e812ed3da3ed5a3d665aa80edf841"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f562255c3d77e7ebbcc47a1d07b6da5b7fb58dd62507bdf547591d4954e8d0ed"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 4533d575b725 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-c9cddb199c143e2778d9e2884ca288256832b00a0f2f4d886a453c679ff0a375)
- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-56e4f399c24481bb24e731d26dfed5cc70a726f98b192c4ca2878e6d4a45267a)
- primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint

<a id="canonical-5c84b2dbc3bc128a5308ba63d84f119482580c7f0156118f3e7472a4f99d34f9"></a>

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

<a id="canonical-1a909fca66b19db9ba7a43380cdeaf97608d53d02092ec983a79f21103231599"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 4533d575b725 / 3

<a id="canonical-c69527b20ecdfd86d6680c0fff40b7cb3febe39304400f7b519eff403ad4e88a"></a>

<a id="canonical-a48aaf186e1c7db126b62c5ce0bb9873e1c3864135815e55fcef60e0ec4531cb"></a>

## fingerprint property — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 4533d575b725 / 4

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

<a id="canonical-b06e4065ec014256a20acbf80dda457331ca8d43a16e3b35dad79410667a3659"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint / 4533d575b725 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-56e4f399c24481bb24e731d26dfed5cc70a726f98b192c4ca2878e6d4a45267a)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-d30a2eb2ae1275a55206b396d1c4ae1912ae459dcc275d595f8ae776b2509cfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-208dd4e9436dad8c04e237dca60b02c103df259b044d1f2fa64b9a3937f3c747"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 61ba3d1e0943 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-c9cddb199c143e2778d9e2884ca288256832b00a0f2f4d886a453c679ff0a375)
- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-56e4f399c24481bb24e731d26dfed5cc70a726f98b192c4ca2878e6d4a45267a)
- primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

<a id="canonical-90f9b4eca7ffcbce57b501b825f9aeda043f270956e838b6073eac5b46c0c466"></a>

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

<a id="canonical-d7122ccdad71d4e4c5d6eca11f422cb22c55662c11a4e9b4e18630b772b22bc0"></a>

## Direct properties — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 61ba3d1e0943 / 3

<a id="canonical-6390f91ea90efb936cbc34fbcc74c133445a50f607a35e79160cf274508dd2a1"></a>

<a id="canonical-1f56aec05d0a27875555f93f0758b2132f096daf4cfac84f0412b90c90825c74"></a>

## fingerprint property — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 61ba3d1e0943 / 4

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

<a id="canonical-3e3abe0d00d7152e4dd6347187c5ca70dfb5739b5c421086584b5abf744cea6a"></a>

## Next pages — primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint / 61ba3d1e0943 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-56e4f399c24481bb24e731d26dfed5cc70a726f98b192c4ca2878e6d4a45267a)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-0db24ebdf992bc441effe1df0838893fc99bd28af63535fadf4301d50c3119d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5eb2d19c1b9f9db01cc46fa5d1ab24375d06b7eef376443f69d5a1bf3c17e8c"></a>

## primary.rr_set_group.rr_set.tlsa_record — primary.rr_set_group.rr_set.tlsa_record / a40ccb847dc3 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.tlsa_record

<a id="canonical-f4b10a1483b2c1444b579ca7ac76fe78c1a92d9dee3b1d3d389c1db2bafa6157"></a>

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

<a id="canonical-0123b35c2484e6c26051875ccb58fca6a576938ff24e203bda29eab05cfde85d"></a>

## Direct properties — primary.rr_set_group.rr_set.tlsa_record / a40ccb847dc3 / 3

<a id="canonical-75d06f0a4375ff2e6c8e6e5606244fd6b8465707069521930a839d682c20b6c5"></a>

<a id="canonical-427df328ce25fd6a5d98b0f7f3a27a77312b042d7e8626c87307c761a0e883f9"></a>

## name property — primary.rr_set_group.rr_set.tlsa_record / a40ccb847dc3 / 4

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

- [values](resources--dns_zone--reference--group-003.md#canonical-df041939faab6227854f3319175f87323ebce92e2f9235c13c0a21ccebe01288): complete subsection reference.

<a id="canonical-ca764376096d97757b5a6754518b596a57dc6f637f71cab1632c1a851b5fca3c"></a>

## Next pages — primary.rr_set_group.rr_set.tlsa_record / a40ccb847dc3 / 5

- [primary.rr_set_group.rr_set.tlsa_record.values](resources--dns_zone--reference--group-003.md#canonical-df041939faab6227854f3319175f87323ebce92e2f9235c13c0a21ccebe01288)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-df041939faab6227854f3319175f87323ebce92e2f9235c13c0a21ccebe01288"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24cacca91aa394ecd67152db557ddcfee49e9898ead1fbc5d53d1a78266a9f3f"></a>

## primary.rr_set_group.rr_set.tlsa_record.values — primary.rr_set_group.rr_set.tlsa_record.values / 167493659ea5 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0db24ebdf992bc441effe1df0838893fc99bd28af63535fadf4301d50c3119d4)
- primary.rr_set_group.rr_set.tlsa_record.values

<a id="canonical-a5bcf0deed728e6c55cd10d5f902eb9866a51aeff88192086e85b4993f0df530"></a>

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

<a id="canonical-2e60403a1941242fedaa15f4f4be9671efba1c73ecf0a1c59f1783a104202836"></a>

## Direct properties — primary.rr_set_group.rr_set.tlsa_record.values / 167493659ea5 / 3

<a id="canonical-85e5d88290496c08fc66922f9e50e30de2766d9a104059f7d94236338887526c"></a>

<a id="canonical-e0f8ae1524caef5642fab74cadb7960ea13af3e265a30071e918fe8aba4e730d"></a>

## certificate_association_data property — primary.rr_set_group.rr_set.tlsa_record.values / 167493659ea5 / 4

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

<a id="canonical-45f20f0c9bd93c8a9743b2fb729702905cda4f7453bf7f13fa8f8e93bc97cea6"></a>

<a id="canonical-d07af899d80489ada0e9013759aa10bb6ec2d97989b34459634a25e7c0797645"></a>

## certificate_usage property — primary.rr_set_group.rr_set.tlsa_record.values / 167493659ea5 / 5

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

<a id="canonical-184f7d6bf598fe1eb725e4c6d8e97384a44bcf706d9cd305af747d68bf2ecfe1"></a>

<a id="canonical-a79181698b8e91daf11b341b9e4c6c01e0e2a5c0a3b136a9f28c99fd08557ebf"></a>

## matching_type property — primary.rr_set_group.rr_set.tlsa_record.values / 167493659ea5 / 6

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

<a id="canonical-6d005b3242721466cb53f22bcc059c0120183f799391761935f3aedbe97f08fc"></a>

<a id="canonical-96336d92b1b67d416ca62e2c6933062a2122c2583e312a8d251491e1e3d35be0"></a>

## selector property — primary.rr_set_group.rr_set.tlsa_record.values / 167493659ea5 / 7

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

<a id="canonical-267480974e1ca3b9634220b8636264e8bb0103e7f45c17905e13c68d83b888e3"></a>

## Next pages — primary.rr_set_group.rr_set.tlsa_record.values / 167493659ea5 / 8

- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0db24ebdf992bc441effe1df0838893fc99bd28af63535fadf4301d50c3119d4)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-59755a283fb35414e8ffd5fcfc23876c4f29dc873fbb94d6d71dbb702cbaf497"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9ae1f35134ae551ccc71df2c4a4d05f4dde3d48250bd137490d5cff99a6c5c7"></a>

## primary.rr_set_group.rr_set.txt_record — primary.rr_set_group.rr_set.txt_record / e05fd71b5904 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-6dbdd49ccf69b9e92099937b34de04d05f8ff0360a2cf6a25a12ce2d2c3f2e94)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- primary.rr_set_group.rr_set.txt_record

<a id="canonical-50f54e8c2155c9e2d42d4031039d03e156041ad2780e0cdec0efb2220ff78c7e"></a>

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

<a id="canonical-d5133c4910fcabc85e41cee14a1900ce8cd73cab6422bdc99e5abf92e8c10b4c"></a>

## Direct properties — primary.rr_set_group.rr_set.txt_record / e05fd71b5904 / 3

<a id="canonical-8c86994d06088715740813022cae3507b28558c19790e9f9be3b9a60abfe2124"></a>

<a id="canonical-a5a750727915d7b9b88ed1a91c659326866aa00a6e859473e38d31c2c4a58961"></a>

## name property — primary.rr_set_group.rr_set.txt_record / e05fd71b5904 / 4

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

<a id="canonical-6789051a40fa4494af829f64d671b6a4ddc6a9468581b91a23293dfc8e4722a4"></a>

<a id="canonical-506ae09bbe8038fa2d5351e4f77e2eb6f700b2aa635d2dab1b86eafba912e637"></a>

## values property — primary.rr_set_group.rr_set.txt_record / e05fd71b5904 / 5

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

<a id="canonical-70713817e97f279e17fcd3621558f8cdd8d6f825be61ebc32d5b29968b3ca243"></a>

## Next pages — primary.rr_set_group.rr_set.txt_record / e05fd71b5904 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-f2a640fc2f00c770aafb80c4e295ede46056e6219982ddf1db27a504cd83c237)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-4888a3369cbc116c1b65198b284eb32389412885bed470f1c01e31f21955301c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e53e9f6c96905ccb4be771afeff414df2a20fc0ba485cee32b1f59f1288a7253"></a>

## primary.soa_parameters — primary.soa_parameters / 32f6c7d6f741 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- primary.soa_parameters

<a id="canonical-99a3950ee58f7323132356a10be24888d9022566671717ccd57139a105d9eaa6"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for soa parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refresh",
    "retry")}
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
soa_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-e9919e69fba5b1c89d0208090ae0b0585e1a982042352c11290d48f4cc4ba92b"></a>

## Direct properties — primary.soa_parameters / 32f6c7d6f741 / 3

<a id="canonical-ec7b000fd4458b9312b9d2a9276a4972791b8aa11a22db7dfda35b52d2eb7235"></a>

<a id="canonical-a251e6623fcfcd808604f6aec0019c15813a3cd019f0c980e53e44ce0bd81141"></a>

## expire property — primary.soa_parameters / 32f6c7d6f741 / 4

Type: `"number"`. Optional.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-b94507c79fdfc2f52cf7e0c1c6a31faa126f8cb37632c5bfd6e210e17cd44309"></a>

<a id="canonical-4cc6b271a180ae83c040c7a7190be5b172a1c52cb0cae5acbdd4ec7f786abf8b"></a>

## negative_ttl property — primary.soa_parameters / 32f6c7d6f741 / 5

Type: `"number"`. Optional.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-b6ea1b18927f2c4348fcd616ee9e4d5afa016462c5d82304c14bf363a0fd2a36"></a>

<a id="canonical-f1bbc1c885b780beb58211ab967149bf036a50a4bfd5a9a9cade61be83ded24d"></a>

## refresh property — primary.soa_parameters / 32f6c7d6f741 / 6

Type: `"number"`. Optional.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(3600, 2147483647),
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
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-1eaf7d30c2e59d3b7afe14b6c2771c7bb412c9cc69779a0b2a1de34e214727e5"></a>

<a id="canonical-22a78f90577980b5f2edc7575d84af94e6c12193143bab2c531e9daf8c3d7c03"></a>

## retry property — primary.soa_parameters / 32f6c7d6f741 / 7

Type: `"number"`. Optional.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

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

<a id="canonical-cd67c0aa5a10d3e7828d7681a8e21e86f36fbefb3684a63ff66fa6bb9777c880"></a>

<a id="canonical-a731e4998acd6710d86728d1ba8d1946470a2aaa07097ef106f1d5ff98b5c056"></a>

## ttl property — primary.soa_parameters / 32f6c7d6f741 / 8

Type: `"number"`. Optional.

TTL. SOA record time to live (in seconds)

Upstream description:

SOA record time to live (in seconds)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-ef3ab3bcf19763cb2a602fcafaf776aac73b968a9ed3f1e18936019f75a63571"></a>

## Next pages — primary.soa_parameters / 32f6c7d6f741 / 9

- [primary](resources--dns_zone--reference--group-001.md#canonical-db8a3235a2c87dc67c7696e4ccf6086b381c99ab00f8349273f1f074faf68db6)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-915ef6a26cadd3cdb1d03e8f1f87b29e6e5292057e8da86defb920ce53ea35c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a1df0c74938791c54c89b0096a9f480dd2d490c518c316e73aa2fa188489100"></a>

## secondary — secondary / e9650fb08382 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- secondary

<a id="canonical-2b11d4bd34678030dd7d0e737680d35191aa7af0d4c737ab4b790cdbc6aeb819"></a>

Type: `"object"`. single nested block, Optional.

SecondaryDNSCreateSpecType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_servers")}
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
secondary {
  # Configure direct properties listed below.
}
```

<a id="canonical-f8647d19046d5ced0c02d25267a7de0496d54d06fffe728db204fcfa2c62c143"></a>

## Direct properties — secondary / e9650fb08382 / 3

<a id="canonical-cae88882b862e89c7e51461eed76abc7cb2cd77780e4ee1b2a2bbbe12b3b00ca"></a>

<a id="canonical-83031bfb1b515267717fdff9c241cc5b909d3d1e0a653a945074dfa7f9eb2e9a"></a>

## primary_servers property — secondary / e9650fb08382 / 4

Type: `["list", "string"]`. Optional.

Configuration parameter for primary servers.

Upstream description:

Configuration parameter for primary servers

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9da4604c27a10e80196fcf20dac84bedfae8c71e6e470be38aa0f4ef965a5ecd"></a>

<a id="canonical-e0762d807c79257d0a7e222ed34c287932c9ecebf22ac00b351b7411a5176867"></a>

## tsig_key_algorithm property — secondary / e9650fb08382 / 5

Type: `"string"`. Optional.

\[Enum: HMAC\_MD5|UNDEFINED|HMAC\_SHA1|HMAC\_SHA224|HMAC\_SHA256|HMAC\_SHA384|HMAC\_SHA512\] TSIG
key value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC\_MD5:
HMAC\_MD5 - HMAC\_SHA1: HMAC\_SHA1 - HMAC\_SHA224: HMAC\_SHA224 - HMAC\_SHA256: HMAC\_SHA256 -
HMAC\_SHA384: HMAC\_SHA384 - HMAC\_SHA512: HMAC\_SHA512. Possible values are \`HMAC\_MD5\`,
\`UNDEFINED\`, \`HMAC\_SHA1\`, \`HMAC\_SHA224\`, \`HMAC\_SHA256\`, \`HMAC\_SHA384\`,
\`HMAC\_SHA512\`. Defaults to \`UNDEFINED\`.

Upstream description:

TSIG key value must be compatible with the specified algorithm

&#8203;- UNDEFINED: UNDEFINED

&#8203;- HMAC\_MD5: HMAC\_MD5

&#8203;- HMAC\_SHA1: HMAC\_SHA1

&#8203;- HMAC\_SHA224: HMAC\_SHA224

&#8203;- HMAC\_SHA256: HMAC\_SHA256

&#8203;- HMAC\_SHA384: HMAC\_SHA384

&#8203;- HMAC\_SHA512: HMAC\_SHA512.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNDEFINED",
  "enum": [
    "HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-765f3c5b1dde3d9a808c80a986177b9032dc3779cd52eda6f054eb82c8689478"></a>

<a id="canonical-531484742649caacaee7262bcecfc09554fd88628aa9e4f59d9d780ea73cc8f8"></a>

## tsig_key_name property — secondary / e9650fb08382 / 6

Type: `"string"`. Optional.

TSIG key name as used in TSIG protocol extension.

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

- [tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-806e5eaa033f23ae9b910e5e539669094e76d85b65b9454633dba4deb4072f07): complete subsection reference.

<a id="canonical-f76e11d82e1589024087edc2cdd480a8cab77cfd985886e5a6fef3fdebf51c08"></a>

## Next pages — secondary / e9650fb08382 / 7

- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-806e5eaa033f23ae9b910e5e539669094e76d85b65b9454633dba4deb4072f07)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-806e5eaa033f23ae9b910e5e539669094e76d85b65b9454633dba4deb4072f07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-783da404c58b02974da2ddaf2135ccc203b22d4246db765e8b700a4399100f34"></a>

## secondary.tsig_key_value — secondary.tsig_key_value / 89a0101089ef / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-915ef6a26cadd3cdb1d03e8f1f87b29e6e5292057e8da86defb920ce53ea35c9)
- secondary.tsig_key_value

<a id="canonical-5ff70a5f1b2717ab7dca34ac281aaf5361a73042d3103d60c24a78c3c61f6579"></a>

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
tsig_key_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2967f6dce210dedcaecc12b381783200d6155a4f8573dd98f6a76036de6f44db"></a>

## Direct properties — secondary.tsig_key_value / 89a0101089ef / 3

- [blindfold_secret_info](resources--dns_zone--reference--group-003.md#canonical-0988da66f8b4d0e00d6f82dedd23e40c4359f07056fbe90f7b892ab49b1dcbbc): complete subsection reference.

- [clear_secret_info](resources--dns_zone--reference--group-003.md#canonical-f2efa92abdbae474af287e5880be9b8ce5caba009376799b99fe01839f3b20b0): complete subsection reference.

<a id="canonical-fdc4917ca31c25e89d906af78276a8e43cc036af6a2aa5ceb8eb311639b96f5f"></a>

## Next pages — secondary.tsig_key_value / 89a0101089ef / 4

- [secondary.tsig_key_value.blindfold_secret_info](resources--dns_zone--reference--group-003.md#canonical-0988da66f8b4d0e00d6f82dedd23e40c4359f07056fbe90f7b892ab49b1dcbbc)
- [secondary.tsig_key_value.clear_secret_info](resources--dns_zone--reference--group-003.md#canonical-f2efa92abdbae474af287e5880be9b8ce5caba009376799b99fe01839f3b20b0)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-915ef6a26cadd3cdb1d03e8f1f87b29e6e5292057e8da86defb920ce53ea35c9)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-0988da66f8b4d0e00d6f82dedd23e40c4359f07056fbe90f7b892ab49b1dcbbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-745be32a5855314f89247e4df48929d681ad8c0be2c7095a67231f2e7cc9c005"></a>

## secondary.tsig_key_value.blindfold_secret_info — secondary.tsig_key_value.blindfold_secret_info / f842023a5c26 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-915ef6a26cadd3cdb1d03e8f1f87b29e6e5292057e8da86defb920ce53ea35c9)
- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-806e5eaa033f23ae9b910e5e539669094e76d85b65b9454633dba4deb4072f07)
- secondary.tsig_key_value.blindfold_secret_info

<a id="canonical-f119a7ff1489aa5b282fc812809501b01cfacca27a44c9cf97537a6c7e23661b"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-6fbaf6f755eaea67a5b5ed3d52c7deccd34aff04659727d17524a50607c7f2ac"></a>

## Direct properties — secondary.tsig_key_value.blindfold_secret_info / f842023a5c26 / 3

<a id="canonical-8f9fa3b635860b2d8146b2e17cab38fb31654c13779a97d610ec5c6ede9eb46f"></a>

<a id="canonical-1f1feea1f4590b6ed0f224b15fa2e18d408fb177b8a5c90fe160bfe839e249ff"></a>

## decryption_provider property — secondary.tsig_key_value.blindfold_secret_info / f842023a5c26 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-cbf02dfc34bd9372c79d9d21d31441817d87117a1df1f7aa8d22c24a73fcaefd"></a>

<a id="canonical-12807a79024fc893773bfed0ee081502bc9d22a1d4cd7ba55c323ae359fb5fa9"></a>

## location property — secondary.tsig_key_value.blindfold_secret_info / f842023a5c26 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2d9fc3d2c35218a5221e72b30b00d0e091afede412db1d46070eb0ef1641cd21"></a>

<a id="canonical-396121a56b72db7c81212e1d22d54433e40a6bdeedca70cff4c554b13a89ba0f"></a>

## store_provider property — secondary.tsig_key_value.blindfold_secret_info / f842023a5c26 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-419e1a95bfa737104fb2951a7c584fc85249b037dfcf7f070cc9712c917eb1dc"></a>

## Next pages — secondary.tsig_key_value.blindfold_secret_info / f842023a5c26 / 7

- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-806e5eaa033f23ae9b910e5e539669094e76d85b65b9454633dba4deb4072f07)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-f2efa92abdbae474af287e5880be9b8ce5caba009376799b99fe01839f3b20b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3958123e4406fa5dbebef78271a97df55ca6acbc2d128c91dde5927148c1e4bf"></a>

## secondary.tsig_key_value.clear_secret_info — secondary.tsig_key_value.clear_secret_info / 011e0b8d1bc6 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-915ef6a26cadd3cdb1d03e8f1f87b29e6e5292057e8da86defb920ce53ea35c9)
- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-806e5eaa033f23ae9b910e5e539669094e76d85b65b9454633dba4deb4072f07)
- secondary.tsig_key_value.clear_secret_info

<a id="canonical-8b79518335a7b1d15a02cace936e0a124b205b59e56b424deeb213b607720d2c"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e5145a94523f118369fcd6b0158c61d063f72d2984137eb1a0ed6f1d2e1e977"></a>

## Direct properties — secondary.tsig_key_value.clear_secret_info / 011e0b8d1bc6 / 3

<a id="canonical-377a887f8fa8ff470cf33d563a948a3112f49014e732aeacc900ce431351f2c7"></a>

<a id="canonical-816c3b3392d8a6d6bc9953cec4d3492a0e92675c1c66f4fa802a77846b21b212"></a>

## provider_ref property — secondary.tsig_key_value.clear_secret_info / 011e0b8d1bc6 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-aa2eae8a5166e7cad4ab8797e66fc837cd46187724615c1e838485e3971be751"></a>

<a id="canonical-390ac41aa13f2dda4f16dc8e03211bccd7f08e64435f8a3154876e565ce21435"></a>

## url property — secondary.tsig_key_value.clear_secret_info / 011e0b8d1bc6 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-e3259c6695d63a51bce764b9faed36d7d473053510ba63d81da78d63f9338d2d"></a>

## Next pages — secondary.tsig_key_value.clear_secret_info / 011e0b8d1bc6 / 6

- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-806e5eaa033f23ae9b910e5e539669094e76d85b65b9454633dba4deb4072f07)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)

<a id="canonical-592d8ee330c4fcf5e49e421377380abc8cd9a954a4313bed2da72b7dc44025bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7ca12def8441ded4c7cfcb0337f768edbdae383e85cc99253974aee7e987dc2"></a>

## timeouts — timeouts / 87ab7ebfd2a9 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- timeouts

<a id="canonical-30e39fcb34f8437dcda977f5d0332283e66222720e8328865080d09fea604511"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-c82a4fde3122c28e514fcb5c7e0c95bd71e00f03716ab11c3e65cf556e06bef1"></a>

## Direct properties — timeouts / 87ab7ebfd2a9 / 3

<a id="canonical-3076d76a9317e888315a3dd82513307bb4f512be658355d81908792a2deea226"></a>

<a id="canonical-fb1f7e0cdc6eceaec27a0d5990784ce40a9b7045659f939bf82b8da0d992a13b"></a>

## create property — timeouts / 87ab7ebfd2a9 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f944a9b66f978d13077ce0332d13de6a271ab5753b9eca7c5fa39514a327b469"></a>

<a id="canonical-3ea0544de2093a286275132b804b7224fe4aa92c372326061f4301040e878484"></a>

## delete property — timeouts / 87ab7ebfd2a9 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-393204c14936e2ae4c4d32a9932b5476d51310e569daab5ed92c4c36cb8865c8"></a>

<a id="canonical-ddf64f7a16ad8034bce442d08ff2cdd1119615a6ddfa9640db9c8396ee30e73e"></a>

## read property — timeouts / 87ab7ebfd2a9 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-fb12c93332b9fccf35b695cd6df6ddb9e37e8cc9fd514e7eee5b3674e8f973b9"></a>

<a id="canonical-ed0d92ec71bfed0a004339618a79ff3fb12e8070886405806b5be983600143fb"></a>

## update property — timeouts / 87ab7ebfd2a9 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1877e8be00b2007cbe35de00032170e0ae13ded56fdd00c6490a7269a31820d7"></a>

## Next pages — timeouts / 87ab7ebfd2a9 / 8

- [Property reference](resources--dns_zone--reference--group-001.md#canonical-02d3f05de9d2c23050ee45d4c5f961ef2fbca233d543b5433a9dcd643291e98d)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-eddf310c10a4a3df7f38c5872428ec47581d1e954090a5ef3951dde988dadcbb)
