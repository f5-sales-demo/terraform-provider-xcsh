---
page_title: "origin_servers.origin_servers"
subcategory: ""
description: "List of origin servers for Proxy."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers origin servers", "upstream servers"], "body_bytes": 4269, "body_sha256": "sha256:3b18df531dbe870e1e1bff93271aaf6c278980961ba7191b8883e66241cce5e3", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_name", "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers", "parent_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers", "path": "documentation/resources/dns_proxy/properties/origin_servers/origin_servers/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0001203120202033-1333333302120132-2213012031011303-0100332112121000-0001000121203032-0031302320212123-0222211212302120-3330033332203222", "registry_path": "docs/guides/resources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "k8s service", "origin servers", "upstream servers"], "anchor": "section", "description": "Specify origin server with K8s service name and site information.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.k8s_service:ConflictingObjectAttributes:outside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:vk8s_networks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.k8s_service:ConflictingObjectAttributes:outside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:vk8s_networks", "type": "conflicts"}], "schema_path": ["origin_servers", "origin_servers", "k8s_service"], "syntax": "block", "type": "object"}, {"aliases": ["no preference"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:no_preference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "no_preference"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "public ip", "upstream servers"], "anchor": "section", "description": "Specify origin server with public IP address.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "public_ip"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "origin servers", "public name", "upstream servers"], "anchor": "section", "description": "Specify origin server with public DNS name.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_name", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_servers--origin_servers--public_name--dns_name", "enforcement": "provider-schema", "group": "origin_servers.origin_servers.public_name:RequiredObjectAttributes:dns_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:public_name", "type": "requires"}], "schema_path": ["origin_servers", "origin_servers", "public_name"], "syntax": "block", "type": "object"}, {"aliases": ["site preferences"], "anchor": "section", "description": "Carries the references to one or more sites.", "document_id": "xcsh-docs:resources:dns_proxy:properties:origin_servers:origin_servers:site_preferences", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "site_preferences"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/origin_servers/origin_servers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of origin servers for Proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- origin_servers.origin_servers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of origin servers for Proxy.

Upstream description:

List of origin servers for Proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("k8s_service",
    "public_ip"),
  validators.ConflictingListObjectAttributes("k8s_service",
    "public_name"),
  validators.ConflictingListObjectAttributes("no_preference",
    "site_preferences"),
  validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

## Direct properties

- [k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/): complete subsection reference.

- [no_preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/no_preference/): complete subsection reference.

- [public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/public_ip/): complete subsection reference.

- [public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/public_name/): complete subsection reference.

- [site_preferences](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/)
- [origin_servers.origin_servers.no_preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/no_preference/)
- [origin_servers.origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/public_ip/)
- [origin_servers.origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/public_name/)
- [origin_servers.origin_servers.site_preferences](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/origin_servers/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
