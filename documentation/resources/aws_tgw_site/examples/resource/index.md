---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1288, "body_sha256": "sha256:a6172325e030f46bb549d8dbf797d2fc21ec044a5b8a9befae1c281e36e7ca15", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c51271321a366128fb5a33731b3df495e257a63396c08d4e66b33465ad7f22a", "source_path": "examples/resources/xcsh_aws_tgw_site/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:aws_tgw_site:example:resource", "parent_id": "xcsh-docs:resources:aws_tgw_site:examples", "path": "documentation/resources/aws_tgw_site/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0220101010212111-3210120100033223-1102030220100330-1303312230023132-3302112231312321-2222223122023021-3113303303302302-1102002323011301", "registry_path": "docs/guides/resources--aws_tgw_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_aws_tgw_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/examples/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
