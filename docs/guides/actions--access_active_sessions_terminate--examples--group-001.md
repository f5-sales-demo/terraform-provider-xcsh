---
page_title: "xcsh_access_active_sessions_terminate examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions_terminate examples."
---

# xcsh_access_active_sessions_terminate examples

<a id="canonical-6f9739f9d132c7c377ebc5696894a038daf04072006a7680a723fd55539a68b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c5eb8af6b581e4c4ca38e14b225fb10cc3ba93c27204158b0eef9e2726b66f7"></a>

## Examples — Examples / 8e4734607a3a / 2

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-aefe1df7253e779a562274f8e664a4a668b95809ac818dc69c1d5be574a4a6d5)
- Examples

<a id="canonical-638b46a81351ea69c0bd7d0a06c6494bd76045e8fe0bc0940651107d6f16afd4"></a>

## Complete configurations — Examples / 8e4734607a3a / 3

- [Action](actions--access_active_sessions_terminate--examples--group-001.md#canonical-3b2a8c0b08dd712ab011d30225746caca7f5258e8ba2b2322cf6c2b408725d75): valid configuration.

<a id="canonical-9d5573fe470cbf6143b27f72fc84ea03527d752abca670cf3b034d2e5cdff3e3"></a>

## Next pages — Examples / 8e4734607a3a / 4

- [Action](actions--access_active_sessions_terminate--examples--group-001.md#canonical-3b2a8c0b08dd712ab011d30225746caca7f5258e8ba2b2322cf6c2b408725d75)
- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-aefe1df7253e779a562274f8e664a4a668b95809ac818dc69c1d5be574a4a6d5)

<a id="canonical-3b2a8c0b08dd712ab011d30225746caca7f5258e8ba2b2322cf6c2b408725d75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30ff0d61c665c99a44124c88186dd70834f55c7bc1d5859938de0f41b83b10f9"></a>

## Action — Action / 4d8dbf98d095 / 2

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-aefe1df7253e779a562274f8e664a4a668b95809ac818dc69c1d5be574a4a6d5)
- [Examples](actions--access_active_sessions_terminate--examples--group-001.md#canonical-6f9739f9d132c7c377ebc5696894a038daf04072006a7680a723fd55539a68b6)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_access_active_sessions_terminate/action.tf`; digest `sha256:e3e7c43f13f1ace8b437c1c8885aa667a9add78ae921c28c7a137377165d7c29`.

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

<a id="canonical-1fcc69475a99780b6f77598b7587e12f1d31f508e1be67e289d1c9a2bc7d573c"></a>

## Next pages — Action / 4d8dbf98d095 / 3

- [Examples](actions--access_active_sessions_terminate--examples--group-001.md#canonical-6f9739f9d132c7c377ebc5696894a038daf04072006a7680a723fd55539a68b6)
- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-aefe1df7253e779a562274f8e664a4a668b95809ac818dc69c1d5be574a4a6d5)
