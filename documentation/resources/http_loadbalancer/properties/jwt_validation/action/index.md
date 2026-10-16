---
page_title: "jwt_validation.action"
subcategory: "Load Balancing"
description: "Action"
xcsh_docs: {"aliases": ["jwt validation action"], "body_bytes": 1298, "body_sha256": "sha256:947cff0b9e0608bc2661d148182297870f8e0c4a5c3355e17b3f613693cc910d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:block", "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:report"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation", "path": "documentation/resources/http_loadbalancer/properties/jwt_validation/action/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "action"], "schema_version": 1, "sections": [{"aliases": ["jwt validation action block"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "action", "block"], "syntax": "attribute", "type": "object"}, {"aliases": ["jwt validation action report"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:report", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jwt_validation", "action", "report"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/action/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Action", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.action

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/)
- jwt_validation.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action

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

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/action/block/): complete subsection reference.

- [report](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/jwt_validation/action/report/): complete subsection reference.
