---
page_title: "proxy_advertisement"
subcategory: ""
description: "Proxy Advertisement Type."
xcsh_docs: {"aliases": ["proxy advertisement"], "body_bytes": 1332, "body_sha256": "sha256:fb4c67acde1f7859ce00d2ea057306ec0a5f965b13f93559113ea97c42506106", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:do_not_advertise"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement", "parent_id": "xcsh-docs:resources:bigip_http_proxy:reference", "path": "documentation/resources/bigip_http_proxy/properties/proxy_advertisement/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise custom"], "anchor": "section", "description": "This defines a way to advertise a VIP on specific sites.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom"], "syntax": "block", "type": "object"}, {"aliases": ["proxy advertisement do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement:do_not_advertise", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "do_not_advertise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_advertisement/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Proxy Advertisement Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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

Additional upstream details:

Proxy Advertisement Type.

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
