---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-c85622e5613bf2f2f67933e440f9c0c45590b5006711bf3298a9ba595a8f3b9a"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.query_params.item / 284418b93dcf / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3303eaee6e6d196b4d787b6b775e24cc8db2decc21dc6c6e052474ea0adcdde5"></a>

<a id="canonical-35a6ec38c360044831baf0e9663c21397df5ddfc36257604cf0383f206afb158"></a>

## regex_values property — policy_based_challenge.rule_list.rules.spec.query_params.item / 284418b93dcf / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-734597fc0f377dc9b2f49e1ae08254d516d797c9fe564a1b985a70080e6d360a"></a>

<a id="canonical-2057b29ae02fe85eab8b4137a60a375e70a088d0b5e02a53bf246b5a0f432aad"></a>

## transformers property — policy_based_challenge.rule_list.rules.spec.query_params.item / 284418b93dcf / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f9399da8b75e118f79f499d4653cb95adb9cf3005b554aa711c94c118c7f7e21"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.query_params.item / 284418b93dcf / 7

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-013.md#canonical-511830efa2a971e782cd6ba252499d4cf40f63d59a6550161b659d22def20550)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ee283ca70e76f514eda54bce44c1789446a7a62c7a862a69afffcfe0cce80cfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6be309fd60352cee2c740c4344793d29b36d518ce33ada40b7ba9ad77e34fee3"></a>

## policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / c24011985cba / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-e7e050a0106b5c243929bc99c717fe27390190239f5f0da6543c11a425e13c71)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-6b6b9bb417be0328d783de67330d570206c65ca27165054e0fe0e20523560695)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-9b2a373a0a5b4919cfd5866358202625f8a727e1c663ac986aba136a77dfd8f1"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-89e2bf72463540326cd748a90ee57526872750f6848194c21076a448a1de4fc5"></a>

## Direct properties — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / c24011985cba / 3

<a id="canonical-3e48f8c2703c0cab7856f33285f2b685223c11f6d35a3088579282c5742d1f5e"></a>

<a id="canonical-35a5e4def51a37806d29d25c99877dbd496042f8e76584516a26b02f2ed6c21b"></a>

## classes property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / c24011985cba / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2e68dd235ceb57595f4c7df9d9f832d647e6033da296c4492ed63856a3835a27"></a>

<a id="canonical-f895d60bf784fe014096842331406455caf6607026f439a912681d5e41556229"></a>

## exact_values property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / c24011985cba / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-910b348086b09610d857dad04b8ffbc4fbe9f9a8517eda7957a1354fc0fd9663"></a>

<a id="canonical-2cf93a7633cc24db402b4e1a503e98fdfb712c4e210a4fd4dcb834d3ab1657e6"></a>

## excluded_values property — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / c24011985cba / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-43ac3dc766ce162a46e71c090e03190f8968422acc1e7ab51aa1cef42478adde"></a>

## Next pages — policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher / c24011985cba / 7

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-c1694f8f786e00f1b8d062766cb92e3219566178dd23899b40bc7c9d82c517a1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-885e305a1d0df444d63d64da6c7a4b0a54459b03c5eb86b881c6ad7b4cb37d06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d40fd1fddc7c7c6dc48e49dbaa574f28636f130fe95af90097784d37217e1608"></a>

## policy_based_challenge.temporary_user_blocking — policy_based_challenge.temporary_user_blocking / 089c50ddde95 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-61af764f539ccfe97dbbe7d8c991f38c0ed8a1359d32925dba0786ebc2bd84e1"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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
temporary_user_blocking {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3ec1a5ba1b6ebdf87d54d70aa34ba3275bbfeae79facafc138b8e37aee94fa8"></a>

## Direct properties — policy_based_challenge.temporary_user_blocking / 089c50ddde95 / 3

<a id="canonical-044f307427b2e843db99be6ea56aa0aa7c211ee872f6e65b976c0f06afe4fac4"></a>

<a id="canonical-d46647afaa88456ee650d9edbdc9be6eff62ebac979229a546dee73393fa2cad"></a>

## custom_page property — policy_based_challenge.temporary_user_blocking / 089c50ddde95 / 4

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in Base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-703ee464a50b3b61db8034144254894d4498702a6e37f082a4e82031527e52a4"></a>

## Next pages — policy_based_challenge.temporary_user_blocking / 089c50ddde95 / 5

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-4ea7912f7023fe07f657cac23ea745a73fb18ab04cb612a4904bf95a7d549a8f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7fec9131c6ed15dc6f3ef07db15d3e8df282bc9d77409c7698448757c3cc30c"></a>

## protected_cookies — protected_cookies / b6dbed857efa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- protected_cookies

<a id="canonical-abb61e2a08c528508d7efd4240b3e35f7368c5b42d8b21f2da3e78b2765dfb02"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("disable_tampering_protection",
    "enable_tampering_protection"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_cookies {
  # Configure direct properties listed below.
}
```

<a id="canonical-60f77fe5f40c03138c4670ee96565b919dbff92a64c351ffd8e2c37db0e15d71"></a>

## Direct properties — protected_cookies / b6dbed857efa / 3

- [add_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-6a7c7d653f7582d7d95f371b369fe1286f3cb609291d021e79b2403024844c37): complete subsection reference.

- [add_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-5e03f62ef2e89c35b7d641d63abb9a7126c0ebfd76e994d3bc0b8c4778b8237a): complete subsection reference.

- [disable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-f0feab8b373671229200139096dfa2286fa0f7b495f279b1fbda24c04535b8bb): complete subsection reference.

- [enable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-ca03d3ad6c00e4b66577fb39e66f1127dd94567eaaf79b68cba1505b5f905912): complete subsection reference.

- [ignore_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-3a1d0bb0f81aee2b32e5979a8a57f8f14c3a14650a2025df9be010ca2f7f667a): complete subsection reference.

- [ignore_max_age](resources--cdn_loadbalancer--reference--group-014.md#canonical-fd64c52cd7ef62bc301fdf25e584ab38c2cb100aca14e8874627e65a6d6d135d): complete subsection reference.

- [ignore_samesite](resources--cdn_loadbalancer--reference--group-014.md#canonical-41d988f83efbf92909f6fcb4a7f8e73a3e12b55cb4857286cfe1cc8c3aa30d46): complete subsection reference.

- [ignore_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-bdc6345a9ac815148169f9ed134356b27b714938c21009b680b94c769cefb104): complete subsection reference.

<a id="canonical-86924383e592034de22e44a18deb4ffbcdcd67a2f5f52c4219ff235445072731"></a>

<a id="canonical-60d8f6f0f103686889cc64411b4948e2e8036cc30511a65257db01c41e82c25c"></a>

## max_age_value property — protected_cookies / b6dbed857efa / 4

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-31f19acc26b4ea3d3cba62ec2463f1c32e31b6eabbdfa2fd5e2be27005fda3f1"></a>

<a id="canonical-8bb1a865d78098954185654eb046a2836ec39f092f8e0dc9d118a6d83695ec5f"></a>

## name property — protected_cookies / b6dbed857efa / 5

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](resources--cdn_loadbalancer--reference--group-014.md#canonical-912be3413bfb4dbe9767fe80771df63c1e1228677cc608ed4917bb9a04aac4ec): complete subsection reference.

- [samesite_none](resources--cdn_loadbalancer--reference--group-014.md#canonical-6e8e17b27d9770efe73908ab6bdce6a8a5f92489df11728cf92dfcc58b105208): complete subsection reference.

- [samesite_strict](resources--cdn_loadbalancer--reference--group-014.md#canonical-03b1cff90b22026550116c9bdef6948f77baab58af152a122554eea171d4274b): complete subsection reference.

<a id="canonical-f10efe357b8acba2f06425c43a0535b2b04758f7464540bd18db874d0e71d737"></a>

## Next pages — protected_cookies / b6dbed857efa / 6

- [protected_cookies.add_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-6a7c7d653f7582d7d95f371b369fe1286f3cb609291d021e79b2403024844c37)
- [protected_cookies.add_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-5e03f62ef2e89c35b7d641d63abb9a7126c0ebfd76e994d3bc0b8c4778b8237a)
- [protected_cookies.disable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-f0feab8b373671229200139096dfa2286fa0f7b495f279b1fbda24c04535b8bb)
- [protected_cookies.enable_tampering_protection](resources--cdn_loadbalancer--reference--group-014.md#canonical-ca03d3ad6c00e4b66577fb39e66f1127dd94567eaaf79b68cba1505b5f905912)
- [protected_cookies.ignore_httponly](resources--cdn_loadbalancer--reference--group-014.md#canonical-3a1d0bb0f81aee2b32e5979a8a57f8f14c3a14650a2025df9be010ca2f7f667a)
- [protected_cookies.ignore_max_age](resources--cdn_loadbalancer--reference--group-014.md#canonical-fd64c52cd7ef62bc301fdf25e584ab38c2cb100aca14e8874627e65a6d6d135d)
- [protected_cookies.ignore_samesite](resources--cdn_loadbalancer--reference--group-014.md#canonical-41d988f83efbf92909f6fcb4a7f8e73a3e12b55cb4857286cfe1cc8c3aa30d46)
- [protected_cookies.ignore_secure](resources--cdn_loadbalancer--reference--group-014.md#canonical-bdc6345a9ac815148169f9ed134356b27b714938c21009b680b94c769cefb104)
- [protected_cookies.samesite_lax](resources--cdn_loadbalancer--reference--group-014.md#canonical-912be3413bfb4dbe9767fe80771df63c1e1228677cc608ed4917bb9a04aac4ec)
- [protected_cookies.samesite_none](resources--cdn_loadbalancer--reference--group-014.md#canonical-6e8e17b27d9770efe73908ab6bdce6a8a5f92489df11728cf92dfcc58b105208)
- [protected_cookies.samesite_strict](resources--cdn_loadbalancer--reference--group-014.md#canonical-03b1cff90b22026550116c9bdef6948f77baab58af152a122554eea171d4274b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6a7c7d653f7582d7d95f371b369fe1286f3cb609291d021e79b2403024844c37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-709c1a69bd1a8d2cf0c39775c8206faf58338235674f465c3a02423909dd309d"></a>

## protected_cookies.add_httponly — protected_cookies.add_httponly / f3891d0d3ba7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.add_httponly

<a id="canonical-142e737292099093aaf8c16a0a6b17499eb4e56d510cac66ddca4286dc517981"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-4af3f4e5287a8469ad8ca989a42b5edb955526abc11cfe782fa813b8291cfcee"></a>

## Direct properties — protected_cookies.add_httponly / f3891d0d3ba7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-708e8df4f3be05bd9825623ee8075b705805238f51c379e80ae00929fadf2cba"></a>

## Next pages — protected_cookies.add_httponly / f3891d0d3ba7 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5e03f62ef2e89c35b7d641d63abb9a7126c0ebfd76e994d3bc0b8c4778b8237a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82d0b410181790ed40958543b1cc1a4b178ae62ab34b2d3a1b18ebcf72ce5fd4"></a>

## protected_cookies.add_secure — protected_cookies.add_secure / b67248f36a33 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.add_secure

<a id="canonical-e43d239bf5098622ccd8157a0af9574f6dee47f99e21f9a956644da6f1bf651a"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
add_secure = {}
```

<a id="canonical-ab4cd7048fd59d53be143766f2df96caec3d6b19af058f6c408e3c34542780bb"></a>

## Direct properties — protected_cookies.add_secure / b67248f36a33 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bab8bd46471bf7971db55a981347f7d4993929d877b71b5784f7603e5fc9b35a"></a>

## Next pages — protected_cookies.add_secure / b67248f36a33 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f0feab8b373671229200139096dfa2286fa0f7b495f279b1fbda24c04535b8bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f347a9b29c7d74de051c974fe2a5decb4e72867c5a0525b75c01885f38c15d5e"></a>

## protected_cookies.disable_tampering_protection — protected_cookies.disable_tampering_protection / 2b912703794a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.disable_tampering_protection

<a id="canonical-cf7965b8459987253be48581cec083cdd2be049e60033367f6988c759ec6c46e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

<a id="canonical-d68209b1c08001ddb08bd3be4529dccfb0a92aa9ca6f8b7237448c660c4bdb35"></a>

## Direct properties — protected_cookies.disable_tampering_protection / 2b912703794a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3e7f9756c0febb468d4ec8f05a1d3f439f2112eb513c30895c071a123ff99d7"></a>

## Next pages — protected_cookies.disable_tampering_protection / 2b912703794a / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ca03d3ad6c00e4b66577fb39e66f1127dd94567eaaf79b68cba1505b5f905912"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cddebd7380352d3e641da2df7a61cad000cdba71a4761484eef070b46b60f3b"></a>

## protected_cookies.enable_tampering_protection — protected_cookies.enable_tampering_protection / 4b4354c09205 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.enable_tampering_protection

<a id="canonical-450a77ea7ad1c30d96e4cc28c25d70fccd99a0ddb935fee315ddcf8e296779b3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tampering protection.

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
enable_tampering_protection = {}
```

<a id="canonical-61237da3484c60a2f6a400e6143b2690c69049a0c578015f4a12abfc95249f3a"></a>

## Direct properties — protected_cookies.enable_tampering_protection / 4b4354c09205 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-660ca179cb9f312696bf20d4b5143acba648730c44e0a9969b8a0712d9fba249"></a>

## Next pages — protected_cookies.enable_tampering_protection / 4b4354c09205 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3a1d0bb0f81aee2b32e5979a8a57f8f14c3a14650a2025df9be010ca2f7f667a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de70260988b01d406a85080a88633aff577ab6721f6cae7dc99eb2c5ca904e7f"></a>

## protected_cookies.ignore_httponly — protected_cookies.ignore_httponly / 6cbbf62254af / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.ignore_httponly

<a id="canonical-36f7b75196ebcd1fcd425a17eebf95a1f61d72f0148e2af3baaac12008b5f862"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-f1aab019695db1feb0d27ec82b3436e2bf8dbeb39fd03113d831ae1132f6cb5d"></a>

## Direct properties — protected_cookies.ignore_httponly / 6cbbf62254af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b608f7171ab4606130a5d0a514c41f7a4a26cbd2e9fa439b85598885e6b6bfcd"></a>

## Next pages — protected_cookies.ignore_httponly / 6cbbf62254af / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fd64c52cd7ef62bc301fdf25e584ab38c2cb100aca14e8874627e65a6d6d135d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97a3b1383537e2af33ca671462e5b9bceb9fa21c04695da4680d01268c92ef5e"></a>

## protected_cookies.ignore_max_age — protected_cookies.ignore_max_age / 40aeedf25e0e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.ignore_max_age

<a id="canonical-e3b44898c81e57843be9ca429d545995be460616e1d6faf3e2ba8d2ed739ad99"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

<a id="canonical-41103c4e4278ac4d477eb3851d5ddff32a4b97bc4f1ce3cdb36cb76c4e03cd7e"></a>

## Direct properties — protected_cookies.ignore_max_age / 40aeedf25e0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8127626a2d5cb565b64c930ae294cd836d56dc4c38ddb0c2ba01457dade3c4c4"></a>

## Next pages — protected_cookies.ignore_max_age / 40aeedf25e0e / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-41d988f83efbf92909f6fcb4a7f8e73a3e12b55cb4857286cfe1cc8c3aa30d46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b16103a86f3039f5755601e469af86cfef052d28ecb6f2b22253fff07f2e79e8"></a>

## protected_cookies.ignore_samesite — protected_cookies.ignore_samesite / c5788f00ac55 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.ignore_samesite

<a id="canonical-2085945b87f21c18cfd7298c7bf8920ef2a7e75c2e81b527cf16ea8976c6e818"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
ignore_samesite = {}
```

<a id="canonical-03ced9fad898274f56055f7e7cc7c3dc39d07ec2fbdefb6f6b2fc28e811d74e0"></a>

## Direct properties — protected_cookies.ignore_samesite / c5788f00ac55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6987e73c463ddba39913f6b19ee38ff4dae5f80573d6db59bef99f80bd4e675"></a>

## Next pages — protected_cookies.ignore_samesite / c5788f00ac55 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bdc6345a9ac815148169f9ed134356b27b714938c21009b680b94c769cefb104"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b09090c59bcbc1fe146a1e6ae38908b80f1a4e4560530633382232edeab67910"></a>

## protected_cookies.ignore_secure — protected_cookies.ignore_secure / 9679f67bbc8a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.ignore_secure

<a id="canonical-718d1af55e52a2b1f5e03b1e757bb5822cf38eead86c36853c518279bad68d6e"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
ignore_secure = {}
```

<a id="canonical-b0f2d2e0010356ee23c28d3f8596f581ac3fa383f22b0db74b441616f50529a5"></a>

## Direct properties — protected_cookies.ignore_secure / 9679f67bbc8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac6f9ea71dc1d89dfe42385f2c84f291df2fbcc842ed540d6d7053d2f87ee154"></a>

## Next pages — protected_cookies.ignore_secure / 9679f67bbc8a / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-912be3413bfb4dbe9767fe80771df63c1e1228677cc608ed4917bb9a04aac4ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea3d7f31df419c9b61affef487b50beadacb847e5742d5d78973cef2112ae721"></a>

## protected_cookies.samesite_lax — protected_cookies.samesite_lax / a0b051452d13 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.samesite_lax

<a id="canonical-165c2199d83b9727baa8a68e9af08022a2d642a194de4277af4c2e946c4beeef"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
samesite_lax = {}
```

<a id="canonical-d8f554991aa8a9d05d743df07b4534583afd7c6e7542721cc34933927aa53fcf"></a>

## Direct properties — protected_cookies.samesite_lax / a0b051452d13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-504fccaced038f265d9b16347a1399cb41d9809e67e149d706aa2ec42b3f36b2"></a>

## Next pages — protected_cookies.samesite_lax / a0b051452d13 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6e8e17b27d9770efe73908ab6bdce6a8a5f92489df11728cf92dfcc58b105208"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-395d85d8005fb7bb690de29ef63b748bce61f389ab24ee3481eea7c72639038d"></a>

## protected_cookies.samesite_none — protected_cookies.samesite_none / 5096d7a74e61 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.samesite_none

<a id="canonical-52507219bf002183163c211b606ad3f6ab4d3593608db515a5229cf93ba22958"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
samesite_none = {}
```

<a id="canonical-206d84892441fb8f753f71b0af1cd595df09c74efff78cb4884a8a70531fd103"></a>

## Direct properties — protected_cookies.samesite_none / 5096d7a74e61 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa2773f44bb85f3b913155af49c41741da8aa30770fd4bf5b9cbadb5ef5669e7"></a>

## Next pages — protected_cookies.samesite_none / 5096d7a74e61 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-03b1cff90b22026550116c9bdef6948f77baab58af152a122554eea171d4274b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d94c19e91f85febfd4d58ff0259ea7fc0864125c3e3432965f91db11ae17183d"></a>

## protected_cookies.samesite_strict — protected_cookies.samesite_strict / 398546cf88c0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- protected_cookies.samesite_strict

<a id="canonical-debaab5a6267950f06a192de30e34dfdc838d0c8eba9032f06a298a692be85f5"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
samesite_strict = {}
```

<a id="canonical-4c7ee5a71f333edcfd1ce0d7a1c1dd84302b115be9a8a7541a9157006bf1a54f"></a>

## Direct properties — protected_cookies.samesite_strict / 398546cf88c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-56f11c014671a556040b523dbff348879140da8b97469cb25b940ac7122b188b"></a>

## Next pages — protected_cookies.samesite_strict / 398546cf88c0 / 4

- [protected_cookies](resources--cdn_loadbalancer--reference--group-014.md#canonical-028d756ed3f6103a40ebc66d6d948f8270ab1a63cc5e5d8b95c8456e68bd391a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb739d2075dd607a4bac280ea1cbc311d8cd9dbfa06563e6d9987635586ac598"></a>

## rate_limit — rate_limit / 077113696bda / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- rate_limit

<a id="canonical-fb642de1cd63afa397778edfca2fdf6d50ff1e631985a951ab1d68f124467a10"></a>

Type: `"object"`. single nested block, Optional.

RateLimitConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("no_policies",
    "policies")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

Terraform syntax:

```terraform
rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-b7223810021e9685d712a62056593e172a6da6213fa4681d9beba2d0d18fdab4"></a>

## Direct properties — rate_limit / 077113696bda / 3

- [custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-e58aeb967eb19c0afe2a6b3d9ffbcceba77b696ffb6fdc93bf380c56fe0fe87d): complete subsection reference.

- [ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-1d559566d15354a99f6c70b3a452757c5f3338da0370df1f491cd9ece900fbdc): complete subsection reference.

- [no_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-a712760b104f81c9da9b46567a5a06a68499e6e4061bf7c5884542149e8267ac): complete subsection reference.

- [no_policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-1e0a9eb7c0aa5ca9217a2c0384d60eacb3da8f812941e72d1a74d2e8053a4b14): complete subsection reference.

- [policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-3e1877c214938ee14368a23018a17892d5f1d966ad97af83048450278cf4e2e7): complete subsection reference.

- [rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5): complete subsection reference.

<a id="canonical-45cdf0e07205d79b7f84150e4328313331c418299105347d1004544b8ae6d487"></a>

## Next pages — rate_limit / 077113696bda / 4

- [rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-e58aeb967eb19c0afe2a6b3d9ffbcceba77b696ffb6fdc93bf380c56fe0fe87d)
- [rate_limit.ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-1d559566d15354a99f6c70b3a452757c5f3338da0370df1f491cd9ece900fbdc)
- [rate_limit.no_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-a712760b104f81c9da9b46567a5a06a68499e6e4061bf7c5884542149e8267ac)
- [rate_limit.no_policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-1e0a9eb7c0aa5ca9217a2c0384d60eacb3da8f812941e72d1a74d2e8053a4b14)
- [rate_limit.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-3e1877c214938ee14368a23018a17892d5f1d966ad97af83048450278cf4e2e7)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e58aeb967eb19c0afe2a6b3d9ffbcceba77b696ffb6fdc93bf380c56fe0fe87d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3af07d08a2b73eb024b4134c68d061d272c60aeedb577a79f24176f69aedb97"></a>

## rate_limit.custom_ip_allowed_list — rate_limit.custom_ip_allowed_list / 3fbeb72ca963 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- rate_limit.custom_ip_allowed_list

<a id="canonical-ded5cd743703686b9a0547f9eef9b227f6031e3ae668c5aec425bcfd4c2ab141"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-5602f291fcf9a7fd69eb957de0d53fc1e28134ff2798610dd3463ae68e99af09"></a>

## Direct properties — rate_limit.custom_ip_allowed_list / 3fbeb72ca963 / 3

- [rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-014.md#canonical-5f07ed44d6ee86376817b45c7e7a545cb67f316f294a482572851b608b40db0a): complete subsection reference.

<a id="canonical-beb7f4d430de7f7c63e8e2466a202ac8373452a50e97765f1402046a87625531"></a>

## Next pages — rate_limit.custom_ip_allowed_list / 3fbeb72ca963 / 4

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-014.md#canonical-5f07ed44d6ee86376817b45c7e7a545cb67f316f294a482572851b608b40db0a)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5f07ed44d6ee86376817b45c7e7a545cb67f316f294a482572851b608b40db0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-717444ce17fe982d92fc76c4aa45b247cb1d7be22745910b152e736202e1e0c0"></a>

## rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 72375ee2e65a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-e58aeb967eb19c0afe2a6b3d9ffbcceba77b696ffb6fdc93bf380c56fe0fe87d)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-c3cb73c1778a5d234cbe7cee7b7cd3b8e11b80d139ccd8a8c30d2cba050b802e"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0a464ff56f24b15fc0da9a5245378d78a5c27b59f27b0c059989a16967af6fa7"></a>

## Direct properties — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 72375ee2e65a / 3

<a id="canonical-11284054b3b6a9eef9a4f0a3b29d67fcd83921dab2a2c538c567a0dab415622f"></a>

<a id="canonical-87089e989dcdabe49abe586ea76798520894c2575442cbf474abd38788b17dc3"></a>

## name property — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 72375ee2e65a / 4

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

<a id="canonical-c12895b92b0f3a7b30a8712ac6f3a032defde5b77823910cdcbc3bf15f8b929f"></a>

<a id="canonical-683cd5e93bf52675e3435f65505c74162df6cb33e981e07af19dc49a261726ee"></a>

## namespace property — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 72375ee2e65a / 5

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

<a id="canonical-91cd4aec929e56ffc6ea6cda8eca3ce1e0e00fd5f53246ad9c3cedd4b1b62669"></a>

<a id="canonical-9c76f159d690d5a140ba77facefe2f54e26ae6fd2a51e943e2df44f0d2d192df"></a>

## tenant property — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 72375ee2e65a / 6

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

<a id="canonical-73601b5e2b67575ce96d2bd7aa3f9d24f19e05aec8eee0a190c9b88a260d0d88"></a>

## Next pages — rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes / 72375ee2e65a / 7

- [rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-e58aeb967eb19c0afe2a6b3d9ffbcceba77b696ffb6fdc93bf380c56fe0fe87d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1d559566d15354a99f6c70b3a452757c5f3338da0370df1f491cd9ece900fbdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aafc98086ba7226723f09a2159f8e34c05a37db85e189b00efcc9720838e5c3f"></a>

## rate_limit.ip_allowed_list — rate_limit.ip_allowed_list / 8ce07d13ed11 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- rate_limit.ip_allowed_list

<a id="canonical-fab23082c9ab6da25a3b276c2c79cfed99d22e315f81abb93c5439d061b2e6b9"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-6162a18206d2dfc36ba438dd389a3af0aae3d96a3188b61c17ea09b3ac87dd19"></a>

## Direct properties — rate_limit.ip_allowed_list / 8ce07d13ed11 / 3

<a id="canonical-3511b514674c3b839fb1d3d8097d6c5ac0b62f774fa06293a725b086c51c4cef"></a>

<a id="canonical-463d8a373a3479bdd11d2ca55aa904ea7ed5869c62ed49b905baa54287d328ab"></a>

## prefixes property — rate_limit.ip_allowed_list / 8ce07d13ed11 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9ca51fc881cdaf61cd77ded3f93b1a6a831f97c177e4f47f321c926dedb2754f"></a>

## Next pages — rate_limit.ip_allowed_list / 8ce07d13ed11 / 5

- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a712760b104f81c9da9b46567a5a06a68499e6e4061bf7c5884542149e8267ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caec4a4349794bf480fee696c050f201ba81545bd6e5df8f9ddf1b2db342b9a6"></a>

## rate_limit.no_ip_allowed_list — rate_limit.no_ip_allowed_list / 7d6e6d33a8fe / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- rate_limit.no_ip_allowed_list

<a id="canonical-2569f2848cf0326415248cfbde0092d2c120be3af1fc65fac48dae756244bd39"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_ip_allowed_list = {}
```

<a id="canonical-582d3208f62f03f2fe442e9a1311f9d9965f51d84f88986bbffe9dcced1b872c"></a>

## Direct properties — rate_limit.no_ip_allowed_list / 7d6e6d33a8fe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-beafe13da2937012740205d81169625f3c23e5dfbb122d164f0c45704a415e7e"></a>

## Next pages — rate_limit.no_ip_allowed_list / 7d6e6d33a8fe / 4

- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1e0a9eb7c0aa5ca9217a2c0384d60eacb3da8f812941e72d1a74d2e8053a4b14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd7147b99369f0b367091c3140947ca000fdb0722ecabf36f9ccc28d4abd5219"></a>

## rate_limit.no_policies — rate_limit.no_policies / 90a6f7f5c488 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- rate_limit.no_policies

<a id="canonical-2162292a08ee19e8b13d38b2a2e7b7a15f5cb62a0a9f975797ad33d24bcbc437"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_policies = {}
```

<a id="canonical-5713dfcfe09cf4239b8aa7275839d3533aa76b193ff9436860898e292693dc5b"></a>

## Direct properties — rate_limit.no_policies / 90a6f7f5c488 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fc0f4dc1fcabb32cb3a237da334f66d1575055718c4595c5ac02ca5e706cca9b"></a>

## Next pages — rate_limit.no_policies / 90a6f7f5c488 / 4

- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3e1877c214938ee14368a23018a17892d5f1d966ad97af83048450278cf4e2e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cb32cf8573e9206313fbfce87e64a01237772e7bb8808e4314b7fa09ad2cd7c"></a>

## rate_limit.policies — rate_limit.policies / d4195bbb8493 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- rate_limit.policies

<a id="canonical-c1d71c834fe26e5986ac09c2afd6de37809cca8a0a4b49b88f255933c5feb0da"></a>

Type: `"object"`. single nested block, Optional.

List of rate limiter policies to be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
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
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-b79c76d1c386a7be2700e835b9bcc9588590fd02ff7067c8ddddbc9fcc9f38fa"></a>

## Direct properties — rate_limit.policies / d4195bbb8493 / 3

- [policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-cd19c5773f7d4c1ae626bcbd44378f5ee8152a8a4c962cb9c579346ca4d20ff3): complete subsection reference.

<a id="canonical-076f523665354b0b1202730cfcf92f1d218718edae4ce3eb2cff15efe1839889"></a>

## Next pages — rate_limit.policies / d4195bbb8493 / 4

- [rate_limit.policies.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-cd19c5773f7d4c1ae626bcbd44378f5ee8152a8a4c962cb9c579346ca4d20ff3)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-cd19c5773f7d4c1ae626bcbd44378f5ee8152a8a4c962cb9c579346ca4d20ff3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c73a451cf41be721a637e03872de18b2f9470612526f0a77308362193e46f35"></a>

## rate_limit.policies.policies — rate_limit.policies.policies / c6e6626eaa22 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-3e1877c214938ee14368a23018a17892d5f1d966ad97af83048450278cf4e2e7)
- rate_limit.policies.policies

<a id="canonical-170397d582207e8b4112e97edffd77a3c18f460d0c07130e65d8f3ca1a0abd51"></a>

Type: `"object"`. list nested block, Optional.

Rate Limiter Policies. Ordered list of rate limiter policies.

Upstream description:

Ordered list of rate limiter policies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1bf0004e2e5d82469d49e62bc1296f9d9f10b6546dfc78f048d2729ab9144319"></a>

## Direct properties — rate_limit.policies.policies / c6e6626eaa22 / 3

<a id="canonical-14e1496cd34375ad1848ca475316d1bcbf9cba200e46f9983353e0342e610e62"></a>

<a id="canonical-e045c24f148665282a2f70a4590d19c6d0b57f499cd6ce591b0cb26436918a27"></a>

## name property — rate_limit.policies.policies / c6e6626eaa22 / 4

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

<a id="canonical-546ff50123633b5b2a62b76cb6904618b22b262336b7d11c28f0abdbc9b3e68f"></a>

<a id="canonical-02c5d93bdff36210605672f869e334250787bba6c3410a1e27a3796f6b9089d4"></a>

## namespace property — rate_limit.policies.policies / c6e6626eaa22 / 5

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

<a id="canonical-c16420f9ac040f35a538e61c23527cfa9365e7660c85ae03fca6f7ae4fe32c0a"></a>

<a id="canonical-8193ecc2c8e53297fa0262c8a689b59d412e4f04ebfb180987a540e82d893b65"></a>

## tenant property — rate_limit.policies.policies / c6e6626eaa22 / 6

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

<a id="canonical-790fb7c1730daf3a66dcec9c2737f23dd3c83187c9e08d27c9e36a682db8a6c1"></a>

## Next pages — rate_limit.policies.policies / c6e6626eaa22 / 7

- [rate_limit.policies](resources--cdn_loadbalancer--reference--group-014.md#canonical-3e1877c214938ee14368a23018a17892d5f1d966ad97af83048450278cf4e2e7)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8d39d480d30556a9a0ab708675c362c8d9d9b63e3c879174520bdf8bb1362bb"></a>

## rate_limit.rate_limiter — rate_limit.rate_limiter / 0fd3d91c6408 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- rate_limit.rate_limiter

<a id="canonical-e2bbb03072131ee296934f868756d6b3bbb04b58940536888e7a95fb2afcaefe"></a>

Type: `"object"`. single nested block, Optional.

Tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Upstream description:

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("total_number"),
  validators.ConflictingObjectAttributes("action_block",
    "disabled"),
  validators.ConflictingObjectAttributes("leaky_bucket",
    "token_bucket")}
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
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

Terraform syntax:

```terraform
rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-767594c52d005634425f3b61862812dceeca08241e9c1afa89aa004cd47894ba"></a>

## Direct properties — rate_limit.rate_limiter / 0fd3d91c6408 / 3

- [action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d): complete subsection reference.

<a id="canonical-990f21c99971636cdbf9b47be96ea9eb3c2d05f3bc9c7a64e24c7f4930b8a16a"></a>

<a id="canonical-9418906a7d4804bafab99574f6fead2e11e4fb7d6a8ad0aeae64be08553bd639"></a>

## burst_multiplier property — rate_limit.rate_limiter / 0fd3d91c6408 / 4

Type: `"number"`. Optional.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](resources--cdn_loadbalancer--reference--group-014.md#canonical-195c3959c0bed9aa24eb4fdd6ea585123bf59051f0e0947ce5429464f3822984): complete subsection reference.

- [leaky_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-99c721be2080364ec24b7b2f29c27d97fb5571fbb9c77c4c611abacb9b118e31): complete subsection reference.

<a id="canonical-39ab9ddc9dc6115d516c56151bd31b3b1c38690d190e6c376e7a2da5f78b3c24"></a>

<a id="canonical-1ed6aa204062454b08ade86ee89c76e0886245038b1a68ff7bf888a7861adb29"></a>

## period_multiplier property — rate_limit.rate_limiter / 0fd3d91c6408 / 5

Type: `"number"`. Optional, Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Upstream description:

This setting, combined with Per Period units, provides a duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-e520a1225186a952c49f5a419213ce1a18c41f913ee1215a398196b7ce0d4200): complete subsection reference.

<a id="canonical-977f190f9f9e927e5697b00cb04099ce5d14f071b1e2a7c8ca5baad43b06cbc1"></a>

<a id="canonical-0a6dbe3a76dad847b9028ef3277926435a60ebc834c17cf408d18d0dbe8ee0bb"></a>

## total_number property — rate_limit.rate_limiter / 0fd3d91c6408 / 6

Type: `"number"`. Optional.

The total number of allowed requests per rate-limiting period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-ab024944b74a28433b3fb17d61baee1bf8520b5ac243a32691966d3ba784c9bd"></a>

<a id="canonical-45443aad6d82349605ba938f0f6306c9c89df6f33121d1f4554887a1dd46c033"></a>

## unit property — rate_limit.rate_limiter / 0fd3d91c6408 / 7

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-89092cd6fbe3eedffe7db73b6525e1ca678b8afb030b2fd4dc7f72fa0161502e"></a>

## Next pages — rate_limit.rate_limiter / 0fd3d91c6408 / 8

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d)
- [rate_limit.rate_limiter.disabled](resources--cdn_loadbalancer--reference--group-014.md#canonical-195c3959c0bed9aa24eb4fdd6ea585123bf59051f0e0947ce5429464f3822984)
- [rate_limit.rate_limiter.leaky_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-99c721be2080364ec24b7b2f29c27d97fb5571fbb9c77c4c611abacb9b118e31)
- [rate_limit.rate_limiter.token_bucket](resources--cdn_loadbalancer--reference--group-014.md#canonical-e520a1225186a952c49f5a419213ce1a18c41f913ee1215a398196b7ce0d4200)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caeda58a5f757fa1e93a46e05642693c70b6151987e859b6e5f1a70410601d2c"></a>

## rate_limit.rate_limiter.action_block — rate_limit.rate_limiter.action_block / bd35342eebd7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- rate_limit.rate_limiter.action_block

<a id="canonical-cd7a095f4bd55836918435c0eb3c43291ac395cb10c40c28c1a3858709b417c5"></a>

Type: `"object"`. single nested block, Optional.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes"),
  validators.ConflictingObjectAttributes("hours",
    "seconds"),
  validators.ConflictingObjectAttributes("minutes",
    "seconds")}
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
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

Terraform syntax:

```terraform
action_block {
  # Configure direct properties listed below.
}
```

<a id="canonical-9d0dc5d0d3bc040967a670e0f82c931680f2d05b083f68b73609a1b1a1169ae0"></a>

## Direct properties — rate_limit.rate_limiter.action_block / bd35342eebd7 / 3

- [hours](resources--cdn_loadbalancer--reference--group-014.md#canonical-01a1d8044fbe72fbffc3f9a09f1c9be431cbad3cdcf0d898b457343737f8df55): complete subsection reference.

- [minutes](resources--cdn_loadbalancer--reference--group-014.md#canonical-ae483960f40f70f6068265d002ef7f51484a3019723a5203c0caded67100787f): complete subsection reference.

- [seconds](resources--cdn_loadbalancer--reference--group-014.md#canonical-789befd9994d2c8980b4a6a66923fa7d188bc8c040c17f6825b9d52ab5ee1ec2): complete subsection reference.

<a id="canonical-e521e8ac6a5e9870bda6227c1bac443afe9ab2b46e95a4695618e5d4c073b898"></a>

## Next pages — rate_limit.rate_limiter.action_block / bd35342eebd7 / 4

- [rate_limit.rate_limiter.action_block.hours](resources--cdn_loadbalancer--reference--group-014.md#canonical-01a1d8044fbe72fbffc3f9a09f1c9be431cbad3cdcf0d898b457343737f8df55)
- [rate_limit.rate_limiter.action_block.minutes](resources--cdn_loadbalancer--reference--group-014.md#canonical-ae483960f40f70f6068265d002ef7f51484a3019723a5203c0caded67100787f)
- [rate_limit.rate_limiter.action_block.seconds](resources--cdn_loadbalancer--reference--group-014.md#canonical-789befd9994d2c8980b4a6a66923fa7d188bc8c040c17f6825b9d52ab5ee1ec2)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-01a1d8044fbe72fbffc3f9a09f1c9be431cbad3cdcf0d898b457343737f8df55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc4d5b91debc1cfdaace7cacdf8d14c5021eb648607db97d7755aa7fd54735d0"></a>

## rate_limit.rate_limiter.action_block.hours — rate_limit.rate_limiter.action_block.hours / f94b2204ef0c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-48293ce3109da5bdca9b9622160c1aabcf817f8bddc371db03a5f6f5ab6e96ff"></a>

Type: `"object"`. single nested block, Optional.

Hours. Input Duration Hours.

Upstream description:

Input Duration Hours.

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
hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-a74dd9636344a2158292bc3658db59dc3602f8de32d2a4c795e032f052f5b6d7"></a>

## Direct properties — rate_limit.rate_limiter.action_block.hours / f94b2204ef0c / 3

<a id="canonical-de99e7e04e83beca52de0f7bbf0a8aa295dedeec3798893b2f438095fc0b5958"></a>

<a id="canonical-bb67b6cc029b7d1769e03d355950767291872ec6274934cfcb0b7987a9d2b3a0"></a>

## duration property — rate_limit.rate_limiter.action_block.hours / f94b2204ef0c / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 48),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-bdab6c0bc7ea0d9c9103f3966d88eab1e5ebed8c75d18f56a782bd5c0b8d6b24"></a>

## Next pages — rate_limit.rate_limiter.action_block.hours / f94b2204ef0c / 5

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ae483960f40f70f6068265d002ef7f51484a3019723a5203c0caded67100787f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f80531079bd8023e291cc377167dc56a699bf6e9d1d0d6c48da2e8c87919de7b"></a>

## rate_limit.rate_limiter.action_block.minutes — rate_limit.rate_limiter.action_block.minutes / 8534a734361a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-6ecd939eb26937713e813d141967dfe7148cc6fc89d2327bf9379137e21f2284"></a>

Type: `"object"`. single nested block, Optional.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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
minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-e103c2320aafba35b43ae55a8e4e904768ea13419fe7f201d9008f0adafe0622"></a>

## Direct properties — rate_limit.rate_limiter.action_block.minutes / 8534a734361a / 3

<a id="canonical-05f1435a95c39b31840f01e8b91e30cb6f80ec4c140d802e2fa161b30b7d4b6f"></a>

<a id="canonical-436bb58e14ebb61d4aac2413442c6aa0555d06aca0e869e1980a8a4418de5be3"></a>

## duration property — rate_limit.rate_limiter.action_block.minutes / 8534a734361a / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-301119603f1c16d25d26d5995f92fe0a57a08d3fd4edc336140a5fd09d111b7c"></a>

## Next pages — rate_limit.rate_limiter.action_block.minutes / 8534a734361a / 5

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-789befd9994d2c8980b4a6a66923fa7d188bc8c040c17f6825b9d52ab5ee1ec2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d60841450d97f1390ec0387d6006a81ca7f0afbd079935c3ad9e0fb8b6a17d90"></a>

## rate_limit.rate_limiter.action_block.seconds — rate_limit.rate_limiter.action_block.seconds / fe957c2d9bdb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-05f7c00f8583e68adc6af69ad275fb694583f04a1900c49b4d257b1ada6ad27f"></a>

Type: `"object"`. single nested block, Optional.

Seconds. Input Duration Seconds.

Upstream description:

Input Duration Seconds.

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
seconds {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe97300f1f99943571882887dcada033b05eb8e611d7d541ffd1ab366aa9621a"></a>

## Direct properties — rate_limit.rate_limiter.action_block.seconds / fe957c2d9bdb / 3

<a id="canonical-6ea21434ea3d1842573dab01cb0e393ec10389345649040146ecf89b3bcd8ade"></a>

<a id="canonical-641928f7bb446928be1e9c63e6b33f8872dd483887ea4c694549936364d7d33c"></a>

## duration property — rate_limit.rate_limiter.action_block.seconds / fe957c2d9bdb / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 300),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-d468a10b30d3501862197bd8d05ea16301fd2ceb0fe78057a15790484531b092"></a>

## Next pages — rate_limit.rate_limiter.action_block.seconds / fe957c2d9bdb / 5

- [rate_limit.rate_limiter.action_block](resources--cdn_loadbalancer--reference--group-014.md#canonical-99e0168a4094af046e77ae3f8e2ec6726ebf574f5cb6a769232f3b2fe81cae1d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-195c3959c0bed9aa24eb4fdd6ea585123bf59051f0e0947ce5429464f3822984"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86497c66e2939c039a7e78cc1f2f98a4ee302f992fbf2689091f89b83cef459e"></a>

## rate_limit.rate_limiter.disabled — rate_limit.rate_limiter.disabled / ac91640b143e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- rate_limit.rate_limiter.disabled

<a id="canonical-59d9074dab1d8033c669616dc029e75afbe695de269366f825569221e2b25ebc"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disabled = {}
```

<a id="canonical-51ac68302f0ae373690e5c477737a2061f0d28f88cdffcda4890506c3555ac0a"></a>

## Direct properties — rate_limit.rate_limiter.disabled / ac91640b143e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-215f4c6cc6607018c9e89037f98f24ff16d941db72dc143e63beb493d371dce3"></a>

## Next pages — rate_limit.rate_limiter.disabled / ac91640b143e / 4

- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-99c721be2080364ec24b7b2f29c27d97fb5571fbb9c77c4c611abacb9b118e31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7da2360218088ca9b66977a695bca377af3009aa93767c2f642ba74e17a2cf26"></a>

## rate_limit.rate_limiter.leaky_bucket — rate_limit.rate_limiter.leaky_bucket / ead732c403b7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-41af4944a5c7f3f91ac383e9b1114d51c40ee6502e9c7cc4a04f018c657f60cb"></a>

Type: `["object", {}]`. Optional.

Leaky-Bucket is the default rate limiter algorithm for F5.

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
leaky_bucket = {}
```

<a id="canonical-fdef7e3b71923f532505ce0d767bf9236dfe2ea462c3c44548bd9b9da8e24087"></a>

## Direct properties — rate_limit.rate_limiter.leaky_bucket / ead732c403b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b76e508346659f0ea28e33dd9f67b20d6e63b903795d8bed3e22065dc607f91"></a>

## Next pages — rate_limit.rate_limiter.leaky_bucket / ead732c403b7 / 4

- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e520a1225186a952c49f5a419213ce1a18c41f913ee1215a398196b7ce0d4200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33be15355231c8f4cce1cc75a98b72d1a100d3dafcf7c2032a2de395499eb416"></a>

## rate_limit.rate_limiter.token_bucket — rate_limit.rate_limiter.token_bucket / 267a30118920 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-254c9fa67a23237cc562ba87c823ba41fb73c9e987ed8ee487963682ba727957)
- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-86b601b9b113948a5a945f826d8aa756dc88cad3097e122f6cdaae4b9e1b629a"></a>

Type: `["object", {}]`. Optional.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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
token_bucket = {}
```

<a id="canonical-a93026c0e3fc88ba99a11436613b0e745ccd309853c38990e6570317c6344efe"></a>

## Direct properties — rate_limit.rate_limiter.token_bucket / 267a30118920 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72ca6bef41ff475406970ed3eaa503bbce74865f56dc72ba65082216780450e9"></a>

## Next pages — rate_limit.rate_limiter.token_bucket / 267a30118920 / 4

- [rate_limit.rate_limiter](resources--cdn_loadbalancer--reference--group-014.md#canonical-a9c43fb4dbcad0b8e924aa6db27987a0f3104a20968159cb3ad72ceef25c15e5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d1c756684a9b6aed13b1e1004dcd60244387d594a41653bbd78e71f0554cae31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0eed7e14b7b2a7a6d832838068a04284b302896a564aa49c764e0682cd010eaf"></a>

## sensitive_data_policy — sensitive_data_policy / af4e8437a80b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- sensitive_data_policy

<a id="canonical-70219ffc73a735b00bea7c5242cebe9a9c3627417eb99b52705567f78aa04aec"></a>

Type: `"object"`. single nested block, Optional.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

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
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-501d64e41e02672ca00da47f3b59c104a28b90f34a6873a1776fe159454aeb09"></a>

## Direct properties — sensitive_data_policy / af4e8437a80b / 3

- [sensitive_data_policy_ref](resources--cdn_loadbalancer--reference--group-014.md#canonical-2109283742f1e745de786fb7f0a323c03e3baffb67105b5e0cf648c0edaf4106): complete subsection reference.

<a id="canonical-74044d7d81dd7cc1858b76c28bcde1066bbb389a3ae5c4c21af050e129b04c84"></a>

## Next pages — sensitive_data_policy / af4e8437a80b / 4

- [sensitive_data_policy.sensitive_data_policy_ref](resources--cdn_loadbalancer--reference--group-014.md#canonical-2109283742f1e745de786fb7f0a323c03e3baffb67105b5e0cf648c0edaf4106)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2109283742f1e745de786fb7f0a323c03e3baffb67105b5e0cf648c0edaf4106"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a32b7c18a4f23dc6fb190990c22aea8d6c78b945c540b89b38bb9fc0866d642"></a>

## sensitive_data_policy.sensitive_data_policy_ref — sensitive_data_policy.sensitive_data_policy_ref / 332393d4ce1f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-d1c756684a9b6aed13b1e1004dcd60244387d594a41653bbd78e71f0554cae31)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-e442294a7e3112d93012db01626a7fc16e3d36a5f3eb466c963d69925e83f360"></a>

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
sensitive_data_policy_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-960243a274066cfa6f7c2ddf7c1df5944aed03302d6233fb9618a9c705451c60"></a>

## Direct properties — sensitive_data_policy.sensitive_data_policy_ref / 332393d4ce1f / 3

<a id="canonical-ecc12102215173780a2265cc19a59e839fb67f99415efaafb12d0edd81145854"></a>

<a id="canonical-e0381c64ee2d30dae5218dd18a7711f87fc45a8749e3015672aa9cee04dcce49"></a>

## name property — sensitive_data_policy.sensitive_data_policy_ref / 332393d4ce1f / 4

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

<a id="canonical-26742219457a3e1e409bcb197454bbd2d5cc9762da98ade04946a38d079383b1"></a>

<a id="canonical-b04a23e35e7c1bb2fe4cfc4df8a615491bf906d496a47eee54c577ab59f000f9"></a>

## namespace property — sensitive_data_policy.sensitive_data_policy_ref / 332393d4ce1f / 5

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

<a id="canonical-3ceb6fba1fe5c19ae7406de9e69ca0e82b510c3fd989a8bc03307c313e4bb180"></a>

<a id="canonical-b68ddb9ce046c71be66abc5397045a7347c9512b1025dae5065318dcae3611a1"></a>

## tenant property — sensitive_data_policy.sensitive_data_policy_ref / 332393d4ce1f / 6

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

<a id="canonical-230900e185c6be6351ebe8f82841604813169acd192002f2913aa7b5a195fd31"></a>

## Next pages — sensitive_data_policy.sensitive_data_policy_ref / 332393d4ce1f / 7

- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-d1c756684a9b6aed13b1e1004dcd60244387d594a41653bbd78e71f0554cae31)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2e87a80d3c690f75e98f2c8ac9c5a16eec12e21cfe3135d22b5ec68445fbc25f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5224bb63ba1dece0ff196fc4741ae23be47506754ef0bd32c53e9e65a04443ef"></a>

## service_policies_from_namespace — service_policies_from_namespace / 894acb4accf5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- service_policies_from_namespace

<a id="canonical-26b6e7b1e65f323aea798b0c5bde785d5c2cbb23a3ce426c939579d8b9288467"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
service_policies_from_namespace = {}
```

<a id="canonical-666ef224587bc83e37f2082bd2400b8ef98bb1d87f8c476d83ffdb6d36622e6f"></a>

## Direct properties — service_policies_from_namespace / 894acb4accf5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3c9fba6940d0f06867e5402855f4ce769c1b2bd9c070536b8423091d2e269040"></a>

## Next pages — service_policies_from_namespace / 894acb4accf5 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6f0c88a559028359020b9885e0a124c2f6941cc75507bfe4ee947054f22f2ed6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a78a60d89c07d1993955ff9a8a8e351159f815daf0bbd98030742547ffbf07f"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / b4713be3f613 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- slow_ddos_mitigation

<a id="canonical-16a70faf79f4d43b89434145ee60dbb1ffbcc06b89da151961834b78c1a4cd1c"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
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
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-16a70faf79f4d43b89434145ee60dbb1ffbcc06b89da151961834b78c1a4cd1c)
- [system_default_timeouts](resources--cdn_loadbalancer--reference--group-014.md#canonical-83a3e747a5ff4b770c0a1cc952833a0bc55b58fa94734ead9638ea979faa2346)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-fbf3d0d923dc875eefeadca721f4e7a7e969dabe99e75cc88b76cf7097f9861c"></a>

## Direct properties — slow_ddos_mitigation / b4713be3f613 / 3

- [disable_request_timeout](resources--cdn_loadbalancer--reference--group-014.md#canonical-e28fec83833668ee65fa55bdf942bc801b242d1d6de9c74212ac10ef2baff497): complete subsection reference.

<a id="canonical-ba6ce42c66b349528fa346139a4023e64ba7f020c21c0c453a657bbc4260901a"></a>

<a id="canonical-1a868d6cc20a5acfc2e993800a733ec363627a175fed93bbd4936190b9aa62f2"></a>

## request_headers_timeout property — slow_ddos_mitigation / b4713be3f613 / 4

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-f92812aaa669521ce4531fbf934e9b03f892509fe2d643310f124977833119a2"></a>

<a id="canonical-0a121051372bca0789624d81a44cecd952aba23ff72e06ea648b7bd8ed7ec136"></a>

## request_timeout property — slow_ddos_mitigation / b4713be3f613 / 5

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-f03cf36073a6524a35400d0580129ce0fa0a96696718d8490c931e519eab542d"></a>

## Next pages — slow_ddos_mitigation / b4713be3f613 / 6

- [slow_ddos_mitigation.disable_request_timeout](resources--cdn_loadbalancer--reference--group-014.md#canonical-e28fec83833668ee65fa55bdf942bc801b242d1d6de9c74212ac10ef2baff497)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e28fec83833668ee65fa55bdf942bc801b242d1d6de9c74212ac10ef2baff497"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dfea9727b8d9ae3ab21427498459b37216fb744043a3007b06c2f03f38cfca5"></a>

## slow_ddos_mitigation.disable_request_timeout — slow_ddos_mitigation.disable_request_timeout / 8190a0f81cc3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-6f0c88a559028359020b9885e0a124c2f6941cc75507bfe4ee947054f22f2ed6)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-fe3816f44111b95e743912eb4b0fbdfe3c39b4324ba1f262b56fc130965f8776"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

<a id="canonical-7de12aef338b0faabbca0ae1c8794a7a1418d8a097d45b9bc40b67dbeff67846"></a>

## Direct properties — slow_ddos_mitigation.disable_request_timeout / 8190a0f81cc3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c00dc6e2243124558638d0af5ee5ee593ea5e9f885c2018d5e7089c527f6f905"></a>

## Next pages — slow_ddos_mitigation.disable_request_timeout / 8190a0f81cc3 / 4

- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-6f0c88a559028359020b9885e0a124c2f6941cc75507bfe4ee947054f22f2ed6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-72cb0a186e7fd30aa2b9bf4141b26dff8238c3cd30698e990babc911234f1e00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1359c804f2ae3ef28141e1f8e25e595eeb33936e32bfd9348a55d0ecd8e12a62"></a>

## system_default_timeouts — system_default_timeouts / 8e8a201b4762 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- system_default_timeouts

<a id="canonical-83a3e747a5ff4b770c0a1cc952833a0bc55b58fa94734ead9638ea979faa2346"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system default timeouts.

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
system_default_timeouts = {}
```

<a id="canonical-7d9c73796d81949a823c428a661593d75286293cda376c3abb8465d414eabc8f"></a>

## Direct properties — system_default_timeouts / 8e8a201b4762 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1265e0abaf31069e430b99d745676d227120971770fc5c5be190f2537ec8a143"></a>

## Next pages — system_default_timeouts / 8e8a201b4762 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-58e15a9529dae4f990093f1f19e485d5d6a8e14c3eee8aa8361abfecf86ca89f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cffde5a1036b8eee795873c25d6182826d479ce5882ac74b7774d678583a8866"></a>

## timeouts — timeouts / 186c411ad402 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- timeouts

<a id="canonical-8ba9426b248afbd179510ac858dc599b2ba1885647d51dc9c2bc7f3e4a104d31"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-29a96276d824ef0b01ae0893b518865f768b2339ff6aed07590a1b9e5c1a6615"></a>

## Direct properties — timeouts / 186c411ad402 / 3

<a id="canonical-6213098928a0d341541202be159886d1748329ee3770b737e97f79f842edca1b"></a>

<a id="canonical-ed56b431271e4b5d01cfb1e8674dc63e45f93be2c87fb1aa68af829d979ef753"></a>

## create property — timeouts / 186c411ad402 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e039e875b266c26a2a6f0742491cad7488acd9b628d2c68b6cb1bec582c6f47a"></a>

<a id="canonical-9083361ca09510f9d2dbd0738d82f25e75eb1b7ff6be8b92530f56e8fed0097c"></a>

## delete property — timeouts / 186c411ad402 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-efa08b30937649b1b7b81af19ffd0542ce567473857254cc87d78d78e9582c4c"></a>

<a id="canonical-07b23d1cbf371e33a947a67846c8f3bcf0d7d6e9797921bdeff2d16a7be52cbc"></a>

## read property — timeouts / 186c411ad402 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-09710611b27974173ca3752010b7ef976b4d1b033466f91396a5a7a5ecaddc24"></a>

<a id="canonical-e8635ea5a333cf2be33edded206479ad198a02702bcfe74cd6c9ef9d2a2a3f89"></a>

## update property — timeouts / 186c411ad402 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-73f74b3162c34b0e7267c9ca935be0cb4c2d8be1dd9faf3b0d8bec6158a11e6e"></a>

## Next pages — timeouts / 186c411ad402 / 8

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3e7f53cd820614599c724ec84ff579f4bf6c723d08bb102ea0b0737f3e6eceb"></a>

## trusted_clients — trusted_clients / 2e6b9478c242 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- trusted_clients

<a id="canonical-3d1d4667538f85cbdfa25b61bf891a13d0506344c8545e2e178229697abb8e12"></a>

Type: `"object"`. list nested block, Optional.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("actions"),
  validators.ConflictingListObjectAttributes("as_number",
    "http_header"),
  validators.ConflictingListObjectAttributes("as_number",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "skip_processing"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ipv6_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("skip_processing",
    "waf_skip_processing")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-809a40d350d2fceb94582caed49b63be20d02b83a087ff870351dbfa1436abc4"></a>

## Direct properties — trusted_clients / 2e6b9478c242 / 3

<a id="canonical-d6800b9cc3690ae31785cad03010a20025158245de41bcd8546da1913a224bb6"></a>

<a id="canonical-f808b859e327abdceba395885abd2f8cd40f7ecc58511092ea10c04f53efefee"></a>

## actions property — trusted_clients / 2e6b9478c242 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-268237a7c1f5a0c85063564585ad269e3204bff7a801a688805e713f154673b3"></a>

<a id="canonical-bd55c3368674a4fed2953066d6e3264d02316ea8437fa4a97c8b4985ddcb83e9"></a>

## as_number property — trusted_clients / 2e6b9478c242 / 5

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 401308),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-8b6b600cc3d689b8b17395540c2964fb7ad194694ad61bcfbff31adf26d46801): complete subsection reference.

<a id="canonical-47f8ecf96b26b7806c73a39d2ae76ae8adf9781dbb8e6c226a69810b221bde08"></a>

<a id="canonical-3770fb5085ad189981a8b7f16dc8d108a64745ecf3c7dbe5f255f6a3d247a0a2"></a>

## expiration_timestamp property — trusted_clients / 2e6b9478c242 / 6

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--cdn_loadbalancer--reference--group-014.md#canonical-61a72c6a44e0cfd917add8b6bc5dc543c617f8b411fc08ead50d09d0c7d037f5): complete subsection reference.

<a id="canonical-ff03538f1d227ae918eca7229d7715269568878387e7961b18e0f5af76769495"></a>

<a id="canonical-2cef25fb12b9f98f1184a92f312311461185cfd0f45eac9c07bd2178c39e86f6"></a>

## ip_prefix property — trusted_clients / 2e6b9478c242 / 7

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3043124ed18e1b725e4b3c615f549af297bdf3f4385b22e2d3487e90a1400b64"></a>

<a id="canonical-178272529fd82a637bf98af3bb24f47a1da50c39783b92d7294588081de6fff7"></a>

## ipv6_prefix property — trusted_clients / 2e6b9478c242 / 8

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-014.md#canonical-f9fcdc0b12591f30ad2ca52a032c51c22a6c7439a8b2a7bf6d7fc3769fed7d90): complete subsection reference.

- [skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-5ff6b870c41310484a064c1fd42818daf4fd1f0df14caba5f5bb5a2dbb2b8717): complete subsection reference.

<a id="canonical-9b88b74cc4d56a33fc587ea37bc8c38c7db35187745d93dc3510d3e50ce242c9"></a>

<a id="canonical-5a2a7a585e962528f48858d76ee8ab9de7fbb93259eeece3573460012dec66b8"></a>

## user_identifier property — trusted_clients / 2e6b9478c242 / 9

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-0f9ae23f1e9fe74d866636337f03d398627580445bc518419c58a4d1761a9340): complete subsection reference.

<a id="canonical-00f3dabc0bd5519be150a8b89c9fa34a8e235a3192e0709dba7b4d29d9c197a9"></a>

## Next pages — trusted_clients / 2e6b9478c242 / 10

- [trusted_clients.bot_skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-8b6b600cc3d689b8b17395540c2964fb7ad194694ad61bcfbff31adf26d46801)
- [trusted_clients.http_header](resources--cdn_loadbalancer--reference--group-014.md#canonical-61a72c6a44e0cfd917add8b6bc5dc543c617f8b411fc08ead50d09d0c7d037f5)
- [trusted_clients.metadata](resources--cdn_loadbalancer--reference--group-014.md#canonical-f9fcdc0b12591f30ad2ca52a032c51c22a6c7439a8b2a7bf6d7fc3769fed7d90)
- [trusted_clients.skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-5ff6b870c41310484a064c1fd42818daf4fd1f0df14caba5f5bb5a2dbb2b8717)
- [trusted_clients.waf_skip_processing](resources--cdn_loadbalancer--reference--group-014.md#canonical-0f9ae23f1e9fe74d866636337f03d398627580445bc518419c58a4d1761a9340)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8b6b600cc3d689b8b17395540c2964fb7ad194694ad61bcfbff31adf26d46801"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fdb7742106b1cf273e7088262fdee531510599837553dd9db64c54c241e4ec6"></a>

## trusted_clients.bot_skip_processing — trusted_clients.bot_skip_processing / e32eb7241763 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- trusted_clients.bot_skip_processing

<a id="canonical-4a18a94464676e854bb0d8e0aa929a0e1874f87f04d59ad706732a7c45380759"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
bot_skip_processing = {}
```

<a id="canonical-54659a4548a614658fd7c5d47389821e3afaae55655be952e32c281f7dbe058a"></a>

## Direct properties — trusted_clients.bot_skip_processing / e32eb7241763 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97f305002a9e8710f53fb8cf9f384a409e4e379f9bcc4a4dc430b3a66e990c59"></a>

## Next pages — trusted_clients.bot_skip_processing / e32eb7241763 / 4

- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-61a72c6a44e0cfd917add8b6bc5dc543c617f8b411fc08ead50d09d0c7d037f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b024777f8beb79036f1752d7a86609cc933477c6c578395a75ff7a8bb8f0640e"></a>

## trusted_clients.http_header — trusted_clients.http_header / 28d9c75e150e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- trusted_clients.http_header

<a id="canonical-3a6d88e910c3214959d4bb9bf868676f3b0d84f6141f6c1e377c9b49b6d5e32c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-f2e5645df6f1c5ba501bf5a1e9cb10ee99fcee92e29271c75db2e180279f0d32"></a>

## Direct properties — trusted_clients.http_header / 28d9c75e150e / 3

- [headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-93d07d7aadc5bc5e619a98cedd7c2dfbaeabae8cbd22d7c63474a2c79d25520f): complete subsection reference.

<a id="canonical-22f0f3c3e144ed0f9431acfaea1899041f0e0aeecfef89a8b48929d403fe82b8"></a>

## Next pages — trusted_clients.http_header / 28d9c75e150e / 4

- [trusted_clients.http_header.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-93d07d7aadc5bc5e619a98cedd7c2dfbaeabae8cbd22d7c63474a2c79d25520f)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-93d07d7aadc5bc5e619a98cedd7c2dfbaeabae8cbd22d7c63474a2c79d25520f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cad8f275a8e55ed1048d1011f1c6ff488eadda76a4c0679d349f901d0c45b59"></a>

## trusted_clients.http_header.headers — trusted_clients.http_header.headers / 6d9f766f2001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- [trusted_clients.http_header](resources--cdn_loadbalancer--reference--group-014.md#canonical-61a72c6a44e0cfd917add8b6bc5dc543c617f8b411fc08ead50d09d0c7d037f5)
- trusted_clients.http_header.headers

<a id="canonical-80e99b0c0aa09656a1bd9a4e0d3b52fc0f3c77333c8a2fe747d2fc92ac0466cf"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-31f1a87b61cfc407b1c88b4eeaffd36aee4312fc4d5d78f47c32eafb8b62652f"></a>

## Direct properties — trusted_clients.http_header.headers / 6d9f766f2001 / 3

<a id="canonical-e84fe314670766f7e4b36e69621b3388f2ca13e4ceef02f3f344c826c4f98c5c"></a>

<a id="canonical-929ed46ea1d60ceb829bb1b98687c5cc177b92899b078ca6f42aefc283595ec3"></a>

## exact property — trusted_clients.http_header.headers / 6d9f766f2001 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-f7a017df59719e5d0f38dba014899ace6ae51d351e77005b34bd35c23df911f0"></a>

<a id="canonical-d8bbb80b41d6c4604b479b3f8ea6d4a8893de77002460fe12235c55b09a0d049"></a>

## invert_match property — trusted_clients.http_header.headers / 6d9f766f2001 / 5

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-a91257fb3b9674955dd1acd60b01835f073416a84c9bcb050c4d9b00a39db930"></a>

<a id="canonical-99d0bb04c662a395fc5bb90533a1d948ade7a7430e9ea8e0dfb41c976da07706"></a>

## name property — trusted_clients.http_header.headers / 6d9f766f2001 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-8d3e3bf7f6dca6e85507d4f85faa487fedc2ca5fe7e748d870821bf8b3b08547"></a>

<a id="canonical-030b1c5199cc461b5dc98724e31319f8af5b8e857a1f6f3e09c2612d21c800a0"></a>

## presence property — trusted_clients.http_header.headers / 6d9f766f2001 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-c423d83d79d0081a0eb2155811ad054bbc6608c2e42d58a92f52ea4c4b5dd4d5"></a>

<a id="canonical-2e40283de58401d90c92fdb7b773fe925cd1e5b7e4025a41f683dc66a65270bb"></a>

## regex property — trusted_clients.http_header.headers / 6d9f766f2001 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-88e97526e48e54cdd39b2abcdca5524f152c71e434bc0c3be86000847eefe936"></a>

## Next pages — trusted_clients.http_header.headers / 6d9f766f2001 / 9

- [trusted_clients.http_header](resources--cdn_loadbalancer--reference--group-014.md#canonical-61a72c6a44e0cfd917add8b6bc5dc543c617f8b411fc08ead50d09d0c7d037f5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f9fcdc0b12591f30ad2ca52a032c51c22a6c7439a8b2a7bf6d7fc3769fed7d90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83eef51a91d76f50133d9ef4919ed56ff0281a00fdd5f28772b4f0e4d03df5c3"></a>

## trusted_clients.metadata — trusted_clients.metadata / 505e4a1c5124 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- trusted_clients.metadata

<a id="canonical-40c55c4076e2b416e0c25aaeab8e21ef5c70d733294c8b2e7470a238107b8735"></a>

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

<a id="canonical-30dd602316c62d571b1d38e39a3a663171a449bd210f03ee8ade73cf8957a948"></a>

## Direct properties — trusted_clients.metadata / 505e4a1c5124 / 3

<a id="canonical-11a48c04be3030795e9493a7ebe55c920ecf07168ecd33e7a6346ddf1103c625"></a>

<a id="canonical-586961dab6c3992e45b459902ee531be87a49ce2ddb97238d3d89741c71ef94a"></a>

## description_spec property — trusted_clients.metadata / 505e4a1c5124 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-a69ec283f08aeb8b387401135e0b4add1ec11cb4d1881842f92d33baa46a68b7"></a>

<a id="canonical-b26de9b11c373a462a57750c6f20b97c80c98257f1be39f9a78b37d92f50e301"></a>

## name property — trusted_clients.metadata / 505e4a1c5124 / 5

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

<a id="canonical-19985f3eed63d0c88d21edef224f5a1307cc535afa92e5e7fd3a52a9acac462a"></a>

## Next pages — trusted_clients.metadata / 505e4a1c5124 / 6

- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5ff6b870c41310484a064c1fd42818daf4fd1f0df14caba5f5bb5a2dbb2b8717"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89d6a2f871dfbcf15b1a07af5fd13185484faf4a310a831c4b2a3fda8c91f712"></a>

## trusted_clients.skip_processing — trusted_clients.skip_processing / ea74e194c18c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- trusted_clients.skip_processing

<a id="canonical-c00f08ec70d667f35250b09dba099a2f401a22fe5de93a6a014736dd0392e2c5"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
skip_processing = {}
```

<a id="canonical-be06daf2a4ff6b8026123d1ebda178382a08a94aad853a8e2a079d1b8923f07c"></a>

## Direct properties — trusted_clients.skip_processing / ea74e194c18c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f19ac7e0c9ab145e0a8178261af5989afa903b11cc911176dc6a6d78406647ae"></a>

## Next pages — trusted_clients.skip_processing / ea74e194c18c / 4

- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0f9ae23f1e9fe74d866636337f03d398627580445bc518419c58a4d1761a9340"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4232e3b8d1c37a4a385063fb9e85050e88a9c2ecff71a1db1e9e0ba6ad6bf1d1"></a>

## trusted_clients.waf_skip_processing — trusted_clients.waf_skip_processing / e5d99043ce5e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- trusted_clients.waf_skip_processing

<a id="canonical-5ad8978180964139ba469a170be2931cbf07761ae6dec627c7674601c3d7e92e"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
waf_skip_processing = {}
```

<a id="canonical-0797949eae36ebe7e2b48904085a5890ea38d7e3cd8c0240fc0f9857de5a6c58"></a>

## Direct properties — trusted_clients.waf_skip_processing / e5d99043ce5e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ce486b62dadeffffda90977036b9e432d901e116f5d40087079ab5be3f38071"></a>

## Next pages — trusted_clients.waf_skip_processing / e5d99043ce5e / 4

- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-fa607075ce0e765101a422c4bb6d254a483e196f2d69d9529aeeeb3395be1872)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f2820c8e4f1688393af0fa5b74f4a574dcfd9a57773098a177d5ad9db86f162d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c470416a4b360d7a36d64b6d891d14925e348e5290376d8d741d1e1d131365bb"></a>

## user_id_client_ip — user_id_client_ip / 9cf8cadaa0c6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- user_id_client_ip

<a id="canonical-76f539ecd4d2f17ca45e9c254fbdb1d21ad673cdb02ed2e208c641fafd4df240"></a>

Type: `["object", {}]`. Optional.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option

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

OneOf alternatives in this subsection:

- [user_id_client_ip](resources--cdn_loadbalancer--reference--group-014.md#canonical-76f539ecd4d2f17ca45e9c254fbdb1d21ad673cdb02ed2e208c641fafd4df240)
- [user_identification](resources--cdn_loadbalancer--reference--group-014.md#canonical-960b48764ebb9c1d3064820ade4c1f868525345397c85ec5334304298ed52cb9)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
user_id_client_ip = {}
```

<a id="canonical-b3dc9353d21f36e453c6f2acdf1ff13523252b0c41e8c66445c4b54fe28d4006"></a>

## Direct properties — user_id_client_ip / 9cf8cadaa0c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b6e17a9506fc94e343f3d3145ce293df6e1e3d6d49ded1279b2b4d620135ce7"></a>

## Next pages — user_id_client_ip / 9cf8cadaa0c6 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-56b11c6aea01027457a9e3e86571f28a84dccc0c36e4c80325f3f4976c5aec1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24f878589945733f5c71a2ffdb8dc49273e1c079a305a2a2bfc8e4166124a111"></a>

## user_identification — user_identification / 418e473e77ff / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- user_identification

<a id="canonical-960b48764ebb9c1d3064820ade4c1f868525345397c85ec5334304298ed52cb9"></a>

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
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-e3e8d210cb0bfce4bf7336b0be5a5ad1f5e349deff52a90c1c86c00aa1681f33"></a>

## Direct properties — user_identification / 418e473e77ff / 3

<a id="canonical-2a07c8a771cd140bc45bda897cd1a28c379d64fc26a257ad00450becdaaabdf3"></a>

<a id="canonical-f8580160295bba03cfe29d077d7c88eaf061af83a6bef8e4fdb5573ffe1a151e"></a>

## name property — user_identification / 418e473e77ff / 4

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

<a id="canonical-47566891c344dc565285d672b82df79db4933d8d98b330a55a15f1b0271b61d7"></a>

<a id="canonical-a26043e399e8149f94ee531e980c3cedc9af369ed968bc7674cef9126828b440"></a>

## namespace property — user_identification / 418e473e77ff / 5

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

<a id="canonical-c983d34c9513aa9165d30c4141078be72150f69d57fc3cc6b85de73288b37f11"></a>

<a id="canonical-0627a71ed3710057e34134837329f1eb73bf614f6f50f84588426e6c4f518b23"></a>

## tenant property — user_identification / 418e473e77ff / 6

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

<a id="canonical-3c3d902adf0ae0ee4d85159c59431daf5cac3efbc8f62f7ec2cabcad933e0ef6"></a>

## Next pages — user_identification / 418e473e77ff / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33e4860b61275a0e6f48dd59aaf3993b64cf9cdad6827536406f381dea5d1b6c"></a>

## waf_exclusion — waf_exclusion / 7addf6d3edd0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- waf_exclusion

<a id="canonical-df6594fe5f1c8ba2dc26e1638d93504b6316752cd7ca5dfa3966c4bd46df0411"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("waf_exclusion_inline_rules",
    "waf_exclusion_policy")}
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
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

Terraform syntax:

```terraform
waf_exclusion {
  # Configure direct properties listed below.
}
```

<a id="canonical-e40b796ae95a9d72ffb5fe8ea6386b1988613e1635e5135920059da2beec0d16"></a>

## Direct properties — waf_exclusion / 7addf6d3edd0 / 3

- [waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1): complete subsection reference.

- [waf_exclusion_policy](resources--cdn_loadbalancer--reference--group-015.md#canonical-50b08193d79b6da13f556973d429b6fab56e31f3341c8fdb5b11cb0509929c25): complete subsection reference.

<a id="canonical-e28eb3e32a99a6c3fdf78a0929701bda4ed9015f377484cde3671ddd0f6b3afb"></a>

## Next pages — waf_exclusion / 7addf6d3edd0 / 4

- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_policy](resources--cdn_loadbalancer--reference--group-015.md#canonical-50b08193d79b6da13f556973d429b6fab56e31f3341c8fdb5b11cb0509929c25)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f682222d76b910b2a848417b8e92ac51188c143aac1c4402e2735d8717e490e"></a>

## waf_exclusion.waf_exclusion_inline_rules — waf_exclusion.waf_exclusion_inline_rules / e4d8a0145fe2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-66042d80067f36db80ba56b5c248f59694292ab7f45c48d285d089806bacb41c"></a>

Type: `"object"`. single nested block, Optional.

List of WAF exclusion rules that will be applied inline.

Upstream description:

A list of WAF exclusion rules that will be applied inline.

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
waf_exclusion_inline_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232780bfcc08ad7102f47bd4afff601db652b7dd3eb7dd1eb3d3618213d44b1"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules / e4d8a0145fe2 / 3

- [rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67): complete subsection reference.

<a id="canonical-f53cebb4b2f775c1310e4d0c59d07b73b5841a9cf5bc3a4f497e5dd52180112a"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules / e4d8a0145fe2 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2eabec0324ea04bd6464b97253fa7b173e2b679688b20e28b830f2eafd5fb58"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-8347f5bac1af181c0e16ba143f4ece4228b24c3fcfc0308ab7f76eef3e33f6ae"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

An ordered list of WAF Exclusions specific to this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex"),
  validators.ConflictingListObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_prefix",
    "path_regex")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-a74189ef22dd235ef74f7136ed11d6fa70d5706504f6fe92f857714596f4e1a2"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-014.md#canonical-cae3a58aff1963bcf00f73a5bb96d833b52d4f7b0eccf45dd6501ad833d0402b): complete subsection reference.

- [any_path](resources--cdn_loadbalancer--reference--group-014.md#canonical-01b4a2d21402c7dc1af7a1c493dcdadce94ba3ab6c83900c99bdedbd792609c8): complete subsection reference.

- [app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089): complete subsection reference.

<a id="canonical-bbbe5166eee9ee651c0a124eee04f224b40491736c5d8243163cf3b784b008cc"></a>

<a id="canonical-d3ec25bffa243a2239503c035d850eb39ced788b749931ecb7d495af68750a2c"></a>

## exact_value property — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-029268fa1710fc1aa655115cbd360e408f92f7143f67be8bf07474da5e374c11"></a>

<a id="canonical-ed1baf3348c741454d854891c7ea78a227cae4d633da70c7d5d50cb1ae62cfa1"></a>

## expiration_timestamp property — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 5

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [metadata](resources--cdn_loadbalancer--reference--group-015.md#canonical-50e38d2b0b9643db571f1b02e9c371c497c660140cd0df1227922ebbc6588996): complete subsection reference.

<a id="canonical-3efbfadf6b079298c92b4f172ab1baf5ae75a0006ea9869691fc2f1b5c942428"></a>

<a id="canonical-acdb200d93f3aa4bf60fed730a63cea3bf3b30a9a2ba21bb8fa99c697443a389"></a>

## methods property — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 6

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c5fdb95c797e860285a00a722a253c672957deeec33d8f59f75170b48c8cf2b7"></a>

<a id="canonical-f64f1f0e4328d2957434a36d4086d11b2d0d2d4eda5ea661cbf98a8bce979634"></a>

## path_prefix property — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3dda6358250c7cd25b25b1d79e324b03d50bf3acda40665599797e7f7072528b"></a>

<a id="canonical-61375c9d6515ef4c2801888ca1e81cd7ffee851b745dd3a6f247ef8c394fc6da"></a>

## path_regex property — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 8

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-14bcacea70ca607ff9a2a8f2ae8a120c1dc4834854f3a151c84855fbc4ea970a"></a>

<a id="canonical-f7859c2085f6aef163ea2a51cb2e92b7d2244df58115deebad346ed213ae99e7"></a>

## suffix_value property — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 9

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-567c741b4fe63e7566354b879d197d6a904549c20408c951f7283eb252326909): complete subsection reference.

<a id="canonical-e06f0d46b43c2c0fa4f59201541d1f039d34262b0c21e81725455baae588aa7a"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules / afdfc192b7ee / 10

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](resources--cdn_loadbalancer--reference--group-014.md#canonical-cae3a58aff1963bcf00f73a5bb96d833b52d4f7b0eccf45dd6501ad833d0402b)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](resources--cdn_loadbalancer--reference--group-014.md#canonical-01b4a2d21402c7dc1af7a1c493dcdadce94ba3ab6c83900c99bdedbd792609c8)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](resources--cdn_loadbalancer--reference--group-015.md#canonical-50e38d2b0b9643db571f1b02e9c371c497c660140cd0df1227922ebbc6588996)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-567c741b4fe63e7566354b879d197d6a904549c20408c951f7283eb252326909)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-cae3a58aff1963bcf00f73a5bb96d833b52d4f7b0eccf45dd6501ad833d0402b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c790e4f35f58f8209e45cefaaa596ecfbf06565423bc07dd2a56bce57113199"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_domain — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / 90030a038776 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-b15d1c49db14d80d926c70627fe31b370682836a383cfb31805987e5cb4103ab"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
any_domain = {}
```

<a id="canonical-a12cbefe7bdae051652e4b6ab82cd51ddb95e5fc59494ed3d8f8213d9c3b1c05"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / 90030a038776 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5121f6dda84b8adaabd2d0e7e609ca076db45b460e3dd96687029cf3c1f40a45"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / 90030a038776 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-01b4a2d21402c7dc1af7a1c493dcdadce94ba3ab6c83900c99bdedbd792609c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6413e39442ced8e2dab44ddfce995883c4c257797658acad5f095608b50218a0"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_path — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / df265bf0b6b3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-127aa3ec8be90a6dc1f5799fea6cd62d01774d298e64120472e82c8e7875dfba"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
any_path = {}
```

<a id="canonical-f3a704bd7b74c9f27d0630275627a9619e252aa6427ad13c8955d00c963fe293"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / df265bf0b6b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3eaccaf9e36256b5866a42dcd3432c2b6473fc1c54cce86825675a12c3bcfac8"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / df265bf0b6b3 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3e7be7ddae02dba204c4bee72baaff07f7f6386b8ddf0a5a3df7e0dc820058a"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 0650b4cdc70b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-881b85eca1e4292512fbaec076abc7849b4269da121a37e9c674935cae828c27"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-44677d181c66bce4a31b798445d1c918aba62848cdd29cee76dcc1103a40900e"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 0650b4cdc70b / 3

- [exclude_attack_type_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-ff5435cc348a8512fb01760e588d542c8660e41e040726efb67d2de4c060ddca): complete subsection reference.

- [exclude_bot_name_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-9e7885e4ae88850cab4ce24c4c7db32450ce48027322e2e6fe0a6de44b626a59): complete subsection reference.

- [exclude_signature_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-b8aa16c4e8126dc92f659e048bbb4c52ceea5d4e42be1f605606b00c0afdb7d8): complete subsection reference.

- [exclude_violation_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-3bfbcea419fc1039a4f52e72a4cbfddff2060653f43dea8dab9123254aada76f): complete subsection reference.

<a id="canonical-0a95e1ddfa9ac29211d2d944cb7856047f1a603baf864641820735b60150a198"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 0650b4cdc70b / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-ff5435cc348a8512fb01760e588d542c8660e41e040726efb67d2de4c060ddca)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-9e7885e4ae88850cab4ce24c4c7db32450ce48027322e2e6fe0a6de44b626a59)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-b8aa16c4e8126dc92f659e048bbb4c52ceea5d4e42be1f605606b00c0afdb7d8)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](resources--cdn_loadbalancer--reference--group-014.md#canonical-3bfbcea419fc1039a4f52e72a4cbfddff2060653f43dea8dab9123254aada76f)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ff5435cc348a8512fb01760e588d542c8660e41e040726efb67d2de4c060ddca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9310cf5cee5a1cf99c46b51115907cc4d855b5a0edf2ba6281d972b0e4cbd388"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 847ae41fe5bc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-5f05e19e4924484e5c44acc9df929ced0e73c005bcd2d8379367a9ac5da4276a"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-ead451b00b18779d299ffda49e94f3a4b53e45df4ec897288b2c620f50894d61"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 847ae41fe5bc / 3

<a id="canonical-375286efd22aaa68e16dd4963942e9e765c79b32ce7bcc244486e247f5ad0fd9"></a>

<a id="canonical-3da9a200ec8e2df8d1637082d1be655d561564271aa2169a1db63a4e9df6474d"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 847ae41fe5bc / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b218f9e440e9406f262c1fc4f4268ca277249854802fffab379c153791cad8fc"></a>

<a id="canonical-9e0ab9344cefaf6cf007623d719ae21dc26bed8165cfd47eabf19d0ce0f0c102"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 847ae41fe5bc / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-86002dfdf71ab43243175379c3ece3b0ee326a3e942ca6f8dfd670564d05aaf1"></a>

<a id="canonical-be375e10cbfded5da242aec3c7829582edbdf3cde3e520e2eac713098a2512f3"></a>

## exclude_attack_type property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 847ae41fe5bc / 6

Type: `"string"`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3b6b3614abb7ffbf463e3f9f7a5e1976ca018380b4e8e9db6676cfd2615fbef7"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 847ae41fe5bc / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9e7885e4ae88850cab4ce24c4c7db32450ce48027322e2e6fe0a6de44b626a59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-292c38fbf50e9662c8f5b0d7501db895bec580b008e0f726f20668094a5685b4"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ee868f7cc939 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-f6bfce2ef417e6e6cba58816fbd3aceb6f43b4836339881ef4a5572aa1af172b"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-6402f11ae16c32a5d0b34e6c03a6147f58969c0bef210ee7819490bd3a83eda3"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ee868f7cc939 / 3

<a id="canonical-04ab0f37e666b21ee59cf6dfe4a59b81b76335974af51021e19860ec3bbc3480"></a>

<a id="canonical-6dff3fed7df3249e88f9299ef1b6aadc2e0fb9c15be1f00ce1dcd11aef65d25e"></a>

## bot_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ee868f7cc939 / 4

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-8537ff2dbf4a9689c97fda28bdab9c4062f459fe1614764477e652dbee0501a8"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ee868f7cc939 / 5

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b8aa16c4e8126dc92f659e048bbb4c52ceea5d4e42be1f605606b00c0afdb7d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-daa0f666327adab216b2bebf2801af70f68ff130d8005ba243d9cb4934f69a70"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / dcacadb93f81 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-af134a62e96f1e39a76089cdec38a4e054c4cb7fd725c336d819b3956676e156"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-dcb0674c8b55576a2541c79bb1d851689a0eb2b6156317bf915d95945ad64d9d"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / dcacadb93f81 / 3

<a id="canonical-f7322fb9beb787a63ecd22f53227f7fa570a760c4183c8bc9b249be6b5f8e381"></a>

<a id="canonical-0325422366893eef4ea3d1f88c80a86f205b6cab824062e76ff74b62bc900009"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / dcacadb93f81 / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2e0ae26ff012e75317ae6f5d57a17f3e4898afd025e185cbec47ba946cc8a397"></a>

<a id="canonical-5126c0bb454e8d3365f1e5ea086333ae60f76cb0a78618e7f5f838b1148403f0"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / dcacadb93f81 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-37d1d09fa1ea783c504bfbf6007ad48da93f1dc4862f2f994d71e8ce0f5d137a"></a>

<a id="canonical-ca98822b5a9de542cf99fe3f591dfe145fc1001962b818f9f0fb94bf4637a980"></a>

## signature_id property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / dcacadb93f81 / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
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
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-74fc9c13e8be0b900a2ca15a7efa3e26e11d2ea13bec914a9351f2db3718743d"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / dcacadb93f81 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3bfbcea419fc1039a4f52e72a4cbfddff2060653f43dea8dab9123254aada76f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d022c7668fac339714ccd1e4e1ecb22b6262f69c89b1d16a0955c50ba87c865"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / d9f37d291000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-49d1159484a5c47de40732ac1f168236175c7884528ef71566d61749b0a9427a)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-a1780d8b0b6577b162b302816275d641da4059ba688f4171f3acd36916647ce1)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-014.md#canonical-70db062a7f2066a92bc75abad2a72d63f3927b31ef162f451df5aafe213dba67)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-014.md#canonical-bca80965c808c9a365380af588d4c86fe89c65909744ede34953320f25b6f089)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-4b4199bb96565d0d97c037deed2b4128b3cb78f27963143f1da9576b1b79f1d9"></a>

Type: `"object"`. list nested block, Optional.

Violations to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-258af2da6ed72919198333e1eaee3ea4eb5024493b1c6839af2677ed5bad4898"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / d9f37d291000 / 3

<a id="canonical-14f76fd6f7dbfe60cc62f2ae666441661947351ee0f0e590a1c6e9f411e8cedd"></a>
