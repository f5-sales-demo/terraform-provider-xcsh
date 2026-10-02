---
page_title: "origin_servers.origin_servers.site_preferences"
subcategory: ""
description: "Carries the references to one or more sites."
xcsh_docs: {"aliases": ["origin servers origin servers site preferences"], "body_bytes": 1762, "body_sha256": "sha256:f9fbf747a26935e5aa307fca94a94f6574a34fa0d6782c4f5120af129bd8e687", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "path": "documentation/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2312211011120111-3323213032322320-2331232323202001-1223333121003231-0131002122113102-2002010310202132-2002130313312330-3113200213003311", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "site_preferences"], "schema_version": 1, "sections": [{"aliases": ["refs"], "anchor": "section", "description": "Reference to one or more sites.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-origin_servers--origin_servers--site_preferences--refs--name", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.site_preferences.refs:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs", "type": "requires"}], "schema_path": ["origin_servers", "origin_servers", "site_preferences", "refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Carries the references to one or more sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.site_preferences

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/)
- origin_servers.origin_servers.site_preferences

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Carries the references to one or more sites.

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
site_preferences {
  # Configure direct properties listed below.
}
```

## Direct properties

- [refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.site_preferences.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
