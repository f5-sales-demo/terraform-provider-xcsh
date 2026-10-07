---
page_title: "csrf_policy"
subcategory: ""
description: "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request"
xcsh_docs: {"aliases": ["csrf policy"], "body_bytes": 2610, "body_sha256": "sha256:4d59eea3f92dd3250075c47d522dddd9d23d989b723a808db41005e1027c9139", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:csrf_policy:all_load_balancer_domains", "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:csrf_policy", "parent_id": "xcsh-docs:resources:virtual_host:reference", "path": "documentation/resources/virtual_host/properties/csrf_policy/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0032333010121002-0332331033300212-2002033020321220-2332202013303113-2313133231100331-0003222333003100-3133200232230203-0121213102202020", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["csrf_policy"], "schema_version": 1, "sections": [{"aliases": ["csrf policy all load balancer domains"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:all_load_balancer_domains", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csrf_policy", "all_load_balancer_domains"], "syntax": "attribute", "type": "object"}, {"aliases": ["csrf policy custom domain list"], "anchor": "section", "description": "List of domain names used for Host header matching.", "document_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-csrf_policy--custom_domain_list--domains", "enforcement": "provider-schema", "group": "csrf_policy.custom_domain_list:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "type": "requires"}], "schema_path": ["csrf_policy", "custom_domain_list"], "syntax": "block", "type": "object"}, {"aliases": ["csrf policy disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["csrf_policy", "disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/csrf_policy/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# csrf_policy

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- csrf_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/all_load_balancer_domains/): complete subsection reference.

- [custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/custom_domain_list/): complete subsection reference.

- [disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/disabled/): complete subsection reference.
