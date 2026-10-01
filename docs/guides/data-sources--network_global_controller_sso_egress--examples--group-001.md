---
page_title: "xcsh_network_global_controller_sso_egress examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_controller_sso_egress examples."
---

# xcsh_network_global_controller_sso_egress examples

<a id="canonical-7c98b3c6732c61ea16a62233454bc840c3ea94a3f443551d55f740e0a4eff02e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d10b2283c8302fcaa82a847c45f5cb32b8d11703f965de02880528c99249ae27"></a>

## Examples — Examples / 4e54f505e998 / 2

Breadcrumbs:

- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-36ec88256389fe8809b5331dd933d33252f115e7f05db401622b88b8bd7836c2)
- Examples

<a id="canonical-0460b9e3be499e0cea70c0ad058ad2b9c52832e60697372d0adfa3b2612342f7"></a>

## Complete configurations — Examples / 4e54f505e998 / 3

- [Data source](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-60947c7f6a9b74c0b24250fec5a945fc9ba7351cfefb4f29e998ae5a37746ca6): valid configuration.

<a id="canonical-435d5e96a3bf58b9050c8613cb836239cd44ab520163e14b23b689022fe44d94"></a>

## Next pages — Examples / 4e54f505e998 / 4

- [Data source](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-60947c7f6a9b74c0b24250fec5a945fc9ba7351cfefb4f29e998ae5a37746ca6)
- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-36ec88256389fe8809b5331dd933d33252f115e7f05db401622b88b8bd7836c2)

<a id="canonical-60947c7f6a9b74c0b24250fec5a945fc9ba7351cfefb4f29e998ae5a37746ca6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a82f8107c7e19c361bb1b951e44d61dce63f85dc37b3b8302a524ef8b74ad4a9"></a>

## Data source — Data source / 154955b5b63b / 2

Breadcrumbs:

- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-36ec88256389fe8809b5331dd933d33252f115e7f05db401622b88b8bd7836c2)
- [Examples](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-7c98b3c6732c61ea16a62233454bc840c3ea94a3f443551d55f740e0a4eff02e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_global_controller_sso_egress/data-source.tf`; digest `sha256:ec1f07213dac86d4f0226f84a49ea20a8bea75367b69b06ef9f30bb6b6ae2e00`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_controller_sso_egress" "https" {}

output "global_controller_sso_https" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_global_controller_sso_egress.https.cidr_blocks
  }
}
```

<a id="canonical-43a337459ebd855017e7bc2dbe3fc0c974e6cd9394d871e5fce021912649f388"></a>

## Next pages — Data source / 154955b5b63b / 3

- [Examples](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-7c98b3c6732c61ea16a62233454bc840c3ea94a3f443551d55f740e0a4eff02e)
- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-36ec88256389fe8809b5331dd933d33252f115e7f05db401622b88b8bd7836c2)
