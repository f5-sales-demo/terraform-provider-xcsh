---
page_title: "xcsh_irule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule landing."
---

# xcsh_irule landing

<a id="canonical-0231101210120303-0011000002331110-3333032201223210-2323130033302300-2222120120302330-2301200112302321-0023231230022212-3231111021211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231013222123013-0300310022130130-0012211110020333-3103310223233222-3323231121322331-2203302333221132-3200202203000300-3333300113212331"></a>

## xcsh_irule — xcsh_irule / 002123201213 / 2

Breadcrumbs:

- xcsh_irule

Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed
Cloud.

<a id="canonical-0332303203231313-1230120123213101-1110103023100121-3012000030300113-2111201130200322-2122132100320110-1110211100332332-0333232202012321"></a>

## Prerequisites — xcsh_irule / 002123201213 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2011111331220121-2231030133200331-3000111302231030-0202231300113203-3132101313230031-3001333003013212-0023111301232121-1232231132333310"></a>

## Minimal configuration — xcsh_irule / 002123201213 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Irule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Irule by name
data "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"
}

output "irule_id" {
  value = data.xcsh_irule.example.id
}
```

<a id="canonical-3132113233113301-3313122012203201-1320100003012001-0000012122110013-2202231333232132-0231210011133021-3013201201020133-2231133200322223"></a>

## Root configuration — xcsh_irule / 002123201213 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2102232201103022-1132120233002202-2131222221230213-2121301222331112-0023132300333231-1132003100230201-3321012330030013-0200210303300023"></a>

## Next pages — xcsh_irule / 002123201213 / 6

- [Property reference](../guides/data-sources--irule--reference--group-001.md#canonical-2012220332011020-0110013123030123-1221233222021110-1212001032200110-0201033020011033-3011302200213220-2332100202322100-1123030132312010)
- [Examples](../guides/data-sources--irule--examples--group-001.md#canonical-0231321322321101-3311033202102201-0101010112300100-0023333101212212-0010201023102311-0133032231011313-2323121302103201-0202212122221131)
