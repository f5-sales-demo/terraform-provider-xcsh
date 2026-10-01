---
page_title: "xcsh_access_active_session landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session landing."
---

# xcsh_access_active_session landing

<a id="canonical-1000030313032121-0103232201301200-2011330102102012-2203212012120203-2301010031212003-1303323312131012-1222202312111300-1032323113103023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113133332312001-3010230002321113-3213210013012230-3232033213232222-2022312333020021-1032203330101002-3032110332300230-3222202311101103"></a>

## xcsh_access_active_session — xcsh_access_active_session / 310110022023 / 2

Breadcrumbs:

- xcsh_access_active_session

Resource retrieval operation.

<a id="canonical-1133332011210013-0003300011321110-0211322202310001-1320201012211312-2210012103223233-2330303112331133-3123303102323222-1313201003320301"></a>

## Prerequisites — xcsh_access_active_session / 310110022023 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0001202111230231-0222221132232010-1201030003223220-3002232213200123-0031102022322103-1122013112122132-3320020210303231-2213130323132003"></a>

## Minimal configuration — xcsh_access_active_session / 310110022023 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSession DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_session" "example" {
  id        = "example-value"
  namespace = "example-value"
}

output "access_active_session_result" {
  value = data.xcsh_access_active_session.example
}
```

<a id="canonical-0222031110132211-3203022132013002-3120231131303131-0303332113203032-2031202122103121-2022322310102210-1132203322230331-3121211000332230"></a>

## Root configuration — xcsh_access_active_session / 310110022023 / 5

Required root properties: `id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0232322323021222-0313323102023022-2020001110302313-2101103311300123-3311103001021012-0012312001132011-3100122320302000-3203110203112310"></a>

## Next pages — xcsh_access_active_session / 310110022023 / 6

- [Property reference](../guides/data-sources--access_active_session--reference--group-001.md#canonical-3321011011311001-0000003030110313-1323202031311111-3202313331200301-3100012032100300-3213321201002033-1302101113321300-3113021210022031)
- [Examples](../guides/data-sources--access_active_session--examples--group-001.md#canonical-0133223210300032-1021113231320322-2030300200002213-0122010310110213-2022102020233203-1100022112100102-0200030100333233-3121121301222132)
