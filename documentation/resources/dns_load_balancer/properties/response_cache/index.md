---
page_title: "response_cache"
subcategory: "DNS"
description: "Response Cache x-required."
xcsh_docs: {"aliases": ["response cache"], "body_bytes": 2017, "body_sha256": "sha256:fb841aa20bd5192a3653c70879c3d805ed04fdd3303d26973a6fb79ee271678e", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "xcsh-docs:resources:dns_load_balancer:properties:response_cache:disable_spec", "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache", "parent_id": "xcsh-docs:resources:dns_load_balancer:reference", "path": "documentation/resources/dns_load_balancer/properties/response_cache/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "response_cache:ConflictingObjectAttributes:default_response_cache_parameters,disable_spec", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cache:ConflictingObjectAttributes:default_response_cache_parameters,response_cache_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cache:ConflictingObjectAttributes:default_response_cache_parameters,disable_spec", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cache:ConflictingObjectAttributes:disable_spec,response_cache_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cache:ConflictingObjectAttributes:default_response_cache_parameters,response_cache_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cache:ConflictingObjectAttributes:disable_spec,response_cache_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache"], "schema_version": 1, "sections": [{"aliases": ["response cache default response cache parameters"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "default_response_cache_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["response cache disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cache", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["response cache response cache parameters"], "anchor": "section", "description": "Configuration parameter for response cache parameters.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-response_cache--response_cache_parameters--cache_cidr_ipv6", "enforcement": "provider-schema", "group": "response_cache.response_cache_parameters:RequiredObjectAttributes:cache_cidr_ipv6", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:response_cache:response_cache_parameters", "type": "requires"}], "schema_path": ["response_cache", "response_cache_parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/response_cache/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Response Cache x-required.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- response_cache

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for response cache.

Additional upstream details:

Response Cache x-required.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "response_cache_parameters"),
  validators.ConflictingObjectAttributes("disable_spec",
    "response_cache_parameters")}
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
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

Terraform syntax:

```terraform
response_cache {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/): complete subsection reference.

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/disable_spec/): complete subsection reference.

- [response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/response_cache/response_cache_parameters/): complete subsection reference.
