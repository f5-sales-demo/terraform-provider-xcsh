---
page_title: "xcsh_aws_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site landing."
---

# xcsh_aws_vpc_site landing

<a id="canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33bc7e36aa1f3951f1dc5b061afcdb8cf7086071b9e503cf94aa25f8fb4a2086"></a>

## xcsh_aws_vpc_site — xcsh_aws_vpc_site / e3f85b66e186 / 2

Breadcrumbs:

- xcsh_aws_vpc_site

Manages a AWS VPC Site resource in F5 Distributed Cloud for deploying F5 sites within AWS VPC
environments.

<a id="canonical-32113dadfd49542bcb95aae872275ee868ef54f83259288e30b1b86a369df2f6"></a>

## Prerequisites — xcsh_aws_vpc_site / e3f85b66e186 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: AWS authentication for deployment

<a id="canonical-ba65190c6eaf2925e0d6028f735bfcf6bea429b564959a4d7e91c262eaf92008"></a>

## Minimal configuration — xcsh_aws_vpc_site / e3f85b66e186 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AWSVPCSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSVPCSite by name
data "xcsh_aws_vpc_site" "example" {
  name      = "example-aws-vpc-site"
  namespace = "staging"
}

output "aws_vpc_site_id" {
  value = data.xcsh_aws_vpc_site.example.id
}
```

<a id="canonical-55dd953d84a637b5b20b0563a16c8d2111156730d29de1999495df3200de0dbd"></a>

## Root configuration — xcsh_aws_vpc_site / e3f85b66e186 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-09254ac09f620eae9f55be85c285657970b3fb2fa3c78c6b39457f34cccb038e"></a>

## Next pages — xcsh_aws_vpc_site / e3f85b66e186 / 6

- [Property reference](../guides/data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [Examples](../guides/data-sources--aws_vpc_site--examples--group-001.md#canonical-5c3437aa50cde77e2eedddc34f6ed350d27b75623455e69714667fa60cf9f156)
