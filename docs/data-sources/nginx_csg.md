---
page_title: "xcsh_nginx_csg landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_csg landing."
---

# xcsh_nginx_csg landing

<a id="canonical-3010010030112133-0023203113011300-1111020333302000-2330320322101020-2330023211231000-3213122030322222-1002103132233101-0213032020002111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331023320223200-1212002232202002-0011100100003300-1121112000132131-3332222102123232-1310111203011012-0233212303332023-2132001333101011"></a>

## xcsh_nginx_csg — xcsh_nginx_csg / 102131300301 / 2

Breadcrumbs:

- xcsh_nginx_csg

Manages a Nginx Csg resource in F5 Distributed Cloud for get nginx csg configuration. configuration.
(read-only data source)

<a id="canonical-1012101120320013-2011021332021330-0222130313331302-3231321111121230-1320231131203303-0120200111113110-0002313021022322-2021233011131202"></a>

## Prerequisites — xcsh_nginx_csg / 102131300301 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3012002302210211-0132003101102332-3031200100220031-1202020000333310-2320120231033013-0023232323102333-1010111333011330-3322032030311202"></a>

## Minimal configuration — xcsh_nginx_csg / 102131300301 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxCsg Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxCsg by name
data "xcsh_nginx_csg" "example" {
  name      = "example-nginx-csg"
  namespace = "staging"
}

output "nginx_csg_id" {
  value = data.xcsh_nginx_csg.example.id
}
```

<a id="canonical-3120121033202323-3200130311003131-3121013021102130-1331121223121002-1111310210200023-3133230001330023-3330302121110312-2232220333013203"></a>

## Root configuration — xcsh_nginx_csg / 102131300301 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0011130030213012-0111102103022232-2120230221231113-0013122330001001-3222200211031201-1313212032300312-2123102222333320-0130202313333023"></a>

## Next pages — xcsh_nginx_csg / 102131300301 / 6

- [Property reference](../guides/data-sources--nginx_csg--reference--group-001.md#canonical-3110131130130312-1330331133222230-1102230130033011-1120202332122213-1330122020313232-2111301013301332-1220122313011023-2213323101223000)
- [Examples](../guides/data-sources--nginx_csg--examples--group-001.md#canonical-0020120112111103-1013021101021323-1310201102232320-0203313220322001-2010023313223123-3332120223110210-3302212201332330-0232321211032312)
