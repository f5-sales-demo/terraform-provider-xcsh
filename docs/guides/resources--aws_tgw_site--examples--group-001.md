---
page_title: "xcsh_aws_tgw_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site examples."
---

# xcsh_aws_tgw_site examples

<a id="canonical-1002300012330103-3110200303023231-1100020132112201-0011320331131032-2101201331101323-1110013001203123-3333023302021313-3333331332223123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- Examples

<a id="canonical-2130133201301210-0021223233112230-0230220313002130-2322301102203130-0103111223023100-3213113123003330-0322220032021123-2302110101131100"></a>

### Complete configurations for `xcsh_aws_tgw_site`

- [Resource](resources--aws_tgw_site--examples--group-001.md#canonical-0220101010212111-3210120100033223-1102030220100330-1303312230023132-3302112231312321-2222223122023021-3113303303302302-1102002323011301): valid configuration.

<a id="canonical-0220101010212111-3210120100033223-1102030220100330-1303312230023132-3302112231312321-2222223122023021-3113303303302302-1102002323011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Examples](resources--aws_tgw_site--examples--group-001.md#canonical-1002300012330103-3110200303023231-1100020132112201-0011320331131032-2101201331101323-1110013001203123-3333023302021313-3333331332223123)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_aws_tgw_site/resource.tf`; digest `sha256:3c51271321a366128fb5a33731b3df495e257a63396c08d4e66b33465ad7f22a`.

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
