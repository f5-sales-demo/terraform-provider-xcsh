---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 983, "body_sha256": "sha256:b9a52492b5753db39016dd0f03ae77bd6663885ad6d6b598b9d83f13badc35e4", "canonical_id": "xcsh-docs:resources:aws_tgw_site:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c51271321a366128fb5a33731b3df495e257a63396c08d4e66b33465ad7f22a", "source_path": "examples/resources/xcsh_aws_tgw_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:aws_tgw_site:example:resource", "parent_id": "xcsh-docs:resources:aws_tgw_site:examples", "path": "docs/guides/resources--aws_tgw_site--example--resource.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Examples](resources--aws_tgw_site--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_aws_tgw_site/resource.tf`; digest `sha256:3c51271321a366128fb5a33731b3df495e257a63396c08d4e66b33465ad7f22a`.

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

## Next pages

- [Examples](resources--aws_tgw_site--examples.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
