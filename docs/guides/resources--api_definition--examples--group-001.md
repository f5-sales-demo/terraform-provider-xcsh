---
page_title: "xcsh_api_definition examples"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition examples."
---

# xcsh_api_definition examples

<a id="canonical-1021200213031100-3331111211312013-1031323303301031-0322221202302021-2220021101122011-3011100220301120-2331131031230210-2001320110010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- Examples

<a id="canonical-1323200133023120-1123012012112322-3113133330003223-1133332021223330-2331202223000320-1333032230310221-3300001001320023-1103232020200010"></a>

### Complete configurations for `xcsh_api_definition`

- [Resource](resources--api_definition--examples--group-001.md#canonical-1213001112132303-3333011121103100-3210233222333023-0331300210320303-1230003312130322-2100032100011211-2201033221322101-2011122021010300): valid configuration.

<a id="canonical-1213001112132303-3333011121103100-3210233222333023-0331300210320303-1230003312130322-2100032100011211-2201033221322101-2011122021010300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md#canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300)
- [Examples](resources--api_definition--examples--group-001.md#canonical-1021200213031100-3331111211312013-1031323303301031-0322221202302021-2220021101122011-3011100220301120-2331131031230210-2001320110010210)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_definition/resource.tf`; digest `sha256:8eaaa98845c6177dba9c28e17da78fe9cbe8a7d10b85b7c4e55e7857d3d8ba64`.

```terraform
# APIDefinition Resource Example
# Manages API Definition in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDefinition configuration
resource "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}
```
