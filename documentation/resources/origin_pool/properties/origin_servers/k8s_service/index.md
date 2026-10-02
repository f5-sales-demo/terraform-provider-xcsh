---
page_title: "origin_servers.k8s_service"
subcategory: "Load Balancing"
description: "Specify origin server with K8s service name and site information."
xcsh_docs: {"aliases": ["origin servers k8s service"], "body_bytes": 5732, "body_sha256": "sha256:c4cc7cbb8149c9c728055bd375b576cf07bcfc5cfb58d92f6ec334b68f54c312", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:inside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:outside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator", "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:snat_pool", "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:vk8s_networks"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "documentation/resources/origin_pool/properties/origin_servers/k8s_service/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service:ConflictingObjectAttributes:outside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service:ConflictingObjectAttributes:inside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:vk8s_networks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service:ConflictingObjectAttributes:outside_network,vk8s_networks", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:vk8s_networks", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "k8s_service"], "schema_version": 1, "sections": [{"aliases": ["inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "k8s_service", "inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:outside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "k8s_service", "outside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol"], "anchor": "schema-origin_servers--k8s_service--protocol", "description": "Type of protocol - PROTOCOL_TCP: TCP - PROTOCOL_UDP: UDP.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "k8s_service", "protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["service name"], "anchor": "schema-origin_servers--k8s_service--service_name", "description": "Exclusive with K8s service name of the origin server will be listed, including the namespace and cluster-ID. For vK8s services, you need to enter a string with the format servicename.namespace:example-namespace\"frontend\", namespace is \"speedtest\" and cluster-ID is \"prod\", then you will enter \"frontend.speedtest:prod\".", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "k8s_service", "service_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["site locator"], "anchor": "section", "description": "This message defines a reference to a site or virtual site object.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:site_locator:virtual_site", "type": "conflicts"}], "schema_path": ["origin_servers", "k8s_service", "site_locator"], "syntax": "block", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:snat_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.k8s_service.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:snat_pool:snat_pool", "type": "conflicts"}], "schema_path": ["origin_servers", "k8s_service", "snat_pool"], "syntax": "block", "type": "object"}, {"aliases": ["vk8s networks"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service:vk8s_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "k8s_service", "vk8s_networks"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/k8s_service/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify origin server with K8s service name and site information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["origin_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.k8s_service

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- origin_servers.k8s_service

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

## Direct properties

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/inside_network/): complete subsection reference.

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/outside_network/): complete subsection reference.

<a id="schema-origin_servers--k8s_service--protocol"></a>

### protocol property

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-origin_servers--k8s_service--service_name"></a>

### service_name property

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/snat_pool/): complete subsection reference.

- [vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/vk8s_networks/): complete subsection reference.

## Next pages

- [origin_servers.k8s_service.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/inside_network/)
- [origin_servers.k8s_service.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/outside_network/)
- [origin_servers.k8s_service.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/)
- [origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/snat_pool/)
- [origin_servers.k8s_service.vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/vk8s_networks/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
