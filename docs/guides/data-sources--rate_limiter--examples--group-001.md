---
page_title: "xcsh_rate_limiter examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter examples."
---

# xcsh_rate_limiter examples

<a id="canonical-d2aa5143f978329a2cb5d3186781df77d0e69f043c28f96ae10250235db9d56e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66ab71747f92e9ba0af43a857eb2e5dbcb3a82e45930aefdb55d0a84675a6662"></a>

## Examples — Examples / efa370199da4 / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- Examples

<a id="canonical-a3cd98a2ec91755b03a2aecfc18afc4d9bbddc2afd89a6e686b6d6047bad55b3"></a>

## Complete configurations — Examples / efa370199da4 / 3

- [Data source](data-sources--rate_limiter--examples--group-001.md#canonical-0aea0a2a1a70412d569ea1701ae8f6150184080383103884abfaf63b80627f93): valid configuration.

<a id="canonical-c3ff3a5b397960e98e572ae70f18be350d16352813d4b4969fc8b46b1b09a2b5"></a>

## Next pages — Examples / efa370199da4 / 4

- [Data source](data-sources--rate_limiter--examples--group-001.md#canonical-0aea0a2a1a70412d569ea1701ae8f6150184080383103884abfaf63b80627f93)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)

<a id="canonical-0aea0a2a1a70412d569ea1701ae8f6150184080383103884abfaf63b80627f93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a0896eacc20d7bd1e2372a5d47285884401dfbe9ba00ff0ff5556593a46ee82"></a>

## Data source — Data source / 06acf99fc2fc / 2

Breadcrumbs:

- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
- [Examples](data-sources--rate_limiter--examples--group-001.md#canonical-d2aa5143f978329a2cb5d3186781df77d0e69f043c28f96ae10250235db9d56e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_rate_limiter/data-source.tf`; digest `sha256:6445461d09bf17fe28d100ae80d4a2ae205fe7e6d895306172e50a8556fc3fc6`.

```terraform
# RateLimiter Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiter by name
data "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}

output "rate_limiter_id" {
  value = data.xcsh_rate_limiter.example.id
}
```

<a id="canonical-3bdd95d250d72d163a7560364de6e5545d187e38e1368fb6021a2e8dddd89285"></a>

## Next pages — Data source / 06acf99fc2fc / 3

- [Examples](data-sources--rate_limiter--examples--group-001.md#canonical-d2aa5143f978329a2cb5d3186781df77d0e69f043c28f96ae10250235db9d56e)
- [xcsh_rate_limiter](../data-sources/rate_limiter.md#canonical-007f7a5d2b174631ab0f6d827b17f892fa0aa9d8bb4058855adff100f838679c)
