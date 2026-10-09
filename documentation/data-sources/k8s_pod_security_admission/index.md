---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "Reads Kubernetes Pod Security Admission information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["k8s pod security admission"], "body_bytes": 1460, "body_sha256": "sha256:d5015ef107e8c51844be14de03e7a2295d9041022eb5fc3c8ceabcaa25b5d6a6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_admission:reference", "xcsh-docs:data-sources:k8s_pod_security_admission:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_admission:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/k8s_pod_security_admission/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013", "registry_path": "docs/data-sources/k8s_pod_security_admission.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reads Kubernetes Pod Security Admission information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_k8s_pod_security_admission

Breadcrumbs:

- xcsh_k8s_pod_security_admission

Reads Kubernetes Pod Security Admission information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityAdmission Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityAdmission by name
data "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}

output "k8s_pod_security_admission_id" {
  value = data.xcsh_k8s_pod_security_admission.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/examples/)
