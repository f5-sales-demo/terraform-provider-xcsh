---
page_title: "xcsh_cdn_purge_command examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command examples."
---

# xcsh_cdn_purge_command examples

<a id="canonical-9a67f832bd5c4a7e9da92f7e6628c30730c414d355cb64c2770d045a240f87ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-183e245b06eba80c8cfa2d6389fee082c8ac01a86a08bc172ea1b4438d1c2196"></a>

## Examples — Examples / da502794f0c9 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
- Examples

<a id="canonical-ed13e648e8bcbb98e4d0591fa32d28dce5b9be4e0dc0848965204f30cd4037d3"></a>

## Complete configurations — Examples / da502794f0c9 / 3

- [Data source](data-sources--cdn_purge_command--examples--group-001.md#canonical-a51f2b5b0b903592552fcc5b385c90dbbb702459af9f1a766de3fc6c0da55d0e): valid configuration.

<a id="canonical-df33c267ce53d20b3b4471cb09b71926638dcbc0dd3417c1fc661ebbdf4c2e3a"></a>

## Next pages — Examples / da502794f0c9 / 4

- [Data source](data-sources--cdn_purge_command--examples--group-001.md#canonical-a51f2b5b0b903592552fcc5b385c90dbbb702459af9f1a766de3fc6c0da55d0e)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)

<a id="canonical-a51f2b5b0b903592552fcc5b385c90dbbb702459af9f1a766de3fc6c0da55d0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f52d8e56a66f3aa68ca0f45e26b390202598175f1802fc975fc795a4f79ef066"></a>

## Data source — Data source / 06f3a2630730 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
- [Examples](data-sources--cdn_purge_command--examples--group-001.md#canonical-9a67f832bd5c4a7e9da92f7e6628c30730c414d355cb64c2770d045a240f87ee)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_purge_command/data-source.tf`; digest `sha256:25cb8095a143a10866ff868fcae4396aaa9eca15b734c77fdd0d7054bddadd8a`.

```terraform
# CDNPurgeCommand Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNPurgeCommand by name
data "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}

output "cdn_purge_command_id" {
  value = data.xcsh_cdn_purge_command.example.id
}
```

<a id="canonical-b0a505c8391a7494bfa15a14d08049c4b2d3ca3d77a0c68ce0175f7db32f8432"></a>

## Next pages — Data source / 06f3a2630730 / 3

- [Examples](data-sources--cdn_purge_command--examples--group-001.md#canonical-9a67f832bd5c4a7e9da92f7e6628c30730c414d355cb64c2770d045a240f87ee)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md#canonical-6df34c3c75523044e1456b37c5cc1e89af059047c2c92713d5005e82a8faf6a1)
