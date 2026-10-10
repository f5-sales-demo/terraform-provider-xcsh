---
page_title: "origin_servers.origin_servers"
subcategory: ""
description: "List of origin servers for Proxy."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers origin servers", "upstream servers"], "body_bytes": 2564, "body_sha256": "sha256:7099c4ef31dec982bf1502a15e4c0dec31dc531f7e70b394c2dbc18c9aa00373", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_name", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1333131311220333-1231120200230111-1310003132330122-1302102303223232-1331131102023222-3003000012212311-2300033331021112-0011131102103213", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers"], "schema_version": 1, "sections": [{"aliases": ["origin servers origin servers k8s service"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers origin servers no preference"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "no_preference"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers origin servers public ip"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers origin servers public name"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:public_name", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "public_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers origin servers site preferences"], "anchor": "section", "description": "Carries the references to one or more sites.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "site_preferences"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of origin servers for Proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- origin_servers.origin_servers

<a id="section"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of origin servers for Proxy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/): complete subsection reference.

- [no_preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/no_preference/): complete subsection reference.

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_name/): complete subsection reference.

- [site_preferences](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/): complete subsection reference.
