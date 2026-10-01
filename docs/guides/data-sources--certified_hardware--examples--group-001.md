---
page_title: "xcsh_certified_hardware examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_certified_hardware examples."
---

# xcsh_certified_hardware examples

<a id="canonical-e51fc21fa808e165a7bb85e9397b7fac5dc3c42b5021e497581f5bd091f4a837"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-612db216b46aa37fad3ed3c6aa384f8febea95ecb4231cc08693cebc27c6adb5"></a>

## Examples — Examples / ece66d867345 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- Examples

<a id="canonical-ff42f8f10f9f522f311279afbbac340bfb5261ba8890453d27a1c915c89388c4"></a>

## Complete configurations — Examples / ece66d867345 / 3

- [Data source](data-sources--certified_hardware--examples--group-001.md#canonical-41ada9d6b9e7c22bc1349adc360019d8a2366f084f0c92a2c9e50ca2b7ce59a7): valid configuration.

<a id="canonical-197c052636304b5a998d14f7bdeadad72ffec2a97b8eb8eb4a1e8f6f6ff6e57e"></a>

## Next pages — Examples / ece66d867345 / 4

- [Data source](data-sources--certified_hardware--examples--group-001.md#canonical-41ada9d6b9e7c22bc1349adc360019d8a2366f084f0c92a2c9e50ca2b7ce59a7)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)

<a id="canonical-41ada9d6b9e7c22bc1349adc360019d8a2366f084f0c92a2c9e50ca2b7ce59a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-808abcd6cf7d4d7f1ae7d26894e0578ea0afbbea272cb7df61752d482f092167"></a>

## Data source — Data source / 0983d5fe20a1 / 2

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
- [Examples](data-sources--certified_hardware--examples--group-001.md#canonical-e51fc21fa808e165a7bb85e9397b7fac5dc3c42b5021e497581f5bd091f4a837)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certified_hardware/data-source.tf`; digest `sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba`.

```terraform
# CertifiedHardware Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertifiedHardware by name
data "xcsh_certified_hardware" "example" {
  name      = "example-certified-hardware"
  namespace = "staging"
}

output "certified_hardware_id" {
  value = data.xcsh_certified_hardware.example.id
}
```

<a id="canonical-ea067736a1630d36e0f9c6555a8e658b62311074b509467d9d924eddf916bc48"></a>

## Next pages — Data source / 0983d5fe20a1 / 3

- [Examples](data-sources--certified_hardware--examples--group-001.md#canonical-e51fc21fa808e165a7bb85e9397b7fac5dc3c42b5021e497581f5bd091f4a837)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-e340f494600d5d419c7e2255805bd0fc2bffe87429d681d8808df8cce9af569a)
