---
page_title: "routes.waf_type.app_firewall"
subcategory: ""
description: "A list of references to the app_firewall configuration objects."
xcsh_docs: {"aliases": ["routes waf type app firewall"], "body_bytes": 1846, "body_sha256": "sha256:7d4b6561d70c211c518bcfc53d6090fcbb3ade35ab62d68d00d1fe10431c26e5", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "parent_id": "xcsh-docs:resources:route:properties:routes:waf_type", "path": "documentation/resources/route/properties/routes/waf_type/app_firewall/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type.app_firewall:RequiredObjectAttributes:app_firewall", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "waf_type", "app_firewall"], "schema_version": 1, "sections": [{"aliases": ["app firewall"], "anchor": "section", "description": "References to an Application Firewall configuration object.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "waf_type", "app_firewall", "app_firewall"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/waf_type/app_firewall/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of references to the app_firewall configuration objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type.app_firewall

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/)
- routes.waf_type.app_firewall

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("app_firewall")}
```

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
app_firewall {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/): complete subsection reference.

## Next pages

- [routes.waf_type.app_firewall.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/)
- [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
