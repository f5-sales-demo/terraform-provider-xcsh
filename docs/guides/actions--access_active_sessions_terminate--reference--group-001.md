---
page_title: "xcsh_access_active_sessions_terminate reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions_terminate reference."
---

# xcsh_access_active_sessions_terminate reference

<a id="canonical-1011310031122232-1101123230020033-3121112303102021-3303203200231201-2030111000030313-0022320120213221-2333100310201032-0210322102310330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-2232333201313313-0211033213132122-1112020213103320-3212121022102212-1220232111200021-2230200120313012-2130013111233211-1310221022123111)
- Property reference

<a id="canonical-2202010200233113-2111003031330233-3221132201210110-2003202020302211-2313033232313233-1102312223222133-1030303313101121-2321112220121310"></a>

### Direct properties for `xcsh_access_active_sessions_terminate`

<a id="canonical-2011011232310322-2220110222223002-1020323022112132-1100003032222332-3321023320313023-1300331302011030-0013211000222223-1303310133320023"></a>

#### `ids` property

Type: `["list", "string"]`. Optional.

List of session IDs to terminate (maximum 100 per request).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

<a id="canonical-1332131032112001-3203333212213222-2020111233121322-3301202300023220-3312111230103021-2111313021200230-0100302033200102-2101130012023011"></a>

<a id="canonical-3000213330211131-3200120112321101-1330000231013213-1123211313320021-1101100203021223-3011311313030221-3103032100020011-0221301101001232"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request (path parameter)

<a id="canonical-3130210013121030-1030130320313200-0122200011032001-2032001303102123-1210102211200110-2322031201123210-3121231203033101-1200233121000303"></a>

### All schema paths for `xcsh_access_active_sessions_terminate`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `ids` | [IDs](actions--access_active_sessions_terminate--reference--group-001.md#canonical-2011011232310322-2220110222223002-1020323022112132-1100003032222332-3321023320313023-1300331302011030-0013211000222223-1303310133320023) |
| `namespace` | [namespace](actions--access_active_sessions_terminate--reference--group-001.md#canonical-1332131032112001-3203333212213222-2020111233121322-3301202300023220-3312111230103021-2111313021200230-0100302033200102-2101130012023011) |
