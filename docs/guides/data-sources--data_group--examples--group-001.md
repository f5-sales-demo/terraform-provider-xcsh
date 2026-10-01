---
page_title: "xcsh_data_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group examples."
---

# xcsh_data_group examples

<a id="canonical-30577a5df82559f5bd7a5798e84dcf0a6240212b5fc1fd6cbfe37d0f6b76bb0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4996792ca6842c99972512b97ad4ef3938611e8f308e25ee56a4598040a36474"></a>

## Examples — Examples / 81c1ac878ff7 / 2

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
- Examples

<a id="canonical-d24d7ff779c2244bb68a42b20a4008673efd87597599a04acc61380139f4dfc5"></a>

## Complete configurations — Examples / 81c1ac878ff7 / 3

- [Data source](data-sources--data_group--examples--group-001.md#canonical-58926b899ebbad90a654ba69cfaafde688e5ec840aa5a18bcf0c56b3cc46af02): valid configuration.

<a id="canonical-1d2b532f9db737bbe764b397b9d48b4c6f34b01d25e1fee4857f011dd3e6d7cc"></a>

## Next pages — Examples / 81c1ac878ff7 / 4

- [Data source](data-sources--data_group--examples--group-001.md#canonical-58926b899ebbad90a654ba69cfaafde688e5ec840aa5a18bcf0c56b3cc46af02)
- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)

<a id="canonical-58926b899ebbad90a654ba69cfaafde688e5ec840aa5a18bcf0c56b3cc46af02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e97120be9018c64a1385d43934207b0442c5e96e4f1309fac3d02dcc544f624c"></a>

## Data source — Data source / 1cb9b7ba6b93 / 2

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
- [Examples](data-sources--data_group--examples--group-001.md#canonical-30577a5df82559f5bd7a5798e84dcf0a6240212b5fc1fd6cbfe37d0f6b76bb0f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_data_group/data-source.tf`; digest `sha256:b9743d4eb532720e8a79b190fba83b8f94d4cbca361dec928076b2b257a940ad`.

```terraform
# DataGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataGroup by name
data "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}

output "data_group_id" {
  value = data.xcsh_data_group.example.id
}
```

<a id="canonical-ed8db531c6405519b00a917a45a84015944d67be9686b1a315aada40ce201c4b"></a>

## Next pages — Data source / 1cb9b7ba6b93 / 3

- [Examples](data-sources--data_group--examples--group-001.md#canonical-30577a5df82559f5bd7a5798e84dcf0a6240212b5fc1fd6cbfe37d0f6b76bb0f)
- [xcsh_data_group](../data-sources/data_group.md#canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2)
