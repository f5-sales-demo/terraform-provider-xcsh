---
page_title: "xcsh_nginx_instance examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_instance examples."
---

# xcsh_nginx_instance examples

<a id="canonical-3301303331101130-2132231102232100-1003302031313130-1320220321332032-1231103011131003-0133223332130313-0113112313121212-3303112113112031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nginx_instance](../data-sources/nginx_instance.md#canonical-2321323302121221-3000233313231030-0313213302223310-1121122112023213-0213032333202311-2032313311010213-0230302320211113-0011100303031122)
- Examples

<a id="canonical-3012010003231200-1121201332330001-0311131232210011-1220200231323003-2220220312223023-1330200020011313-2001300302021013-3300302133130133"></a>

### Complete configurations for `xcsh_nginx_instance`

- [Data source](data-sources--nginx_instance--examples--group-001.md#canonical-3232112231331013-0121032112132013-1032301313133203-2131310322022232-0132100203131033-3213123232130311-3013310002300202-0121110320133133): valid configuration.

<a id="canonical-3232112231331013-0121032112132013-1032301313133203-2131310322022232-0132100203131033-3213123232130311-3013310002300202-0121110320133133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_nginx_instance](../data-sources/nginx_instance.md#canonical-2321323302121221-3000233313231030-0313213302223310-1121122112023213-0213032333202311-2032313311010213-0230302320211113-0011100303031122)
- [Examples](data-sources--nginx_instance--examples--group-001.md#canonical-3301303331101130-2132231102232100-1003302031313130-1320220321332032-1231103011131003-0133223332130313-0113112313121212-3303112113112031)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_instance/data-source.tf`; digest `sha256:772405375dce9fbbb8eabe6970ee7d9b3a944178189b317c587b8f2983eaf8fd`.

```terraform
# NginxInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxInstance by name
data "xcsh_nginx_instance" "example" {
  name      = "example-nginx-instance"
  namespace = "staging"
}

output "nginx_instance_id" {
  value = data.xcsh_nginx_instance.example.id
}
```
