---
page_title: "xcsh_k8s_pod_security_admission landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission landing."
---

# xcsh_k8s_pod_security_admission landing

<a id="canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06a1dc0e03bfa920db6b33abcf53ee76bf6584a4623372f5e14949e1d55c0f03"></a>

## xcsh_k8s_pod_security_admission — xcsh_k8s_pod_security_admission / da950a4438f4 / 2

Breadcrumbs:

- xcsh_k8s_pod_security_admission

Manages k8s\_pod\_security\_admission will create the object in the storage backend in F5
Distributed Cloud.

<a id="canonical-160b785f2aca27f72e3332d2e87f526fa341089da4c7bd3d5b1ad5137450b057"></a>

## Prerequisites — xcsh_k8s_pod_security_admission / da950a4438f4 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-fed4bdc8ebc2d4e10e79a47c6a910a25df0aa03bce3c6d85e94391ca58b2a311"></a>

## Minimal configuration — xcsh_k8s_pod_security_admission / da950a4438f4 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityAdmission Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityAdmission by name
data "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}

output "k8s_pod_security_admission_id" {
  value = data.xcsh_k8s_pod_security_admission.example.id
}
```

<a id="canonical-d12882eaaec18a47f0f01180cdf65493b6d91b3e534b0a038cc1a681f59ac516"></a>

## Root configuration — xcsh_k8s_pod_security_admission / da950a4438f4 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-df9466e905367202e675b0e91c9bb527550464fa6173887455be912b731b31af"></a>

## Next pages — xcsh_k8s_pod_security_admission / da950a4438f4 / 6

- [Property reference](../guides/data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-5fd005062b9320a2068df985e51783fcea0f0a0befdb9d79d80f0e2c81436219)
- [Examples](../guides/data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-5a85d8d1b4c7cefa592434cedd4b2c9208658b616ce6e1b1a3785f38c70c6370)
