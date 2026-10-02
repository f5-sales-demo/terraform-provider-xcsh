---
page_title: "endpoint"
subcategory: "Security"
description: "Shape of the endpoint choices for a view."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 3562, "body_sha256": "sha256:1b9220de45777fa137f6a308ff0afa9d7d712b87b6dbfbcd8228da152e30378f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:properties:endpoint:any", "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:endpoint", "parent_id": "xcsh-docs:resources:network_policy:reference", "path": "documentation/resources/network_policy/properties/endpoint/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2230123131001221-0003020123323130-2210111031030032-0120322100010011-3122331012333110-1213102313210220-0223333202132233-0212103232303013", "registry_path": "docs/guides/resources--network_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,inside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,label_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,outside_endpoints", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:any,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:inside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:label_selector,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "endpoint:ConflictingObjectAttributes:outside_endpoints,prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "sections": [{"aliases": ["any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:any", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-endpoint--label_selector--expressions", "enforcement": "provider-schema", "group": "endpoint.label_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "type": "requires"}], "schema_path": ["endpoint", "label_selector"], "syntax": "block", "type": "object"}, {"aliases": ["outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint", "prefix_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Shape of the endpoint choices for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [endpoint.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/any/)
- [endpoint.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/inside_endpoints/)
- [endpoint.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/label_selector/)
- [endpoint.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/outside_endpoints/)
- [endpoint.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/prefix_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
