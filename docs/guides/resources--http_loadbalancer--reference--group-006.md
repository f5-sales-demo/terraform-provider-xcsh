---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-efaca7e7c9232678d003a67330ac68119ef91f191c41fdbe9362defcece25fa1"></a>

## api_protection_rules.api_groups_rules.request_matcher — api_protection_rules.api_groups_rules.request_matcher / ca82eb55784d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- api_protection_rules.api_groups_rules.request_matcher

<a id="canonical-e258d659de3a8da5fe6eaed20625d47476867b15880b2c01aaac8d81e30b0b84"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-992892672d1c3876db3fc7725ce36298cbd90aaa3cef20d1240ae04f2134b2c6"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher / ca82eb55784d / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29): complete subsection reference.

<a id="canonical-f2b92853e6cea2d30bca0db0b78cc2fad5ad32fc7a811ee1c2d941a4c35e9d44"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher / ca82eb55784d / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69ff4c5e36699c5ab276f79f9305f8ef1631fc5fb1a913ccf8710e46b76f8361"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / d81e9929f8c6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers

<a id="canonical-ff36fa5389e5998a7b30b0551294faa6abd68840cc9aa01cd50f23c00e424ba9"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-26f8c06ac287ed1d0290b0459c8da00db31845e81482cc36d96adc6366de309c"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / d81e9929f8c6 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-150ecfc21eb9ff546f77d67f0447138cee591cf280ac1ef747200545f4e85e43): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-4d4be9f07eaebe4250104a57d9813f3c7ac18ffa6a425c160014a8108247d168): complete subsection reference.

<a id="canonical-4e5245c8677f323ec06b538b70b165a439e551ab0fa201f54527c9e69956fea7"></a>

<a id="canonical-61ca0dc73a5a47fc9a9abc603cf46d8bfaa2ec42c65c8bdcd540f90b7a2336cc"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / d81e9929f8c6 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-ef4ec75ef44fa15efd9874b0df1daae4e7715a19dbc796e7cf3dd8b65369e43e): complete subsection reference.

<a id="canonical-b84df0b0ee0e26389203d462bd143e7c8d8bf7e8115287d8f0e35c4e6fe0d865"></a>

<a id="canonical-b6db1b60f2f0a69b7f8fd314d77972a7636d45b139c396d4a083f1fc90cd6edb"></a>

## name property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / d81e9929f8c6 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-b135df3eb6ecc4c264d3ffc8f669de44a286a9c9b58fffd08671a5c635325ac9"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers / d81e9929f8c6 / 6

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-150ecfc21eb9ff546f77d67f0447138cee591cf280ac1ef747200545f4e85e43)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-4d4be9f07eaebe4250104a57d9813f3c7ac18ffa6a425c160014a8108247d168)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-006.md#canonical-ef4ec75ef44fa15efd9874b0df1daae4e7715a19dbc796e7cf3dd8b65369e43e)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-150ecfc21eb9ff546f77d67f0447138cee591cf280ac1ef747200545f4e85e43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0720e83908b964d1871d6f81a44d536418c3de75d45dcf0eed9513edc6a08b62"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_ / 7d9c72ae4cdf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-7611278a37e55b7f69081eec5dda0b5171fabe95774a79db59acc1b90a393e17"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-a2f3f6aff8e39506a560078d848395c1a33a45c107f716e12e9cd3e21e0d8d59"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_ / 7d9c72ae4cdf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122aa320abfaf2adebc56c415423ea2e070f9b2c2831e144cbac64abe0dad9f"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_ / 7d9c72ae4cdf / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4d4be9f07eaebe4250104a57d9813f3c7ac18ffa6a425c160014a8108247d168"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2de4b6bcd2adc24fefae6cbfb06f6fef5c059d184fc96dfdeb48baefe2e5ea07"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_pres / 0bf6a688fdc8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-46b69612f761db50805f1867b5e66d46d82f733d442e3b4e40e80282243900b5"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-2fb6ab0c2e9e6c5abd258ccce9df51efe06d8d05eddc2d6557affbcf7108158b"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_pres / 0bf6a688fdc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74f72eabe6b35b761f60f3200413cc31c0b8feb9ca5d47388f1a8697bef8396e"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_pres / 0bf6a688fdc8 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ef4ec75ef44fa15efd9874b0df1daae4e7715a19dbc796e7cf3dd8b65369e43e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-764a9e499b3b04786c1474ef22686053642242dbe97396d14893dad0e99cd4fa"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / dc4144bc81f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item

<a id="canonical-4f9fab06c057519c17854236b271dbc9926522927cd77eea44d2b57132e060fa"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-b46a5f802463c17e06c40da34549abc1456ada8cff81830fda46776629838f70"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / dc4144bc81f3 / 3

<a id="canonical-adba6b47ecac88073182f629c9d707e1b86295dd6a3d5b47ffacd7583e953736"></a>

<a id="canonical-18a1dc3d75919c7184689fcc7d87ab08d490800c70ae0cf4ca79c4c99df6ee2f"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / dc4144bc81f3 / 4

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

<a id="canonical-57beb06b9372be686dedab460534737357f4587f2542389055a3c91e6557055c"></a>

<a id="canonical-4d9257207ac81eb2b1ae07909a4510ddbd02b426b37b0459f82058311a8e6d2e"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / dc4144bc81f3 / 5

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

<a id="canonical-3211b2df5c2dd961cfe5f67a64a185ecc705985ed18121e42cd07b77169afcf9"></a>

<a id="canonical-240ea48b3b620394d03125189c74c58568585875ddd0f903ad3a396b9ba96420"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / dc4144bc81f3 / 6

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

<a id="canonical-7882e74b358fa52adefd9eb334a85775be642ed9976dedd585cec5c6792f0ea8"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item / dc4144bc81f3 / 7

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-bcac43a893f76f71c7ebc21b51a30abf4bc3b85f23cbe70ea0018fd25ea7ef46)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-868ad025b7ed04ad11664486fa292b46d1a7e1dbe429c0177a8bbd2851aad023"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers — api_protection_rules.api_groups_rules.request_matcher.headers / 2b7ad1396c2b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- api_protection_rules.api_groups_rules.request_matcher.headers

<a id="canonical-b99842dc94564ab74c6b83e2b0a3d44cbd874a7db4af1db9637221242924a748"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
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

<a id="canonical-e108577d83f3b548002da2defce9cd01d8a7212b4fc40e014124a9d1489d719e"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers / 2b7ad1396c2b / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-15342515ea6755a0a8c6c85808fefa91e733baa151b79ff065f8579ecb5b8818): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-eccd86d708898f9799b08b3a2c1888ce117318cb04d0b727a0f4d12aa2d2970c): complete subsection reference.

<a id="canonical-ad542f3828e80cb049becdbee832e2ac313fbd443a1f593c01b2aa933a9621ff"></a>

<a id="canonical-381b6f1f99ca8a92fc1e3b6cb31093ee58d080514d7e99730ea2e9799f3ff2e3"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.headers / 2b7ad1396c2b / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-96609d35d11eee691ba840c532864af80fa0cf4e96a1670ae99c1f0ce12cf0ce): complete subsection reference.

<a id="canonical-3ec5e7f0ba140f3d8d5635fc9c4d125638fd23c577ceec1709519a48afaf5892"></a>

<a id="canonical-c72e3b93179b88da42906e428bf9b99acf87395704bfe3a2594d4302003eec4a"></a>

## name property — api_protection_rules.api_groups_rules.request_matcher.headers / 2b7ad1396c2b / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-aca77181ef0ef132843c8cea009ec3fb2ffdedb2ec7a0659f8b1d3b2012962da"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers / 2b7ad1396c2b / 6

- [api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-15342515ea6755a0a8c6c85808fefa91e733baa151b79ff065f8579ecb5b8818)
- [api_protection_rules.api_groups_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-eccd86d708898f9799b08b3a2c1888ce117318cb04d0b727a0f4d12aa2d2970c)
- [api_protection_rules.api_groups_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-006.md#canonical-96609d35d11eee691ba840c532864af80fa0cf4e96a1670ae99c1f0ce12cf0ce)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-15342515ea6755a0a8c6c85808fefa91e733baa151b79ff065f8579ecb5b8818"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84f5429d4e961b96a15ff443285027732121a66b5c3f594e072093edc2a2b846"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present — api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present / f189a4363b68 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present

<a id="canonical-8a0ad491ae8a0502aeed9ab25c53583259ad348bdee91ce7a050ba68d4199c11"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-60f30870748d4a68da9dd59e2dce2ab45649ad9340c5b633ca28ea050253a70d"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present / f189a4363b68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5fb41b1696dbe9fa459359d85838ddfa5d66d91971c350d9c70c15adffefb2c1"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present / f189a4363b68 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eccd86d708898f9799b08b3a2c1888ce117318cb04d0b727a0f4d12aa2d2970c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a017d4f3c56654624938ddf1d63cc7771f5f45a2f51a7a85213455091307b94e"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_present — api_protection_rules.api_groups_rules.request_matcher.headers.check_present / 3f8202f14f01 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_present

<a id="canonical-529953ea5ea261ae8cbfc9432ce1006c55d73d8b7dc7b637b7fcd88af8108ada"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-3b9d6e73628bb612b3932dd0dda7b38690fb2a23f72ba4825b8d71a0e2ebbcab"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers.check_present / 3f8202f14f01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ac667daa00ef32de1bcde2e6919f1200933122c549303a6e2823b4831d027739"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers.check_present / 3f8202f14f01 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-96609d35d11eee691ba840c532864af80fa0cf4e96a1670ae99c1f0ce12cf0ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa3250b42bd3680ba9b129b460183c11e1f41b646257a26670938c589aac7903"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.item — api_protection_rules.api_groups_rules.request_matcher.headers.item / 1bcbd3d0c53d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec)
- api_protection_rules.api_groups_rules.request_matcher.headers.item

<a id="canonical-b3768363613583a6e27f844f86c46975bba3df0ae19f5c1cec4adb89738f640c"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-87b385c62c99d72cf37a6b030770271799de2c458cb269c3379c4c2883b40b4b"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.headers.item / 1bcbd3d0c53d / 3

<a id="canonical-9367a120c23054659017997c19cf429133103dd8d2ce309915a7f12539266d44"></a>

<a id="canonical-90a49946d05e79dbb3d977f1ac89a638a30cb07508b9c7e0a1698c5a607d7c82"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.headers.item / 1bcbd3d0c53d / 4

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

<a id="canonical-5a46f808ce86b6355ed283109a14d33082bdc0adb59ad87388893e0eda72c224"></a>

<a id="canonical-e7d3f33b0e5c207a8b0a9d6a439bb5848ffa4cc3f9dfa119e15ced62762e593d"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.headers.item / 1bcbd3d0c53d / 5

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

<a id="canonical-c7052bdcd2a158d8b271a967c42a83fa18329db98fac9c7626ace631113507c3"></a>

<a id="canonical-fd4933645acf5f8459f9b4abcef7479fb263a4a62b443272e779c0f9ab548b2f"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.headers.item / 1bcbd3d0c53d / 6

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

<a id="canonical-0bcad5fb252776719cde93c322dcb99f1bd3a943ae2eb9ae3099f638dd8331a8"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.headers.item / 1bcbd3d0c53d / 7

- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-525cb052c36d4a74c2b60313f2433a16806838c2d0c4b08500592cc30ca08fec)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc695ca69da63f86c297bfd4d31c87d0712f9176f4cae9f148d425a4f26db03b"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / 156ef539a817 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims

<a id="canonical-6e1dfcb0704699eddad1f0fce662fc0a63e5a7e41e82f7db21846dc26d16828a"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-171bcd9f7b101c970c440783fdd281017fc5f3d149631415c4ef1088784acc98"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / 156ef539a817 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-692b94c5bb994d89f1df5937b72d7d5a28736e80c80d4c616590c6fc572d5766): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-1699b1b29ff21273e341662b9d752c9b7cb361818eb11c1c973ac406d87429f2): complete subsection reference.

<a id="canonical-e5e3e8ba9a13ac4820ff736397051bc9598c34add67aaf13a1cb3a8682f345c7"></a>

<a id="canonical-8c4a6821704b2a1d576ea1a39c42859a98c6677833eed355bab0c04a1dca7fd3"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / 156ef539a817 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-e985eade3a1c885408da47038fadb109f1ee5de0a75188e94faebc7bfeeeaeb9): complete subsection reference.

<a id="canonical-ea796c7ea14a81fc9430496c2b7a787aaec6e59767491ab97abc9404639a1249"></a>

<a id="canonical-5e9a5b2ab8ee2201e47d614459095ff9a355c963130373486ed21fe218522cce"></a>

## name property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / 156ef539a817 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-6494cff126688c16360aab41b176c1d6eddf188151283abd6900c633f1613633"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims / 156ef539a817 / 6

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-692b94c5bb994d89f1df5937b72d7d5a28736e80c80d4c616590c6fc572d5766)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-1699b1b29ff21273e341662b9d752c9b7cb361818eb11c1c973ac406d87429f2)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-006.md#canonical-e985eade3a1c885408da47038fadb109f1ee5de0a75188e94faebc7bfeeeaeb9)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-692b94c5bb994d89f1df5937b72d7d5a28736e80c80d4c616590c6fc572d5766"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a664691f4ab09715dc39385bc4c323dbd8d8d47abf2d8069b9390b3b3a04f35e"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_prese / 5b632f893291 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-ce3b26f217f3c515dd533d6b2d349e51575564923719f28a34e7544e87e13dfc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-a80fabd5d6d88d3b8f30b6b96cfaa3b47bad3e089b57bcbd581c4a8bf9940ad2"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_prese / 5b632f893291 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e02c3f60ff169cb68ff89dc87698e5ec8120d2b81f255b3eb0ca525c5a347c0"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_prese / 5b632f893291 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1699b1b29ff21273e341662b9d752c9b7cb361818eb11c1c973ac406d87429f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-677f571b76bb928bfe5a55d8ad44c1ab92e6c6d55bd1b8b2febe39e2347c9226"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present / 194638d9c839 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present

<a id="canonical-f9e865ec6a54252d0eff7540e3dede379f2759f8b6f9eddda632d84019bcf4e1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-7095b000b3ae7f891b5a8db629614095d62c7bd0b7cb7e6de077caaa94c2d676"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present / 194638d9c839 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25eeae2fa32809d3d0164b5f2741b2627f8ddec6a03cb0102b844cde30bd3cc7"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present / 194638d9c839 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e985eade3a1c885408da47038fadb109f1ee5de0a75188e94faebc7bfeeeaeb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6e37ebfbc490d66d438552e14f6ade98672ca90f4453cfe566febf3470e5ef8"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / ab346ad9851e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item

<a id="canonical-a7dfc0c3b3fe987b18eebdf0f9b7d2f87e8b618b8d85a94548cab7d696952bc8"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ba23c67366025503a70047d63a273ff29ec1ea095ee1bf747ab266803b43771"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / ab346ad9851e / 3

<a id="canonical-77f9e0249b4af90d83eda93291c797cc81bfe0dc877232115ed71e9bdfb7ea01"></a>

<a id="canonical-649af0ef67647b1ea902bf0f6b2e805d271e368454c727cd341af062f2ebdf77"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / ab346ad9851e / 4

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

<a id="canonical-d58ae4ec9a98176503cded1b1392eff270e22f241ab2bd4d0c7f8fe017c2ba67"></a>

<a id="canonical-0a8f3624c89bd4cd1ef3da5c20e0fdaf46eb0128d992f54f284dad5d32e3c4b4"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / ab346ad9851e / 5

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

<a id="canonical-e5d9e096868794d54797af321dd7ca8a5dec8287af57a4beb863584a86078328"></a>

<a id="canonical-bf1f4687de3120231b0a24200cd0a2decf9dfc4b8155c5748973d0bc2ea7d75e"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / ab346ad9851e / 6

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

<a id="canonical-8dff0ba758a2a4c0b29d63924d32bd7c4b968efc496211756733ef3a9d801775"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item / ab346ad9851e / 7

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-40ba267d28e7b815a58a1fbe380c0ad4ba10f5f28e7ce00f482edab41816d374)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f53f9368673719ed38c764d0f9d32e437d78bb6895f9b8a5e4b6b5cba4231ec8"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params — api_protection_rules.api_groups_rules.request_matcher.query_params / 16ec639495ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- api_protection_rules.api_groups_rules.request_matcher.query_params

<a id="canonical-6aeb77b8198554c134f2acbdf06427c726e27fab44857c7bfbf64feac349ac3b"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-638ce2a9730c873887a580795cdd40075d9865278faaac389851024a1a5c5cd6"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params / 16ec639495ef / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-b1c6d728caf188daf668666b45250c35059afb03ac71a95b24d125a7becfbc36): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-b2817914fb69f877ae68f219e41cca41c86310ad5f940dcf0f747e5d250081b1): complete subsection reference.

<a id="canonical-4bc8e379e8e380d22420bdde08e33f23184f26b9a475717ae96aaa424536c280"></a>

<a id="canonical-1885e9fc42435ace1af197426ae55f94aa17d16ff3ed9af498b9d140cd0b2c2d"></a>

## invert_matcher property — api_protection_rules.api_groups_rules.request_matcher.query_params / 16ec639495ef / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-162c865e4fe28d3034bfba6689393fce3e1caa7dd53d6e85651c2e8eab0170be): complete subsection reference.

<a id="canonical-8ef035f7590dfcd714185035093d449c5f1435e7c497b06be2a6621a18af6cd4"></a>

<a id="canonical-f271145002a6d47e8d74d752257d4fc60231452a8c83960805b755fb5c6c2a77"></a>

## key property — api_protection_rules.api_groups_rules.request_matcher.query_params / 16ec639495ef / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-93242f2e5d6dae454ce33d7defbf2596bde1c785a693967f4eaf31ea062305aa"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params / 16ec639495ef / 6

- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-b1c6d728caf188daf668666b45250c35059afb03ac71a95b24d125a7becfbc36)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-b2817914fb69f877ae68f219e41cca41c86310ad5f940dcf0f747e5d250081b1)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-006.md#canonical-162c865e4fe28d3034bfba6689393fce3e1caa7dd53d6e85651c2e8eab0170be)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b1c6d728caf188daf668666b45250c35059afb03ac71a95b24d125a7becfbc36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cad6945f8ff5875f88e9511fd4f8d0716a5a432789ee0112023e32915cec0c7"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present — api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_pre / e2fd94b40ca5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present

<a id="canonical-84507377aa93a96ef0ed6db61bf9fad45c5a980bb045c0264fd373fb8ee1a217"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-5d82055df5d08ef1740cb917a375a03da57f11f35d49c5d71e76cd5e5f5357a5"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_pre / e2fd94b40ca5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb263ac4b02b8222cffcbe766a42bd2c0d1014297a4f3bbadece71dd714f47fd"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_pre / e2fd94b40ca5 / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b2817914fb69f877ae68f219e41cca41c86310ad5f940dcf0f747e5d250081b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f57d2e7f6345a75395c5f8e3768b8c6e4be9c6744b0b1ccee2beec4f3864441"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_present — api_protection_rules.api_groups_rules.request_matcher.query_params.check_present / 8d433b190993 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_present

<a id="canonical-5538e84509e85dd8dc067f277ae54741cee24a0edf1a5af7f744710a6c7b14b7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-1c45104b94c8218ddf01e24507c7a2e6f364fe36d0ba34a83ceca11b7b02b2fd"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params.check_present / 8d433b190993 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9f419a443c95f52bce8755eeec57bf6804ae19bfd6d2bec00b25adac3c1dd98"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params.check_present / 8d433b190993 / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-162c865e4fe28d3034bfba6689393fce3e1caa7dd53d6e85651c2e8eab0170be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97b032049f8f0923cc1601382bb900fe3576ed3dfb1e0201e785dfd59f58a605"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.item — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 2043120643f1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-3b1381d64777b4f5a89a689f219cfa952e950413ffd1efdf2b31880c6aa4a85c)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-7204ea846ba92086800e02d67a46281be189a5c2fb8889c5ac06a6d3ec5a33c5)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2829dee6801699124f354f49d1b1d6b69abd8bdd76c05d8cef9252c8466a5546)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29)
- api_protection_rules.api_groups_rules.request_matcher.query_params.item

<a id="canonical-c3e4196d6d793eca14ec0321de879d559a8fc0a39074c165c07e4d247d763127"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-89e22aa3abb6f0ff922b2f090bf623b599d7b407f2d62450d7c610e5e72fea99"></a>

## Direct properties — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 2043120643f1 / 3

<a id="canonical-66e13e5d7a00d9a9a2f60d61be17556df7ef912c746b01d71fe2840dabdf32c6"></a>

<a id="canonical-53539bf8c50d6fb6b96ddce87ac5f894db2235cf0b83c4048b61c4c370ced4d8"></a>

## exact_values property — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 2043120643f1 / 4

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

<a id="canonical-69d434a7c3600fc70fbe57c34f0a298950d048fc8621616a02f4de4931650b87"></a>

<a id="canonical-7c95dc69dc9baec28ac424287a65682e0c86d7e46ecf3f61072e41c445d97c1f"></a>

## regex_values property — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 2043120643f1 / 5

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

<a id="canonical-f584e18baee768c1c4cb510fc9b52e16301f8c0af228dbc7423c016a78c08831"></a>

<a id="canonical-bff1ddb59aeaf43d57b0964a5a31023963292f06d6b44967d159d1c09621dc30"></a>

## transformers property — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 2043120643f1 / 6

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

<a id="canonical-329bf50b549dac29f102f3f41b088055cdbda3b7892d0bb47245d4f2bba5f05a"></a>

## Next pages — api_protection_rules.api_groups_rules.request_matcher.query_params.item / 2043120643f1 / 7

- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-d40954a5fe1a9942fb52f4a3b66dca1b84ae5e60876dbfc3bba08d7d7aa0dc29)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e644227cb5c65d3d1460f8f482f03792c7c929d078b4c009409fbec0f1cda055"></a>

## api_rate_limit — api_rate_limit / 5b028acf10bd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- api_rate_limit

<a id="canonical-2f37149c5cb28324f062cd3e822babeb8b40ef0d2bcca60cf0a928878534d4e0"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\] Path-
or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and choose
inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Upstream description:

Path- or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and
choose inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "custom_ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-2f37149c5cb28324f062cd3e822babeb8b40ef0d2bcca60cf0a928878534d4e0)
- [disable_rate_limit](resources--http_loadbalancer--reference--group-017.md#canonical-47adf84a77fa93ae7d8206d6c0c59c560f0728d1b6ff16f895e03e19b5d03eca)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-66b6e5d20091beacb9f66b456b93933288390f6d274966deea27b7ff44fd2239)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-b1db265907a35489dedcc7dded56c880a739939dd3e4d557d128f05f60dc74c2"></a>

## Direct properties — api_rate_limit / 5b028acf10bd / 3

- [api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8): complete subsection reference.

- [bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-83438b206ab5978b7c2a91280ed12652707aa076d66ece929570a47a3bb96981): complete subsection reference.

- [custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-b7b28053e6c70126c76df78f5c8d05a459dce130378187bcac769f9645feebfa): complete subsection reference.

- [ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-61932de3a5c2e6783be5d0bb3d6ba270ac39bdc0fc21e0b0cc57a72fdcd06ef0): complete subsection reference.

- [no_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-4cc948a3cab3c32c9bbacd1a87339619d2654cd958da12277ae5daeb5cbcf3f9): complete subsection reference.

- [server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393): complete subsection reference.

<a id="canonical-3eb1e794ef875d097a085602b9e21540b6b0d856e0faae6a0f79ffb74a4359b0"></a>

## Next pages — api_rate_limit / 5b028acf10bd / 4

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-83438b206ab5978b7c2a91280ed12652707aa076d66ece929570a47a3bb96981)
- [api_rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-b7b28053e6c70126c76df78f5c8d05a459dce130378187bcac769f9645feebfa)
- [api_rate_limit.ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-61932de3a5c2e6783be5d0bb3d6ba270ac39bdc0fc21e0b0cc57a72fdcd06ef0)
- [api_rate_limit.no_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-4cc948a3cab3c32c9bbacd1a87339619d2654cd958da12277ae5daeb5cbcf3f9)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-b3c1098ae74095c96f3a42d34809c92a0f6dea8fd1e1dd8f153177ac55845393)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c59bf0dc1e895852d44cf0f731a124d43958826262f6186673f2ecdba3de5cc9"></a>

## api_rate_limit.api_endpoint_rules — api_rate_limit.api_endpoint_rules / b20ef317eac0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- api_rate_limit.api_endpoint_rules

<a id="canonical-5e757fd38a712e0aac517f261e3574186d46a92ff25a9126103b6d5a7b1ef4d7"></a>

Type: `"object"`. list nested block, Optional.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("api_endpoint_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_endpoint_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-e50dcc71192e426ec0a8cbdfe5a1629895291b7f440c61c0571f4f68b3d462a3"></a>

## Direct properties — api_rate_limit.api_endpoint_rules / b20ef317eac0 / 3

- [any_domain](resources--http_loadbalancer--reference--group-006.md#canonical-52a79f1b56dad8f9a6e1c43488ea83a9beb14644d6cb6e26cae6b8b8c94275fe): complete subsection reference.

- [api_endpoint_method](resources--http_loadbalancer--reference--group-006.md#canonical-e5a2fe6b2d40d38a86ac06636e1341caa4b79acfc07fd9094bb0e5a482fd0913): complete subsection reference.

<a id="canonical-bd2addc4970ca79aa4b76c54c3ec94b7061430eb825e336c4d63d9ad7601749e"></a>

<a id="canonical-2473753c9c4ea66541b060ea2bc2958fad5d8b7b7a1a96dc27e11e7c5dc6bb81"></a>

## api_endpoint_path property — api_rate_limit.api_endpoint_rules / b20ef317eac0 / 4

Type: `"string"`. Optional.

API Endpoint. The endpoint (path) of the request.

Upstream description:

The endpoint (path) of the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-59bcb76843568a451fd355c0497676cdfb304f30208274496f31c3bf07e90d8d): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-cd6549792aae039f6b4e5d8b46aa14043b8ad700bb87ce546abb5272e28b558f): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a): complete subsection reference.

<a id="canonical-e3675d2c90ffa913edc61198ba455f5417fc14655f8ec9efdb909e4b0a1aea18"></a>

<a id="canonical-82b29d18dc377768d2a9315c14ce55d920ae7be0fb1a825b458dfe9abf4f49a9"></a>

## specific_domain property — api_rate_limit.api_endpoint_rules / b20ef317eac0 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

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
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-d9731447ab4883bfb51c6808c3c2174df8ee052b1d8d0d97d4a8bd2ea678af6d"></a>

## Next pages — api_rate_limit.api_endpoint_rules / b20ef317eac0 / 6

- [api_rate_limit.api_endpoint_rules.any_domain](resources--http_loadbalancer--reference--group-006.md#canonical-52a79f1b56dad8f9a6e1c43488ea83a9beb14644d6cb6e26cae6b8b8c94275fe)
- [api_rate_limit.api_endpoint_rules.api_endpoint_method](resources--http_loadbalancer--reference--group-006.md#canonical-e5a2fe6b2d40d38a86ac06636e1341caa4b79acfc07fd9094bb0e5a482fd0913)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-59bcb76843568a451fd355c0497676cdfb304f30208274496f31c3bf07e90d8d)
- [api_rate_limit.api_endpoint_rules.ref_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-cd6549792aae039f6b4e5d8b46aa14043b8ad700bb87ce546abb5272e28b558f)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-52a79f1b56dad8f9a6e1c43488ea83a9beb14644d6cb6e26cae6b8b8c94275fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fb2a17c89ad6fbbce82403c46fe936fb6be99b7eb8b1fae7bfbcb2944a549e9"></a>

## api_rate_limit.api_endpoint_rules.any_domain — api_rate_limit.api_endpoint_rules.any_domain / a517b287ae91 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-54cdeaad2414495472db25088411e941004eaf5d6da044a0015bad899bf9fc40"></a>

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

<a id="canonical-e04198b08b6cdcba44f8dab96ccf09904795d1cbe8389e2594ae916be5990583"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.any_domain / a517b287ae91 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da66732508892591c05331435126d67c87681c88d8b9931f65c64dd36d21a63e"></a>

## Next pages — api_rate_limit.api_endpoint_rules.any_domain / a517b287ae91 / 4

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e5a2fe6b2d40d38a86ac06636e1341caa4b79acfc07fd9094bb0e5a482fd0913"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93687493ce128ff4e393a6dabb7919d37276641ecff652835e266130d49ecd1c"></a>

## api_rate_limit.api_endpoint_rules.api_endpoint_method — api_rate_limit.api_endpoint_rules.api_endpoint_method / 6a8a97942ca9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-dff2a429db821f9ab190ee0ea8bb22101c0fb846ec84ff1e04b1b83e4d69e095"></a>

Type: `"object"`. single nested block, Optional.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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
api_endpoint_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac93e2da8a2fbd36cdc48c2a724bb4def09cb1b10700a78aef13e5a108798c43"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.api_endpoint_method / 6a8a97942ca9 / 3

<a id="canonical-7ede9bc932c5c15ca40f27b029cf35a87fd5189b5820613af9fb58c753afa73c"></a>

<a id="canonical-3159f8fb2dd858fec3bded691a0762a397cff3006f93099c978235d843763351"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.api_endpoint_method / 6a8a97942ca9 / 4

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-c3952b779baf1612a0753c63da1cb1f3be136883cb09ce8ac1a7b67b8210b41b"></a>

<a id="canonical-2c25929cfd522b1fa549f399119a7e71888db40e59660d65072e2cf2fb4050c4"></a>

## methods property — api_rate_limit.api_endpoint_rules.api_endpoint_method / 6a8a97942ca9 / 5

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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

<a id="canonical-724df09ddce2899e7ea4d9e2881124d59f1e42610211bba2029b5b89ba97bdf7"></a>

## Next pages — api_rate_limit.api_endpoint_rules.api_endpoint_method / 6a8a97942ca9 / 6

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34653abf798c45a1783402ad7db350ad87f41006434e663d60c40ed4eb6b9a35"></a>

## api_rate_limit.api_endpoint_rules.client_matcher — api_rate_limit.api_endpoint_rules.client_matcher / ec931a466564 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-d228b209f995a84c7c1bf6ac4212a05aa1db536a1dbb9147ee4d457980ada128"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-31a7e60aebe59e4fc5046dde8389584a3729449e641ff9f787aeb24105181602"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher / ec931a466564 / 3

- [any_client](resources--http_loadbalancer--reference--group-006.md#canonical-06a5dddfdbc149dce55a0ddcc0d96b318a2debe76102e1799eabb48ea516d411): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-006.md#canonical-d5b5e28e86867715e46bb6539c502525039dfe97124540ad48b7ab370f686d27): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-006.md#canonical-2d7529119f1d39dc395d7cc524e50b142652a53063d56db0ccfb7425c17be5b2): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-9f4539315fc253f8a8e1d29603298b0912d2c6ea0287c27332a5721b3d4f03db): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-006.md#canonical-d0e1539a96c89754a7217a139bd2a0e0784e4668a324d56007bd51912f74289a): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-e0ebaa2657df665e0a212f94039066fe2e87820b7d0d13545f56f53e4ee6681c): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-006.md#canonical-768bb219bc6964e792f66f3d35d516763da6025f954856eb150b24f68cd9e840): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-006.md#canonical-6d7050a121a2478a62145eb8eaa6c16568dbbc11573eebb070715ecc1718e3c9): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-e188cf56d2496d5aab5305cfad06ace5c1e3076fe2aed29e677fd190da2323db): complete subsection reference.

<a id="canonical-0652f0ec49f7a1a2a4dad208c7c19904eac420d9680b4552971607d55ccf6838"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher / ec931a466564 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.any_client](resources--http_loadbalancer--reference--group-006.md#canonical-06a5dddfdbc149dce55a0ddcc0d96b318a2debe76102e1799eabb48ea516d411)
- [api_rate_limit.api_endpoint_rules.client_matcher.any_ip](resources--http_loadbalancer--reference--group-006.md#canonical-d5b5e28e86867715e46bb6539c502525039dfe97124540ad48b7ab370f686d27)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_list](resources--http_loadbalancer--reference--group-006.md#canonical-2d7529119f1d39dc395d7cc524e50b142652a53063d56db0ccfb7425c17be5b2)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-9f4539315fc253f8a8e1d29603298b0912d2c6ea0287c27332a5721b3d4f03db)
- [api_rate_limit.api_endpoint_rules.client_matcher.client_selector](resources--http_loadbalancer--reference--group-006.md#canonical-d0e1539a96c89754a7217a139bd2a0e0784e4668a324d56007bd51912f74289a)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-e0ebaa2657df665e0a212f94039066fe2e87820b7d0d13545f56f53e4ee6681c)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--reference--group-006.md#canonical-768bb219bc6964e792f66f3d35d516763da6025f954856eb150b24f68cd9e840)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--reference--group-006.md#canonical-6d7050a121a2478a62145eb8eaa6c16568dbbc11573eebb070715ecc1718e3c9)
- [api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-e188cf56d2496d5aab5305cfad06ace5c1e3076fe2aed29e677fd190da2323db)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-06a5dddfdbc149dce55a0ddcc0d96b318a2debe76102e1799eabb48ea516d411"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26b7b3bc24754d178e675afb4165f3e6c42d099823982d1b0c96881dbda27467"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_client — api_rate_limit.api_endpoint_rules.client_matcher.any_client / fea15cf7e348 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-fb108eb0372ccfa02a8e6409294aa81557e5229343d7784e404e91a7054ceded"></a>

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
any_client = {}
```

<a id="canonical-577d7151024015a44e509885a16cb18c0764781f63df67781df414ebed7d96e8"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.any_client / fea15cf7e348 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-77fcc15eb1d5c6c30f4a9bacaa6906277fc20bfbedea8437ae32b191b894df9c"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.any_client / fea15cf7e348 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d5b5e28e86867715e46bb6539c502525039dfe97124540ad48b7ab370f686d27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ffdd8cdb632a94183e0958dc912fe8c08ae7538caf4559edf46af56475ab21a"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_ip — api_rate_limit.api_endpoint_rules.client_matcher.any_ip / b2e3d6940e54 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-6e4c14fb9b41f0eb7421d54e62d53b32cda25bbe288d924fbd4891e36835b266"></a>

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
any_ip = {}
```

<a id="canonical-728b620a32698e51a3cc1601590c2ca113283251011adaa43ff38a7faf146b58"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.any_ip / b2e3d6940e54 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-438636ffe240f6ec3667fe8fdb096e811355aa87899f930e80c7f27274f963f5"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.any_ip / b2e3d6940e54 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2d7529119f1d39dc395d7cc524e50b142652a53063d56db0ccfb7425c17be5b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c5e447bb755c91af01e8fd4b19dd3c15c8789a42c032e68ae17e71b780cb122"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_list — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 24f6a4d6f079 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-03feeb666ef387c6ee493c17970170ed608f36c21336b3efadb11d10230b69ca"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-8458b517c1117b6b4869ea967c1a958ada3da048b4795cbaa7a2d703a9a8c29f"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 24f6a4d6f079 / 3

<a id="canonical-6067c70842de1f7b96f6bb822cd5d2492affbff165290d04ce022a66ced9280e"></a>

<a id="canonical-38f13e86118d855fbabcb8bf922c2266b55387b36060bdd49fe03db9aa58e727"></a>

## as_numbers property — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 24f6a4d6f079 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c55f2a04e76d7c2cec4ebb48473209a18f02efa9c6163c19ff64b513494c94a6"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.asn_list / 24f6a4d6f079 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9f4539315fc253f8a8e1d29603298b0912d2c6ea0287c27332a5721b3d4f03db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f538a762115007555b0d8d6fe0cd33cd18eb2147301ddb8f31630a82a0269b4a"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher / 69ea30d50173 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-8f5d99595cfd6de9fa199008b2ae1067291f073c7fb99e3820473c7a73e5b000"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-bb227d81d92cdeed8c7462e0fb22e37e1064587661470d6aaf814fd65ca74c68"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher / 69ea30d50173 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-006.md#canonical-6093852534770a9806fde46cebf7a562f0866094211a4acc22edc80247e8bc7b): complete subsection reference.

<a id="canonical-908a7f976fccc96a0e9419244f18e1a98b926748aa32400e9bc5458821743713"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher / 69ea30d50173 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-006.md#canonical-6093852534770a9806fde46cebf7a562f0866094211a4acc22edc80247e8bc7b)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6093852534770a9806fde46cebf7a562f0866094211a4acc22edc80247e8bc7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4e63a2b1de19793eddb13c36f2f550084f74e3490c34145f0e84f445577c989"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-9f4539315fc253f8a8e1d29603298b0912d2c6ea0287c27332a5721b3d4f03db)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-a3e9a88e742faa0ed3c8adeb4dd7ce986932ee5ba145e5b8c62f5b2793fe525c"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-4f8c708de2ebb9921dc64cc9efb6b8933347deeeb22f4001f0876a391ff778ee"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 3

<a id="canonical-c4183fd41be3ef354cf3c54b1fbfa426256fa0adb74345d44fd8b9f9e4a91818"></a>

<a id="canonical-8fca0c22b3919510655c212c6350e1a2f87d6b19cb3176656299cf07f25b072d"></a>

## kind property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-cf1eb1ee5b13f2840edac96cf1e979acffc68a9c7b4b673c02956de7c41b1661"></a>

<a id="canonical-6e81d5b2b68c329c8fe9eb4be6368ab88d6f7142c685da89a9d634b0d2fc540d"></a>

## name property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-d6a620f48c8953ae8ba7bad108fddf4aff2b7d53095d7dc615b44057f6dfecc7"></a>

<a id="canonical-273e67fc6cf79e6d2f86ea4148581608394f659145f229800a73f924bd959506"></a>

## namespace property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-6d89da829c511e4c06bea5f9269eb5f1540fb26645699621f80334580a68d5e2"></a>

<a id="canonical-8bc2ffee2a0405c84b73b2973eb5d0bd55e2aa724741ad18037c412bd8614810"></a>

## tenant property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-886f86a76f6c7fb9f1fe11a0f816ed6c8ecda129efc6cb5f2f2fc692ba066bf5"></a>

<a id="canonical-233cb13df6e4f49c75ae7847bae9f60d8392f18ca1394978bb86ec42da4b8eb3"></a>

## uid property — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-1ae148ad4f0025c527e328a03a6be1a35492364b53ea51eec6e31a80bf790be0"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets / 5729ed7021bc / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-9f4539315fc253f8a8e1d29603298b0912d2c6ea0287c27332a5721b3d4f03db)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d0e1539a96c89754a7217a139bd2a0e0784e4668a324d56007bd51912f74289a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d282c704bc10fce27c4d349d8d396b69b12dce4a755554107d561526ff9abf6d"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.client_selector — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / d6e53a3abe44 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-f8df13069f2ab41ab633b7f27d7f9ff7d220657679c00d750d750ee0cffb5f9f"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-58a12c0cb985348c8aacdbb670f82458dd4febcadeda7399a10ccc7ec1434001"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / d6e53a3abe44 / 3

<a id="canonical-d06228aba2dc8b4a0ff1b19c9a0f80bbec99baf1c70ecaa20843345c99db26a0"></a>

<a id="canonical-cf1128821ba00a8da865329eb2264186ba6f079bd9a3e3c60ca173062daaa229"></a>

## expressions property — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / d6e53a3abe44 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-dbcf372f678724fbbadab21fa5b939cdf555124f967ebf00120c76ef8971855b"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.client_selector / d6e53a3abe44 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e0ebaa2657df665e0a212f94039066fe2e87820b7d0d13545f56f53e4ee6681c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a6fbd44f6c4e99693d3239df31f241e061b7b6c0db4b637765230e1811ef775"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 4632c6268dc4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-f92fc2221805da9f034fd52d34a6baeee823361a9681cb7cf1c7b863243985f8"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc834f3ad84aa62f75c0ffc948425fba93159d1b868692745a0a60e046c79f87"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 4632c6268dc4 / 3

<a id="canonical-e0079d5517e2fcdef4020e3827207fd3b2628b1345fcda708616b85b1dd52ee6"></a>

<a id="canonical-9247449b53708c1c9f1c198af5450f773036020105e8b3580e27eda2762e272d"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 4632c6268dc4 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-006.md#canonical-8b67e9b209c55031caaada00a498846eefdb742aa40565ad469df0373492be1e): complete subsection reference.

<a id="canonical-33ff9d2dae9a1b8b029ba9dde2601aefda74f0d2454cd16ef4ed0f1293e36821"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher / 4632c6268dc4 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-006.md#canonical-8b67e9b209c55031caaada00a498846eefdb742aa40565ad469df0373492be1e)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8b67e9b209c55031caaada00a498846eefdb742aa40565ad469df0373492be1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12a0f454621183ecba9bd5120df5ad3d35e020af14cb6c7dcf7acf03759cc289"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-e0ebaa2657df665e0a212f94039066fe2e87820b7d0d13545f56f53e4ee6681c)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-eb697a42e352324258d92f1b95f05ddf9c919072f8c451a45a04dcefe760be55"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-720b37cd5fa1f2834e1e493165d954049fb9f339e1a58993658db7c35271c3f1"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 3

<a id="canonical-60ad3fd615f3c163e1699fa62f4fe9adc2ddf39933d241c5252198d8cece4f52"></a>

<a id="canonical-14f7b84df359c46c8968cb991dadeed112fc76fbad74f2abd46d4d052688d84e"></a>

## kind property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-ed26eb9f6bd04057d732e466fe0240c93c394a35a0e486914d9c225521415298"></a>

<a id="canonical-10559f5eb41d85e981c8929d202f2fc4d88a88beb537a824bcedf220f0f972ba"></a>

## name property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-297a1be999cdc753f12cf688a401fd72e3fb6bfd63e1e8e5397484d63ae4bf6b"></a>

<a id="canonical-2a0650e689dd342539d312defe756b2f0bbdb215729de909b9da31c65be89374"></a>

## namespace property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-979e783273639a6d380a8c23e443b78fe14bf7b26937e8f39a64993d35baddb7"></a>

<a id="canonical-f8c979ce8e2dfe3ec29b68d1d765c899b23fd90d3ea3fb11f9d8bec99ebee76e"></a>

## tenant property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-9f76373cd51c19f2fc4bbfcbc8593fe55c3f4d8a87bfd9ea709bb41af42dfe3c"></a>

<a id="canonical-a86b04b9b80e0d338bafdb84fb669e504720ce854eccaba2899f3194102659d4"></a>

## uid property — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-169b0174229b1d3b92ec4940692c7de4779735cecfeeac9f75ed29016c734cbd"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets / 64542a9c4adc / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-e0ebaa2657df665e0a212f94039066fe2e87820b7d0d13545f56f53e4ee6681c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-768bb219bc6964e792f66f3d35d516763da6025f954856eb150b24f68cd9e840"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82f6ddd59ff58a851aecaea1082de21a99c53020683453e517f8b36ba12ef000"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 950d3d39ca62 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-7ceb14f8c226b8f3c53d434d6783ab4d989040c3fbcb75ae4f15738728afce0c"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-54d032c42c6afac417b10cd29db576cc306a88f31fcd201b71426270ee35aefa"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 950d3d39ca62 / 3

<a id="canonical-13e0c40b4e7e186ac2c8cf6f49c61b31efb724961057d075bb5682d9079a516a"></a>

<a id="canonical-c88beb6d250f54911480fb1e7cb52e28c32f8559ec9a3ca7f4c00639876168a1"></a>

## invert_match property — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 950d3d39ca62 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-9e76355c9a59fb8cbdee919224f1441ea73889045950f1a415bf966ab1c6b6cb"></a>

<a id="canonical-172eae4575e4d06f97653d6bceba5e61ada942059c5c2bcf916823e215cfce05"></a>

## ip_prefixes property — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 950d3d39ca62 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-70066e0e239bdea5bbfd6f80dc249e6c5389bd34ae6d2fb46db0a85810f1c078"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list / 950d3d39ca62 / 6

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6d7050a121a2478a62145eb8eaa6c16568dbbc11573eebb070715ecc1718e3c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc502ecd0aa4ae6bdd3c6a4cbd523350802e8742098b1b4adf9d9223b03581fd"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / f9b09eb59c76 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-0167b4d3f07191092101ebcf80084c7343ee878199cf181d1e02d81d74a903a5"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-86a5f1764a335b7508ed5ef2b902689390c5d7ea5537e10a695f34b7a32b92ca"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / f9b09eb59c76 / 3

<a id="canonical-d64ec6784c631451c031b17e23eeaf0de3fcd0f4d349da6272a40d0af7f03030"></a>

<a id="canonical-671131f2936d59a685be43d48750efcbf121e479576f5d3188d95c36325d582c"></a>

## ip_threat_categories property — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / f9b09eb59c76 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f1b0b136e4dc79fd796b1027fe522e8e66366257b55682ecdcf951a4cdffa57a"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list / f9b09eb59c76 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e188cf56d2496d5aab5305cfad06ace5c1e3076fe2aed29e677fd190da2323db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25c81d3722ae7903745482aefb99fdf79546a9c05e7280db50f26d6229822497"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / d555ed34a5b0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-483240730fa4bc4d34e47ff5e7b26b4235f1f740e6f69096358feff6b2e4f07a"></a>

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

<a id="canonical-19b485828fed751fc081d84bdc9d2bb57ada66840b467af7a112a8ba891f8bf7"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / d555ed34a5b0 / 3

<a id="canonical-1ee708d4630aff4b3697890c75cafe1e231d7b33279037dd872981f69c0e04cc"></a>

<a id="canonical-6651a441300bcd711e1589785a515bc0ea99a838689cead7069af199f6e70d59"></a>

## classes property — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / d555ed34a5b0 / 4

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

<a id="canonical-94d9354e15c1fded5aad5740bdf9259152e540dbcaee39deadbd8ccb90ddd58e"></a>

<a id="canonical-a5a83767523503f41c24ac9bc98be7118c64ea4bd7ca0fe067e17f4d77f81115"></a>

## exact_values property — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / d555ed34a5b0 / 5

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

<a id="canonical-6c52a2ce226fa84e6dfe6e7d4c3db856a8e36165d2d9eb25980b9d071564f023"></a>

<a id="canonical-3edfa2c597b6a146959538f2015f78863a0a36b764d15c4ccb1c76be87dffdd7"></a>

## excluded_values property — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / d555ed34a5b0 / 6

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

<a id="canonical-a2e3cd3c3aeb1a276ef966d9d404df5f6e74fb1ade66ba146de2b988c95bb22d"></a>

## Next pages — api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher / d555ed34a5b0 / 7

- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-1586f4e464c1aa1376a1fb7d8adf451f9246049d4e94e542933ce57faa836e4a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-59bcb76843568a451fd355c0497676cdfb304f30208274496f31c3bf07e90d8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbad0f618c53da2ed0c8319b675001eeb0b1da9dd819e86914474869d1bcb970"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter — api_rate_limit.api_endpoint_rules.inline_rate_limiter / 15cab4bcedb5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-ce7c94af7d31037acc8a82576be90845d4c6d09e11125291900370a82260d68f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inline rate limiter.

Upstream description:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
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
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe199bcc61fd9565f70973f2bc0fbb9b0b9fc992546e3ea190a5e956ed2ec528"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.inline_rate_limiter / 15cab4bcedb5 / 3

- [ref_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-7a02471b17d744dfd68dd91938cbf53517357802eb676ca057bef8eac0127883): complete subsection reference.

<a id="canonical-75bab838f1a39e6556e1b6f0172c00b3466384c06e259a650257fc4eed880ed5"></a>

<a id="canonical-ac96717fa3729799a6aa9e5467d9c91f4d7e81ef8b0a4b0acdbfbf3bf74d0369"></a>

## threshold property — api_rate_limit.api_endpoint_rules.inline_rate_limiter / 15cab4bcedb5 / 4

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

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
    "minimum": 1,
    "multipleOf": 1
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

<a id="canonical-9c9620125a74a340725e086e28ec1b1d3cf87236ce9610f34201df571402d736"></a>

<a id="canonical-6e99badd9416d4567acac9a0e1ebf64cfcd5fec787ae3f08f4fe84cb6e7ad0ac"></a>

## unit property — api_rate_limit.api_endpoint_rules.inline_rate_limiter / 15cab4bcedb5 / 5

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

- [use_http_lb_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-fc350215cc8ef7a940e5f82379ddd103da9e57c4f097321e60370c1da4388d79): complete subsection reference.

<a id="canonical-59e2dd90aac236fa5d32459087e3b96541d8d1e6d60c356891ec3ebd161d2863"></a>

## Next pages — api_rate_limit.api_endpoint_rules.inline_rate_limiter / 15cab4bcedb5 / 6

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-7a02471b17d744dfd68dd91938cbf53517357802eb676ca057bef8eac0127883)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id](resources--http_loadbalancer--reference--group-006.md#canonical-fc350215cc8ef7a940e5f82379ddd103da9e57c4f097321e60370c1da4388d79)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7a02471b17d744dfd68dd91938cbf53517357802eb676ca057bef8eac0127883"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1513ed26e6856f128acbafb3124056e13ff754babcb0b8599d1ad5102890adb"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 4c1177f1788c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-59bcb76843568a451fd355c0497676cdfb304f30208274496f31c3bf07e90d8d)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-2ef560d932d51d7dbddd88f1667b7d0f620ff876b04fb9c10faf0a843250e1aa"></a>

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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-6acbd8d23fca4b4f807c61dae37cf99dceaac9387ab2d96ab4c3abbb7ca4e4f8"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 4c1177f1788c / 3

<a id="canonical-7710748f9c239e7564bd05ced3ece2d4bf10b646ec7deae801851b6014efac9d"></a>

<a id="canonical-04b1e1d64656dd30f35afcb3d6beb0219c2affad80153f66e5200d2d5751f546"></a>

## name property — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 4c1177f1788c / 4

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

<a id="canonical-189d373261131eb169232ed9a50372fe899e969aadd93b085f27a94175d4f407"></a>

<a id="canonical-80051ab841e1b4f49042628d49a950e76a8198307880616d3273742cf90630b0"></a>

## namespace property — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 4c1177f1788c / 5

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

<a id="canonical-52303a54ce19e48ae2c7c7ba245f1a678d5096ace979ee559101e478a2c66c8e"></a>

<a id="canonical-3c3c32f57fd74cbf316835f8b56f413db442f91b94ff690c0517ba8b29fd54a4"></a>

## tenant property — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 4c1177f1788c / 6

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

<a id="canonical-c822abe4542e96030d487e798c64197f72cc0147e9429b5560c28f6d899d1d0f"></a>

## Next pages — api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id / 4c1177f1788c / 7

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-59bcb76843568a451fd355c0497676cdfb304f30208274496f31c3bf07e90d8d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fc350215cc8ef7a940e5f82379ddd103da9e57c4f097321e60370c1da4388d79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f8a501c92d49fcf950dfacd5077347bfce6675af79755ab54bd494a58db6d28"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id — api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id / 66ad1873fb6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-59bcb76843568a451fd355c0497676cdfb304f30208274496f31c3bf07e90d8d)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-a3514f7ace1a2bfb5727f1256e924a397e8b3adf01ed22ca294c8d74fd0121f3"></a>

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
use_http_lb_user_id = {}
```

<a id="canonical-93ad13b0217f499db48ae4fc5c2288c16dba7b388352a80931077a98b0e2753d"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id / 66ad1873fb6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9863358e7f5a186ef99c4c244b6a789b295ccc34622015a380a1953f973c6d9c"></a>

## Next pages — api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id / 66ad1873fb6c / 4

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-006.md#canonical-59bcb76843568a451fd355c0497676cdfb304f30208274496f31c3bf07e90d8d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cd6549792aae039f6b4e5d8b46aa14043b8ad700bb87ce546abb5272e28b558f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9cc403ce1512b662dd7ae2e88c20d525ebacd4739ac6c9b728c1804db3204b0"></a>

## api_rate_limit.api_endpoint_rules.ref_rate_limiter — api_rate_limit.api_endpoint_rules.ref_rate_limiter / 7df3461256ea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-d69efdf099f0f4043379f331f808b0aaaea08cc1f0021a9aa41e7164cfe50f61"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd25c3206aac1023c379d0848a29654428bca1f5426a41f364710fd7aa595241"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.ref_rate_limiter / 7df3461256ea / 3

<a id="canonical-906f6eca9ca1b81ee19e8705da4d64f04e31ea2ee60c149f63ba63a9343d6afa"></a>

<a id="canonical-6b432d041a3324b9aee4ea35a8950ef1837790ecc5e7d22c0292c786f3078ef6"></a>

## name property — api_rate_limit.api_endpoint_rules.ref_rate_limiter / 7df3461256ea / 4

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

<a id="canonical-723d0a2a9d133979656a98b56e58f47c86e82ee1c473eeb9f1bbe7959afd1acd"></a>

<a id="canonical-41a787d7cc2f1b684df5525734c6b7e5512bf508a63442e417b8e160347bb59b"></a>

## namespace property — api_rate_limit.api_endpoint_rules.ref_rate_limiter / 7df3461256ea / 5

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

<a id="canonical-6afa46b0d8d04feb8e26c96eed6dd04815f477b8674c840b2756dc4e55da7f39"></a>

<a id="canonical-044d571bc9bd5b775987ab58a636f2204bd72d6b2abbb1bc35624c906867efe7"></a>

## tenant property — api_rate_limit.api_endpoint_rules.ref_rate_limiter / 7df3461256ea / 6

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

<a id="canonical-246a45f8848bc8467c9c58916a78b9996d0988ab760a58ce3c3c4cb7fe90e253"></a>

## Next pages — api_rate_limit.api_endpoint_rules.ref_rate_limiter / 7df3461256ea / 7

- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06445ff8b607839d5b230a0256653052c81dd3efc6734bf5789cf514095ec16d"></a>

## api_rate_limit.api_endpoint_rules.request_matcher — api_rate_limit.api_endpoint_rules.request_matcher / d5f01258676a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-7820795e8d1c208763873ddc6783814f9387a5721f6c41aee7a91837f2b77f99"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-332ffcc6783e52c5673efb4cc21d7420bc1a922f1bef62afefbed0335e5f2e89"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher / d5f01258676a / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-006.md#canonical-6ff66b29b731b6791474789a14899fdf91ab0c5aa20154b672949b7c1f9f2da9): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-1d761f2e1fe328370ff787e6073fdc546da194cd10ff19df9f845fc90ade7516): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-007.md#canonical-c5f6fefeaf8ed3232b55c08de311fdf82f6b1281c33b8dd2075248ffd642c9b9): complete subsection reference.

<a id="canonical-fca8fc623856fd262e248a52ac768362d8a3388e1b09353b36ae33debdb45c4f"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher / d5f01258676a / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-6ff66b29b731b6791474789a14899fdf91ab0c5aa20154b672949b7c1f9f2da9)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-1d761f2e1fe328370ff787e6073fdc546da194cd10ff19df9f845fc90ade7516)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-c5f6fefeaf8ed3232b55c08de311fdf82f6b1281c33b8dd2075248ffd642c9b9)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88db8a933e3f2367b773f2aae29c1ee8be2a3a52135c0e63568f52876106475b"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / b315acd4fe17 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-76b8232ccd2e3c83aff67557ecdafede3a3658ff776237d14bba50495e6a6a61"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-d47b072d41ec626f7a0891e9aa1071ba2a4f9f964ce82914602ac1b2a4f12e7a"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / b315acd4fe17 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-9d2cab907736c4977af7c16f8340498769342f3e9cbcfc9071405bbb6ecc18e5): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-9a45b9fcefb68b827e1cbe422b2af817fe0925f77c00e554e0a2b31aa0ee1705): complete subsection reference.

<a id="canonical-87ea04c9be182d1cf221c5f4743a8a6d9f43979d3ad150ebbf1fe983974c09c0"></a>

<a id="canonical-c146276d142f84e8a822480b7c35080995cfe8f617d9fc71e9db8d0bdd7241c1"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / b315acd4fe17 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-4600bc3590d308913c81fc589ebb84c7aac106e5ae84141a6fb0888646327805): complete subsection reference.

<a id="canonical-1e71a58b01cd17f4f7aee6708b8df03942308b285c64be577d4bbfa38beeddb9"></a>

<a id="canonical-8650d2bb401c76b1b4dde27ca44b1aba9d6db3f8f2eaea6232d8b0eb06963b19"></a>

## name property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / b315acd4fe17 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-91fbd080ec107810160da91769d431b5a723093a942660377e1ccbc93a4d303b"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers / b315acd4fe17 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-9d2cab907736c4977af7c16f8340498769342f3e9cbcfc9071405bbb6ecc18e5)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-006.md#canonical-9a45b9fcefb68b827e1cbe422b2af817fe0925f77c00e554e0a2b31aa0ee1705)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-006.md#canonical-4600bc3590d308913c81fc589ebb84c7aac106e5ae84141a6fb0888646327805)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9d2cab907736c4977af7c16f8340498769342f3e9cbcfc9071405bbb6ecc18e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c53bf0049412b7e312265bfe2746dbb40ca61956f16c00a7c9875f4943f87356"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_pres / ab321838a439 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2b34eaa885a9893b3792ae76fe219486cbbe8d2affb7b6bacd5197dc0e6feb1f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-7197550c214fae39ba6ce92d3430336b5545a65d40a225cf29176191c3b7319f"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_pres / ab321838a439 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9385b7e41ebfbd257336ebe172bd97c7bc782661acce041bdd1e14468c3625d"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_pres / ab321838a439 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9a45b9fcefb68b827e1cbe422b2af817fe0925f77c00e554e0a2b31aa0ee1705"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b5f8488ce563671e2d8bdaee8ccf86b853176860156c68b47bb9ea4ecc1cfa3"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present / 974fee40c810 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-29a5952505e0e6c9de96d5964d51ae19b21916e88d388cb80e63bc1954212f98"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-a7fb9f7b438d5452a77c89f3de09fab10ddef682f83b182e533711ade104a253"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present / 974fee40c810 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed7c173410ab077eb0ebfc12a6b5f30f1632af599be30aacb52929111e09a7b7"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present / 974fee40c810 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4600bc3590d308913c81fc589ebb84c7aac106e5ae84141a6fb0888646327805"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e74ee16a0af804bc1838094317f5b022fca8e7b434cb0bddb3aa6140f0a7b076"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / 2d39098245e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-fbc2f81cea3219799277899da34ff48467cde036dec42ffbdc6d24865aa83671"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-f29e9033c205f0d9eaef6b4724744fda4ee80e894e72378fa10575ff1031b27e"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / 2d39098245e4 / 3

<a id="canonical-499936a79445ab9d61f4d0aee49c20f0a898f7585f9a710e4f3bbe8123946237"></a>

<a id="canonical-7eaa7bd3e2d19c768f4ad43b74a6a1cc80eaea1fefd3515b02c839914990b906"></a>

## exact_values property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / 2d39098245e4 / 4

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

<a id="canonical-1cd342ba36e7d5e2d6c1ce75506b0a2ee1b7ee97814b196fbabc8645c06993c7"></a>

<a id="canonical-5d1b7043b80a9778ede6acc9b6a47c86844baf04ed8788d16f38501dc13c7159"></a>

## regex_values property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / 2d39098245e4 / 5

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

<a id="canonical-ed44caa10342e9479e32c8afd561c46b174e47107d9170619ac9afbbf2b3d8b3"></a>

<a id="canonical-7cc4ddb7e90365775d0b77ecfb0468ac36aef2604b9ec88cfcd565ccd27abb85"></a>

## transformers property — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / 2d39098245e4 / 6

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

<a id="canonical-4f96aaaa2fbb2d86d016d75ec098520da33c318a85ae981422d9924f203b4e7c"></a>

## Next pages — api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item / 2d39098245e4 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-cc9a16936a22d22d0ec1478509464e0fc2b14436b321399475f4ae061d6e85bd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6ff66b29b731b6791474789a14899fdf91ab0c5aa20154b672949b7c1f9f2da9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6537410e9cffb935c4f5736de58cce176296eed721cb72baee55e405368ffd4e"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers — api_rate_limit.api_endpoint_rules.request_matcher.headers / 267c2e7caaad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-1b01b31912bf2f089b589729fddac8f3f71d81998021584e3f8cf211ba1d9592)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-e6adecb5e52d41afb4154ad68a3a0905e616c1f8e4179602acd94099d5c22fe8)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-a3304dc0dc7f999a4c9ecfc207e39de29e2beb366f3940f5465976cbe2c7089a)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-5638accd74e7eedef10328d8280c9da6d66e2d72b870bfb01b96e078639eea03"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
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

<a id="canonical-970d72f0bf665cc299cfa893a74685803f48c60a75cd948f66e0b7258da45f07"></a>

## Direct properties — api_rate_limit.api_endpoint_rules.request_matcher.headers / 267c2e7caaad / 3

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-94e818ffe4235204def0be660815384886e4de59ff2a0953015f3bba9d19d8dd): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-1d171fe96738f6a6df68c0ebfcb02a61043c5e532a25545a7fdaca6d789b8768): complete subsection reference.

<a id="canonical-03d73851c1c5d928177ed2d710b3b191e24bd6e4dd732ec2191a5434e3f61508"></a>

<a id="canonical-924d7d05282ccd936ff5a7feea768375d8375a97e9a79714b564c78a21cca058"></a>

## invert_matcher property — api_rate_limit.api_endpoint_rules.request_matcher.headers / 267c2e7caaad / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-54e22f66d187caff861e046b268c2f3b3868edd7b1965ba3223252be353044cb): complete subsection reference.

<a id="canonical-25dc748b24802b179e58e146e0548434b58bbed04350b1ad55420dcf48cdd59b"></a>

<a id="canonical-b3fb287787809f78e35d74d4adc2c2acba874c931937f7c1bd19d8a0490f1f17"></a>

## name property — api_rate_limit.api_endpoint_rules.request_matcher.headers / 267c2e7caaad / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```
