---
page_title: "xcsh_aws_tgw_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site landing."
---

# xcsh_aws_tgw_site landing

<a id="canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210011101023330-1210032021331220-3200103232102230-2232101120101103-0123202332232033-0320032110330012-0230113330022301-0201332122211211"></a>

## xcsh_aws_tgw_site — xcsh_aws_tgw_site / 323213300021 / 2

Breadcrumbs:

- xcsh_aws_tgw_site

Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS
Transit Gateway.

<a id="canonical-2320122000232023-0322100211322203-2230020333313102-1310313223000302-1020330220031032-1222013101320033-2202012033312313-0123013211032130"></a>

## Prerequisites — xcsh_aws_tgw_site / 323213300021 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3310111020111123-3231033200320023-2012221011023112-2303123022222112-2102311130111321-0122303022332003-3303032002020032-0133000022212133"></a>

## Minimal configuration — xcsh_aws_tgw_site / 323213300021 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AWSTGWSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSTGWSite by name
data "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}

output "aws_tgw_site_id" {
  value = data.xcsh_aws_tgw_site.example.id
}
```

<a id="canonical-2022212102310031-1103223112023311-0030101112111113-2313220121002021-0010221210012303-0303300113112231-3202311113311131-3320322003102110"></a>

## Root configuration — xcsh_aws_tgw_site / 323213300021 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1310303221100302-2213302302022033-1010211130111123-2033130001032311-3003001013002220-3020311311103322-2130013303021310-1333133012133333"></a>

## Next pages — xcsh_aws_tgw_site / 323213300021 / 6

- [Property reference](../guides/data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [Examples](../guides/data-sources--aws_tgw_site--examples--group-001.md#canonical-3133121303000003-2123223103111022-1103223113300032-0031121220130230-2202132310101300-1032021013320110-0330233311313032-0312022011210233)
