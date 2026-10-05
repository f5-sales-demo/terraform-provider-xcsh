---
page_title: "cloudflare.protected_endpoints.mobile_client.continue"
subcategory: ""
description: "Continue mitigation action."
xcsh_docs: {"aliases": ["cloudflare protected endpoints mobile client continue"], "body_bytes": 2915, "body_sha256": "sha256:2feb612543dd657ea8771c090d81375c518d00cd76230f84ba10e253363248f1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:add_header", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:no_header"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client", "path": "documentation/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.mobile_client.continue:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:add_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudflare.protected_endpoints.mobile_client.continue:ConflictingObjectAttributes:add_header,no_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:no_header", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "continue"], "schema_version": 1, "sections": [{"aliases": ["cloudflare protected endpoints mobile client continue add header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:add_header", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "continue", "add_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudflare protected endpoints mobile client continue no header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:mobile_client:continue:no_header", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "protected_endpoints", "mobile_client", "continue", "no_header"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Continue mitigation action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.mobile_client.continue

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/)
- [cloudflare.protected_endpoints.mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/)
- cloudflare.protected_endpoints.mobile_client.continue

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/add_header/): complete subsection reference.

- [no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/no_header/): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.mobile_client.continue.add_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/add_header/)
- [cloudflare.protected_endpoints.mobile_client.continue.no_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/continue/no_header/)
- [cloudflare.protected_endpoints.mobile_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/protected_endpoints/mobile_client/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
