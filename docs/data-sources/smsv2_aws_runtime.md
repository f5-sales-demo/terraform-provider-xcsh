---
page_title: "xcsh_smsv2_aws_runtime landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_aws_runtime landing."
---

# xcsh_smsv2_aws_runtime landing

<a id="canonical-1120113232301113-0213023023132120-1120122010310202-0111213100202201-0123213030213220-2303201220111233-2220132332120003-3013101222121310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032322011133101-2310311223000023-2022033010033133-3211032233312020-3321001130201111-3222221030010132-1120230232211023-3103322112333221"></a>

## xcsh_smsv2_aws_runtime — xcsh_smsv2_aws_runtime / 333312302301 / 2

Breadcrumbs:

- xcsh_smsv2_aws_runtime

Correlates AWS ENI identities with SMSv2 configuration, site provisioning and published
physical-link status.

<a id="canonical-1202032013301300-1133013300331222-2233303000311320-2331020310322030-0013113023201332-2032020000202032-1213113020201030-3121111301110303"></a>

## Prerequisites — xcsh_smsv2_aws_runtime / 333312302301 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3130020213132030-0331020100002223-0111310132023121-2000122101132020-0302003021120223-0212003122123312-2310011001000123-2332201210213303"></a>

## Minimal configuration — xcsh_smsv2_aws_runtime / 333312302301 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Correlate stable logical node keys and AWS-authoritative ENI MAC addresses
# with the SMSv2 interface configuration and runtime health observed by F5 XC.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_aws_runtime" "site" {
  namespace = "system"
  site      = "example-smsv2-site"

  nodes = {
    node_0_slo = {
      node = "node-0"
      role = "slo"
      mac  = "02:00:00:00:00:10"
    }
    node_0_sli = {
      node = "node-0"
      role = "sli"
      mac  = "02:00:00:00:00:11"
    }
  }
}

output "smsv2_interfaces" {
  value = data.xcsh_smsv2_aws_runtime.site.interfaces
}

output "smsv2_healthy" {
  value = data.xcsh_smsv2_aws_runtime.site.healthy
}
```

<a id="canonical-0210112132330302-2111132303230102-2130232100033311-2102303023311310-2203303201203301-0132321322330303-0132203001021213-2212110103021022"></a>

## Root configuration — xcsh_smsv2_aws_runtime / 333312302301 / 5

Required root properties: `namespace`, `nodes`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-2332133000003333-3132330232030213-0330031030302231-0300202203311202-0213132003122231-0213111321313003-2132233012023311-1020123321010003"></a>

## Next pages — xcsh_smsv2_aws_runtime / 333312302301 / 6

- [Property reference](../guides/data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-1131203033100032-2003330011020220-3022010022132201-2022013302210203-1333333332220120-3012021230132301-0322030310002231-3122323132101032)
- [Examples](../guides/data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-0133013233122312-1321111213003103-0302002301312333-3233212133322132-1120031213220132-2300020203302100-0031003332102011-2112330331200013)
