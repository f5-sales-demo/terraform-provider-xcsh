---
page_title: "xcsh_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site landing."
---

# xcsh_site landing

<a id="canonical-3103121123130212-2311101033212231-0302320211230210-2003212033020122-0110200112201132-1223021030222031-3113123130220130-2002110333013120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131211222110123-2121201321322023-1211122322201333-1031120120020210-0323231100031020-2021010112211233-3011330231021113-3120112120002233"></a>

## xcsh_site — xcsh_site / 022331033330 / 2

Breadcrumbs:

- xcsh_site

Manages a Site resource in F5 Distributed Cloud for get of site. configuration. (read-only data
source)

<a id="canonical-2231300030113231-1313303003203330-0231003022201110-3213122210333211-2011020000112223-2221323310201202-3121202100120132-2021211001121220"></a>

## Prerequisites — xcsh_site / 022331033330 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `virtual_site`.

- virtual_site: Logical grouping of physical sites

<a id="canonical-2010321332213020-2320132211122330-1233300130211212-2132332313311033-3001033103333013-2211130120220033-0313312102302333-1001213311122220"></a>

## Minimal configuration — xcsh_site / 022331033330 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Site Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Site by name
data "xcsh_site" "example" {
  name      = "example-site"
  namespace = "staging"
}

output "site_id" {
  value = data.xcsh_site.example.id
}
```

<a id="canonical-1311001313230332-0220021302302020-0122302301000102-2231133103011022-2130213233230022-3133321320011312-3231332123310103-2022130120232022"></a>

## Root configuration — xcsh_site / 022331033330 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1212130021311133-1222033110002010-1121212223212002-1312030020021333-2122102021022310-1323320202122321-2030011321330101-0110232121312301"></a>

## Next pages — xcsh_site / 022331033330 / 6

- [Property reference](../guides/data-sources--site--reference--group-001.md#canonical-3311133011300230-0321223031200101-3103113120002330-0012013133321103-2130230313302033-2301110320110320-0033300331010032-0312301003103313)
- [Examples](../guides/data-sources--site--examples--group-001.md#canonical-1131302300230133-1313031231101112-3002302221232002-2310311022331312-3010032300202333-3211222122213210-3303330020123231-2121302003130122)
