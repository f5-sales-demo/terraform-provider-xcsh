---
page_title: "xcsh_fast_acl examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl examples."
---

# xcsh_fast_acl examples

<a id="canonical-6f43747349ebe04c37600937c82d7383a504fe9eb260f386303fa0fec99920fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33bc5bcdb31fb5a7578d3d98baa16f0c85e84fd3a39804d012408b34d44df1e7"></a>

## Examples — Examples / 9433dfe9b9ca / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- Examples

<a id="canonical-be1b6157898c092c69e0cbd7ff13fd62cdd7b73ed66577128fd40af5264057b9"></a>

## Complete configurations — Examples / 9433dfe9b9ca / 3

- [Data source](data-sources--fast_acl--examples--group-001.md#canonical-cfc985905d9c7aab711e6790514afc4ad80dc00d61cf555a42fa38f9bbe5836a): valid configuration.

<a id="canonical-69f7e4754744d4adfcdafd4f075b496cef406aeb6bcff09ec11754729a94e514"></a>

## Next pages — Examples / 9433dfe9b9ca / 4

- [Data source](data-sources--fast_acl--examples--group-001.md#canonical-cfc985905d9c7aab711e6790514afc4ad80dc00d61cf555a42fa38f9bbe5836a)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)

<a id="canonical-cfc985905d9c7aab711e6790514afc4ad80dc00d61cf555a42fa38f9bbe5836a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b794b730a806e05118570a8980d36a4b0873a43e2e82bea1cf1a71f5b4c0034c"></a>

## Data source — Data source / b3d7867eba60 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
- [Examples](data-sources--fast_acl--examples--group-001.md#canonical-6f43747349ebe04c37600937c82d7383a504fe9eb260f386303fa0fec99920fb)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fast_acl/data-source.tf`; digest `sha256:8fad5fdbb88c4dc30f2314eb3a4717e4d1278f83525eb02233ddf9ff899b963a`.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```

<a id="canonical-c542c881e05651c07872428082f28960a77fc73612d4d2031cd83d76c9f01d46"></a>

## Next pages — Data source / b3d7867eba60 / 3

- [Examples](data-sources--fast_acl--examples--group-001.md#canonical-6f43747349ebe04c37600937c82d7383a504fe9eb260f386303fa0fec99920fb)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-889642fbc1d44571aa2571a7f1beb95cf72d830b6f043b2d78e83cf7f144a4a4)
