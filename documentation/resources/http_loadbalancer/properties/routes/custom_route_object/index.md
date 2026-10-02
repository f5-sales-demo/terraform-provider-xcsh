---
page_title: "routes.custom_route_object"
subcategory: "Load Balancing"
description: "A custom route uses a route object created outside of this view."
xcsh_docs: {"aliases": ["routes custom route object"], "body_bytes": 2665, "body_sha256": "sha256:8dce6657f4b74b97820ed58ba929a9484d0b2cb8667ca3b4df16bb4bd26ff849", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_disable", "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:route_ref"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "path": "documentation/resources/http_loadbalancer/properties/routes/custom_route_object/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.custom_route_object:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.custom_route_object:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom_route_object"], "schema_version": 1, "sections": [{"aliases": ["caching disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_disable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom_route_object", "caching_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["caching inherit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom_route_object", "caching_inherit"], "syntax": "attribute", "type": "object"}, {"aliases": ["route ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:route_ref", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--custom_route_object--route_ref--name", "enforcement": "provider-schema", "group": "routes.custom_route_object.route_ref:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:route_ref", "type": "requires"}], "schema_path": ["routes", "custom_route_object", "route_ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A custom route uses a route object created outside of this view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom_route_object

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- routes.custom_route_object

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

## Direct properties

- [caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/caching_disable/): complete subsection reference.

- [caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/caching_inherit/): complete subsection reference.

- [route_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/route_ref/): complete subsection reference.

## Next pages

- [routes.custom_route_object.caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/caching_disable/)
- [routes.custom_route_object.caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/caching_inherit/)
- [routes.custom_route_object.route_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/route_ref/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
