---
page_title: "xcsh_forward_proxy_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy examples."
---

# xcsh_forward_proxy_policy examples

<a id="canonical-3301002023232221-2123112022013300-2222032223232202-2230303310220002-1112212003003010-2213333013122132-2120321333320331-3201301212021132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- Examples

<a id="canonical-2300300332032213-2300032023212020-2023011312223222-1113222320102113-0121001033102330-3310202021310321-1322003102321230-0330322021331102"></a>

### Complete configurations for `xcsh_forward_proxy_policy`

- [Data source](data-sources--forward_proxy_policy--examples--group-001.md#canonical-2122121120310111-2113013013313300-3021323301232200-2100311203213301-3333302133221220-1003221331200212-3123221231023322-2132331023111001): valid configuration.

<a id="canonical-2122121120310111-2113013013313300-3021323301232200-2100311203213301-3333302133221220-1003221331200212-3123221231023322-2132331023111001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Examples](data-sources--forward_proxy_policy--examples--group-001.md#canonical-3301002023232221-2123112022013300-2222032223232202-2230303310220002-1112212003003010-2213333013122132-2120321333320331-3201301212021132)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_forward_proxy_policy/data-source.tf`; digest `sha256:3fff12bb997c990a591800c80ace62a7790c01337e27ba68116856af1dfaf950`.

```terraform
# ForwardProxyPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardProxyPolicy by name
data "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}

output "forward_proxy_policy_id" {
  value = data.xcsh_forward_proxy_policy.example.id
}
```
