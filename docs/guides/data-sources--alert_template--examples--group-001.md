---
page_title: "xcsh_alert_template examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_template examples."
---

# xcsh_alert_template examples

<a id="canonical-3213301303320332-3211020211111123-0131003133300303-0130333120222113-3302133011010030-2202002310102011-0320232133100203-2233310233010231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-1003131020323200-2102322313132332-1330100113320011-3213311001122330-0122322231213130-1212212131231012-1332020002133101-0232021330332020)
- Examples

<a id="canonical-0210311231323201-1112133021222110-1023212200301033-2011133230333102-0132032331313131-3133323010221313-2103323312123212-1011321021031030"></a>

### Complete configurations for `xcsh_alert_template`

- [Data source](data-sources--alert_template--examples--group-001.md#canonical-2123331100131133-2222022211322210-3111100201232122-2320003033322022-1130001200303213-2132022130203100-1222110321223111-3100230233222001): valid configuration.

<a id="canonical-2123331100131133-2222022211322210-3111100201232122-2320003033322022-1130001200303213-2132022130203100-1222110321223111-3100230233222001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_alert_template](../data-sources/alert_template.md#canonical-1003131020323200-2102322313132332-1330100113320011-3213311001122330-0122322231213130-1212212131231012-1332020002133101-0232021330332020)
- [Examples](data-sources--alert_template--examples--group-001.md#canonical-3213301303320332-3211020211111123-0131003133300303-0130333120222113-3302133011010030-2202002310102011-0320232133100203-2233310233010231)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_template/data-source.tf`; digest `sha256:38bcb6351cb774c27b9a7119ac259c1eaff5879d005e5da296e0eb47b5f478f7`.

```terraform
# AlertTemplate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertTemplate by name
data "xcsh_alert_template" "example" {
  name      = "example-alert-template"
  namespace = "staging"
}

output "alert_template_id" {
  value = data.xcsh_alert_template.example.id
}
```
