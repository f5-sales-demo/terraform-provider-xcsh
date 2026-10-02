---
page_title: "xcsh_k8s_pod_security_policy"
subcategory: ""
description: "Manages k8s_pod_security_policy will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "k8s pod security policy", "origin servers", "upstream servers"], "body_bytes": 1764, "body_sha256": "sha256:a1d4fc7e6db9e2bcece932bff1af0d990fc3c09333f23137db8c5718e0e5f901", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_pod_security_policy:reference", "xcsh-docs:resources:k8s_pod_security_policy:examples", "xcsh-docs:resources:k8s_pod_security_policy:import", "xcsh-docs:resources:k8s_pod_security_policy:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_pod_security_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/k8s_pod_security_policy/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112", "registry_path": "docs/resources/k8s_pod_security_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages k8s_pod_security_policy will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_k8s_pod_security_policy

Breadcrumbs:

- xcsh_k8s_pod_security_policy

Manages k8s\_pod\_security\_policy will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityPolicy Resource Example
# Manages k8s_pod_security_policy will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityPolicy configuration
resource "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_pod_security_policy/lifecycle/timeouts/)
