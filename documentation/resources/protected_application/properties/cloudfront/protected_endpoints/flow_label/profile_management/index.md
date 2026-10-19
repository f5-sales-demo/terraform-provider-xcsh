---
page_title: "cloudfront.protected_endpoints.flow_label.profile_management"
subcategory: ""
description: "Bot Defense Flow Label Profile Management Category."
xcsh_docs: {"aliases": ["cloudfront protected endpoints flow label profile management"], "body_bytes": 2423, "body_sha256": "sha256:9b9ef2726e1cef9f7d08715e427b63371581d0538038de22a644d1be8b632b30", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:create", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:update", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:view"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3211102101003312-3032211031321310-2211210232311300-1131013131331012-1303012230231101-0112030102132033-1131131132220101-1221202011233013", "registry_path": "docs/guides/resources--protected_application--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.flow_label.profile_management:ConflictingObjectAttributes:create,update", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:create", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.flow_label.profile_management:ConflictingObjectAttributes:create,view", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:create", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.flow_label.profile_management:ConflictingObjectAttributes:create,update", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:update", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.flow_label.profile_management:ConflictingObjectAttributes:update,view", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:update", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.flow_label.profile_management:ConflictingObjectAttributes:create,view", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:view", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.protected_endpoints.flow_label.profile_management:ConflictingObjectAttributes:update,view", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:view", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "profile_management"], "schema_version": 1, "sections": [{"aliases": ["cloudfront protected endpoints flow label profile management create"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:create", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "profile_management", "create"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints flow label profile management update"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:update", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "profile_management", "update"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront protected endpoints flow label profile management view"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:profile_management:view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "profile_management", "view"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Bot Defense Flow Label Profile Management Category.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.profile_management

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- cloudfront.protected_endpoints.flow_label.profile_management

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "update"),
  validators.ConflictingObjectAttributes("create",
    "view"),
  validators.ConflictingObjectAttributes("update",
    "view")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

## Direct properties

- [create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/create/): complete subsection reference.

- [update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/update/): complete subsection reference.

- [view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/profile_management/view/): complete subsection reference.
