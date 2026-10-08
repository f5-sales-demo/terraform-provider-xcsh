---
page_title: "xcsh_swagger_object examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_swagger_object examples."
---

# xcsh_swagger_object examples

<a id="canonical-0310331101333000-0131302031302323-0231230320030331-0132320333022212-3100123310213201-2103312121102123-2033200122203232-3103210002013122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_swagger_object](../resources/swagger_object.md#canonical-0113122122131303-3312011120012020-2331301002012223-3033302331132230-0311201302013200-2313211231330002-3123103203213300-2123202121003133)
- Examples

<a id="canonical-1132021031003331-3001033302303112-0320002201011221-3130223113221232-1330203211212303-0221332321031310-3100020213322101-0110213021202001"></a>

### Complete configurations for `xcsh_swagger_object`

- [Resource](resources--swagger_object--examples--group-001.md#canonical-0112212111233001-3021103220033002-3101233330202101-1102101332032210-3312030123003320-2010310212111112-2032111312312230-3313132120222133): valid configuration.

<a id="canonical-0112212111233001-3021103220033002-3101233330202101-1102101332032210-3312030123003320-2010310212111112-2032111312312230-3313132120222133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_swagger_object](../resources/swagger_object.md#canonical-0113122122131303-3312011120012020-2331301002012223-3033302331132230-0311201302013200-2313211231330002-3123103203213300-2123202121003133)
- [Examples](resources--swagger_object--examples--group-001.md#canonical-0310331101333000-0131302031302323-0231230320030331-0132320333022212-3100123310213201-2103312121102123-2033200122203232-3103210002013122)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_swagger_object/resource.tf`; digest `sha256:ac6ca388f618a440a05eeef1c5d8d8a3ed3b1914604de5ad200997bcadce42a7`.

```terraform
terraform {
  required_version = ">= 1.14"
  required_providers {
    xcsh = { source = "f5-sales-demo/xcsh", version = ">= 15.3.0" }
  }
}

resource "xcsh_swagger_object" "example" {
  namespace = "demo"
  name      = "schema"
  content = jsonencode({
    openapi = "3.0.3"
    info    = { title = "Synthetic demo", version = "1" }
    paths   = {}
  })
}

output "swagger_path" {
  value = xcsh_swagger_object.example.path
}
```
