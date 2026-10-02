---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_policy_based_routing."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1381, "body_sha256": "sha256:fdf33dc197d18e625f99f759df0a828ad1b0ee9ebca617c9ffc2c1b926e28ea8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968", "source_path": "examples/resources/xcsh_policy_based_routing/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:policy_based_routing:example:resource", "parent_id": "xcsh-docs:resources:policy_based_routing:examples", "path": "documentation/resources/policy_based_routing/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3310100020110201-2313301030311313-2123101220300301-3311301012232032-0302121032101013-3003300102331202-1220111310000022-2200130313332003", "registry_path": "docs/guides/resources--policy_based_routing--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_policy_based_routing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policy_based_routing/resource.tf`; digest `sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968`.

```terraform
# PolicyBasedRouting Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PolicyBasedRouting configuration
resource "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/examples/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/)
