---
page_title: "origin_servers.consul_service"
subcategory: "Load Balancing"
description: "Specify origin server with HashiCorp Consul service name and site information."
xcsh_docs: {"aliases": ["origin servers consul service"], "body_bytes": 2877, "body_sha256": "sha256:112062b2d62cc4ed7db43bd20a55ba87c724a7a952a03f793c6a68c3e5185626", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:inside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:outside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:site_locator", "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "documentation/resources/origin_pool/properties/origin_servers/consul_service/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:outside_network", "type": "conflicts"}, {"anchor": "schema-origin_servers--consul_service--service_name", "enforcement": "provider-schema", "group": "origin_servers.consul_service:RequiredObjectAttributes:service_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "consul_service"], "schema_version": 1, "sections": [{"aliases": ["origin servers consul service inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:inside_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "consul_service", "inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers consul service outside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:outside_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "consul_service", "outside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["origin servers consul service service name"], "anchor": "schema-origin_servers--consul_service--service_name", "description": "Consul service name of this origin server will be listed, including cluster-ID. The format is servicename:cluster-ID.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "consul_service", "service_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["origin servers consul service site locator"], "anchor": "section", "description": "This message defines a reference to a site or virtual site object.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:site_locator", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:site_locator:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:site_locator:virtual_site", "type": "conflicts"}], "schema_path": ["origin_servers", "consul_service", "site_locator"], "syntax": "block", "type": "object"}, {"aliases": ["origin servers consul service snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.consul_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service:snat_pool:snat_pool", "type": "conflicts"}], "schema_path": ["origin_servers", "consul_service", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/consul_service/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Specify origin server with HashiCorp Consul service name and site information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.consul_service

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- origin_servers.consul_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with HashiCorp Consul service name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

Terraform syntax:

```terraform
consul_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/inside_network/): complete subsection reference.

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/outside_network/): complete subsection reference.

<a id="schema-origin_servers--consul_service--service_name"></a>

### service_name property

Type: `"string"`. Optional.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/): complete subsection reference.
