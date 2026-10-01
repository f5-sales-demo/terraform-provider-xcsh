---
page_title: "xcsh_policer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer examples."
---

# xcsh_policer examples

<a id="canonical-1ef5a3f8e94ec063f1bd6e0f9cd0fd48a7ac2bcac02a6bd4e066dc01077cf67d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6a94e1543fc0d3fd715cb51d89d700a9515a025af6a28251664e0a52a1829cb"></a>

## Examples — Examples / b4371ffa29ca / 2

Breadcrumbs:

- [xcsh_policer](../data-sources/policer.md#canonical-fc38c1c976cffbbb7d1001c95e37ff2eb2f3d9753461e9f18f25729eae4994c0)
- Examples

<a id="canonical-0a5fe25d00520112b1076bfc52f66a60a106ae89e61c95bf765038cec9c52ff2"></a>

## Complete configurations — Examples / b4371ffa29ca / 3

- [Data source](data-sources--policer--examples--group-001.md#canonical-4224eb4b578d2e2251fa67d880c23c5737d6c6d52722682d57a42fff13ba17da): valid configuration.

<a id="canonical-33623f5b15283577c5caa106f835b657f0184045064fde3ea00a7dd79c820f2d"></a>

## Next pages — Examples / b4371ffa29ca / 4

- [Data source](data-sources--policer--examples--group-001.md#canonical-4224eb4b578d2e2251fa67d880c23c5737d6c6d52722682d57a42fff13ba17da)
- [xcsh_policer](../data-sources/policer.md#canonical-fc38c1c976cffbbb7d1001c95e37ff2eb2f3d9753461e9f18f25729eae4994c0)

<a id="canonical-4224eb4b578d2e2251fa67d880c23c5737d6c6d52722682d57a42fff13ba17da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c096d823efa3471976820922ffd56caed2d152dee054ee68c66fb40bf9156f48"></a>

## Data source — Data source / 1698f1544e2e / 2

Breadcrumbs:

- [xcsh_policer](../data-sources/policer.md#canonical-fc38c1c976cffbbb7d1001c95e37ff2eb2f3d9753461e9f18f25729eae4994c0)
- [Examples](data-sources--policer--examples--group-001.md#canonical-1ef5a3f8e94ec063f1bd6e0f9cd0fd48a7ac2bcac02a6bd4e066dc01077cf67d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_policer/data-source.tf`; digest `sha256:bf10d06a3b33dc40f3d87d76f2625b6d1d04a36df75db4e7eceed8ef9f19362a`.

```terraform
# Policer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Policer by name
data "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"
}

output "policer_id" {
  value = data.xcsh_policer.example.id
}
```

<a id="canonical-0225c85f70cbbc6cec1ad61cca5bf1f00e92af1fe3aef872f6f27647c1fd5c5c"></a>

## Next pages — Data source / 1698f1544e2e / 3

- [Examples](data-sources--policer--examples--group-001.md#canonical-1ef5a3f8e94ec063f1bd6e0f9cd0fd48a7ac2bcac02a6bd4e066dc01077cf67d)
- [xcsh_policer](../data-sources/policer.md#canonical-fc38c1c976cffbbb7d1001c95e37ff2eb2f3d9753461e9f18f25729eae4994c0)
