---
page_title: "xcsh_tmm_session_metrics landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tmm_session_metrics landing."
---

# xcsh_tmm_session_metrics landing

<a id="canonical-0001230110001322-3111312333111112-1132331032100110-2331112112213012-0233311131222231-0011030021132123-2021311322012101-1333301310111303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313012001123323-0031223112301120-0021201323111301-2210222020123320-3331101121201332-3201112030121213-1020333010310131-2221022232101132"></a>

## xcsh_tmm_session_metrics — xcsh_tmm_session_metrics / 213331222022 / 2

Breadcrumbs:

- xcsh_tmm_session_metrics

Resource creation operation.

<a id="canonical-1023122002213312-3130330001113020-0212100211320232-0131120103322012-0131133002020231-0320222300201200-1021022220011032-1232300220200221"></a>

## Prerequisites — xcsh_tmm_session_metrics / 213331222022 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3031210331211113-0020220020132230-3300312310211122-3001311130120212-1112010132023321-1120031322010313-3100220110120333-3310300000201013"></a>

## Minimal configuration — xcsh_tmm_session_metrics / 213331222022 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1330132221032210-0100031320231101-3321212121021003-3322303023202032-0313202230001120-1303202103102123-1203011231010313-3030222232321231"></a>

## Root configuration — xcsh_tmm_session_metrics / 213331222022 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2220102200101030-3212201021112020-3210003233113312-0013232002013123-0031311203112201-3233200311102012-1231203301232023-0222322031211100"></a>

## Next pages — xcsh_tmm_session_metrics / 213331222022 / 6

- [Property reference](../guides/data-sources--tmm_session_metrics--reference--group-001.md#canonical-1102133100011312-3223013330330031-0122231013233021-3302021011031130-2012012221022001-0001110032023221-2202110303032301-2312300333322010)
- [Examples](../guides/data-sources--tmm_session_metrics--examples--group-001.md#canonical-1233211223130102-1300331332213330-1230330121133120-1002001001303333-0120223112101120-0033020020330023-1032001313021023-2023222120223131)
