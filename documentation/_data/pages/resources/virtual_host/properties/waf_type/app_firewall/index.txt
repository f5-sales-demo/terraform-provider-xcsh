---
page_title: "waf_type.app_firewall"
subcategory: ""
description: "A list of references to the app_firewall configuration objects."
xcsh_docs: {"aliases": ["waf type app firewall"], "body_bytes": 1743, "body_sha256": "sha256:be0b952502cc25d9b524e152cc3eca7c92f6ad8792e5b21c0bd89de5468efd0c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall", "parent_id": "xcsh-docs:resources:virtual_host:properties:waf_type", "path": "documentation/resources/virtual_host/properties/waf_type/app_firewall/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_type.app_firewall:RequiredObjectAttributes:app_firewall", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_type", "app_firewall"], "schema_version": 1, "sections": [{"aliases": ["waf type app firewall app firewall"], "anchor": "section", "description": "References to an Application Firewall configuration object.", "document_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall:app_firewall", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["waf_type", "app_firewall", "app_firewall"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/waf_type/app_firewall/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A list of references to the app_firewall configuration objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_type.app_firewall

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/)
- waf_type.app_firewall

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

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/): complete subsection reference.

## Next pages

- [waf_type.app_firewall.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
