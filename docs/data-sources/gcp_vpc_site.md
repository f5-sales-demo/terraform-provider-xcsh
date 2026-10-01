---
page_title: "xcsh_gcp_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site landing."
---

# xcsh_gcp_vpc_site landing

<a id="canonical-1132222133321101-1232102200201102-2111112102132122-0010023221133123-2333213300031302-1001011022311003-3121103220313010-2103321200303202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302320013302232-2012123233300212-2102030203333303-3331211001030212-2010223320200031-1031232310311231-3113123232131012-3111001213233210"></a>

## xcsh_gcp_vpc_site — xcsh_gcp_vpc_site / 303132231013 / 2

Breadcrumbs:

- xcsh_gcp_vpc_site

Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud
VPC environments.

<a id="canonical-1203103233031203-2001013021213221-0220201210300303-1131123322112101-2033120222013221-3030313302100120-3323232332102100-2123123111032131"></a>

## Prerequisites — xcsh_gcp_vpc_site / 303132231013 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: GCP authentication for deployment

<a id="canonical-3010020020020322-1023301002031130-0112000223232222-0033320002100231-1223001331333110-1221320123230232-1310320032111331-3003200023111222"></a>

## Minimal configuration — xcsh_gcp_vpc_site / 303132231013 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GCPVPCSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GCPVPCSite by name
data "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"
}

output "gcp_vpc_site_id" {
  value = data.xcsh_gcp_vpc_site.example.id
}
```

<a id="canonical-0110323220321212-3323020000330233-1232303321020121-3222130032132130-1002230301021222-1213222110222113-3231012100312001-3011123120202020"></a>

## Root configuration — xcsh_gcp_vpc_site / 303132231013 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0133301013010133-3301330003213222-2022313102310103-1223313233322333-2331012000103021-2011112301112033-2002021102131332-3112233000133123"></a>

## Next pages — xcsh_gcp_vpc_site / 303132231013 / 6

- [Property reference](../guides/data-sources--gcp_vpc_site--reference--group-001.md#canonical-0023231003100101-1010210121112300-0303103132302303-0331000323201313-3003311210331022-3000312312212303-0030202320102022-1232013113103201)
- [Examples](../guides/data-sources--gcp_vpc_site--examples--group-001.md#canonical-1123010130110323-3232133001133023-1103301112123013-1111321332121122-1123230100031132-0222323133321122-1103110122123312-2031022003220131)
