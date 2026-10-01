---
page_title: "xcsh_access_active_sessions_terminate reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions_terminate reference."
---

# xcsh_access_active_sessions_terminate reference

<a id="canonical-45d0d6ae516ec20fd95b3489f38e0b618c5403370ae189e9bf43484e24e92d3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2120bd7950cdf2fe97a191483888ca5b73eedef52daba9f4ccf7459b95a8674"></a>

## Property reference — Property reference / 00db6982da46 / 2

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-aefe1df7253e779a562274f8e664a4a668b95809ac818dc69c1d5be574a4a6d5)
- Property reference

<a id="canonical-c09fc95de0616e517c02d1e75b977e095142326bc5d77329d339020529c5106e"></a>

## Direct properties — Property reference / 00db6982da46 / 3

<a id="canonical-8516ed3aa852aac248eca59e500ceabef92f8dcb70f7214c07940aab73d1fe0b"></a>

<a id="canonical-dc90764c4c738de01a8053818e07349b644a5814ba3616e4d9b633d160bd9033"></a>

## ids property — Property reference / 00db6982da46 / 4

Type: `["list", "string"]`. Optional.

List of session IDs to terminate (maximum 100 per request).

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

<a id="canonical-7e74e581e3fe69ea8856f67af18b02e8f656c4c995dc982c10c8f812917062c5"></a>

<a id="canonical-ab1a89d039006f4c4dccaf3b3bd9e21f692a43741981acdcd2b6c85a6a0b5b43"></a>

## namespace property — Property reference / 00db6982da46 / 5

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request (path parameter)

<a id="canonical-228c5b566852ae54ed02153563b1bc263811eb35b0214370eb593e16634dacf2"></a>

## All schema paths — Property reference / 00db6982da46 / 6

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `ids` | [ids](actions--access_active_sessions_terminate--reference--group-001.md#canonical-8516ed3aa852aac248eca59e500ceabef92f8dcb70f7214c07940aab73d1fe0b) |
| `namespace` | [namespace](actions--access_active_sessions_terminate--reference--group-001.md#canonical-7e74e581e3fe69ea8856f67af18b02e8f656c4c995dc982c10c8f812917062c5) |

<a id="canonical-c9205f10c200ac1c0d319ca7f46a1c4b63852bcda9030bfa1c04f386c9eb251c"></a>

## Next pages — Property reference / 00db6982da46 / 7

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-aefe1df7253e779a562274f8e664a4a668b95809ac818dc69c1d5be574a4a6d5)
