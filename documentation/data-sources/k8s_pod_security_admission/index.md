---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "Reads Kubernetes Pod Security Admission information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["k8s pod security admission"], "body_bytes": 1460, "body_sha256": "sha256:d5015ef107e8c51844be14de03e7a2295d9041022eb5fc3c8ceabcaa25b5d6a6", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_admission:reference", "xcsh-docs:data-sources:k8s_pod_security_admission:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_admission:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/k8s_pod_security_admission/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013", "registry_path": "docs/data-sources/k8s_pod_security_admission.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Kubernetes Pod Security Admission information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
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
