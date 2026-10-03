---
page_title: "origin_servers"
subcategory: ""
description: "List of origin Servers for the DNS proxy."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "upstream servers"], "body_bytes": 1793, "body_sha256": "sha256:a4f28f65f10f4e3cdb7e574c04dfa855f46a4c967fb18fe0b15bdb6c934a86c0", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers", "parent_id": "xcsh-docs:resources:dns_proxy:reference", "path": "documentation/resources/dns_proxy/properties/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2110031213123010-1131131232222132-0120033311220002-3003331331202213-2211032102302033-0210000331223310-3131302232100111-2330320211332111", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:RequiredObjectAttributes:origin_servers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers"], "schema_version": 1, "sections": [{"aliases": ["origin servers health checks"], "anchor": "section", "description": "Origin Server Health Checks.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_servers--health_checks--healthy_threshold", "enforcement": "provider-schema", "group": "origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "type": "requires"}, {"anchor": "schema-origin_servers--health_checks--interval", "enforcement": "provider-schema", "group": "origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "type": "requires"}, {"anchor": "schema-origin_servers--health_checks--timeout", "enforcement": "provider-schema", "group": "origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "type": "requires"}, {"anchor": "schema-origin_servers--health_checks--unhealthy_threshold", "enforcement": "provider-schema", "group": "origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:health_checks:health_check", "type": "requires"}], "schema_path": ["origin_servers", "health_checks"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "origin servers", "origin servers origin servers", "upstream servers"], "anchor": "section", "description": "List of origin servers for Proxy.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:k8s_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:k8s_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:no_preference,site_preferences", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:k8s_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:k8s_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers:ConflictingListObjectAttributes:no_preference,site_preferences", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "type": "conflicts"}], "schema_path": ["origin_servers", "origin_servers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of origin Servers for the DNS proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- origin_servers

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of origin Servers for the DNS proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers")}
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
origin_servers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/): complete subsection reference.

- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/): complete subsection reference.

## Next pages

- [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/health_checks/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
