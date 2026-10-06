---
page_title: "xcsh_dns_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy examples."
---

# xcsh_dns_proxy examples

<a id="canonical-3113202003002233-1222122230113102-3003330131001222-3210112003232032-0312101003300232-3331100110201213-1312211131003200-1203202103332012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- Examples

<a id="canonical-3033020021210120-3311103312100103-3332132320231210-2200010210231220-2033001011010033-1203110123030000-3020322133122203-2303222323020210"></a>

### Complete configurations for `xcsh_dns_proxy`

- [Data source](data-sources--dns_proxy--examples--group-001.md#canonical-3232113333222221-2310022121033100-1232222032132013-3021132202211022-2103103301221103-2221012022020230-0332310321000001-3322132022103120): valid configuration.

<a id="canonical-3232113333222221-2310022121033100-1232222032132013-3021132202211022-2103103301221103-2221012022020230-0332310321000001-3322132022103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Examples](data-sources--dns_proxy--examples--group-001.md#canonical-3113202003002233-1222122230113102-3003330131001222-3210112003232032-0312101003300232-3331100110201213-1312211131003200-1203202103332012)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_proxy/data-source.tf`; digest `sha256:8864854e827d63992024dc1c04e3c3a45871e2f9bca0bf3baae302eae31eb760`.

```terraform
# DNSProxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSProxy by name
data "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}

output "dns_proxy_id" {
  value = data.xcsh_dns_proxy.example.id
}
```
