---
page_title: "xcsh_protocol_policer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer examples."
---

# xcsh_protocol_policer examples

<a id="canonical-67a0ff220c34dd46fc92c2f52ddfe0d0c4c17068f875b188345d25d5061dbbab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea61298d8707c11483783afdc7628424124e283090df992ac09fa416a2c514c7"></a>

## Examples — Examples / 8f209415bac4 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- Examples

<a id="canonical-671908506a5796852ff5722cf04f14082ce25758b267d2b1d02dddf38b332fbd"></a>

## Complete configurations — Examples / 8f209415bac4 / 3

- [Data source](data-sources--protocol_policer--examples--group-001.md#canonical-ca66ff1bb85e0477dac5bdf1b2e9b3031a1112590e2eb917a172cb999e35359e): valid configuration.

<a id="canonical-0333ccdfa4e40baea74810ad1b4f8eda747b9cc5415a4f4e39f9fe83fb80c155"></a>

## Next pages — Examples / 8f209415bac4 / 4

- [Data source](data-sources--protocol_policer--examples--group-001.md#canonical-ca66ff1bb85e0477dac5bdf1b2e9b3031a1112590e2eb917a172cb999e35359e)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-ca66ff1bb85e0477dac5bdf1b2e9b3031a1112590e2eb917a172cb999e35359e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d32e158a10a55ed7110585b99c5db53953eb264a622d98ef48496e7650cf6241"></a>

## Data source — Data source / c741d3584f70 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Examples](data-sources--protocol_policer--examples--group-001.md#canonical-67a0ff220c34dd46fc92c2f52ddfe0d0c4c17068f875b188345d25d5061dbbab)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protocol_policer/data-source.tf`; digest `sha256:fc068c58763daba5264bf1b0ce09146319c11b6d3fcae995dc7232b3a411a2d4`.

```terraform
# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
```

<a id="canonical-eed3ececfed74644e03fd8bb77bfa97d0938a393bad74ab8c76e7d0ac0c081ab"></a>

## Next pages — Data source / c741d3584f70 / 3

- [Examples](data-sources--protocol_policer--examples--group-001.md#canonical-67a0ff220c34dd46fc92c2f52ddfe0d0c4c17068f875b188345d25d5061dbbab)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
