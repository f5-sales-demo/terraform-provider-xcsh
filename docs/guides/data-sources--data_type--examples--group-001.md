---
page_title: "xcsh_data_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type examples."
---

# xcsh_data_type examples

<a id="canonical-165dbfcb731f3d0db618fbe9d30ce1e3d10aad80ef7d2255d69f075394000722"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fae531f8717eb44863100fec0307550f8ed493b3b9ad3a1c1e13e93b642d3215"></a>

## Examples — Examples / ac3f5538eaff / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- Examples

<a id="canonical-02bbbf6c8d9c6de6658c7fd92454b25f6405c61df0a6788f517949de85136f4a"></a>

## Complete configurations — Examples / ac3f5538eaff / 3

- [Data source](data-sources--data_type--examples--group-001.md#canonical-f07421d2fc5908f1338fbbcdf0be2bbe30a5be86407448d3bb6f05532522f6d0): valid configuration.

<a id="canonical-ee609e09dd36cb7b885cc8fe18d013f2a27b42813c3ff839abbb775301faeeaa"></a>

## Next pages — Examples / ac3f5538eaff / 4

- [Data source](data-sources--data_type--examples--group-001.md#canonical-f07421d2fc5908f1338fbbcdf0be2bbe30a5be86407448d3bb6f05532522f6d0)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)

<a id="canonical-f07421d2fc5908f1338fbbcdf0be2bbe30a5be86407448d3bb6f05532522f6d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd9cab16eb169c83e53b0768c4843393aac807e4a527e4c3095e6acaa52db8fc"></a>

## Data source — Data source / c2f9d40ac471 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
- [Examples](data-sources--data_type--examples--group-001.md#canonical-165dbfcb731f3d0db618fbe9d30ce1e3d10aad80ef7d2255d69f075394000722)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_data_type/data-source.tf`; digest `sha256:2ab6eae5bd6da2f1828ca8b3f7f9c368730657551b0232c70f3bd6f7dc6e3501`.

```terraform
# DataType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataType by name
data "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}

output "data_type_id" {
  value = data.xcsh_data_type.example.id
}
```

<a id="canonical-827a8dc4cec027662cf6f0da38f0e549744753e33648982f6410bd7526c11e66"></a>

## Next pages — Data source / c2f9d40ac471 / 3

- [Examples](data-sources--data_type--examples--group-001.md#canonical-165dbfcb731f3d0db618fbe9d30ce1e3d10aad80ef7d2255d69f075394000722)
- [xcsh_data_type](../data-sources/data_type.md#canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f)
