---
page_title: "xcsh_api_testing landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing landing."
---

# xcsh_api_testing landing

<a id="canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001122300002330-2200001003233232-0230333213300220-2012020233301200-0031310330202230-2300032201301033-1033200132311110-3122200130220323"></a>

## xcsh_api_testing — xcsh_api_testing / 033213001120 / 2

Breadcrumbs:

- xcsh_api_testing

Manages a API Testing resource in F5 Distributed Cloud.

<a id="canonical-2210023101223203-0332211201200223-1001010010213200-1133312013230331-3333101233321321-0103200012301121-1031212110333010-2203012201113313"></a>

## Prerequisites — xcsh_api_testing / 033213001120 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0103332011301022-0013210233011130-1010200332302301-0210310231202121-0130133103212123-2333332130211102-0023133013233103-1313330222212122"></a>

## Minimal configuration — xcsh_api_testing / 033213001120 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```

<a id="canonical-1230211310321201-2002221020331112-2011122233020322-2203013301200120-3300231002031113-2111123211312202-3000311100321321-1310020131103221"></a>

## Root configuration — xcsh_api_testing / 033213001120 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2201012003020023-0111320210122012-0123203302120020-0230110232112001-2223302100331230-1131100022312133-0303010033323001-3222033011311121"></a>

## Next pages — xcsh_api_testing / 033213001120 / 6

- [Property reference](../guides/resources--api_testing--reference--group-001.md#canonical-1331011221020002-0101113010020201-3000223003122333-2221100022203120-0003123322222021-1111021213002232-3230122010113202-3110332000132333)
- [Examples](../guides/resources--api_testing--examples--group-001.md#canonical-3122233230003213-2011003033202121-2331112103320233-2321012312111213-1010101031002022-2203213121202220-1032111212002233-2230021233210122)
- [Import](../guides/resources--api_testing--lifecycle--group-001.md#canonical-1332020310032312-1023111302320331-2210032310200113-0103202212210300-0222230213012020-2322222331000203-3330033323010002-2110023100013010)
- [Timeouts](../guides/resources--api_testing--lifecycle--group-001.md#canonical-2210300101312311-3000121202121211-1320203322312213-0313033311122002-3020201201102011-3300011003232003-1121202332022020-0311010333203301)
