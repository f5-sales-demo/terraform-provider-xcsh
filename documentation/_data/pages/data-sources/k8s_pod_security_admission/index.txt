---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["k8s pod security admission"], "body_bytes": 1477, "body_sha256": "sha256:1dbf18595f23c15bde22d544eefb4215ff4b93e87f84d2fe8d4e5dbf9866cc27", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_pod_security_admission:reference", "xcsh-docs:data-sources:k8s_pod_security_admission:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_pod_security_admission:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/k8s_pod_security_admission/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013", "registry_path": "docs/data-sources/k8s_pod_security_admission.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/examples/)
