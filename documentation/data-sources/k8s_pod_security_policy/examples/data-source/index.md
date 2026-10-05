---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1418, "body_sha256": "sha256:b8a61208e2d11c23c2544f60f5cfb43e5b243c60140be317085cfd60090c2bd3", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:466f0cfc6eb4d12538a2c33e13b91f9270feaa389452e6596f117ab99a547c9b", "source_path": "examples/data-sources/xcsh_k8s_pod_security_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_pod_security_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:examples", "path": "documentation/data-sources/k8s_pod_security_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3232233311113333-1301121233203113-1311330100220222-3321230200332022-2330232213110220-0231030131021333-3300301112003201-0113032212200122", "registry_path": "docs/guides/data-sources--k8s_pod_security_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_k8s_pod_security_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_pod_security_policy/data-source.tf`; digest `sha256:466f0cfc6eb4d12538a2c33e13b91f9270feaa389452e6596f117ab99a547c9b`.

```terraform
# K8SPodSecurityPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityPolicy by name
data "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}

output "k8s_pod_security_policy_id" {
  value = data.xcsh_k8s_pod_security_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/examples/)
- [xcsh_k8s_pod_security_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_policy/)
