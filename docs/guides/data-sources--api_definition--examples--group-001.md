---
page_title: "xcsh_api_definition examples"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition examples."
---

# xcsh_api_definition examples

<a id="canonical-1033121102020132-2320020110333032-2102300020133012-3123220022131120-3112301032313011-1100003031333322-3012330310222003-3302320212221331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- Examples

<a id="canonical-1232122132030331-0232221300120121-3301122213233002-3221002210332031-1011233021103233-3233112231002020-0101002101101123-2001133031031013"></a>

### Complete configurations for `xcsh_api_definition`

- [Data source](data-sources--api_definition--examples--group-001.md#canonical-1011110003001222-0002133011213111-2121022013100123-1332320322213012-3013322122122210-0020023020120223-0230020032300213-1212100213312313): valid configuration.

<a id="canonical-1011110003001222-0002133011213111-2121022013100123-1332320322213012-3013322122122210-0020023020120223-0230020032300213-1212100213312313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- [Examples](data-sources--api_definition--examples--group-001.md#canonical-1033121102020132-2320020110333032-2102300020133012-3123220022131120-3112301032313011-1100003031333322-3012330310222003-3302320212221331)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_definition/data-source.tf`; digest `sha256:00117931ab83e7521384d3c585bd071aa24022def862d24108a8273097450108`.

```terraform
# APIDefinition Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDefinition by name
data "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}

output "api_definition_id" {
  value = data.xcsh_api_definition.example.id
}
```
