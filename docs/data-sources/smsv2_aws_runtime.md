---
page_title: "xcsh_smsv2_aws_runtime landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_smsv2_aws_runtime landing."
---

# xcsh_smsv2_aws_runtime landing

<a id="canonical-585eec57272cb79858684d22159d08a11b9cc9e8b386856fa87be603c746a674"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ee857d1b4d6b00b8a3c43dfe53afd88f905c855eaa4c11e58b2e94bd3e96fe9"></a>

## xcsh_smsv2_aws_runtime — xcsh_smsv2_aws_runtime / 84a901ff6cb1 / 2

Breadcrumbs:

- xcsh_smsv2_aws_runtime

Correlates AWS ENI identities with SMSv2 configuration, site provisioning and published
physical-link status.

<a id="canonical-62387c705f1f0f6aafcc0d78bd234e8c075cb87e8e20088e675c884cd9571533"></a>

## Prerequisites — xcsh_smsv2_aws_runtime / 84a901ff6cb1 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-dc22778c3d2100ab15d1e2d980691788320c962b260da6f6b414101bbe8649f3"></a>

## Minimal configuration — xcsh_smsv2_aws_runtime / 84a901ff6cb1 / 4

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

<a id="canonical-2459ef32957b3b129cb903f592ccbd74a3ce18f11ee7af331e8c1267a651324a"></a>

## Root configuration — xcsh_smsv2_aws_runtime / 84a901ff6cb1 / 5

Required root properties: `namespace`, `nodes`, `site`. Full root flags and choices appear in the property reference.

<a id="canonical-be7c00ffdef2e3273c34ccad308a3d62277836ad27579dc39ebc62f5486f9103"></a>

## Next pages — xcsh_smsv2_aws_runtime / 84a901ff6cb1 / 6

- [Property reference](../guides/data-sources--smsv2_aws_runtime--reference--group-001.md#canonical-5d8cf40e83f05228ca10a7a18a1f29237fffea18c626c7b13a3340addaede44e)
- [Examples](../guides/data-sources--smsv2_aws_runtime--examples--group-001.md#canonical-1f1ef6b6795670d3320b1dbfef99fe9e58367a1eb0223c900d0fe48596f3d807)
