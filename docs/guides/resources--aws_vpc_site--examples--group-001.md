---
page_title: "xcsh_aws_vpc_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site examples."
---

# xcsh_aws_vpc_site examples

<a id="canonical-c9a8749ca7e9910893bcefc40a95b7ad0057213fb843066fc240448a0674c2d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0c28f903c49f1dcf1aaa492de4298ea8782a073ac97e1191835018b5d9f8236"></a>

## Examples — Examples / 1cb869cce6f4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- Examples

<a id="canonical-4d26391e64569d163aae95d69893fd927d610a43658313cac6a82c0310a3689f"></a>

## Complete configurations — Examples / 1cb869cce6f4 / 3

- [Resource](resources--aws_vpc_site--examples--group-001.md#canonical-a96d4410acc9dc482871113ddb785831fee7256682e644d6f30532fe43cc6b02): valid configuration.

<a id="canonical-f99f0daf77ae15abf1311081fbf014e5476694889e059fefdae050c30a4ff6ca"></a>

## Next pages — Examples / 1cb869cce6f4 / 4

- [Resource](resources--aws_vpc_site--examples--group-001.md#canonical-a96d4410acc9dc482871113ddb785831fee7256682e644d6f30532fe43cc6b02)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a96d4410acc9dc482871113ddb785831fee7256682e644d6f30532fe43cc6b02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-456755cd0f9c2d0e7f9b5af1b7228274ca806932161947405cf46bbeabf92bf8"></a>

## Resource — Resource / 6359ab1eafd1 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Examples](resources--aws_vpc_site--examples--group-001.md#canonical-c9a8749ca7e9910893bcefc40a95b7ad0057213fb843066fc240448a0674c2d2)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_aws_vpc_site/resource.tf`; digest `sha256:50f8f72e662dc6823bf6f43a5de7129e0be922fc8a2f83bf9623196d54167ed5`.

```terraform
# AWSVPCSite Resource Example
# Manages a AWS VPC Site resource in F5 Distributed Cloud for deploying F5 sites within AWS VPC environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AWSVPCSite configuration
resource "xcsh_aws_vpc_site" "example" {
  name      = "example-aws-vpc-site"
  namespace = "staging"

  aws_region    = "example-value"
  instance_type = "example-value"
  ssh_key       = "example-value"
}
```

<a id="canonical-17e1b9bcd550e0ec7fb9adba7268ef98d7b69ab88b15e94bff022e8b821495a2"></a>

## Next pages — Resource / 6359ab1eafd1 / 3

- [Examples](resources--aws_vpc_site--examples--group-001.md#canonical-c9a8749ca7e9910893bcefc40a95b7ad0057213fb843066fc240448a0674c2d2)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
