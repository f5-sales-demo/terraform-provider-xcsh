---
page_title: "routes.waf_type"
subcategory: ""
description: "WAF instance will be pointing to an app_firewall object."
xcsh_docs: {"aliases": ["routes waf type"], "body_bytes": 2419, "body_sha256": "sha256:0825d6e921c40f0993bf9a9097bf5cd930224e03a80b17b0c8d8f8a429245bdf", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:waf_type", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/waf_type/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "waf_type"], "schema_version": 1, "sections": [{"aliases": ["app firewall"], "anchor": "section", "description": "A list of references to the app_firewall configuration objects.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type.app_firewall:RequiredObjectAttributes:app_firewall", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "type": "requires"}], "schema_path": ["routes", "waf_type", "app_firewall"], "syntax": "block", "type": "object"}, {"aliases": ["disable waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "waf_type", "disable_waf"], "syntax": "attribute", "type": "object"}, {"aliases": ["inherit waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "waf_type", "inherit_waf"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/waf_type/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "WAF instance will be pointing to an app_firewall object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.waf_type

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- routes.waf_type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

WAF instance will be pointing to an app\_firewall object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherit_waf"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherit_waf")}
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
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

Terraform syntax:

```terraform
waf_type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/): complete subsection reference.

- [disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/disable_waf/): complete subsection reference.

- [inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/inherit_waf/): complete subsection reference.

## Next pages

- [routes.waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/)
- [routes.waf_type.disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/disable_waf/)
- [routes.waf_type.inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/inherit_waf/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
