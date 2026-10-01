---
page_title: "xcsh_nginx_csg examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_csg examples."
---

# xcsh_nginx_csg examples

<a id="canonical-086165534725127b74852bb823de8e81842f7adbfe62b524f29a1fbc2ee653b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab654f0025373c224246c305428430e0f2e5eccdefffb24e71762d2a7125d5ff"></a>

## Examples — Examples / 8aaad82b38db / 2

Breadcrumbs:

- [xcsh_nginx_csg](../data-sources/nginx_csg.md#canonical-c410c59f0b8d71705523fc80bce3a448bc2e5b40e768ceaa424debd127388095)
- Examples

<a id="canonical-647bba058d15481a06206fc1088f563d9e533a5bbe76a6dccbc10509a1bc4841"></a>

## Complete configurations — Examples / 8aaad82b38db / 3

- [Data source](data-sources--nginx_csg--examples--group-001.md#canonical-a396990b6778cfb0c3854d6d4d1738a0b4df851bb154b0dba8160c901797aba8): valid configuration.

<a id="canonical-c2486b9a0b6aba54d1a687c5ea21bddc31248ec857b3f1c938c9faf13a458b78"></a>

## Next pages — Examples / 8aaad82b38db / 4

- [Data source](data-sources--nginx_csg--examples--group-001.md#canonical-a396990b6778cfb0c3854d6d4d1738a0b4df851bb154b0dba8160c901797aba8)
- [xcsh_nginx_csg](../data-sources/nginx_csg.md#canonical-c410c59f0b8d71705523fc80bce3a448bc2e5b40e768ceaa424debd127388095)

<a id="canonical-a396990b6778cfb0c3854d6d4d1738a0b4df851bb154b0dba8160c901797aba8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a47ec7d26a5d0d7b5c342fe517a9fc3e3e05772097248a2d88f024c2d9bebfd"></a>

## Data source — Data source / 934afd65232b / 2

Breadcrumbs:

- [xcsh_nginx_csg](../data-sources/nginx_csg.md#canonical-c410c59f0b8d71705523fc80bce3a448bc2e5b40e768ceaa424debd127388095)
- [Examples](data-sources--nginx_csg--examples--group-001.md#canonical-086165534725127b74852bb823de8e81842f7adbfe62b524f29a1fbc2ee653b6)
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

<a id="canonical-36e03d4d7d4185baaae8b1e7a0968638e50ee6efff3a172f4cf87f2c25d3cea3"></a>

## Next pages — Data source / 934afd65232b / 3

- [Examples](data-sources--nginx_csg--examples--group-001.md#canonical-086165534725127b74852bb823de8e81842f7adbfe62b524f29a1fbc2ee653b6)
- [xcsh_nginx_csg](../data-sources/nginx_csg.md#canonical-c410c59f0b8d71705523fc80bce3a448bc2e5b40e768ceaa424debd127388095)
