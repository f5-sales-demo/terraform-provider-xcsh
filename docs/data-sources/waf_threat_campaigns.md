---
page_title: "xcsh_waf_threat_campaigns landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threat_campaigns landing."
---

# xcsh_waf_threat_campaigns landing

<a id="canonical-7be59cb54e54a57b58790a6d698af9a944d99024e95e82c6c13245f1022ce4c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff47f473f2bd737b16019aff9046b3289469e69410445030988ffdfdaf2e7314"></a>

## xcsh_waf_threat_campaigns — xcsh_waf_threat_campaigns / c74bdce1806d / 2

Breadcrumbs:

- xcsh_waf_threat_campaigns

Resource retrieval operation.

<a id="canonical-b17c422ea14f7de2d06b6bf8dcd281320c07c8b30a96c8b12b2fb44834829ba4"></a>

## Prerequisites — xcsh_waf_threat_campaigns / c74bdce1806d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c18d9e03d2c454ec246e72431dfdb73553cf6260d21d9422ca1d380e987e9560"></a>

## Minimal configuration — xcsh_waf_threat_campaigns / c74bdce1806d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFThreatCampaigns DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_threat_campaigns" "example" {
}

output "waf_threat_campaigns_result" {
  value = data.xcsh_waf_threat_campaigns.example
}
```

<a id="canonical-770c9245d22b3c81d8a09100f758619f418b4a3321b17a97403938018179727a"></a>

## Root configuration — xcsh_waf_threat_campaigns / c74bdce1806d / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-5b2a19c65ddda50b2c331292fb5aa435f3ba372f016271cace83befbc45eff31"></a>

## Next pages — xcsh_waf_threat_campaigns / c74bdce1806d / 6

- [Property reference](../guides/data-sources--waf_threat_campaigns--reference--group-001.md#canonical-b6e27744b7fb85106a6f572c628de8e4a957ca299bbde4a98a1b3f663513454a)
- [Examples](../guides/data-sources--waf_threat_campaigns--examples--group-001.md#canonical-29903504357da819dc7add84f48a72f236128ef7415d2078166b99480be82358)
