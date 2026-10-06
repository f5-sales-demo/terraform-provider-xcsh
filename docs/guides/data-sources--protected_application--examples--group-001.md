---
page_title: "xcsh_protected_application examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application examples."
---

# xcsh_protected_application examples

<a id="canonical-2201033120231331-3003011110101323-2002010000323102-3000312102110323-2131311201000033-2021230001203130-1330202023130023-2223332012123102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- Examples

<a id="canonical-1123021020112102-0103200020230003-0130230300321030-0330010003000110-2203320320231300-1322220230032112-2131331230101230-1220230223013203"></a>

### Complete configurations for `xcsh_protected_application`

- [Data source](data-sources--protected_application--examples--group-001.md#canonical-2302323202133100-1210020003010131-3221021002120333-3303321302132100-2013231130003331-1031000313023100-0032230130121211-1030222020202023): valid configuration.

<a id="canonical-2302323202133100-1210020003010131-3221021002120333-3303321302132100-2013231130003331-1031000313023100-0032230130121211-1030222020202023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Examples](data-sources--protected_application--examples--group-001.md#canonical-2201033120231331-3003011110101323-2002010000323102-3000312102110323-2131311201000033-2021230001203130-1330202023130023-2223332012123102)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protected_application/data-source.tf`; digest `sha256:a72cc964b57acaba96990190e74ffa0a1d2b62d15327b6423caa90a5e44e936b`.

```terraform
# ProtectedApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedApplication by name
data "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}

output "protected_application_id" {
  value = data.xcsh_protected_application.example.id
}
```
