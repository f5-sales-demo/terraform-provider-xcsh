---
page_title: "xcsh_ip_prefix_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set examples."
---

# xcsh_ip_prefix_set examples

<a id="canonical-1101000022102203-0000201123101023-1320212232000013-0331231023232211-2210323020013101-0001101321233310-0201121321201012-1033323211201001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130)
- Examples

<a id="canonical-3203233222322322-1212223120132312-1130130233131130-2020032231120112-0012133003002122-2302232230132011-0011133121300300-0330210331223322"></a>

### Complete configurations for `xcsh_ip_prefix_set`

- [Data source](data-sources--ip_prefix_set--examples--group-001.md#canonical-0122032310030323-2130201101200121-3322131133110023-2131031230011120-3230303011320302-2110201030222033-1313000232102002-2200031021020000): valid configuration.

<a id="canonical-0122032310030323-2130201101200121-3322131133110023-2131031230011120-3230303011320302-2110201030222033-1313000232102002-2200031021020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_ip_prefix_set](../data-sources/ip_prefix_set.md#canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130)
- [Examples](data-sources--ip_prefix_set--examples--group-001.md#canonical-1101000022102203-0000201123101023-1320212232000013-0331231023232211-2210323020013101-0001101321233310-0201121321201012-1033323211201001)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ip_prefix_set/data-source.tf`; digest `sha256:e21d61ff8d692544ea192645c0aec9c60f1857a80969caf8df2d970c69dd061b`.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```
