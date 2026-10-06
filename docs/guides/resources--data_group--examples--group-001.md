---
page_title: "xcsh_data_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group examples."
---

# xcsh_data_group examples

<a id="canonical-0131022103021013-3130012323223332-0301313022112003-2121303301303310-1011331110202103-3203133023212013-3322012222320133-0213210000232313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031)
- Examples

<a id="canonical-2020320203130320-1122133333133113-3213310322112222-2120133030031111-2033001213213202-1302331023101111-0133130322220030-1000210022323230"></a>

### Complete configurations for `xcsh_data_group`

- [Resource](resources--data_group--examples--group-001.md#canonical-2203123231131002-1010333110013022-0002230320220313-2231031310211333-1211333111103021-2310133322223011-1021233301312012-0111201321002133): valid configuration.

<a id="canonical-2203123231131002-1010333110013022-0002230320220313-2231031310211333-1211333111103021-2310133322223011-1021233301312012-0111201321002133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_data_group](../resources/data_group.md#canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031)
- [Examples](resources--data_group--examples--group-001.md#canonical-0131022103021013-3130012323223332-0301313022112003-2121303301303310-1011331110202103-3203133023212013-3322012222320133-0213210000232313)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_group/resource.tf`; digest `sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690`.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```
