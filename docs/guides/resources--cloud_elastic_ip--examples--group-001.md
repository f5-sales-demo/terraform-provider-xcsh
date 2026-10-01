---
page_title: "xcsh_cloud_elastic_ip examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip examples."
---

# xcsh_cloud_elastic_ip examples

<a id="canonical-1d4686fe67fa6921671b8abc615eb2e9b7327b95438af7b56d449703a9ec52fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e40a5355a722162cf61802c2ce74dd0fd1399188fc43c1dbaa916e53ff235b9e"></a>

## Examples — Examples / 9e0fecd51bcb / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)
- Examples

<a id="canonical-a594c069315d645888c8a19e1495562f4c2c65ab084c32b9ef56a03c3e868596"></a>

## Complete configurations — Examples / 9e0fecd51bcb / 3

- [Resource](resources--cloud_elastic_ip--examples--group-001.md#canonical-b31d3e77368047b88db7cb2e877efcd20424504b8a3343daa738d4c571c6090a): valid configuration.

<a id="canonical-c16340e10330736f081e558f8a74d10778947f02f0668b4c2baae75e5235cd31"></a>

## Next pages — Examples / 9e0fecd51bcb / 4

- [Resource](resources--cloud_elastic_ip--examples--group-001.md#canonical-b31d3e77368047b88db7cb2e877efcd20424504b8a3343daa738d4c571c6090a)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)

<a id="canonical-b31d3e77368047b88db7cb2e877efcd20424504b8a3343daa738d4c571c6090a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4919c25400330d424f02bc90099e0dde051a4e6820eb05a4b9f3995833d698a"></a>

## Resource — Resource / 77c1507195ce / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)
- [Examples](resources--cloud_elastic_ip--examples--group-001.md#canonical-1d4686fe67fa6921671b8abc615eb2e9b7327b95438af7b56d449703a9ec52fe)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_elastic_ip/resource.tf`; digest `sha256:e8be04df02915d54850afcd099a22da22eee85dc886b468e9bf65ca321094b18`.

```terraform
# CloudElasticIP Resource Example
# Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudElasticIP configuration
resource "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"

  item_count = 1
}
```

<a id="canonical-a7b56f56023ce30885b8470df0bcf60ecff150f1b30f6fcaa3f2d905319bd587"></a>

## Next pages — Resource / 77c1507195ce / 3

- [Examples](resources--cloud_elastic_ip--examples--group-001.md#canonical-1d4686fe67fa6921671b8abc615eb2e9b7327b95438af7b56d449703a9ec52fe)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c)
