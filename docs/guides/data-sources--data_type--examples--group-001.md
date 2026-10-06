---
page_title: "xcsh_data_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type examples."
---

# xcsh_data_type examples

<a id="canonical-0112113123333023-1303013303310031-2312012033233221-3103003032013203-3101002222312000-3233133102021111-3112213300131103-2110000000130202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- Examples

<a id="canonical-3322321103013320-1301133223101020-1203010000333230-0003001311110033-2032311021032303-2321223103220130-0132010332210323-1210023103020111"></a>

### Complete configurations for `xcsh_data_type`

- [Data source](data-sources--data_type--examples--group-001.md#canonical-3300131002013102-3330112100203301-0303203323233031-3300233202232332-0300221123322012-1000131010203103-2323123300111103-0211020233123100): valid configuration.

<a id="canonical-3300131002013102-3330112100203301-0303203323233031-3300233202232332-0300221123322012-1000131010203103-2323123300111103-0211020233123100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Examples](data-sources--data_type--examples--group-001.md#canonical-0112113123333023-1303013303310031-2312012033233221-3103003032013203-3101002222312000-3233133102021111-3112213300131103-2110000000130202)
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
