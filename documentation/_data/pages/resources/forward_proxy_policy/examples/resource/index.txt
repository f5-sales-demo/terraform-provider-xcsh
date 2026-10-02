---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1366, "body_sha256": "sha256:2c9ee2cc23ce873c55fab705163c2612d355735b8cc5345fba1419965be6ebd1", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bcc594a880ce1beadc720b2bd1c76066016eb6d5489e15cf1b1295e1e8d40584", "source_path": "examples/resources/xcsh_forward_proxy_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:forward_proxy_policy:example:resource", "parent_id": "xcsh-docs:resources:forward_proxy_policy:examples", "path": "documentation/resources/forward_proxy_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1333002001312121-0210033011100310-2331003122231131-2003002003022312-3113033121200010-1120120033200012-0021311013010101-2303222112331101", "registry_path": "docs/guides/resources--forward_proxy_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_forward_proxy_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/examples/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
