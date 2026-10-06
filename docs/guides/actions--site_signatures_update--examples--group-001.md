---
page_title: "xcsh_site_signatures_update examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_signatures_update examples."
---

# xcsh_site_signatures_update examples

<a id="canonical-1133211000213003-1211333131332233-3010121333101232-3320302222201103-1303011220210213-1330031203320202-3303112331111132-1010002322013121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_signatures_update](../actions/site_signatures_update.md#canonical-2121133033301232-0011100131212313-0013203021322030-2021023332100333-1321010310023012-3332003100213031-1333010220333010-1320230310323133)
- Examples

<a id="canonical-0103312203133233-0203321311021012-3212302203003231-3323001330233023-3131122211101032-0113210011302330-1130101130200311-3021223001100102"></a>

### Complete configurations for `xcsh_site_signatures_update`

- [Action](actions--site_signatures_update--examples--group-001.md#canonical-3020300320201011-0002300300301300-2133123121323132-3231032221131331-1133221011312332-0201133122101132-2303110212031300-0323022321011103): valid configuration.

<a id="canonical-3020300320201011-0002300300301300-2133123121323132-3231032221131331-1133221011312332-0201133122101132-2303110212031300-0323022321011103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_site_signatures_update](../actions/site_signatures_update.md#canonical-2121133033301232-0011100131212313-0013203021322030-2021023332100333-1321010310023012-3332003100213031-1333010220333010-1320230310323133)
- [Examples](actions--site_signatures_update--examples--group-001.md#canonical-1133211000213003-1211333131332233-3010121333101232-3320302222201103-1303011220210213-1330031203320202-3303112331111132-1010002322013121)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_signatures_update/action.tf`; digest `sha256:b8fc93938ea82d6d38f39cc7c48dce93f107714b65d36ac9a46938636cbed4f9`.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```
