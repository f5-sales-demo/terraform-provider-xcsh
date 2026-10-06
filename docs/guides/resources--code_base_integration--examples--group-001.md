---
page_title: "xcsh_code_base_integration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_code_base_integration examples."
---

# xcsh_code_base_integration examples

<a id="canonical-3223311312113011-0300001102331331-3332332231201211-3013102320020123-3131313102303221-3001032000113130-3330213021233132-2201232013021000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- Examples

<a id="canonical-1030012021121220-1110222131133221-3003031332213111-2323031313031103-0213023311133213-1303012003221031-2321222123232213-3321032031313300"></a>

### Complete configurations for `xcsh_code_base_integration`

- [Resource](resources--code_base_integration--examples--group-001.md#canonical-0302131003030332-2210232300310230-2102122221132322-3031133031321130-3332332002121301-0322313030311313-3101101132020113-2300310303302110): valid configuration.

<a id="canonical-0302131003030332-2210232300310230-2102122221132322-3031133031321130-3332332002121301-0322313030311313-3101101132020113-2300310303302110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md#canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002)
- [Examples](resources--code_base_integration--examples--group-001.md#canonical-3223311312113011-0300001102331331-3332332231201211-3013102320020123-3131313102303221-3001032000113130-3330213021233132-2201232013021000)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_code_base_integration/resource.tf`; digest `sha256:96e41779f04da03e1e24b87e34c40ee5955713fd86df36344fb3cbec779ab18f`.

```terraform
# CodeBaseIntegration Resource Example
# Manages integration details in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CodeBaseIntegration configuration
resource "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}
```
