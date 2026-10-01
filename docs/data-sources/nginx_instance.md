---
page_title: "xcsh_nginx_instance landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_instance landing."
---

# xcsh_nginx_instance landing

<a id="canonical-2321323302121221-3000233313231030-0313213302223310-1121122112023213-0213032333202311-2032313311010213-0230302320211113-0011100303031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210013210202330-1111101310331201-1231113330101210-2211120203300321-3212022010220031-1310311111033120-1322232320212032-3020213230203121"></a>

## xcsh_nginx_instance — xcsh_nginx_instance / 203211121103 / 2

Breadcrumbs:

- xcsh_nginx_instance

Manages a Nginx Instance resource in F5 Distributed Cloud for get nginx instance configuration.
configuration. (read-only data source)

<a id="canonical-2221113301320101-1022130211101311-1311101020231303-2310023300232033-1301311012000321-0233121021012200-3022312030030301-1011101233132033"></a>

## Prerequisites — xcsh_nginx_instance / 203211121103 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1231321223210011-3321223210111331-2310011303003000-3102131313330330-2012022212132121-2100123123030230-3003202231333003-3113013122123013"></a>

## Minimal configuration — xcsh_nginx_instance / 203211121103 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxInstance by name
data "xcsh_nginx_instance" "example" {
  name      = "example-nginx-instance"
  namespace = "staging"
}

output "nginx_instance_id" {
  value = data.xcsh_nginx_instance.example.id
}
```

<a id="canonical-2213331303230331-3211111323132201-1021012033112211-2130323003332210-2021013002111312-3310003032001022-1011322122010313-1210300303111020"></a>

## Root configuration — xcsh_nginx_instance / 203211121103 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2323022133032202-2100330031311012-0022021031031011-3331303020303220-3312003301202301-3103230113000130-0130013203300101-2233322203231122"></a>

## Next pages — xcsh_nginx_instance / 203211121103 / 6

- [Property reference](../guides/data-sources--nginx_instance--reference--group-001.md#canonical-3131110011123322-2133331020011210-3000121001113220-1122300010232133-0220211102033013-2121023120220133-1222111233001023-2003132120300201)
- [Examples](../guides/data-sources--nginx_instance--examples--group-001.md#canonical-3301303331101130-2132231102232100-1003302031313130-1320220321332032-1231103011131003-0133223332130313-0113112313121212-3303112113112031)
