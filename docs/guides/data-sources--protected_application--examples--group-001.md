---
page_title: "xcsh_protected_application examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application examples."
---

# xcsh_protected_application examples

<a id="canonical-a13d8b7dc315447b82100ed2c0d9253b9dd6100f89b018dc7c88b70babf866d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b24859213808b031cb30e4c3c103014a3e38b707aa2c3969df6c46c68b2b1e3"></a>

## Examples — Examples / 76cf40457103 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- Examples

<a id="canonical-6e543dc68b78d1c9704a6ce776e2a4d474f7d461a528cd37ba2713bf85e1d754"></a>

## Complete configurations — Examples / 76cf40457103 / 3

- [Data source](data-sources--protected_application--examples--group-001.md#canonical-b2ee27d06420311de924263ff3e7279087b5c0fd4d0372d00eb1c6654ca8888b): valid configuration.

<a id="canonical-2179ff8475579b4eff14dedb1d198c52f2510d24e36b2108b12c39aa95b612ad"></a>

## Next pages — Examples / 76cf40457103 / 4

- [Data source](data-sources--protected_application--examples--group-001.md#canonical-b2ee27d06420311de924263ff3e7279087b5c0fd4d0372d00eb1c6654ca8888b)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)

<a id="canonical-b2ee27d06420311de924263ff3e7279087b5c0fd4d0372d00eb1c6654ca8888b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c9c4d208b70b1052bcc7c777879d770da6c1879fd7f6c7bfe3845a22824b43d"></a>

## Data source — Data source / 26e55ccc6698 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
- [Examples](data-sources--protected_application--examples--group-001.md#canonical-a13d8b7dc315447b82100ed2c0d9253b9dd6100f89b018dc7c88b70babf866d2)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protected_application/data-source.tf`; digest `sha256:a72cc964b57acaba96990190e74ffa0a1d2b62d15327b6423caa90a5e44e936b`.

```terraform
# ProtectedApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedApplication by name
data "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}

output "protected_application_id" {
  value = data.xcsh_protected_application.example.id
}
```

<a id="canonical-077f22c7892504c3577369cbb304769777d9be03479bb443361dece4e4c7489a"></a>

## Next pages — Data source / 26e55ccc6698 / 3

- [Examples](data-sources--protected_application--examples--group-001.md#canonical-a13d8b7dc315447b82100ed2c0d9253b9dd6100f89b018dc7c88b70babf866d2)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-3945f996a7833227c6453f2920e3792db4dce69f7e356ae3b248dfefac25ae0a)
