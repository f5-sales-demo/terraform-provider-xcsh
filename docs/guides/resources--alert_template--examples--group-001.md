---
page_title: "xcsh_alert_template examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template examples."
---

# xcsh_alert_template examples

<a id="canonical-1332302023131123-1023300023220300-1231133301313001-2112102020003301-2112101230321132-2020110303233322-1213031323133120-1132031130303123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-2113333112122122-3132222003103103-2101331030303103-1200332320133001-1230313330112233-3120111103131333-0332111212113231-2112301302213113)
- Examples

<a id="canonical-2121333000331123-2302203330330322-2110221003122302-2300030230100023-2301010023111032-2203022222213223-3112002320011113-2213323302022120"></a>

### Complete configurations for `xcsh_alert_template`

- [Resource](resources--alert_template--examples--group-001.md#canonical-0121320133030232-0210111213113303-1023020123030013-3101102121221121-0210121030213133-0130000311123201-2131310233122023-0200133320302103): valid configuration.

<a id="canonical-0121320133030232-0210111213113303-1023020123030013-3101102121221121-0210121030213133-0130000311123201-2131310233122023-0200133320302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_alert_template](../resources/alert_template.md#canonical-2113333112122122-3132222003103103-2101331030303103-1200332320133001-1230313330112233-3120111103131333-0332111212113231-2112301302213113)
- [Examples](resources--alert_template--examples--group-001.md#canonical-1332302023131123-1023300023220300-1231133301313001-2112102020003301-2112101230321132-2020110303233322-1213031323133120-1132031130303123)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_template/resource.tf`; digest `sha256:fca951c776be93b0c4b2e24ed4f4a814383784763f4abf35cbade1403b91638c`.

```terraform
# AlertTemplate Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertTemplate configuration
resource "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"

  alert_message         = "example-value"
  alert_message_details = "example-value"
  alert_name            = "example-value"
}
```
