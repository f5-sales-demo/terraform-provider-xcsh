---
page_title: "xcsh_k8s_pod_security_admission examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission examples."
---

# xcsh_k8s_pod_security_admission examples

<a id="canonical-5a85d8d1b4c7cefa592434cedd4b2c9208658b616ce6e1b1a3785f38c70c6370"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15ccd47a19998b3f87b5ed6bc576dd2ea1129e42f7539a5ffb608c169e2cf221"></a>

## Examples — Examples / dfa6d7f810da / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- Examples

<a id="canonical-d637df7862c477ee8db3581a6b568ef46742a0dddebd57e47ac2b049cbde6167"></a>

## Complete configurations — Examples / dfa6d7f810da / 3

- [Data source](data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-e9ac93d54938afb36772a96c244e147dc5763295f0998d35950c572c3edc3188): valid configuration.

<a id="canonical-5989fd05b50acf1b4b6b9879d99cf9e2f0feeb548c56df10be67137c94ba1ab7"></a>

## Next pages — Examples / dfa6d7f810da / 4

- [Data source](data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-e9ac93d54938afb36772a96c244e147dc5763295f0998d35950c572c3edc3188)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)

<a id="canonical-e9ac93d54938afb36772a96c244e147dc5763295f0998d35950c572c3edc3188"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb11299bc2134c844790bb0b5516c7d593c28914593fa98bc0bc1abfe0c8b607"></a>

## Data source — Data source / 74c4630865ed / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
- [Examples](data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-5a85d8d1b4c7cefa592434cedd4b2c9208658b616ce6e1b1a3785f38c70c6370)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_pod_security_admission/data-source.tf`; digest `sha256:d1bc83d3e3b733a106b0eb3f1736fa294f7fa3ace28ea4c63d278edef289369e`.

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

<a id="canonical-3d8658a6f3e4071f296872542d3ed5aae03b2da303040188998b4e2e48d029d3"></a>

## Next pages — Data source / 74c4630865ed / 3

- [Examples](data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-5a85d8d1b4c7cefa592434cedd4b2c9208658b616ce6e1b1a3785f38c70c6370)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-88ae5d168b412044bc8417c7ce5983a32c7859a06953fb6b8556ed0a1d095187)
