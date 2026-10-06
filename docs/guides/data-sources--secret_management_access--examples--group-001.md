---
page_title: "xcsh_secret_management_access examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access examples."
---

# xcsh_secret_management_access examples

<a id="canonical-2233211101011210-0200323221322013-3111333211300333-1012012001332032-3313222213113030-0203213230312332-1221130132002332-0111031330020001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- Examples

<a id="canonical-2212111113100111-2223100020121213-1111133020031032-0123312232030102-2333302312000332-0123300321132020-0220212130033210-2123112022232203"></a>

### Complete configurations for `xcsh_secret_management_access`

- [Data source](data-sources--secret_management_access--examples--group-001.md#canonical-3302210133323002-1000031112203132-2012120131121021-3210112320101210-3003101021323310-3212230000320000-2320123132012231-0101100201333230): valid configuration.

<a id="canonical-3302210133323002-1000031112203132-2012120131121021-3210112320101210-3003101021323310-3212230000320000-2320123132012231-0101100201333230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Examples](data-sources--secret_management_access--examples--group-001.md#canonical-2233211101011210-0200323221322013-3111333211300333-1012012001332032-3313222213113030-0203213230312332-1221130132002332-0111031330020001)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_secret_management_access/data-source.tf`; digest `sha256:ed6c81d9e9def7ced3f618a92fe6932b7593d44a1f95e72e70fe1cd4aab5ace2`.

```terraform
# SecretManagementAccess Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecretManagementAccess by name
data "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"
}

output "secret_management_access_id" {
  value = data.xcsh_secret_management_access.example.id
}
```
