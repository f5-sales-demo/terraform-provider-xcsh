---
page_title: "xcsh_allowed_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain examples."
---

# xcsh_allowed_domain examples

<a id="canonical-1302323331312313-3102131232131300-1221010023033210-3232331202213030-1130131020111011-1011011203223203-0321101210111222-2303301203132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-3031012320322023-2113120332301001-0010123230112302-0303132031320332-1032231003302322-1130311111131220-1203223310030220-1032110013113220)
- Examples

<a id="canonical-2221220121113123-3323200132223100-3213300302101213-0102232302301100-3113310333103210-1011000200322323-2212220301303320-2332213113330300"></a>

### Complete configurations for `xcsh_allowed_domain`

- [Data source](data-sources--allowed_domain--examples--group-001.md#canonical-0100000003211222-1011030202302300-1223321113203003-2011211123233321-2112020030000123-0001300223021013-2131200103302111-3020112332210010): valid configuration.

<a id="canonical-0100000003211222-1011030202302300-1223321113203003-2011211123233321-2112020030000123-0001300223021013-2131200103302111-3020112332210010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-3031012320322023-2113120332301001-0010123230112302-0303132031320332-1032231003302322-1130311111131220-1203223310030220-1032110013113220)
- [Examples](data-sources--allowed_domain--examples--group-001.md#canonical-1302323331312313-3102131232131300-1221010023033210-3232331202213030-1130131020111011-1011011203223203-0321101210111222-2303301203132121)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_allowed_domain/data-source.tf`; digest `sha256:e22d7f6871a55bc1c9dd4658347afbf10c80e91173e9c9082eacc38fb5713006`.

```terraform
# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
```
