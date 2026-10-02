---
page_title: "bot_defense.policy.protected_app_endpoints.web_mobile"
subcategory: "Load Balancing"
description: "Web and Mobile traffic type."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints web mobile"], "body_bytes": 2440, "body_sha256": "sha256:3c4b35af0be1335493d07ac2ec4ac31468a93540294b747fa2259c68de75804e", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web_mobile", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/web_mobile/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1033303310333201-2322202210330231-1303323202223200-0121131311331201-2331312332001013-3100331211333112-3213330112110230-3230321011112100", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "web_mobile"], "schema_version": 1, "sections": [{"aliases": ["mobile identifier"], "anchor": "schema-bot_defense--policy--protected_app_endpoints--web_mobile--mobile_identifier", "description": "Mobile identifier type - HEADERS: Headers Headers.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:web_mobile", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "web_mobile", "mobile_identifier"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/web_mobile/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Web and Mobile traffic type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.web_mobile

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile traffic type. Web and Mobile traffic type.

Upstream description:

Web and Mobile traffic type.

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
web_mobile {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bot_defense--policy--protected_app_endpoints--web_mobile--mobile_identifier"></a>

### mobile_identifier property

Type: `"string"`. Optional.

\[Enum: HEADERS\] Mobile identifier type - HEADERS: Headers Headers. The only possible value is
\`HEADERS\`. Defaults to \`HEADERS\`.

Upstream description:

Mobile identifier type

&#8203;- HEADERS: Headers

Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HEADERS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HEADERS",
  "enum": [
    "HEADERS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
