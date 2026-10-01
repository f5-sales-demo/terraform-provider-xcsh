---
page_title: "xcsh_rate_limiter examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter examples."
---

# xcsh_rate_limiter examples

<a id="canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-818b10bf0f14d39136e069e3fcecbafa4f8235338edc54c1917c740918a37401"></a>

## Examples — Examples / a09fffe47453 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- Examples

<a id="canonical-11f256f4986fc69273ec477caad52bf82310a284c88de999eb14123242df6918"></a>

## Complete configurations — Examples / a09fffe47453 / 3

- [All attributes](resources--rate_limiter--examples--group-001.md#canonical-40234919bdee57ecbfc8a492cdd4d075c9dcb73f6159bc073007becb8448f62b): valid configuration.

- [Resource](resources--rate_limiter--examples--group-001.md#canonical-125f068700f49a24cdd17b3239fb7650f2ab30e76e485c6d3504e9ca85645afd): valid configuration.

- [Token bucket](resources--rate_limiter--examples--group-001.md#canonical-8fa1eb1e24a41081eae6531a1f17a55a3a9f0e5b0b773fff04b009d65a2ea549): valid configuration.

- [Unit](resources--rate_limiter--examples--group-001.md#canonical-52e691df8e611f2dd0479be4b45aa60ba9cdeabf94cebc7c758d60533ff9148a): valid configuration.

- [With annotations](resources--rate_limiter--examples--group-001.md#canonical-c9cf8485f2c691d2cd3fe36786d7229b650905efe79acc747a2e4d2353cd5e93): valid configuration.

- [With description](resources--rate_limiter--examples--group-001.md#canonical-7fa1fbb1fe5d938b3b34039af0dd9617ad5a5dd63528971c957c113d4bfc2014): valid configuration.

- [With labels](resources--rate_limiter--examples--group-001.md#canonical-cb1e9cec3997d3897797ffc2a5a5e88c737f0e178699378841ab8f560c94d47e): valid configuration.

- [With limits](resources--rate_limiter--examples--group-001.md#canonical-ee58c928c30e4c47da61b2f6bd52a9f2b55e79fdd34ac6e3957bcc4e89feea34): valid configuration.

<a id="canonical-f014de8dc633e0e5f455e1372cf1325a1dd637cc86403631cbca2cabbae67fcc"></a>

## Next pages — Examples / a09fffe47453 / 4

- [All attributes](resources--rate_limiter--examples--group-001.md#canonical-40234919bdee57ecbfc8a492cdd4d075c9dcb73f6159bc073007becb8448f62b)
- [Resource](resources--rate_limiter--examples--group-001.md#canonical-125f068700f49a24cdd17b3239fb7650f2ab30e76e485c6d3504e9ca85645afd)
- [Token bucket](resources--rate_limiter--examples--group-001.md#canonical-8fa1eb1e24a41081eae6531a1f17a55a3a9f0e5b0b773fff04b009d65a2ea549)
- [Unit](resources--rate_limiter--examples--group-001.md#canonical-52e691df8e611f2dd0479be4b45aa60ba9cdeabf94cebc7c758d60533ff9148a)
- [With annotations](resources--rate_limiter--examples--group-001.md#canonical-c9cf8485f2c691d2cd3fe36786d7229b650905efe79acc747a2e4d2353cd5e93)
- [With description](resources--rate_limiter--examples--group-001.md#canonical-7fa1fbb1fe5d938b3b34039af0dd9617ad5a5dd63528971c957c113d4bfc2014)
- [With labels](resources--rate_limiter--examples--group-001.md#canonical-cb1e9cec3997d3897797ffc2a5a5e88c737f0e178699378841ab8f560c94d47e)
- [With limits](resources--rate_limiter--examples--group-001.md#canonical-ee58c928c30e4c47da61b2f6bd52a9f2b55e79fdd34ac6e3957bcc4e89feea34)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-40234919bdee57ecbfc8a492cdd4d075c9dcb73f6159bc073007becb8448f62b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38ce30e83b6facf192ff336b9ff7c722f278f00967d0f7f340f4ae02f40b549d"></a>

## All attributes — All attributes / f7affbb71eb9 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/all-attributes.tf`; digest `sha256:0fa6c0ef5d6c3a83afac2353f4434ae5a9fdbf3ec95f2509bae99373fcdd2e92`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test rate limiter with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "testing"
  }
}
```

<a id="canonical-b39f6bac2020d216b43699bf6772f130d0c484d067e0089f4cb00eb703d8c0d4"></a>

## Next pages — All attributes / f7affbb71eb9 / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-125f068700f49a24cdd17b3239fb7650f2ab30e76e485c6d3504e9ca85645afd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1516460a49e9d527e058dd81226bd0e308669d33aa8feffaaf2892da5824de85"></a>

## Resource — Resource / 1978e4eb76d7 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/resource.tf`; digest `sha256:a62d0a0f1c23b001202acba9d26c2e3e80886dc9974d49dd6b47b6c5169f8a7c`.

```terraform
# RateLimiter Resource Example
# Manages rate_limiter creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiter configuration
resource "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}
```

<a id="canonical-0292dc5d815b38688d0e1d2f48aff7757359dd94f804c6b1251c04ee402f5b80"></a>

## Next pages — Resource / 1978e4eb76d7 / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-8fa1eb1e24a41081eae6531a1f17a55a3a9f0e5b0b773fff04b009d65a2ea549"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78110940f1b879e5746fb03dd16399fec489cd7cdd97e188fff87a7211dfe3f1"></a>

## Token bucket — Token bucket / 01ee6393d045 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- Token bucket

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/token-bucket.tf`; digest `sha256:12c5b85092b1f6303b0c05874db596bd776622def34147f387a7ecad55d82778`.

```terraform
# TokenBucket — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name      = "example"
  namespace = "system"

  limits {
    total_number     = 50
    unit             = "SECOND"
    burst_multiplier = 5

    token_bucket = {}
  }
}
```

<a id="canonical-d4fc1c6cc3fb62a78114ec92e364ac9735fab5b3a38ccf33419c8d964917dfbe"></a>

## Next pages — Token bucket / 01ee6393d045 / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-52e691df8e611f2dd0479be4b45aa60ba9cdeabf94cebc7c758d60533ff9148a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-791e2feca9c114e7726f1262b2c6a9988d193f8fba5515fb4c314f358a08ba42"></a>

## Unit — Unit / 433368b44d42 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- Unit

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/unit.tf`; digest `sha256:93b1c54c089411766714ddbc2d2a6c1769c9446b2d2ef86dfb9ea72876548c98`.

```terraform
# Unit — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name      = "example"
  namespace = "system"

  limits {
    total_number     = 3
    unit             = "MINUTE"
    burst_multiplier = 2

    leaky_bucket = {}
  }
}
```

<a id="canonical-fd0189f22905b94cf80dc582a0749ca0d616431147cb40aad4a9558f657d786b"></a>

## Next pages — Unit / 433368b44d42 / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-c9cf8485f2c691d2cd3fe36786d7229b650905efe79acc747a2e4d2353cd5e93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c50f10251d3b6de7593ddb15dfb29f501bf663558df469de50fa566dace398f"></a>

## With annotations — With annotations / c7f6e79ae50a / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-annotations.tf`; digest `sha256:b278928b01f70454355f22fa0b9f5a1fb258167336f683c4e01b990a20527c97`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name      = "example"
  namespace = "system"

  annotations = {
    example-key = "example-value"
  }
}
```

<a id="canonical-ea75f1fda971f1a7e1eb59d3573ba2ab9180d9f938b02aa7a69ff61a7e066579"></a>

## Next pages — With annotations / c7f6e79ae50a / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-7fa1fbb1fe5d938b3b34039af0dd9617ad5a5dd63528971c957c113d4bfc2014"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8aae797a70cda5134e43c020ad6a09e9269c3a5666102c6499b938b3c977cc9"></a>

## With description — With description / f5a5dbf1c263 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-description.tf`; digest `sha256:87a15fe7b89798f0857db8053a2641296d700d8b857fd78fbb80b8d1c50c50a9`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name        = "example"
  namespace   = "system"
  description = "example-value"
}
```

<a id="canonical-5302747306eb2987f62123f44a6dd0e00fbf6694fb1bf831ca5608e3c103831f"></a>

## Next pages — With description / f5a5dbf1c263 / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-cb1e9cec3997d3897797ffc2a5a5e88c737f0e178699378841ab8f560c94d47e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f06f33adcd2211e07b76c3205c8f6130e8adc4c0357b932e6cb51a4ae9869fd9"></a>

## With labels — With labels / 3195234e28b1 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-labels.tf`; digest `sha256:5609348e52d35b34f6d604be53430c08f2fce483ae5ffb0505d915c66d61db51`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    example-key = "example-value"
  }
}
```

<a id="canonical-c0dc200379fca9b0d8ac42456856c58c4051966b80ea13b1bbf023ea84c15175"></a>

## Next pages — With labels / 3195234e28b1 / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)

<a id="canonical-ee58c928c30e4c47da61b2f6bd52a9f2b55e79fdd34ac6e3957bcc4e89feea34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c636551799b01393880eab1c2108ab787af28cdb169b625e34b046c8853f4a47"></a>

## With limits — With limits / 84a926c563e3 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- With limits

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-limits.tf`; digest `sha256:851a79af35fd3d7c3025e31d8945aa3a5098e4fad6d2cccc80ec89bdb883b6c4`.

```terraform
# WithLimits — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name        = "example"
  namespace   = "system"
  description = "Rate limiter with limits configuration"

  limits {
    total_number      = 100
    unit              = "MINUTE"
    burst_multiplier  = 2
    period_multiplier = 1

    leaky_bucket = {}
  }
}
```

<a id="canonical-402787179fa73660583269d266f43fdd9abd2ac214af23976794d9438138cac3"></a>

## Next pages — With limits / 84a926c563e3 / 3

- [Examples](resources--rate_limiter--examples--group-001.md#canonical-9ea7fc79d045c404a4ce2b5ecf5c8943c947fcdddc514443d6bd7a84f955b6a9)
- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-b2c6bc1b321c698ee5c9ba1c5a551db32a4d4f24b7a856632c892693b9535555)
