---
page_title: "xcsh_registration"
subcategory: ""
description: "Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users. configuration."
xcsh_docs: {"aliases": ["registration"], "body_bytes": 1688, "body_sha256": "sha256:01e81f16a967ecbb423208cc93e59b11e62487313b5b44c39b2f1ce1f87ccefa", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration:reference", "xcsh-docs:resources:registration:examples", "xcsh-docs:resources:registration:import", "xcsh-docs:resources:registration:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/registration/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323", "registry_path": "docs/resources/registration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_registration

Breadcrumbs:

- xcsh_registration

Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this
message, never used by users. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `token`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/lifecycle/timeouts/)
