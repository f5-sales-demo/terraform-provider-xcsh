---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_k8s_pod_security_admission."
xcsh_docs: {"aliases": [], "body_bytes": 1250, "body_sha256": "sha256:561e01bb6d74c224664eac043da18308dc47ee0dc8c08a7c93f288232871a0a2", "canonical_id": "xcsh-docs:data-sources:k8s_pod_security_admission:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_pod_security_admission:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d1bc83d3e3b733a106b0eb3f1736fa294f7fa3ace28ea4c63d278edef289369e", "source_path": "examples/data-sources/xcsh_k8s_pod_security_admission/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_pod_security_admission:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_pod_security_admission:examples", "path": "docs/guides/data-sources--k8s_pod_security_admission--example--data-source.md", "provider_name": "k8s_pod_security_admission", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_pod_security_admission/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_k8s_pod_security_admission.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_pod_security_admissionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md)
- [Examples](data-sources--k8s_pod_security_admission--examples.md)
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

- [Examples](data-sources--k8s_pod_security_admission--examples.md)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md)
