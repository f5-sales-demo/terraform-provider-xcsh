---
page_title: "xcsh_application_profiles"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles."
---

# xcsh_application_profiles

<a id="canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_application_profiles

Reads Application Profiles information from F5 Distributed Cloud.

<a id="canonical-2131132111100302-3130203010211213-2112200101113333-1233221322012231-3103021203203031-3200200312221302-1030103102212123-3200320103022033"></a>

### Prerequisites for `xcsh_application_profiles`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3232032322233212-2332202013010022-3032133013230001-3023331022003030-1003121223101102-1201133331331023-0313002230323211-0300003200302313"></a>

### Minimal configuration for `xcsh_application_profiles`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ApplicationProfiles Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ApplicationProfiles by name
data "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}

output "application_profiles_id" {
  value = data.xcsh_application_profiles.example.id
}
```

<a id="canonical-1321022020012303-2012032313122011-2231132313200021-3100222212031312-3033010312223223-3000121023330132-2002203323211120-2121103332323323"></a>

### Root configuration for `xcsh_application_profiles`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3022311201013032-0001013100213223-1330001033210222-2103211031213320-1210221123032013-1302122121311112-1231000210201233-0003121103213000"></a>

### Explore this collection for `xcsh_application_profiles`

- [Property reference](../guides/data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [Examples](../guides/data-sources--application_profiles--examples--group-001.md#canonical-1033011113233320-1301333330113200-2100233022211032-0313022133301123-2122221113330300-0232313020112220-3023003000001232-2122011022223110)
