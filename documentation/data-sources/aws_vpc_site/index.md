---
page_title: "xcsh_aws_vpc_site"
subcategory: "Infrastructure"
description: "xcsh_aws_vpc_site for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1388, "body_sha256": "sha256:c420fe29be3748b0e5fecee87a4dd092cd6c1765daa0ed1e7818f510e3f5050d", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:reference", "xcsh-docs:data-sources:aws_vpc_site:examples"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:fundamentals", "parent_id": null, "path": "documentation/data-sources/aws_vpc_site/index.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_aws_vpc_site for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_aws_vpc_site

Breadcrumbs:

- xcsh_aws_vpc_site

Manages a AWS VPC Site resource in F5 Distributed Cloud for deploying F5 sites within AWS VPC
environments.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: AWS authentication for deployment

## Minimal configuration

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/examples/)
