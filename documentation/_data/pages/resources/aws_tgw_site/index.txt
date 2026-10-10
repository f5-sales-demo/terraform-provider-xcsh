---
page_title: "xcsh_aws_tgw_site"
subcategory: ""
description: "Manages an AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS Transit Gateway."
xcsh_docs: {"aliases": ["aws tgw site"], "body_bytes": 1623, "body_sha256": "sha256:1b98f1dcb056e005790b27e60f0d77bd105904e3c84ea76ea6a92ddf72cfdcbb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:reference", "xcsh-docs:resources:aws_tgw_site:examples", "xcsh-docs:resources:aws_tgw_site:import", "xcsh-docs:resources:aws_tgw_site:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/aws_tgw_site/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112", "registry_path": "docs/resources/aws_tgw_site.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages an AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS Transit Gateway.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_aws_tgw_site

Breadcrumbs:

- xcsh_aws_tgw_site

Manages an AWS TGW Site resource in F5 Distributed Cloud for deploying F5 sites connected via AWS
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/lifecycle/timeouts/)
