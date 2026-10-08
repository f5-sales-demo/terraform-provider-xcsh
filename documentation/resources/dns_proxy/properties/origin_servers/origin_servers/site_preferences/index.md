---
page_title: "origin_servers.origin_servers.site_preferences"
subcategory: ""
description: "Carries the references to one or more sites."
xcsh_docs: {"aliases": ["origin servers origin servers site preferences"], "body_bytes": 1299, "body_sha256": "sha256:cf6e682231e612f5a00caca8f87501b00f8c17efb10085b9533b823032dc6bf7", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "path": "documentation/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2312211011120111-3323213032322320-2331232323202001-1223333121003231-0131002122113102-2002010310202132-2002130313312330-3113200213003311", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "site_preferences"], "schema_version": 1, "sections": [{"aliases": ["origin servers origin servers site preferences refs"], "anchor": "section", "description": "Reference to one or more sites.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-origin_servers--origin_servers--site_preferences--refs--name", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.site_preferences.refs:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences:refs", "type": "requires"}], "schema_path": ["origin_servers", "origin_servers", "site_preferences", "refs"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Carries the references to one or more sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
