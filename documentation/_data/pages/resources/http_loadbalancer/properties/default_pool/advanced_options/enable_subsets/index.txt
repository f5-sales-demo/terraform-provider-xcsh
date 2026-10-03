---
page_title: "default_pool.advanced_options.enable_subsets"
subcategory: "Load Balancing"
description: "Configure subset OPTIONS for origin pool."
xcsh_docs: {"aliases": ["default pool advanced options enable subsets"], "body_bytes": 3642, "body_sha256": "sha256:f033571fc677d9afb5c64fe343da7be40d4cbafb5a1fad5ce70300b27e851eea", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:any_endpoint", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "path": "documentation/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3122011012203130-2032023233101333-0322112222003112-2311302231102332-2113303332212233-1102000111031122-0303332031001202-2221302030020230", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-015.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,default_subset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:any_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:any_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,default_subset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets:ConflictingObjectAttributes:default_subset,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets:ConflictingObjectAttributes:default_subset,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets:RequiredObjectAttributes:endpoint_subsets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "advanced_options", "enable_subsets"], "schema_version": 1, "sections": [{"aliases": ["default pool advanced options enable subsets any endpoint"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:any_endpoint", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "enable_subsets", "any_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["default pool advanced options enable subsets default subset"], "anchor": "section", "description": "Default Subset definition.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:default_subset", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_pool", "advanced_options", "enable_subsets", "default_subset"], "syntax": "block", "type": "object"}, {"aliases": ["default pool advanced options enable subsets endpoint subsets"], "anchor": "section", "description": "List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset within the class.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-default_pool--advanced_options--enable_subsets--endpoint_subsets--keys", "enforcement": "provider-schema", "group": "default_pool.advanced_options.enable_subsets.endpoint_subsets:RequiredListObjectAttributes:keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:endpoint_subsets", "type": "requires"}], "schema_path": ["default_pool", "advanced_options", "enable_subsets", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["default pool advanced options enable subsets fail request"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets:fail_request", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "advanced_options", "enable_subsets", "fail_request"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configure subset OPTIONS for origin pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options.enable_subsets

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- default_pool.advanced_options.enable_subsets

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

- [any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/any_endpoint/): complete subsection reference.

- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/endpoint_subsets/): complete subsection reference.

- [fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/fail_request/): complete subsection reference.

## Next pages

- [default_pool.advanced_options.enable_subsets.any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/any_endpoint/)
- [default_pool.advanced_options.enable_subsets.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/default_subset/)
- [default_pool.advanced_options.enable_subsets.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/endpoint_subsets/)
- [default_pool.advanced_options.enable_subsets.fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/enable_subsets/fail_request/)
- [default_pool.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/advanced_options/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
