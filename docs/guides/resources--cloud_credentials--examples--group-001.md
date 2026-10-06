---
page_title: "xcsh_cloud_credentials examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials examples."
---

# xcsh_cloud_credentials examples

<a id="canonical-3120020011033110-3030221321302133-3001201311303003-0200222331003321-3311303322032013-0001302102303201-3100011302032313-0333323311023021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- Examples

<a id="canonical-1203300113010233-1111102133120030-3023103110233333-1020223122233300-2233101013002122-2220221220313223-2222302011033133-2322011201213323"></a>

### Complete configurations for `xcsh_cloud_credentials`

- [Resource](resources--cloud_credentials--examples--group-001.md#canonical-1311302123221221-1320030133321200-2312101213223331-1230210011313132-1110231031310201-0133030030003102-2212201012002100-0200233100032300): valid configuration.

<a id="canonical-1311302123221221-1320030133321200-2312101213223331-1230210011313132-1110231031310201-0133030030003102-2212201012002100-0200233100032300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md#canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120)
- [Examples](resources--cloud_credentials--examples--group-001.md#canonical-3120020011033110-3030221321302133-3001201311303003-0200222331003321-3311303322032013-0001302102303201-3100011302032313-0333323311023021)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_credentials/resource.tf`; digest `sha256:59ca9345886180c4c5420989ec1dac0bd8bd93a36d37acaaca6a81918a8c91f1`.

```terraform
# CloudCredentials Resource Example
# Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudCredentials configuration
resource "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}
```
