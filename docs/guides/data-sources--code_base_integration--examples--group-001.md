---
page_title: "xcsh_code_base_integration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration examples."
---

# xcsh_code_base_integration examples

<a id="canonical-0130211200321301-3120330003030130-0333331123321300-1033122113102103-3121121133011111-1000213313233233-3202013201321223-2331132333232303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- Examples

<a id="canonical-2100300201332133-1301321023222002-1100221203021303-2213202330023210-2332033323301000-0211000100232013-1021031020001333-0110232121332213"></a>

### Complete configurations for `xcsh_code_base_integration`

- [Data source](data-sources--code_base_integration--examples--group-001.md#canonical-2313002030023133-0033012021330030-1333322022112320-2122213022302232-2333110212101030-2103111100231103-0302312010003330-2103332012021212): valid configuration.

<a id="canonical-2313002030023133-0033012021330030-1333322022112320-2122213022302232-2333110212101030-2103111100231103-0302312010003330-2103332012021212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_code_base_integration](../data-sources/code_base_integration.md#canonical-3133311203100320-1222120302232013-1212300133132122-1333320201213132-3221131302130130-2102313133230121-0230321221021103-1211303123231111)
- [Examples](data-sources--code_base_integration--examples--group-001.md#canonical-0130211200321301-3120330003030130-0333331123321300-1033122113102103-3121121133011111-1000213313233233-3202013201321223-2331132333232303)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_code_base_integration/data-source.tf`; digest `sha256:170c39b82f55f5ee2a4de8f66b829b7a3cbefe8bbffc6b09199912a373cb3b88`.

```terraform
# CodeBaseIntegration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CodeBaseIntegration by name
data "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}

output "code_base_integration_id" {
  value = data.xcsh_code_base_integration.example.id
}
```
