---
page_title: "xcsh_cloud_link examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link examples."
---

# xcsh_cloud_link examples

<a id="canonical-eb8da44addd6c73094389b886d916b0452553228175d7b5178810d9b504760d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18e33118fd3863a78f38115b35704db62886e0ecf275e2b316350753872bfebd"></a>

## Examples — Examples / 3ee395a17208 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- Examples

<a id="canonical-b012e7097048d195687ee6a4d0899b1580b7b03d5f27b35cf772a78f6bb8b9d1"></a>

## Complete configurations — Examples / 3ee395a17208 / 3

- [Resource](resources--cloud_link--examples--group-001.md#canonical-44bd258358309df48714bcc8e1a93b83a57626a484fcbefdad2adb9acb60e28c): valid configuration.

<a id="canonical-519d81a1287774d9407cb887554cae3915b0cbbfc52858c851f97cf425f94c6e"></a>

## Next pages — Examples / 3ee395a17208 / 4

- [Resource](resources--cloud_link--examples--group-001.md#canonical-44bd258358309df48714bcc8e1a93b83a57626a484fcbefdad2adb9acb60e28c)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)

<a id="canonical-44bd258358309df48714bcc8e1a93b83a57626a484fcbefdad2adb9acb60e28c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac8f10d9b25f6dc196f33a4d04737fd5c482817bcccdfed9080b3e3137d2b0c3"></a>

## Resource — Resource / 2dfdf6c8228f / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
- [Examples](resources--cloud_link--examples--group-001.md#canonical-eb8da44addd6c73094389b886d916b0452553228175d7b5178810d9b504760d0)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_link/resource.tf`; digest `sha256:1be93a99c9a3175561f8ebf72f83a5762c7f54a0a05e770026c4c08e2199ba00`.

```terraform
# CloudLink Resource Example
# Manages new CloudLink with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudLink configuration
resource "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}
```

<a id="canonical-df87b5fa5a37e116380b239f043f89dd83bfaa1753a66bd97b7088303b4777ef"></a>

## Next pages — Resource / 2dfdf6c8228f / 3

- [Examples](resources--cloud_link--examples--group-001.md#canonical-eb8da44addd6c73094389b886d916b0452553228175d7b5178810d9b504760d0)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-e2bde7dd950eff273e290001dfe725260a3cdddc3ece9f8132c786c2291ffa6f)
