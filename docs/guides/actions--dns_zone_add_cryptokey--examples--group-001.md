---
page_title: "xcsh_dns_zone_add_cryptokey examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_add_cryptokey examples."
---

# xcsh_dns_zone_add_cryptokey examples

<a id="canonical-3333230033331101-1020132231121332-1311232012321312-3122211311211132-0021303011320210-3002031101120303-1302112000321320-2321202101122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-2011333313231130-3200122021220331-0010211322333131-1302203232313332-2231111130321232-3311022020312311-2322131130200322-0032011221212010)
- Examples

<a id="canonical-2110223311110322-2231002012131211-1133120211230331-3201221220032103-1002301031331010-1012032020021310-0330212231011122-0021232331123320"></a>

### Complete configurations for `xcsh_dns_zone_add_cryptokey`

- [Action](actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-1321110312233313-2202012033000310-1120121203112113-2212220301023123-0021320210303212-0123121213322322-2033231303210313-2300312122230022): valid configuration.

<a id="canonical-1321110312233313-2202012033000310-1120121203112113-2212220301023123-0021320210303212-0123121213322322-2033231303210313-2300312122230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-2011333313231130-3200122021220331-0010211322333131-1302203232313332-2231111130321232-3311022020312311-2322131130200322-0032011221212010)
- [Examples](actions--dns_zone_add_cryptokey--examples--group-001.md#canonical-3333230033331101-1020132231121332-1311232012321312-3122211311211132-0021303011320210-3002031101120303-1302112000321320-2321202101122100)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_dns_zone_add_cryptokey/action.tf`; digest `sha256:41e2e4052bac9b33dbdf74fdf0a795bf2458f8e723d6c49a082ace67a79ec803`.

```terraform
# DNSZoneAddCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_add_cryptokey" "example" {
  config {
  }
}
```
