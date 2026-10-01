---
page_title: "xcsh_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery examples."
---

# xcsh_discovery examples

<a id="canonical-32c6db1798df2a0749a45f5581c13347d4a7697806d0a15a0d82c83f1477df2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-704c451de7468a116910c16ad0cd6bbf41de8cf1672d540b6f299576344b2304"></a>

## Examples — Examples / b1ea5507c5bd / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- Examples

<a id="canonical-80a608b4e7dbccde388f2d9f3b6ecc0bbc8966ee7dc4c3d39efa9d4752bcd132"></a>

## Complete configurations — Examples / b1ea5507c5bd / 3

- [Resource](resources--discovery--examples--group-001.md#canonical-edfbcf9fc629c4492d16aaf3c1a4cb8ace73021368106a188fbe027ceaadad8f): valid configuration.

<a id="canonical-1654c98a57d3646be4a2d8501606b46cc28ba4581fc344ffbc937b7740677f73"></a>

## Next pages — Examples / b1ea5507c5bd / 4

- [Resource](resources--discovery--examples--group-001.md#canonical-edfbcf9fc629c4492d16aaf3c1a4cb8ace73021368106a188fbe027ceaadad8f)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)

<a id="canonical-edfbcf9fc629c4492d16aaf3c1a4cb8ace73021368106a188fbe027ceaadad8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7909516382bbee05cf9da42304ea20a47d712f19d99da15953faad6d2a8ec349"></a>

## Resource — Resource / fdf35c5b8ad3 / 2

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
- [Examples](resources--discovery--examples--group-001.md#canonical-32c6db1798df2a0749a45f5581c13347d4a7697806d0a15a0d82c83f1477df2f)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_discovery/resource.tf`; digest `sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c`.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```

<a id="canonical-49756b3baf33a7a00a88363a64db611cf197fc29108e8b0c2d189d3d6612cced"></a>

## Next pages — Resource / fdf35c5b8ad3 / 3

- [Examples](resources--discovery--examples--group-001.md#canonical-32c6db1798df2a0749a45f5581c13347d4a7697806d0a15a0d82c83f1477df2f)
- [xcsh_discovery](../resources/discovery.md#canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2)
