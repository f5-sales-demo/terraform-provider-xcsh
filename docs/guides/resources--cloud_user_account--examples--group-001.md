---
page_title: "xcsh_cloud_user_account examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account examples."
---

# xcsh_cloud_user_account examples

<a id="canonical-1201010100032201-3121233112123000-2101330232023112-1331201022311202-1331232332031200-3132300021332320-3012231020101202-0111122311221301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- Examples

<a id="canonical-3213032121121200-3030321122013012-1101021033233331-0331022033212202-1311321231130221-0331023212121033-1301111012232203-0230002223122210"></a>

### Complete configurations for `xcsh_cloud_user_account`

- [Resource](resources--cloud_user_account--examples--group-001.md#canonical-3320102201013303-0323232301333013-3022133202021203-2001003031233102-1100302300311030-2013113123333033-2021303002213021-1123010131030010): valid configuration.

<a id="canonical-3320102201013303-0323232301333013-3022133202021203-2001003031233102-1100302300311030-2013113123333033-2021303002213021-1123010131030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Examples](resources--cloud_user_account--examples--group-001.md#canonical-1201010100032201-3121233112123000-2101330232023112-1331201022311202-1331232332031200-3132300021332320-3012231020101202-0111122311221301)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_user_account/resource.tf`; digest `sha256:10ccdaeaddaf605a3535a274963418653193279aea9eade41b140aeacaaa4d59`.

```terraform
# CloudUserAccount Resource Example
# Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create specifications.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudUserAccount configuration
resource "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}
```
