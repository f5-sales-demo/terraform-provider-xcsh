---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_segment."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1186, "body_sha256": "sha256:c420287c618c42d7203dab2a850183cb4d66a0748e8426388fc5050a44c081a3", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d784ed069273afd06c3680c3f92fabd8acd51be0aa4d921e6bfa8799ea0d6232", "source_path": "examples/resources/xcsh_segment/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:segment:example:resource", "parent_id": "xcsh-docs:resources:segment:examples", "path": "documentation/resources/segment/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0231123230002013-3232012203311111-1201213111100111-3033333033132300-0301320201032003-2322333203120010-2022201300222220-3203323130021100", "registry_path": "docs/guides/resources--segment--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_segment.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["segmentCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_segment/resource.tf`; digest `sha256:d784ed069273afd06c3680c3f92fabd8acd51be0aa4d921e6bfa8799ea0d6232`.

```terraform
# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Segment configuration
resource "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/examples/)
- [xcsh_segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/)
