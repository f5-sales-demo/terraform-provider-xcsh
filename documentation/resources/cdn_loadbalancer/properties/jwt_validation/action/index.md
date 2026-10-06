---
page_title: "jwt_validation.action"
subcategory: "Load Balancing"
description: "Action"
xcsh_docs: {"aliases": ["jwt validation action"], "body_bytes": 1483, "body_sha256": "sha256:e3074d9a076eea4da8cec1e2105f07d184777931343ae3d7e925e6f40f245977", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:block", "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:report"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation", "path": "documentation/resources/cdn_loadbalancer/properties/jwt_validation/action/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-011.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.action:ConflictingObjectAttributes:block,report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "jwt_validation.action:ConflictingObjectAttributes:block,report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:report", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "action"], "schema_version": 1, "sections": [{"aliases": ["jwt validation action block"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "action", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt validation action report"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:report", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "action", "report"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/action/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Action", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.action

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/)
- jwt_validation.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/action/block/): complete subsection reference.

- [report](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/jwt_validation/action/report/): complete subsection reference.
