---
page_title: "xcsh_alert_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_receiver examples."
---

# xcsh_alert_receiver examples

<a id="canonical-9c611f5c98034d913775613ac7408603296d5248701a996dd4d98f6a2bd352b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-629dbf9f3ece97cbb2601b8b85d13122064c553aae1bcff4c46c349b99673768"></a>

## Examples — Examples / 2b89af0e70db / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- Examples

<a id="canonical-e5e160d24165814e3efd37c6cbe8b028731719c20650a1e57f93d2c870f1f928"></a>

## Complete configurations — Examples / 2b89af0e70db / 3

- [Data source](data-sources--alert_receiver--examples--group-001.md#canonical-e9538215f650702b7eaa4c3a4fc73df65a5aad7573fa1a170dd3a5b4936230c1): valid configuration.

<a id="canonical-cec41308a5898e85d8aee23621b7092a06fb5e826a4fd6688fb3ed8a63037890"></a>

## Next pages — Examples / 2b89af0e70db / 4

- [Data source](data-sources--alert_receiver--examples--group-001.md#canonical-e9538215f650702b7eaa4c3a4fc73df65a5aad7573fa1a170dd3a5b4936230c1)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)

<a id="canonical-e9538215f650702b7eaa4c3a4fc73df65a5aad7573fa1a170dd3a5b4936230c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbef5281e622d7dc433a1d0398b57a96829606645102233027a655bba2881b0c"></a>

## Data source — Data source / e44203c301ec / 2

Breadcrumbs:

- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
- [Examples](data-sources--alert_receiver--examples--group-001.md#canonical-9c611f5c98034d913775613ac7408603296d5248701a996dd4d98f6a2bd352b4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_receiver/data-source.tf`; digest `sha256:96027faa6721d178ff8fb480c66120ec2d45036c783323915216275133e6a1d6`.

```terraform
# AlertReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertReceiver by name
data "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}

output "alert_receiver_id" {
  value = data.xcsh_alert_receiver.example.id
}
```

<a id="canonical-16e753ee710c213c587bd4bafb0790295990a7bb22c3a745bc8e5666007b72f9"></a>

## Next pages — Data source / e44203c301ec / 3

- [Examples](data-sources--alert_receiver--examples--group-001.md#canonical-9c611f5c98034d913775613ac7408603296d5248701a996dd4d98f6a2bd352b4)
- [xcsh_alert_receiver](../data-sources/alert_receiver.md#canonical-c67ce5a0c445e137f6b5fed733115616ef3a219c23d6e98b458886aead3078d2)
