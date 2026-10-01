---
page_title: "xcsh_bot_infrastructure examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure examples."
---

# xcsh_bot_infrastructure examples

<a id="canonical-f81d994c8dee00bb9e231d904fced88e7af9e8c8d882c731cb033b9b31919ca3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-608e4187d58486a1f8adb68fbcec61d36115e2433a611d5ffccb4ab7610c5a5b"></a>

## Examples — Examples / b482d5db384c / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
- Examples

<a id="canonical-a2da1e1a430989b8d26ea43b65593c0c274f29eba238238c1098e7314bc82e2a"></a>

## Complete configurations — Examples / b482d5db384c / 3

- [Data source](data-sources--bot_infrastructure--examples--group-001.md#canonical-9acc3f3696d38b0835ae3399a0c053bfab88295b66bd08e1a31755a8b30df597): valid configuration.

<a id="canonical-11c7f9346ae6c9470e8f4908b4dceaa597b24d3611fbbfb38dc1971b11393369"></a>

## Next pages — Examples / b482d5db384c / 4

- [Data source](data-sources--bot_infrastructure--examples--group-001.md#canonical-9acc3f3696d38b0835ae3399a0c053bfab88295b66bd08e1a31755a8b30df597)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)

<a id="canonical-9acc3f3696d38b0835ae3399a0c053bfab88295b66bd08e1a31755a8b30df597"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7ae1eccbbd1f1b549bf3a66dcac323ee1481541a4f8f61ff428ba6b94837bc0"></a>

## Data source — Data source / 30a02f9583e3 / 2

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
- [Examples](data-sources--bot_infrastructure--examples--group-001.md#canonical-f81d994c8dee00bb9e231d904fced88e7af9e8c8d882c731cb033b9b31919ca3)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_infrastructure/data-source.tf`; digest `sha256:f72b753447890fa3999ee0487a20af9414eac1431bcf20d1ef50048bd9025f06`.

```terraform
# BotInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotInfrastructure by name
data "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}

output "bot_infrastructure_id" {
  value = data.xcsh_bot_infrastructure.example.id
}
```

<a id="canonical-5a38a1871430cfa41c49ebabd8c9fb1d54ddbc24db9e57df8d585768582ff28a"></a>

## Next pages — Data source / 30a02f9583e3 / 3

- [Examples](data-sources--bot_infrastructure--examples--group-001.md#canonical-f81d994c8dee00bb9e231d904fced88e7af9e8c8d882c731cb033b9b31919ca3)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md#canonical-877bbb7d337b693d557733ad6caa7f00c7f8b096aeb76ab4bf994720688abb64)
