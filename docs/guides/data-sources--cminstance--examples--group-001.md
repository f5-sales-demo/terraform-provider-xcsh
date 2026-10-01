---
page_title: "xcsh_cminstance examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance examples."
---

# xcsh_cminstance examples

<a id="canonical-a7fb1bdcf2f4a1339b3dacfc81a14bb5d9562839eee61c351ac173a8891f74e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df2ea5a1c819f05b7afd8e55fd569dc834f7560be14876b7fb465f114e2f46a5"></a>

## Examples — Examples / 9adf48a12f48 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- Examples

<a id="canonical-3e81e4e71272763aaeaa6eed97052b11813dc1a66b4ea676c10c15ab2146beda"></a>

## Complete configurations — Examples / 9adf48a12f48 / 3

- [Data source](data-sources--cminstance--examples--group-001.md#canonical-7305c52ef58c55b23de36242aafa3dd270aedf7d9a1646519918cebd8c533e85): valid configuration.

<a id="canonical-277b644051f426ad8f11eb901778eb42687fa5b99ecfb6265505a945383c88d9"></a>

## Next pages — Examples / 9adf48a12f48 / 4

- [Data source](data-sources--cminstance--examples--group-001.md#canonical-7305c52ef58c55b23de36242aafa3dd270aedf7d9a1646519918cebd8c533e85)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-7305c52ef58c55b23de36242aafa3dd270aedf7d9a1646519918cebd8c533e85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7822d72703e053670230bceee24f02465bb2b50138d0be7ba1bc1558c8e99827"></a>

## Data source — Data source / b8b9f1a64a04 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Examples](data-sources--cminstance--examples--group-001.md#canonical-a7fb1bdcf2f4a1339b3dacfc81a14bb5d9562839eee61c351ac173a8891f74e3)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cminstance/data-source.tf`; digest `sha256:833c4b202ec741b02e9ef290bb6512a37d3156dd48242949d7ad995d973193c7`.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```

<a id="canonical-6dae3ab0c81faf8aa113133524ea9da69e4725c72475a7bb94543584aa8d84d0"></a>

## Next pages — Data source / b8b9f1a64a04 / 3

- [Examples](data-sources--cminstance--examples--group-001.md#canonical-a7fb1bdcf2f4a1339b3dacfc81a14bb5d9562839eee61c351ac173a8891f74e3)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
