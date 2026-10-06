---
page_title: "xcsh_nfv_service examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service examples."
---

# xcsh_nfv_service examples

<a id="canonical-0113201012132103-2303210110031133-1100201210101112-3112311201021033-0132302110030230-2000330102332211-0130111220020003-0013200013312112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- Examples

<a id="canonical-0110121003021023-2333000032113001-0000033311001001-1222133011122312-1112030030100312-3023021320031211-3130003111030323-0033322202113201"></a>

### Complete configurations for `xcsh_nfv_service`

- [Data source](data-sources--nfv_service--examples--group-001.md#canonical-2323330232101000-2311330213121130-0232012023331321-0302310130122220-1322233212010303-0133031033103223-3002231002023331-1022303312123022): valid configuration.

<a id="canonical-2323330232101000-2311330213121130-0232012023331321-0302310130122220-1322233212010303-0133031033103223-3002231002023331-1022303312123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Examples](data-sources--nfv_service--examples--group-001.md#canonical-0113201012132103-2303210110031133-1100201210101112-3112311201021033-0132302110030230-2000330102332211-0130111220020003-0013200013312112)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nfv_service/data-source.tf`; digest `sha256:e1bcf01b10f5271939ad8fda84b9d63610e3af5498aba69a8e2b34b2c60983a2`.

```terraform
# NfvService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NfvService by name
data "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}

output "nfv_service_id" {
  value = data.xcsh_nfv_service.example.id
}
```
