---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1286, "body_sha256": "sha256:3a5d422879b735a8f10b08d9fa0ce2228048a95d78603e63aabcab3c74ad4dfe", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:50f8f72e662dc6823bf6f43a5de7129e0be922fc8a2f83bf9623196d54167ed5", "source_path": "examples/resources/xcsh_aws_vpc_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:aws_vpc_site:example:resource", "parent_id": "xcsh-docs:resources:aws_vpc_site:examples", "path": "documentation/resources/aws_vpc_site/examples/resource/index.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/examples/)
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/examples/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
