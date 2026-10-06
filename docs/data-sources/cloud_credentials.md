---
page_title: "xcsh_cloud_credentials"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_cloud_credentials."
---

# xcsh_cloud_credentials

<a id="canonical-2110132231330011-2002032233012103-3311330333333101-3203113311300130-0003233333122123-0313221211032211-0002232311213001-2011013113113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_cloud_credentials

Reads Cloud Credentials information from F5 Distributed Cloud.

<a id="canonical-1330203331000232-2010033311313223-3313303201023030-3020301030001330-0020103123310212-3300101122130131-0003332202301030-2212311303033223"></a>

### Prerequisites for `xcsh_cloud_credentials`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1210121131011232-1013001231222002-3201330212133013-2123232313332320-3200102122110031-3202230020220310-0210303113313013-1131011133312012"></a>

### Minimal configuration for `xcsh_cloud_credentials`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudCredentials Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudCredentials by name
data "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}

output "cloud_credentials_id" {
  value = data.xcsh_cloud_credentials.example.id
}
```

<a id="canonical-0211122030001123-1300300312122001-0120131203212221-1132231112211313-3023123012133100-1110333332231023-0301003003002123-1322200200020132"></a>

### Root configuration for `xcsh_cloud_credentials`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3023313303021132-0212200301331020-3002322022222133-3313212010230012-1111022310033100-1023133011011022-0303111000032333-1000302323102221"></a>

### Explore this collection for `xcsh_cloud_credentials`

- [Property reference](../guides/data-sources--cloud_credentials--reference--group-001.md#canonical-1222110032121311-2010213320102211-0010322021330012-3032221320121332-0111232230211002-0322333203023001-1112333021213333-0303111020330003)
- [Examples](../guides/data-sources--cloud_credentials--examples--group-001.md#canonical-1311212303012203-1120120003000323-0002010120133202-0033322333331231-0123120001201200-2223210332033123-3332013103003310-1033310232212120)
