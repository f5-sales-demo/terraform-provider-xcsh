---
page_title: "proxy_advertisement"
subcategory: ""
description: "Proxy Advertisement Type."
xcsh_docs: {"aliases": ["proxy advertisement"], "body_bytes": 2101, "body_sha256": "sha256:81043b2855d760382afcc37c2e64e2a0f79584e3699ffb9bf19158d043dcd113", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:do_not_advertise"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "documentation/resources/bigip_http_proxy/properties/proxy_advertisement/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_advertisement:ConflictingObjectAttributes:advertise_custom,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_advertisement:ConflictingObjectAttributes:advertise_custom,do_not_advertise", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:do_not_advertise", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement"], "schema_version": 1, "sections": [{"aliases": ["advertise custom"], "anchor": "section", "description": "This defines a way to advertise a VIP on specific sites.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_advertisement.advertise_custom:RequiredObjectAttributes:advertise_where", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "type": "requires"}], "schema_path": ["proxy_advertisement", "advertise_custom"], "syntax": "block", "type": "object"}, {"aliases": ["do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:do_not_advertise", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "do_not_advertise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Proxy Advertisement Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- proxy_advertisement

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
proxy_advertisement {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/): complete subsection reference.

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/do_not_advertise/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/)
- [proxy_advertisement.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_advertisement/do_not_advertise/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
