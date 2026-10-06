---
page_title: "xcsh_forwarding_class examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class examples."
---

# xcsh_forwarding_class examples

<a id="canonical-1311111311300112-2102033310212222-0300103200221301-3230203003323211-2012012101313220-3103132200331032-1112102303131321-0100100130030332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- Examples

<a id="canonical-3322113201113330-2210313310221010-3203101311300011-0112233300030102-2121131230102103-3331201032101220-3212301032011203-0131202210333220"></a>

### Complete configurations for `xcsh_forwarding_class`

- [Data source](data-sources--forwarding_class--examples--group-001.md#canonical-1211130101120122-3130023111212113-1301010213200202-2030102011033201-3201111120220223-0032001333031030-3311023311310211-2023233100000331): valid configuration.

<a id="canonical-1211130101120122-3130023111212113-1301010213200202-2030102011033201-3201111120220223-0032001333031030-3311023311310211-2023233100000331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_forwarding_class](../data-sources/forwarding_class.md#canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203)
- [Examples](data-sources--forwarding_class--examples--group-001.md#canonical-1311111311300112-2102033310212222-0300103200221301-3230203003323211-2012012101313220-3103132200331032-1112102303131321-0100100130030332)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_forwarding_class/data-source.tf`; digest `sha256:12747b7c8fc5a62067b4a0cf603f5a2dd940ca02738f214a627abb11e786118d`.

```terraform
# ForwardingClass Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardingClass by name
data "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}

output "forwarding_class_id" {
  value = data.xcsh_forwarding_class.example.id
}
```
