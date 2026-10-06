---
page_title: "routes.waf_type"
subcategory: ""
description: "WAF instance will be pointing to an app_firewall object."
xcsh_docs: {"aliases": ["routes waf type"], "body_bytes": 1774, "body_sha256": "sha256:3b13699543a395043e661e2adaeb4a92342a8485c5a62caeb91cd7f1a76dd3c2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:waf_type", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/waf_type/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301", "registry_path": "docs/guides/resources--route--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "waf_type"], "schema_version": 1, "sections": [{"aliases": ["routes waf type app firewall"], "anchor": "section", "description": "A list of references to the app_firewall configuration objects.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type.app_firewall:RequiredObjectAttributes:app_firewall", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall:app_firewall", "type": "requires"}], "schema_path": ["routes", "waf_type", "app_firewall"], "syntax": "block", "type": "object"}, {"aliases": ["routes waf type disable waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "waf_type", "disable_waf"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes waf type inherit waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "waf_type", "inherit_waf"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/waf_type/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "WAF instance will be pointing to an app_firewall object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
