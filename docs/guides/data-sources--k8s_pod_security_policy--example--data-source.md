---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_k8s_pod_security_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1212, "body_sha256": "sha256:58af445ec63801b7c6f89d19588452e2c83ad8c1b5660f4eec129f01c8f851a5", "canonical_id": "xcsh-docs:data-sources:k8s_pod_security_policy:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_pod_security_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:466f0cfc6eb4d12538a2c33e13b91f9270feaa389452e6596f117ab99a547c9b", "source_path": "examples/data-sources/xcsh_k8s_pod_security_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_pod_security_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_policy:examples", "path": "docs/guides/data-sources--k8s_pod_security_policy--example--data-source.md", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_k8s_pod_security_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
- [Examples](data-sources--k8s_pod_security_policy--examples.md)
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

- [Examples](data-sources--k8s_pod_security_policy--examples.md)
- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md)
