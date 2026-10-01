---
page_title: "xcsh_device_intelligence_device_history landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_history landing."
---

# xcsh_device_intelligence_device_history landing

<a id="canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e2f6d9849a6c9de05598bb81aebfd304d148a21a2d7dfc3e70f0f8401d4a80c"></a>

## xcsh_device_intelligence_device_history — xcsh_device_intelligence_device_history / 5dc83232b064 / 2

Breadcrumbs:

- xcsh_device_intelligence_device_history

Resource creation operation.

<a id="canonical-99bd464bad059cd9acb6834f774762a70fff2dcb6bcd278ed27353c0f8f77db5"></a>

## Prerequisites — xcsh_device_intelligence_device_history / 5dc83232b064 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a068ab3015ff8ddaccb518c39a21790620ae25ae1c5bc4eb5145342677044c76"></a>

## Minimal configuration — xcsh_device_intelligence_device_history / 5dc83232b064 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDeviceHistory DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_history" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_history_result" {
  value = data.xcsh_device_intelligence_device_history.example
}
```

<a id="canonical-bde3a038a8a081f5397ccac1402830f547566e12b9ec7b5758093c4ea52e8662"></a>

## Root configuration — xcsh_device_intelligence_device_history / 5dc83232b064 / 5

Required root properties: `device_id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-27f82b78343e386344acf6330c92a7d97547b09bffb16310a0eb340a45af9577"></a>

## Next pages — xcsh_device_intelligence_device_history / 5dc83232b064 / 6

- [Property reference](../guides/data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- [Examples](../guides/data-sources--device_intelligence_device_history--examples--group-001.md#canonical-1951b16ac543cc0787ec0e6e895a4bba92897dd508464ba4197dac078a4fb569)
