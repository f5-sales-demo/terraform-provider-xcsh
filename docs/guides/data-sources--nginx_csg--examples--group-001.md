---
page_title: "xcsh_nginx_csg examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_csg examples."
---

# xcsh_nginx_csg examples

<a id="canonical-0020120112111103-1013021101021323-1310201102232320-0203313220322001-2010023313223123-3332120223110210-3302212201332330-0232321211032312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nginx_csg](../data-sources/nginx_csg.md#canonical-3010010030112133-0023203113011300-1111020333302000-2330320322101020-2330023211231000-3213122030322222-1002103132233101-0213032020002111)
- Examples

<a id="canonical-2223121110330000-0211031303300202-1002101230030011-1002201003003200-3302321132303031-3233333323021032-1301131202310222-1301021131113333"></a>

### Complete configurations for `xcsh_nginx_csg`

- [Data source](data-sources--nginx_csg--examples--group-001.md#canonical-2203211221210023-1213132030332300-3003201110311231-1031011303202200-2310313320110123-2301111023003123-2220011200302100-0113211322232220): valid configuration.

<a id="canonical-2203211221210023-1213132030332300-3003201110311231-1031011303202200-2310313320110123-2301111023003123-2220011200302100-0113211322232220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_nginx_csg](../data-sources/nginx_csg.md#canonical-3010010030112133-0023203113011300-1111020333302000-2330320322101020-2330023211231000-3213122030322222-1002103132233101-0213032020002111)
- [Examples](data-sources--nginx_csg--examples--group-001.md#canonical-0020120112111103-1013021101021323-1310201102232320-0203313220322001-2010023313223123-3332120223110210-3302212201332330-0232321211032312)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_csg/data-source.tf`; digest `sha256:c9552491d1b728d4d6e84172cf532690c2cae1b286159e0a14dfd5d2f7845eb3`.

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
