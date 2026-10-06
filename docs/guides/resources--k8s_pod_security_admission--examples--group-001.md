---
page_title: "xcsh_k8s_pod_security_admission examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission examples."
---

# xcsh_k8s_pod_security_admission examples

<a id="canonical-3303212110021111-0310330201010013-3130321323000001-2103313102022233-1213301323300113-2112033122032222-1032310313122120-3020321110323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- Examples

<a id="canonical-3032120130111210-1230212222131012-1031303021022132-1022300112012120-3122300222102101-2303103303300020-1110301101300201-3231312312102011"></a>

### Complete configurations for `xcsh_k8s_pod_security_admission`

- [Resource](resources--k8s_pod_security_admission--examples--group-001.md#canonical-1202203123123302-0101221202003123-2133202023323322-3231332220120210-1031221233313023-1122123230121120-3313032301323000-1323230331320202): valid configuration.

<a id="canonical-1202203123123302-0101221202003123-2133202023323322-3231332220120210-1031221233313023-1122123230121120-3313032301323000-1323230331320202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Examples](resources--k8s_pod_security_admission--examples--group-001.md#canonical-3303212110021111-0310330201010013-3130321323000001-2103313102022233-1213301323300113-2112033122032222-1032310313122120-3020321110323322)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_pod_security_admission/resource.tf`; digest `sha256:55b7bf85163231f6d1cb9fd77f147ddf87ba9bdf58b1aa2d6af5bb2885858dc2`.

```terraform
# K8SPodSecurityAdmission Resource Example
# Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityAdmission configuration
resource "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}
```
