---
page_title: "xcsh_mitigated_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain examples."
---

# xcsh_mitigated_domain examples

<a id="canonical-1230120121021001-2100330101330313-0000131010233210-3233021320010013-2322121011013130-2330100102101222-0132023121130001-2233232331220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-3233303323320020-3201003222322030-3103103212101100-1321230323023131-3310220303032323-0211212333301333-1330120302113020-2012202211332310)
- Examples

<a id="canonical-0233123001300301-3130120022001212-1022020100112323-3121310211133121-2221322121101030-3123132323022302-0213001321011301-2112001231022310"></a>

### Complete configurations for `xcsh_mitigated_domain`

- [Data source](data-sources--mitigated_domain--examples--group-001.md#canonical-3313030103121120-3301201332001313-0103102033021113-3202232130333030-1211001200203032-1130030301302313-3303232310100111-0003021032202212): valid configuration.

<a id="canonical-3313030103121120-3301201332001313-0103102033021113-3202232130333030-1211001200203032-1130030301302313-3303232310100111-0003021032202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-3233303323320020-3201003222322030-3103103212101100-1321230323023131-3310220303032323-0211212333301333-1330120302113020-2012202211332310)
- [Examples](data-sources--mitigated_domain--examples--group-001.md#canonical-1230120121021001-2100330101330313-0000131010233210-3233021320010013-2322121011013130-2330100102101222-0132023121130001-2233232331220210)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_mitigated_domain/data-source.tf`; digest `sha256:bc0c9edeff4c50185f9f5d99bbe4c09ceec7f010ca7ea7432a8786d884607a16`.

```terraform
# MitigatedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MitigatedDomain by name
data "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"
}

output "mitigated_domain_id" {
  value = data.xcsh_mitigated_domain.example.id
}
```
