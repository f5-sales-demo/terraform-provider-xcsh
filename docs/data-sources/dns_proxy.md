---
page_title: "xcsh_dns_proxy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy landing."
---

# xcsh_dns_proxy landing

<a id="canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321330131023102-1003011121012231-2230222233012223-1322230100010112-2010203211112212-0001200103223131-1211322320303103-0203233333203121"></a>

## xcsh_dns_proxy — xcsh_dns_proxy / 013210132132 / 2

Breadcrumbs:

- xcsh_dns_proxy

Manages DNS Proxy in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-1031122213100303-1233220333333101-0222211320033013-1201002100123223-1021332232132333-1221223332011130-2333102213202210-1303212121230033"></a>

## Prerequisites — xcsh_dns_proxy / 013210132132 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0100332113310221-2323200211001032-3110110021022222-3310230022212302-3232111020211003-2132101310002121-0212103212212321-1011010232102320"></a>

## Minimal configuration — xcsh_dns_proxy / 013210132132 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2332111231012131-1003000103120022-3113130222321330-3011223330003031-1301022111112020-1221322203320121-1301233013020112-0230112202023303"></a>

## Root configuration — xcsh_dns_proxy / 013210132132 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1101221202223131-3302021000010310-2330033001110120-3333022120313101-1120321102213111-1321220303131312-0220331331002131-3013223112212130"></a>

## Next pages — xcsh_dns_proxy / 013210132132 / 6

- [Property reference](../guides/data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [Examples](../guides/data-sources--dns_proxy--examples--group-001.md#canonical-3113202003002233-1222122230113102-3003330131001222-3210112003232032-0312101003300232-3331100110201213-1312211131003200-1203202103332012)
