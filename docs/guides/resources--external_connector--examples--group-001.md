---
page_title: "xcsh_external_connector examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector examples."
---

# xcsh_external_connector examples

<a id="canonical-1023020021113213-0322331233001322-0012100022302130-1013123302233230-0232032201001213-3203001323100231-3330210233012201-1001111112222333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- Examples

<a id="canonical-2132122233113020-1032221111210110-3121032031120121-1223102000331020-1111311233303030-3321330133331332-3201003332311330-2231233311021002"></a>

### Complete configurations for `xcsh_external_connector`

- [Resource](resources--external_connector--examples--group-001.md#canonical-1311210323213000-2020122012123210-0121232212103321-2100123111010301-2020231230132333-0000023203230331-3033131002100022-0013331232230331): valid configuration.

<a id="canonical-1311210323213000-2020122012123210-0121232212103321-2100123111010301-2020231230132333-0000023203230331-3033131002100022-0013331232230331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Examples](resources--external_connector--examples--group-001.md#canonical-1023020021113213-0322331233001322-0012100022302130-1013123302233230-0232032201001213-3203001323100231-3330210233012201-1001111112222333)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_external_connector/resource.tf`; digest `sha256:e6d206bd3e4355ffab92fe542ac16d79135373b63491d27d87ac9997db3c291a`.

```terraform
# ExternalConnector Resource Example
# Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ExternalConnector configuration
resource "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}
```
