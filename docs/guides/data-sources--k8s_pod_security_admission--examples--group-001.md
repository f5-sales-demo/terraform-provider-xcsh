---
page_title: "xcsh_k8s_pod_security_admission examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission examples."
---

# xcsh_k8s_pod_security_admission examples

<a id="canonical-1122201131203101-2310301330323322-1121021003103032-3131102302302102-0020121120231201-1230321232012301-2203132011330320-3013003012031300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- Examples

<a id="canonical-0111303031101322-0121212120230333-2013231132311223-3011131231310232-2201010221321002-3313110321221133-3323120020300112-2132023033020201"></a>

### Complete configurations for `xcsh_k8s_pod_security_admission`

- [Data source](data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-3221223021033111-1021032022332303-1213130222211230-0210103201101331-3011131203022111-3300212120310311-2111003011130230-0332313003012020): valid configuration.

<a id="canonical-3221223021033111-1021032022332303-1213130222211230-0210103201101331-3011131203022111-3300212120310311-2111003011130230-0332313003012020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Examples](data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-1122201131203101-2310301330323322-1121021003103032-3131102302302102-0020121120231201-1230321232012301-2203132011330320-3013003012031300)
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
