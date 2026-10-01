---
page_title: "xcsh_access_active_sessions_terminate landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions_terminate landing."
---

# xcsh_access_active_sessions_terminate landing

<a id="canonical-aefe1df7253e779a562274f8e664a4a668b95809ac818dc69c1d5be574a4a6d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb4ba510549e1cd50d006ed59d31ae41b2e1186b921e84faf2765b7622a0be06"></a>

## xcsh_access_active_sessions_terminate — xcsh_access_active_sessions_terminate / 2ffda1b8ad5c / 2

Breadcrumbs:

- xcsh_access_active_sessions_terminate

Resource creation operation.

<a id="canonical-c5e5a3b1e0d0d0b326a82d3353fd64c33e7086c06c94c2297bad38aadedbdd39"></a>

## Prerequisites — xcsh_access_active_sessions_terminate / 2ffda1b8ad5c / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e07ba608f91b495f0d3865200bbf4e54561e50a4e5d3e4249a07936c085105cc"></a>

## Minimal configuration — xcsh_access_active_sessions_terminate / 2ffda1b8ad5c / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionsTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_sessions_terminate" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-726f98c63a6f7460aa530b1b4cbd93874c50ae819cbbee22ec106bad58d99b72"></a>

## Root configuration — xcsh_access_active_sessions_terminate / 2ffda1b8ad5c / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-fe74a1c6d5ebb78f0f8c74d0eda3dd2b5c54b014aff46b912c448dc6e1164cc7"></a>

## Next pages — xcsh_access_active_sessions_terminate / 2ffda1b8ad5c / 6

- [Property reference](../guides/actions--access_active_sessions_terminate--reference--group-001.md#canonical-45d0d6ae516ec20fd95b3489f38e0b618c5403370ae189e9bf43484e24e92d3c)
- [Examples](../guides/actions--access_active_sessions_terminate--examples--group-001.md#canonical-6f9739f9d132c7c377ebc5696894a038daf04072006a7680a723fd55539a68b6)
- [Lifecycle](../guides/actions--access_active_sessions_terminate--lifecycle--group-001.md#canonical-f6605ac137239869cf06cb89257bf1c43bdd02b0562f3ea3df9a9296a7e478e9)
