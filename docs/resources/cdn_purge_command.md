---
page_title: "xcsh_cdn_purge_command landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command landing."
---

# xcsh_cdn_purge_command landing

<a id="canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101101133032332-3211203010332031-0202330101223113-0012220210311111-0333001321032233-0023302122001021-0212311220033220-2021110301230211"></a>

## xcsh_cdn_purge_command — xcsh_cdn_purge_command / 210032111330 / 2

Breadcrumbs:

- xcsh_cdn_purge_command

Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.
configuration.

<a id="canonical-3013332021111010-2122332013112012-3320030311103220-3221212201010330-1303220011110030-0302200010333021-3032323102033111-1310302231313311"></a>

## Prerequisites — xcsh_cdn_purge_command / 210032111330 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1211021320202302-3011221213103122-2312030321121033-3023231220321232-1003122112013130-1303333232203313-1012103032312013-0101010301033223"></a>

## Minimal configuration — xcsh_cdn_purge_command / 210032111330 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNPurgeCommand Resource Example
# Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNPurgeCommand configuration
resource "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}
```

<a id="canonical-2012031213031321-3123210120313223-1303320112302322-0112133330310333-2313303033302102-0033321033311023-0311100123101320-3312231312023323"></a>

## Root configuration — xcsh_cdn_purge_command / 210032111330 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1232323312303123-0103300230300122-1212001121002220-1212111231321200-2213230122030322-1133201200010222-3320122103301211-2331301113300112"></a>

## Next pages — xcsh_cdn_purge_command / 210032111330 / 6

- [Property reference](../guides/resources--cdn_purge_command--reference--group-001.md#canonical-0031133333032031-1202232013003002-3122211020121123-0122133210330011-2311331103310102-1022210232331223-3030132130110223-1122200112310302)
- [Examples](../guides/resources--cdn_purge_command--examples--group-001.md#canonical-1213321202020020-1110201131212332-3220110122310332-0101202001133221-3021211132222120-3032102012001031-2302002331310030-1110002322303332)
- [Import](../guides/resources--cdn_purge_command--lifecycle--group-001.md#canonical-3303030121222012-1321313321113120-0103303012220020-0202332011012301-2013130212202000-3020222222110312-0010202121123222-3312300302301321)
- [Timeouts](../guides/resources--cdn_purge_command--lifecycle--group-001.md#canonical-0103230121303002-2321303320211100-1223133220032230-3330333113322313-0030110012130021-1311000212233321-3300231200233332-3312033133310100)
