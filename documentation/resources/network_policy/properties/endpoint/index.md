---
page_title: "endpoint"
subcategory: "Security"
description: "Shape of the endpoint choices for a view."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 2640, "body_sha256": "sha256:d58a9d4acfc520188f52305a39b8459173bdebcfac29264ecdac26a998251ade", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:properties:endpoint:any", "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:endpoint", "parent_id": "xcsh-docs:resources:network_policy:reference", "path": "documentation/resources/network_policy/properties/endpoint/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "sections": [{"aliases": ["endpoint any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-endpoint--label_selector--expressions", "enforcement": "provider-schema", "group": "endpoint.label_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "requires"}], "schema_path": ["endpoint", "label_selector"], "syntax": "block", "type": "object"}, {"aliases": ["endpoint outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint", "prefix_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/endpoint/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Shape of the endpoint choices for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["network_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- endpoint

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Shape of the endpoint choices for a view.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingObjectAttributes("outside_endpoints",
    "prefix_list")}
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
  "x-ves-oneof-field-endpoint_choice": "[\"any\",\"inside_endpoints\",\"label_selector\",\"outside_endpoints\",\"prefix_list\"]"
}
```

Terraform syntax:

```terraform
endpoint {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/any/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/inside_endpoints/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/label_selector/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/prefix_list/): complete subsection reference.
