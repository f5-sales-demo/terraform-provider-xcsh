---
page_title: "xcsh_securemesh_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site landing."
---

# xcsh_securemesh_site landing

<a id="canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230202111012300-2222323031300101-0310031302313133-2021210113021233-1020022320010130-2002333011221220-1333031330320301-1233133003000301"></a>

## xcsh_securemesh_site — xcsh_securemesh_site / 032000121112 / 2

Breadcrumbs:

- xcsh_securemesh_site

Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with
distributed security.

<a id="canonical-2032103012102323-2312321323312212-2221230000021102-0021233033111032-3030220133300333-1013130113133132-1223111200332133-3002123112201310"></a>

## Prerequisites — xcsh_securemesh_site / 032000121112 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0333112222103112-3031100202001321-1020311010211032-0000112330203202-0133002013000120-3011312020203131-2013002233122113-1321110100323131"></a>

## Minimal configuration — xcsh_securemesh_site / 032000121112 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSite by name
data "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"
}

output "securemesh_site_id" {
  value = data.xcsh_securemesh_site.example.id
}
```

<a id="canonical-2231333313212220-1003203301020123-3302223103031223-3212312011100301-0221002211010030-1122123223113012-2320110333033102-3121203311330222"></a>

## Root configuration — xcsh_securemesh_site / 032000121112 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0303202331000010-2311002201220120-1030103331330030-3221111021322303-2313033033310220-0330300321233333-2300011332020120-2123332320132131"></a>

## Next pages — xcsh_securemesh_site / 032000121112 / 6

- [Property reference](../guides/data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [Examples](../guides/data-sources--securemesh_site--examples--group-001.md#canonical-3221333130330333-2121202112311131-0231203330030212-0030012021000200-1303232120310123-1223201020200123-0233303123030322-3212220303223132)
