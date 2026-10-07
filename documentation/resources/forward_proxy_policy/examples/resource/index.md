---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1120, "body_sha256": "sha256:8e3bcec2df5a7fd96230355ec2265dca6f1accbe003503f89f1f5a1cc17a874b", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bcc594a880ce1beadc720b2bd1c76066016eb6d5489e15cf1b1295e1e8d40584", "source_path": "examples/resources/xcsh_forward_proxy_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:forward_proxy_policy:example:resource", "parent_id": "xcsh-docs:resources:forward_proxy_policy:examples", "path": "documentation/resources/forward_proxy_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1333002001312121-0210033011100310-2331003122231131-2003002003022312-3113033121200010-1120120033200012-0021311013010101-2303222112331101", "registry_path": "docs/guides/resources--forward_proxy_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Resource for xcsh_forward_proxy_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forward_proxy_policy/resource.tf`; digest `sha256:bcc594a880ce1beadc720b2bd1c76066016eb6d5489e15cf1b1295e1e8d40584`.

```terraform
# ForwardProxyPolicy Resource Example
# Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardProxyPolicy configuration
resource "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}
```
