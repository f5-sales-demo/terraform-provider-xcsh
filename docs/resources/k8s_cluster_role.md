---
page_title: "xcsh_k8s_cluster_role landing"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_k8s_cluster_role landing."
---

# xcsh_k8s_cluster_role landing

<a id="canonical-4e8d2df4747cdd8e8e9686cbca47d61ff08b2031a918d8da2f09af43bd55e66b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c2578ceaf5d9f0667ac9cdabef64a88c507b957c2a57bf990924bc66e65a4ef"></a>

## xcsh_k8s_cluster_role — xcsh_k8s_cluster_role / 310ebcdb2072 / 2

Breadcrumbs:

- xcsh_k8s_cluster_role

Manages k8s\_cluster\_role will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-6cccea18c43f864ea98b6fac2b46a5063d5eb979e263d37f77dfe5c990a269b4"></a>

## Prerequisites — xcsh_k8s_cluster_role / 310ebcdb2072 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-f1f383ecaef765948d6c031214bb6d0b441986900a57c43dbf1bb38d0cfefc81"></a>

## Minimal configuration — xcsh_k8s_cluster_role / 310ebcdb2072 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SClusterRole Resource Example
# Manages k8s_cluster_role will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SClusterRole configuration
resource "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}
```

<a id="canonical-91fe4c83e687303efedc31d87b6932068c8ece514434ded67fab2ef94c574d79"></a>

## Root configuration — xcsh_k8s_cluster_role / 310ebcdb2072 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-765fed01d6f84eac3c3fb649e014ff4e5631c5a2eeb95cbc0748e4af9d892f46"></a>

## Next pages — xcsh_k8s_cluster_role / 310ebcdb2072 / 6

- [Property reference](../guides/resources--k8s_cluster_role--reference--group-001.md#canonical-9f8cba3c529fa9bc1e4ba8c8f635e5cd17a64356b4e2c0920e65d8384a68de2c)
- [Examples](../guides/resources--k8s_cluster_role--examples--group-001.md#canonical-ce3a232593bab97db089ecb898a23d4a4451bfae8bbddb21d8687777b9b1d0e0)
- [Import](../guides/resources--k8s_cluster_role--lifecycle--group-001.md#canonical-75409b56538408d86ccab6f08f31d82e11a455d18c722da77bc0c922acc45ecb)
- [Timeouts](../guides/resources--k8s_cluster_role--lifecycle--group-001.md#canonical-0659617a6c0ee175253adc808b2f682c45968122c1fd2aa9c04e894e5ed49401)
