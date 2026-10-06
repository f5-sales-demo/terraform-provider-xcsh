---
page_title: "xcsh_data_group"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group."
---

# xcsh_data_group

<a id="canonical-2033221010300023-0310201310213312-1111123303023302-1002002103313121-0023323222202020-3323233223120112-1122030002313100-3103222111013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_data_group

Manages data group in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-1303012102001233-0320311113023113-2122020021211213-2332020030332022-3311130320333331-3102200212213002-2230021333300300-3202130021212001"></a>

### Prerequisites for `xcsh_data_group`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2013012200230231-2010023203203223-2213332123131013-1200021133103331-1331231323133123-1111210222310033-0210012102023130-1012010220012022"></a>

### Minimal configuration for `xcsh_data_group`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```

<a id="canonical-3022033000012331-1012103202013121-1122112223331201-0000300320223311-3111002312232213-2223233211330011-0320300012102332-0312312230002203"></a>

### Root configuration for `xcsh_data_group`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3311030320201211-2001210013031200-1031031120303031-2023031212030332-1012222302100112-0010322333221323-2330303223113020-0233302121022120"></a>

### Explore this collection for `xcsh_data_group`

- [Property reference](../guides/resources--data_group--reference--group-001.md#canonical-3231031010332101-1122313002132323-2311122213023231-0030023331302232-0222213302220303-1210001020031210-3202112323203130-1223123303002132)
- [Examples](../guides/resources--data_group--examples--group-001.md#canonical-0131022103021013-3130012323223332-0301313022112003-2121303301303310-1011331110202103-3203133023212013-3322012222320133-0213210000232313)
- [Import](../guides/resources--data_group--lifecycle--group-001.md#canonical-2111123311200222-3203112331122322-1220233301001222-3103133230001300-0133003221020111-2031030310131100-1331233001021222-2101103022011212)
- [Timeouts](../guides/resources--data_group--lifecycle--group-001.md#canonical-2213321113330302-0020013301200300-2001002123202303-2103313113322331-1120120300302231-0010211300321201-0013102013223130-3201302332221103)
