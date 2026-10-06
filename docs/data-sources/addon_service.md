---
page_title: "xcsh_addon_service"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service."
---

# xcsh_addon_service

<a id="canonical-2330122101113210-0031213211000013-0113320231112112-2023300013333011-1030201010131303-2332021031023321-1030302210322030-1132200333322121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_addon_service

Retrieves information about an F5 Distributed Cloud Addon Service.

Addon services are system-managed resources that provide additional functionality such as Bot
Defense, Client Side Defense, and other security features. This data source allows you to query
addon service details including tier requirements and activation type.

~&gt; \*\*Note:\*\* Addon services cannot be created or modified via Terraform. To activate or
subscribe to an addon service, please use the F5 Distributed Cloud Console or contact your account
team.

<a id="canonical-0222300023323320-1120303213101123-0000333102301300-0011210212103011-1123123011233123-2003021323100112-0321002211020032-2323301303120033"></a>

### Prerequisites for `xcsh_addon_service`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1300323231110133-0332002223001011-0000321103111201-3310121012312022-2322303030101123-1022030103330031-3002321220312013-3233221223003010"></a>

### Minimal configuration for `xcsh_addon_service`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```

<a id="canonical-1002011313023013-2323210222313112-1000132332033311-0313101202113302-3102232211110122-1331213233011203-1332103020102030-1022211103213202"></a>

### Root configuration for `xcsh_addon_service`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3103223022031110-0323130123112021-2231010000103322-0113211302222311-0023130033201112-3000122033332303-2123102113303013-3312110003300203"></a>

### Explore this collection for `xcsh_addon_service`

- [Property reference](../guides/data-sources--addon_service--reference--group-001.md#canonical-0101311200030200-2033233113311222-3123301302321122-2232020112101123-3323000031131011-2103311130212112-2210113200221131-3330200122001111)
- [Examples](../guides/data-sources--addon_service--examples--group-001.md#canonical-2200032121332322-2232122133010113-2030311301331333-0012110031130233-2230232102303223-1021131330013322-3112031330020320-3332012230232033)
