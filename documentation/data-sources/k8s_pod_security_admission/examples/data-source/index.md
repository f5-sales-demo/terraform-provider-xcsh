---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1456, "body_sha256": "sha256:60fab613eda45361f47d026f3bb8bd18b1b18e9ba6f3658b2662c76b85b4c7f8", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d1bc83d3e3b733a106b0eb3f1736fa294f7fa3ace28ea4c63d278edef289369e", "source_path": "examples/data-sources/xcsh_k8s_pod_security_admission/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_pod_security_admission:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_admission:examples", "path": "documentation/data-sources/k8s_pod_security_admission/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3221223021033111-1021032022332303-1213130222211230-0210103201101331-3011131203022111-3300212120310311-2111003011130230-0332313003012020", "registry_path": "docs/guides/data-sources--k8s_pod_security_admission--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_k8s_pod_security_admission.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_pod_security_admission/data-source.tf`; digest `sha256:d1bc83d3e3b733a106b0eb3f1736fa294f7fa3ace28ea4c63d278edef289369e`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/examples/)
- [xcsh_k8s_pod_security_admission](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_pod_security_admission/)
