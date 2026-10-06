---
page_title: "xcsh_k8s_pod_security_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy examples."
---

# xcsh_k8s_pod_security_policy examples

<a id="canonical-3333322100323002-0130321130311000-2221323221320232-2220202010130003-1033331121221112-2023311231013122-1000003231033210-3020133031320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- Examples

<a id="canonical-1113203303313011-3333013232330130-0000313213031130-2031302120023123-2323233230132313-2322221110232321-2122333003012313-0031001001012301"></a>

### Complete configurations for `xcsh_k8s_pod_security_policy`

- [Resource](resources--k8s_pod_security_policy--examples--group-001.md#canonical-0001121231032000-3333100002011011-0332202000130110-3220332300212333-3000330112021320-3022220321303212-2321201323000331-3113010032023233): valid configuration.

<a id="canonical-0001121231032000-3333100002011011-0332202000130110-3220332300212333-3000330112021320-3022220321303212-2321201323000331-3113010032023233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../resources/k8s_pod_security_policy.md#canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112)
- [Examples](resources--k8s_pod_security_policy--examples--group-001.md#canonical-3333322100323002-0130321130311000-2221323221320232-2220202010130003-1033331121221112-2023311231013122-1000003231033210-3020133031320020)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_pod_security_policy/resource.tf`; digest `sha256:e574c58014b7e9d4bf5b442007d145c5a6006853ecb4a1aebacdb4b005ac5fa9`.

```terraform
# K8SPodSecurityPolicy Resource Example
# Manages k8s_pod_security_policy will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityPolicy configuration
resource "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}
```
