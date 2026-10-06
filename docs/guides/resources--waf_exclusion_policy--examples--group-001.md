---
page_title: "xcsh_waf_exclusion_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_exclusion_policy examples."
---

# xcsh_waf_exclusion_policy examples

<a id="canonical-0110122321102222-3111203320123211-3222003212013303-2333230331313031-2003120230233113-3123231103322202-1020122311021101-3011330131323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- Examples

<a id="canonical-2113031213100123-2013001110211312-0033111133300212-0231201302002021-0302133202331212-0220232121120023-1231323223302110-3303121132131110"></a>

### Complete configurations for `xcsh_waf_exclusion_policy`

- [Resource](resources--waf_exclusion_policy--examples--group-001.md#canonical-2230322233122230-1101312320013330-3203011100003232-1113232203023323-1012202020311211-1322303313201112-3213313330311311-2322010330201213): valid configuration.

<a id="canonical-2230322233122230-1101312320013330-3203011100003232-1113232203023323-1012202020311211-1322303313201112-3213313330311311-2322010330201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_waf_exclusion_policy](../resources/waf_exclusion_policy.md#canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021)
- [Examples](resources--waf_exclusion_policy--examples--group-001.md#canonical-0110122321102222-3111203320123211-3222003212013303-2333230331313031-2003120230233113-3123231103322202-1020122311021101-3011330131323000)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_waf_exclusion_policy/resource.tf`; digest `sha256:873fdc1e55067d3a5c09e53e53d5c6c7bcd1d7d8f83e917fef809f43dd1b0325`.

```terraform
# WAFExclusionPolicy Resource Example
# Manages WAF exclusion policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WAFExclusionPolicy configuration
resource "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}
```
