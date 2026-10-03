---
page_title: "advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "Configure subset OPTIONS for origin pool."
xcsh_docs: {"aliases": ["advanced options enable subsets"], "body_bytes": 3191, "body_sha256": "sha256:11841b58f71f3e68cb236ac0297db458d3169cab5a1b233ea355fa055847e8a2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:fail_request"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "documentation/resources/origin_pool/properties/advanced_options/enable_subsets/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0122100002013231-2001203221223231-3110322321022003-1321210303333102-0312031012103102-3213111211122031-3001333210101120-3103332213211103", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,default_subset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,default_subset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:default_subset,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:fail_request", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:default_subset,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:fail_request", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:RequiredObjectAttributes:endpoint_subsets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options", "enable_subsets"], "schema_version": 1, "sections": [{"aliases": ["advanced options enable subsets any endpoint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "any_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["advanced options enable subsets default subset"], "anchor": "section", "description": "Default Subset definition.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "default_subset"], "syntax": "block", "type": "object"}, {"aliases": ["advanced options enable subsets endpoint subsets"], "anchor": "section", "description": "List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset within the class.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-advanced_options--enable_subsets--endpoint_subsets--keys", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets.endpoint_subsets:RequiredListObjectAttributes:keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "type": "requires"}], "schema_path": ["advanced_options", "enable_subsets", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["advanced options enable subsets fail request"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:fail_request", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "enable_subsets", "fail_request"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure subset OPTIONS for origin pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.enable_subsets

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- advanced_options.enable_subsets

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure subset OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("endpoint_subsets"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "default_subset"),
  validators.ConflictingObjectAttributes("any_endpoint",
    "fail_request"),
  validators.ConflictingObjectAttributes("default_subset",
    "fail_request")}
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
  "x-ves-oneof-field-fallback_policy_choice": "[\"any_endpoint\",\"default_subset\",\"fail_request\"]"
}
```

Terraform syntax:

```terraform
enable_subsets {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/any_endpoint/): complete subsection reference.

- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/default_subset/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/endpoint_subsets/): complete subsection reference.

- [fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/fail_request/): complete subsection reference.

## Next pages

- [advanced_options.enable_subsets.any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/any_endpoint/)
- [advanced_options.enable_subsets.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/default_subset/)
- [advanced_options.enable_subsets.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/endpoint_subsets/)
- [advanced_options.enable_subsets.fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/fail_request/)
- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
