---
page_title: "xcsh_token examples"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token examples."
---

# xcsh_token examples

<a id="canonical-48e55ab18f6d9f5efd96fc6d6200ed21e196eeca9c0b4d74d36b64f40d067778"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f317a43de04612f2453f8831c50edc51669b5eedfc07ccbe3c9f192f85436682"></a>

## Examples — Examples / c9d9474ae3da / 2

Breadcrumbs:

- [xcsh_token](../data-sources/token.md#canonical-807087d4571bee22ab642d9c81935b55fbfa12fbb563a3245dda111f3a3e04a2)
- Examples

<a id="canonical-6748b24b60e666eb31cc52c27653ef2f18ba717a4b055c462a4995630d158013"></a>

## Complete configurations — Examples / c9d9474ae3da / 3

- [Data source](data-sources--token--examples--group-001.md#canonical-128b23b09d8414c867798f690d11b81018427d49497ed605388bb86bbb56d549): valid configuration.

<a id="canonical-b2722c1b80b88a94a1244fcbbb580e24a9ea08db04b8dc7d08fc03c396335650"></a>

## Next pages — Examples / c9d9474ae3da / 4

- [Data source](data-sources--token--examples--group-001.md#canonical-128b23b09d8414c867798f690d11b81018427d49497ed605388bb86bbb56d549)
- [xcsh_token](../data-sources/token.md#canonical-807087d4571bee22ab642d9c81935b55fbfa12fbb563a3245dda111f3a3e04a2)

<a id="canonical-128b23b09d8414c867798f690d11b81018427d49497ed605388bb86bbb56d549"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5280e4124947673909f09593925d9066924b6506988a6784c0d93fd8edb3ef39"></a>

## Data source — Data source / 62d30be68ee1 / 2

Breadcrumbs:

- [xcsh_token](../data-sources/token.md#canonical-807087d4571bee22ab642d9c81935b55fbfa12fbb563a3245dda111f3a3e04a2)
- [Examples](data-sources--token--examples--group-001.md#canonical-48e55ab18f6d9f5efd96fc6d6200ed21e196eeca9c0b4d74d36b64f40d067778)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_token/data-source.tf`; digest `sha256:2be846025946a447e16fd9c1be64b48387d966481fa919712ec2c2311d266aee`.

```terraform
# Token Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Token by name
data "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
}

output "token_id" {
  value = data.xcsh_token.example.id
}
```

<a id="canonical-3f636cd6c9561a3f7bc6b920596518de716193d1384c78a344d14125626da775"></a>

## Next pages — Data source / 62d30be68ee1 / 3

- [Examples](data-sources--token--examples--group-001.md#canonical-48e55ab18f6d9f5efd96fc6d6200ed21e196eeca9c0b4d74d36b64f40d067778)
- [xcsh_token](../data-sources/token.md#canonical-807087d4571bee22ab642d9c81935b55fbfa12fbb563a3245dda111f3a3e04a2)
