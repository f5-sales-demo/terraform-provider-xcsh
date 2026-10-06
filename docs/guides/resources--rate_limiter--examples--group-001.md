---
page_title: "xcsh_rate_limiter examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter examples."
---

# xcsh_rate_limiter examples

<a id="canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- Examples

<a id="canonical-2001202301002333-0033011031032101-0312320012213203-3330323023223322-1033200203110303-2032313011103001-2101133013100021-0120220313100001"></a>

### Complete configurations for `xcsh_rate_limiter`

- [All attributes](resources--rate_limiter--examples--group-001.md#canonical-1000020310210121-2331323211133230-2333302022102102-3031311031001311-3021313023130333-1201112123300013-0300001323323023-2010102033120223): valid configuration.

- [Resource](resources--rate_limiter--examples--group-001.md#canonical-0102113300122013-0000331021220210-3031310113230302-0321332313121100-3302222303003213-1232102011301231-0311001032213022-2011121011223331): valid configuration.

- [Token bucket](resources--rate_limiter--examples--group-001.md#canonical-2033220132230132-0210221001002001-3222321211030122-0133011322111122-0322213300321123-0023131303333333-0010230000213112-1122023222111021): valid configuration.

- [Unit](resources--rate_limiter--examples--group-001.md#canonical-1102321221013133-2032120101330231-3100101321233210-2310112222120023-2221303132222333-2110303223301330-1311203112001103-0333332101102022): valid configuration.

- [With annotations](resources--rate_limiter--examples--group-001.md#canonical-3021303320102011-3302301221013102-3031033332031213-2012311302022123-1211002100113233-3213212230301310-1322023210310203-1103303111322103): valid configuration.

- [With description](resources--rate_limiter--examples--group-001.md#canonical-1333220133232301-3332113121032023-0323031000032122-3300313121120113-2231112211313112-0311022021130130-2111133001010331-1023333002000110): valid configuration.

- [With labels](resources--rate_limiter--examples--group-001.md#canonical-3023013221303230-0321211331032021-1313211333333002-2211221132202030-1303133300320113-2012212103132020-1001222320331112-0030211031101332): valid configuration.

- [With limits](resources--rate_limiter--examples--group-001.md#canonical-3232112030210220-3003003210301013-3122120123023312-2331110222213302-2311113213213331-3103102230123203-2111132330301032-2021333232220310): valid configuration.

<a id="canonical-1000020310210121-2331323211133230-2333302022102102-3031311031001311-3021313023130333-1201112123300013-0300001323323023-2010102033120223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## All attributes example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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

<a id="canonical-0102113300122013-0000331021220210-3031310113230302-0321332313121100-3302222303003213-1232102011301231-0311001032213022-2011121011223331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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

<a id="canonical-2033220132230132-0210221001002001-3222321211030122-0133011322111122-0322213300321123-0023131303333333-0010230000213112-1122023222111021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Token bucket example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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

<a id="canonical-1102321221013133-2032120101330231-3100101321233210-2310112222120023-2221303132222333-2110303223301330-1311203112001103-0333332101102022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Unit example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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

<a id="canonical-3021303320102011-3302301221013102-3031033332031213-2012311302022123-1211002100113233-3213212230301310-1322023210310203-1103303111322103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With annotations example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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

<a id="canonical-1333220133232301-3332113121032023-0323031000032122-3300313121120113-2231112211313112-0311022021130130-2111133001010331-1023333002000110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With description example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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

<a id="canonical-3023013221303230-0321211331032021-1313211333333002-2211221132202030-1303133300320113-2012212103132020-1001222320331112-0030211031101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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

<a id="canonical-3232112030210220-3003003210301013-3122120123023312-2331110222213302-2311113213213331-3103102230123203-2111132330301032-2021333232220310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With limits example

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md#canonical-2302301223300123-0302013012212032-3211302123220130-1122111101312303-0222103110330210-2313222011121203-0230202102122103-2321110311111111)
- [Examples](resources--rate_limiter--examples--group-001.md#canonical-2132221333301321-3100101130100010-2210303202231132-3033113020211003-3021101333303131-3130110110101003-3112233113222010-3321111123122221)
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
