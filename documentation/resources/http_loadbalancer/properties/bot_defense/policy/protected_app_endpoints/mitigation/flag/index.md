---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation.flag"
subcategory: "Load Balancing"
description: "Flag mitigation action."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints mitigation flag"], "body_bytes": 3120, "body_sha256": "sha256:bc4a782ec85bbe685163b28ecdd4161265df71da502b351d64f17cabfcea31c4", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-013.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag:ConflictingObjectAttributes:append_headers,no_headers", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag:ConflictingObjectAttributes:append_headers,no_headers", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag"], "schema_version": 1, "sections": [{"aliases": ["append headers"], "anchor": "section", "description": "Append flag mitigation headers to forwarded request.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers--auto_type_header_name", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers:RequiredObjectAttributes:auto_type_header_name,inference_header_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "type": "requires"}, {"anchor": "schema-bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers--inference_header_name", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers:RequiredObjectAttributes:auto_type_header_name,inference_header_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "type": "requires"}], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag", "append_headers"], "syntax": "block", "type": "object"}, {"aliases": ["no headers"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag", "no_headers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Flag mitigation action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation.flag

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Flag Bot Mitigation Action. Flag mitigation action.

Upstream description:

Flag mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_headers",
    "no_headers")}
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
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

Terraform syntax:

```terraform
flag {
  # Configure direct properties listed below.
}
```

## Direct properties

- [append_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/append_headers/): complete subsection reference.

- [no_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/no_headers/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/append_headers/)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/no_headers/)
- [bot_defense.policy.protected_app_endpoints.mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
