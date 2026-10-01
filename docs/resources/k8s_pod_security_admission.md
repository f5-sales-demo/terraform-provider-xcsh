---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "xcsh_k8s_pod_security_admission for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": [], "body_bytes": 1531, "body_sha256": "sha256:ffabf0e2f3f387d83c28b9c0a9d03b5104c53091fb95a039b39109be7aa563fd", "canonical_id": "xcsh-docs:resources:k8s_pod_security_admission:fundamentals", "child_ids": ["xcsh-docs:resources:k8s_pod_security_admission:reference", "xcsh-docs:resources:k8s_pod_security_admission:examples", "xcsh-docs:resources:k8s_pod_security_admission:import", "xcsh-docs:resources:k8s_pod_security_admission:timeouts"], "collection_id": "xcsh-docs:resources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_pod_security_admission:fundamentals", "parent_id": null, "path": "docs/resources/k8s_pod_security_admission.md", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_pod_security_admission/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_k8s_pod_security_admission for xcsh_k8s_pod_security_admission.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_k8s_pod_security_admission

Breadcrumbs:

- xcsh_k8s_pod_security_admission

Manages k8s\_pod\_security\_admission will create the object in the storage backend in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityAdmission Resource Example
# Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityAdmission configuration
resource "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--k8s_pod_security_admission--reference.md)
- [Examples](../guides/resources--k8s_pod_security_admission--examples.md)
- [Import](../guides/resources--k8s_pod_security_admission--import.md)
- [Timeouts](../guides/resources--k8s_pod_security_admission--timeouts.md)
