---
page_title: "routes.match.incoming_port.no_port_match"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes match incoming port no port match"], "body_bytes": 1490, "body_sha256": "sha256:c7f9360680d5185dd5286dd7a03c113af6b5fab3deb43d4b62d411660588817d", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "parent_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "path": "documentation/resources/route/properties/routes/match/incoming_port/no_port_match/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3130033300012001-2302020302203032-0000302003213101-2332122331012333-0103130103322232-3223301101200100-0120231213231201-3003032220012030", "registry_path": "docs/guides/resources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "match", "incoming_port", "no_port_match"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/match/incoming_port/no_port_match/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.match.incoming_port.no_port_match

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/)
- [routes.match.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/)
- routes.match.incoming_port.no_port_match

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
no_port_match = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.match.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
