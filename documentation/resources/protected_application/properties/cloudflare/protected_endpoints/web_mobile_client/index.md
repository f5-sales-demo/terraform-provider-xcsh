---
page_title: "cloudflare.protected_endpoints.web_mobile_client"
subcategory: ""
description: "Web and Mobile client configuration OPTIONS."
xcsh_docs: {"aliases": ["cloudflare protected endpoints web mobile client"], "body_bytes": 4268, "body_sha256": "sha256:cfc07be7969ed100581fa4aea9b05699e9ff13e3a4ba4c3568c2e9b148a1763b", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_mobile", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_web", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints", "path": "documentation/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_mobile,continue_mobile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_mobile", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,continue_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_mobile,continue_mobile", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,continue_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:continue_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:block_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client:ConflictingObjectAttributes:continue_web,redirect_web", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client"], "schema_version": 1, "sections": [{"aliases": ["block mobile"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_mobile", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "block_mobile"], "syntax": "block", "type": "object"}, {"aliases": ["block web"], "anchor": "section", "description": "Block Response.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:block_web", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "block_web"], "syntax": "block", "type": "object"}, {"aliases": ["continue mobile"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client.continue_mobile:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile:add_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client.continue_mobile:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile:no_header", "type": "conflicts"}], "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "continue_mobile"], "syntax": "block", "type": "object"}, {"aliases": ["continue web"], "anchor": "section", "description": "Continue mitigation action.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client.continue_web:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web:add_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client.continue_web:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_web:no_header", "type": "conflicts"}], "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "continue_web"], "syntax": "block", "type": "object"}, {"aliases": ["redirect web"], "anchor": "section", "description": "Redirect.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-cloudflare--protected_endpoints--web_mobile_client--redirect_web--location", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.web_mobile_client.redirect_web:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:redirect_web", "type": "requires"}], "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "redirect_web"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Web and Mobile client configuration OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_mobile_client

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/)
- cloudflare.protected_endpoints.web_mobile_client

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block_mobile",
    "continue_mobile"),
  validators.ConflictingObjectAttributes("block_web",
    "continue_web"),
  validators.ConflictingObjectAttributes("block_web",
    "redirect_web"),
  validators.ConflictingObjectAttributes("continue_web",
    "redirect_web")}
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
  "x-ves-oneof-field-mobile_mitigation": "[\"block_mobile\",\"continue_mobile\"]",
  "x-ves-oneof-field-web_mitigation": "[\"block_web\",\"continue_web\",\"redirect_web\"]"
}
```

Terraform syntax:

```terraform
web_mobile_client {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/block_mobile/): complete subsection reference.

- [block_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/block_web/): complete subsection reference.

- [continue_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/): complete subsection reference.

- [continue_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_web/): complete subsection reference.

- [redirect_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/redirect_web/): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_mobile_client.block_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/block_mobile/)
- [cloudflare.protected_endpoints.web_mobile_client.block_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/block_web/)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_web/)
- [cloudflare.protected_endpoints.web_mobile_client.redirect_web](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/redirect_web/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
