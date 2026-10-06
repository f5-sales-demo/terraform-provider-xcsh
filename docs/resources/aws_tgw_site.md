---
page_title: "xcsh_aws_tgw_site"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site."
---

# xcsh_aws_tgw_site

<a id="canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_aws_tgw_site

Manages an AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS
Transit Gateway.

<a id="canonical-0112110331323201-2220321321033221-0300301323130022-2033312232113100-2000211323030211-1020121223232311-1122221001101320-1002230213002122"></a>

### Prerequisites for `xcsh_aws_tgw_site`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0303010103311321-3200022101300213-0111130331223031-1210300032011000-0202331200101030-0220323321110221-2122230223210212-3022032323333001"></a>

### Minimal configuration for `xcsh_aws_tgw_site`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AWSTGWSite Resource Example
# Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS Transit Gateway.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AWSTGWSite configuration
resource "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}
```

<a id="canonical-2102203020320022-0012320313322232-3213312200032000-2200111100130113-3122232103120123-1103031210211103-3102032210123010-2203320133020302"></a>

### Root configuration for `xcsh_aws_tgw_site`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0220020331101313-1231320313303220-0222230122211310-3311213311302223-2120003213133033-1131310130202001-1112013112233102-2133320230203303"></a>

### Explore this collection for `xcsh_aws_tgw_site`

- [Property reference](../guides/resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [Examples](../guides/resources--aws_tgw_site--examples--group-001.md#canonical-1002300012330103-3110200303023231-1100020132112201-0011320331131032-2101201331101323-1110013001203123-3333023302021313-3333331332223123)
- [Import](../guides/resources--aws_tgw_site--lifecycle--group-001.md#canonical-0013122201011232-2031011220302322-1022103030112132-3112330301133321-3323133323022032-2111230303232202-0213302301312120-2201123031000010)
- [Timeouts](../guides/resources--aws_tgw_site--lifecycle--group-001.md#canonical-0310122022133233-3210103321301203-1133220123012010-1301332012101001-1112330301110302-3100000321212131-0132032323103301-2311032232121030)
