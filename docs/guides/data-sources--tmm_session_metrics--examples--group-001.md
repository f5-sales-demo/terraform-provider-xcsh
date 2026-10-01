---
page_title: "xcsh_tmm_session_metrics examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tmm_session_metrics examples."
---

# xcsh_tmm_session_metrics examples

<a id="canonical-6f96b71270f7e9fc6cf197d842041cff18ad64580f208f0b4e07724b8ba98add"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6ea824b6306ee9e3cc7fcd57777e864a580e59d521a2d55fca1dd4c5da63e67"></a>

## Examples — Examples / 1fd23269b4ae / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- Examples

<a id="canonical-bda03ad8647d7dd1a86154328b9efce48cc351695cdd5a71fc13729952caa0b2"></a>

## Complete configurations — Examples / 1fd23269b4ae / 3

- [Data source](data-sources--tmm_session_metrics--examples--group-001.md#canonical-898cfbfc61199cc0e8d65c3e349c5987f8ce23aaf810637970f00e48be1d946e): valid configuration.

<a id="canonical-317e04032e15578a11edd51deacc178a833bc1c86ea793ccf6c890157b14d320"></a>

## Next pages — Examples / 1fd23269b4ae / 4

- [Data source](data-sources--tmm_session_metrics--examples--group-001.md#canonical-898cfbfc61199cc0e8d65c3e349c5987f8ce23aaf810637970f00e48be1d946e)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)

<a id="canonical-898cfbfc61199cc0e8d65c3e349c5987f8ce23aaf810637970f00e48be1d946e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23394bc3f2a8bbb5976b9deb936c1122560597a3152b1efecb741c745eceab69"></a>

## Data source — Data source / 6f0aa6ed1cb6 / 2

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
- [Examples](data-sources--tmm_session_metrics--examples--group-001.md#canonical-6f96b71270f7e9fc6cf197d842041cff18ad64580f208f0b4e07724b8ba98add)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tmm_session_metrics/data-source.tf`; digest `sha256:68c10b51b9e48b65d8a384096eef45f5c9fac00e135d661c65b0101b25c350c0`.

```terraform
# TmmSessionMetrics DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_tmm_session_metrics" "example" {
  namespace = "example-value"
}

output "tmm_session_metrics_result" {
  value = data.xcsh_tmm_session_metrics.example
}
```

<a id="canonical-1a2af6ec664f3b6d7d01ca89647da408bb3247633f0343f7119377865f21396a"></a>

## Next pages — Data source / 6f0aa6ed1cb6 / 3

- [Examples](data-sources--tmm_session_metrics--examples--group-001.md#canonical-6f96b71270f7e9fc6cf197d842041cff18ad64580f208f0b4e07724b8ba98add)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573)
