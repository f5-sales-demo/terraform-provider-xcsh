---
page_title: "xcsh_global_log_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver examples."
---

# xcsh_global_log_receiver examples

<a id="canonical-cbd6e8041fc1a2b6d29872815c4d2cbae71d5823e5f69ef6270a7d868231f463"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35f51e9d114a0d431558e8617dd30b31801075abbfe6bbfb3dc26d7a50c71686"></a>

## Examples — Examples / 02a6b54c73eb / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- Examples

<a id="canonical-cc755aa75425eea9619d28953347cec3fa8c8ca86f2445d80aee0a875c149f13"></a>

## Complete configurations — Examples / 02a6b54c73eb / 3

- [Data source](data-sources--global_log_receiver--examples--group-001.md#canonical-5ae2746333dde0b30823369e7f44491d31c08d84447a95107417a9d57798ff83): valid configuration.

<a id="canonical-0a0513023586e499f69d3a727536dedef79775032d87bc305b0beb9cb64092f0"></a>

## Next pages — Examples / 02a6b54c73eb / 4

- [Data source](data-sources--global_log_receiver--examples--group-001.md#canonical-5ae2746333dde0b30823369e7f44491d31c08d84447a95107417a9d57798ff83)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)

<a id="canonical-5ae2746333dde0b30823369e7f44491d31c08d84447a95107417a9d57798ff83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fef5cc780f728a0f19c93b2ad4b898312d8392e92686853fd9789ea2b27919f"></a>

## Data source — Data source / 4fc891b6b648 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
- [Examples](data-sources--global_log_receiver--examples--group-001.md#canonical-cbd6e8041fc1a2b6d29872815c4d2cbae71d5823e5f69ef6270a7d868231f463)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_global_log_receiver/data-source.tf`; digest `sha256:9f73258c5164e08d5a0687330200706260aacebe3599fa365152a1976e19940b`.

```terraform
# GlobalLogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GlobalLogReceiver by name
data "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}

output "global_log_receiver_id" {
  value = data.xcsh_global_log_receiver.example.id
}
```

<a id="canonical-91ccb2bcaa67547b5bc0ab1f4979d8eb6d8022381058f7b3e531d472ed876009"></a>

## Next pages — Data source / 4fc891b6b648 / 3

- [Examples](data-sources--global_log_receiver--examples--group-001.md#canonical-cbd6e8041fc1a2b6d29872815c4d2cbae71d5823e5f69ef6270a7d868231f463)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-bbe33d0388be4510b21f5849c51ab7aaec189b2464d934d6c0530c9b932bbb7c)
