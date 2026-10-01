---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-48bd28db7d751439b20a3e6797ae465331034c7efa2073c09fe09a66523dcca8"></a>

## rule_list.rules.tls_list.tls_list — rule_list.rules.tls_list.tls_list / 4670370f577a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-39a0f8b5db164c87e09341ae6ccd864e837f7f0cf9e256ada0b7c3e7a50705c3)
- rule_list.rules.tls_list.tls_list

<a id="canonical-38a0d7f71c86d0936d8585817d2e61748feab88e31d617336e9c0fe038ef3522"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0c8ee5277b12ee54ae80014e0bd519bc14118c477095b768d408341a9906b5fe"></a>

## Direct properties — rule_list.rules.tls_list.tls_list / 4670370f577a / 3

<a id="canonical-a122bc266eba651fe3240fca7153413a642a86d9aa642610c8ace4e073b78997"></a>

<a id="canonical-2f409dce73069d5ebe913646984451fa83515cfbcd1ab403364168015860409a"></a>

## exact_value property — rule_list.rules.tls_list.tls_list / 4670370f577a / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-bfdaec6f9be54cf9b0b26e13488f20214083a84d69741f2ea98113ab6a7b346f"></a>

<a id="canonical-39eb07d7c36e8b5d91b8e8502fc61db1558f5213f954ab7b54a5f4c570aa342a"></a>

## regex_value property — rule_list.rules.tls_list.tls_list / 4670370f577a / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-9de97188dfde27ae9aa81f90d399368453e839b4a6e55fe8b51fcc1882aa206a"></a>

<a id="canonical-04c6ec93c1c2acb0d8ddaa25adc1ff47d309a8aaf1feb59f0809d58ed762477d"></a>

## suffix_value property — rule_list.rules.tls_list.tls_list / 4670370f577a / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-120f3ff5e4c871d8d13240b48068ad4fd7afd244cf463979a32f1032389cc71f"></a>

## Next pages — rule_list.rules.tls_list.tls_list / 4670370f577a / 7

- [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-39a0f8b5db164c87e09341ae6ccd864e837f7f0cf9e256ada0b7c3e7a50705c3)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-df60768b4402f4efd096157d5efcfd1041e5802507dd445736e3aef70d2a8cae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a91c1066791ed78a101f12714affbe9bf1e5db32973131107914f124db69f47"></a>

## rule_list.rules.url_category_list — rule_list.rules.url_category_list / 80e4a0d63a4a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-a0abaa83b85e600384d61df268d4542ea7f5cc5e45839a8897aa9bc8bca4af18)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- rule_list.rules.url_category_list

<a id="canonical-519967dc066ccee4814bea113875a753827b12b86f0140d4249eb6c0367221ec"></a>

Type: `"object"`. single nested block, Optional.

URL Category List Type. List of URL categories.

Upstream description:

List of URL categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url_categories")}
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
url_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-6ce95ee16941f2c9719c280d3288994fcbcad2afb9171ee881d3fbf2ece1f26e"></a>

## Direct properties — rule_list.rules.url_category_list / 80e4a0d63a4a / 3

<a id="canonical-a439204415809db48b3272540f93588fa1f033e75052c373d17ad693fccfb50c"></a>

<a id="canonical-6dcb18f3ab1b4c30c12939978ed08caa9ca1b112698e6c3ddff0b72a7808d469"></a>

## url_categories property — rule_list.rules.url_category_list / 80e4a0d63a4a / 4

Type: `["list", "string"]`. Optional.

\[Enum:
UNCATEGORIZED|REAL\_ESTATE|COMPUTER\_AND\_INTERNET\_SECURITY|FINANCIAL\_SERVICES|BUSINESS\_AND\_ECONOMY|COMPUTER\_AND\_INTERNET\_INFO|AUCTIONS|SHOPPING|CULT\_AND\_OCCULT|TRAVEL|ABUSED\_DRUGS|ADULT\_AND\_PORNOGRAPHY|HOME\_AND\_GARDEN|MILITARY|SOCIAL\_NETWORKING|DEAD\_SITES|INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS|TRAINING\_AND\_TOOLS|DATING|SEX\_EDUCATION|RELIGION|ENTERTAINMENT\_AND\_ARTS|PERSONAL\_SITES\_AND\_BLOGS|LEGAL|LOCAL\_INFORMATION|STREAMING\_MEDIA|JOB\_SEARCH|GAMBLING|TRANSLATION|REFERENCE\_AND\_RESEARCH|SHAREWARE\_AND\_FREEWARE|PEER\_TO\_PEER|MARIJUANA|HACKING|GAMES|PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY|WEAPONS|PAY\_TO\_SURF|HUNTING\_AND\_FISHING|SOCIETY|EDUCATIONAL\_INSTITUTIONS|ONLINE\_GREETING\_CARDS|SPORTS|SWIMSUITS\_AND\_INTIMATE\_APPAREL|QUESTIONABLE|KIDS|HATE\_AND\_RACISM|PERSONAL\_STORAGE|VIOLENCE|KEYLOGGERS\_AND\_MONITORING|SEARCH\_ENGINES|INTERNET\_PORTALS|WEB\_ADVERTISEMENTS|CHEATING|GROSS|WEB\_BASED\_EMAIL|MALWARE\_SITES|PHISHING\_AND\_OTHER\_FRAUDS|PROXY\_AVOIDANCE\_AND\_ANONYMIZERS|SPYWARE\_AND\_ADWARE|MUSIC|GOVERNMENT|NUDITY|NEWS\_AND\_MEDIA|ILLEGAL|CONTENT\_DELIVERY\_NETWORKS|INTERNET\_COMMUNICATIONS|BOT\_NETS|ABORTION|HEALTH\_AND\_MEDICINE|CONFIRMED\_SPAM\_SOURCES|SPAM\_URLS|UNCONFIRMED\_SPAM\_SOURCES|OPEN\_HTTP\_PROXIES|DYNAMICALLY\_GENERATED\_CONTENT|PARKED\_DOMAINS|ALCOHOL\_AND\_TOBACCO|PRIVATE\_IP\_ADDRESSES|IMAGE\_AND\_VIDEO\_SEARCH|FASHION\_AND\_BEAUTY|RECREATION\_AND\_HOBBIES|MOTOR\_VEHICLES|WEB\_HOSTING\]
URL Categories. List of URL categories to be selected. Possible values are \`UNCATEGORIZED\`,
\`REAL\_ESTATE\`, \`COMPUTER\_AND\_INTERNET\_SECURITY\`, \`FINANCIAL\_SERVICES\`,
\`BUSINESS\_AND\_ECONOMY\`, \`COMPUTER\_AND\_INTERNET\_INFO\`, \`AUCTIONS\`, \`SHOPPING\`,
\`CULT\_AND\_OCCULT\`, \`TRAVEL\`, \`ABUSED\_DRUGS\`, \`ADULT\_AND\_PORNOGRAPHY\`,
\`HOME\_AND\_GARDEN\`, \`MILITARY\`, \`SOCIAL\_NETWORKING\`, \`DEAD\_SITES\`,
\`INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS\`, \`TRAINING\_AND\_TOOLS\`, \`DATING\`, \`SEX\_EDUCATION\`,
\`RELIGION\`, \`ENTERTAINMENT\_AND\_ARTS\`, \`PERSONAL\_SITES\_AND\_BLOGS\`, \`LEGAL\`,
\`LOCAL\_INFORMATION\`, \`STREAMING\_MEDIA\`, \`JOB\_SEARCH\`, \`GAMBLING\`, \`TRANSLATION\`,
\`REFERENCE\_AND\_RESEARCH\`, \`SHAREWARE\_AND\_FREEWARE\`, \`PEER\_TO\_PEER\`, \`MARIJUANA\`,
\`HACKING\`, \`GAMES\`, \`PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY\`, \`WEAPONS\`, \`PAY\_TO\_SURF\`,
\`HUNTING\_AND\_FISHING\`, \`SOCIETY\`, \`EDUCATIONAL\_INSTITUTIONS\`, \`ONLINE\_GREETING\_CARDS\`,
\`SPORTS\`, \`SWIMSUITS\_AND\_INTIMATE\_APPAREL\`, \`QUESTIONABLE\`, \`KIDS\`,
\`HATE\_AND\_RACISM\`, \`PERSONAL\_STORAGE\`, \`VIOLENCE\`, \`KEYLOGGERS\_AND\_MONITORING\`,
\`SEARCH\_ENGINES\`, \`INTERNET\_PORTALS\`, \`WEB\_ADVERTISEMENTS\`, \`CHEATING\`, \`GROSS\`,
\`WEB\_BASED\_EMAIL\`, \`MALWARE\_SITES\`, \`PHISHING\_AND\_OTHER\_FRAUDS\`,
\`PROXY\_AVOIDANCE\_AND\_ANONYMIZERS\`, \`SPYWARE\_AND\_ADWARE\`, \`MUSIC\`, \`GOVERNMENT\`,
\`NUDITY\`, \`NEWS\_AND\_MEDIA\`, \`ILLEGAL\`, \`CONTENT\_DELIVERY\_NETWORKS\`,
\`INTERNET\_COMMUNICATIONS\`, \`BOT\_NETS\`, \`ABORTION\`, \`HEALTH\_AND\_MEDICINE\`,
\`CONFIRMED\_SPAM\_SOURCES\`, \`SPAM\_URLS\`, \`UNCONFIRMED\_SPAM\_SOURCES\`,
\`OPEN\_HTTP\_PROXIES\`, \`DYNAMICALLY\_GENERATED\_CONTENT\`, \`PARKED\_DOMAINS\`,
\`ALCOHOL\_AND\_TOBACCO\`, \`PRIVATE\_IP\_ADDRESSES\`, \`IMAGE\_AND\_VIDEO\_SEARCH\`,
\`FASHION\_AND\_BEAUTY\`, \`RECREATION\_AND\_HOBBIES\`, \`MOTOR\_VEHICLES\`, \`WEB\_HOSTING\`.
Defaults to \`UNCATEGORIZED\`.

Upstream description:

List of URL categories to be selected.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6fe44e505b01602bc736f73ac9915bd45b1b12b0f3d5456004950f02e74bbd40"></a>

## Next pages — rule_list.rules.url_category_list / 80e4a0d63a4a / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-afab4a047205d697bdab226cdf941e709470e98e95e0b5159240eeff46cd6ebc)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)

<a id="canonical-0a1dba132dcaf5b5fd87600bcd2b4d3970caf054a75249a1bd899cb20db29093"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-164bebdafb7771a628db45869edd310d70aa297734c4ad8fa6376f7e1f55f1ba"></a>

## timeouts — timeouts / 9b2dcf154c3a / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- timeouts

<a id="canonical-36fe3e0d81ed9c7197e152bf6d69498796c93e2393100a0aa02999a235072f54"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e9b29fd482a5fbe5e22797afa24bb414de8e34d0ec2aa1d4db316f968e87593a"></a>

## Direct properties — timeouts / 9b2dcf154c3a / 3

<a id="canonical-a76dd655832c1d0aef519c9ffe274d7f6bbaba8da8fa3b024cf6e174a4492e2b"></a>

<a id="canonical-7809b7cecde00fcb302174d8acc0690dcd917ed6131f6bf6fe9e648bf7a98eef"></a>

## create property — timeouts / 9b2dcf154c3a / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-30e90edf21657ac494731b546330f941755f75abbcb9cb72f496349bcd1e2ffa"></a>

<a id="canonical-ae0b9b0878ef76dc78d8789c6b89a9afe0cbcb1879f212b5798d80a5da19354e"></a>

## delete property — timeouts / 9b2dcf154c3a / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-735b7252d854f651e3dae862f6611c2a0a562d8992067b87ca04be39e735a565"></a>

<a id="canonical-8cebdd919fbc8fbb7c763fb107e9b1be9e41455b194a93cd66507218e0370769"></a>

## read property — timeouts / 9b2dcf154c3a / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-5ff281d344f432b0876813f9cf2bed3b04c6fc22167fbf486d9ac627f176f550"></a>

<a id="canonical-54829a00bf1188f4693fe044b4178953269c006e479a775a8e8bdf545f9e8765"></a>

## update property — timeouts / 9b2dcf154c3a / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8d2f7af7ed3d2e6ebab28fbe11f6c1f3396a6af7ff1603d27f401f972a4a4bd3"></a>

## Next pages — timeouts / 9b2dcf154c3a / 8

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-485809bf0384a6a4cb7d90f6ecd5664258e1f1b256117117772a4b654fd70056)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-e83cf4454367ab2f5b7f6845514a2b10ce97ff6d01fb4809e7b01c220f14541b)
