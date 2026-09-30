---
page_title: "xcsh_aws_tgw_site"
subcategory: ""
description: "xcsh_aws_tgw_site for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1510, "body_sha256": "sha256:5ee0fcab06f400692d8aebff91b20ac8da38ae25c8fee79f228330c861235d58", "child_ids": ["xcsh-docs:resources:aws_tgw_site:reference", "xcsh-docs:resources:aws_tgw_site:examples", "xcsh-docs:resources:aws_tgw_site:import", "xcsh-docs:resources:aws_tgw_site:timeouts"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:fundamentals", "parent_id": null, "path": "documentation/resources/aws_tgw_site/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_aws_tgw_site for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_aws_tgw_site

Breadcrumbs:

- xcsh_aws_tgw_site

Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS
Transit Gateway.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AWSTGWSite Resource Example
# Manages a AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS Transit Gateway.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AWSTGWSite configuration
resource "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/lifecycle/timeouts/)
