---
page_title: "origin_servers.origin_servers.k8s_service"
subcategory: ""
description: "Specify origin server with K8s service name and site information."
xcsh_docs: {"aliases": ["origin servers origin servers k8s service"], "body_bytes": 5592, "body_sha256": "sha256:ae7cef5bb788823e9fac352304905e4d5feb5ec82d1de8659f2cb01dc707fc81", "capabilities": ["dns", "load-balancing.backend-servers"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:inside_network", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:outside_network", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:vk8s_networks"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "parent_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers", "path": "documentation/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3013321012300130-2001201333112023-2331221100330200-3212311300200201-0022113122003202-1313310102332230-0201213331323330-2230122103332312", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "origin_servers", "k8s_service"], "schema_version": 1, "sections": [{"aliases": ["inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:inside_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:outside_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "outside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol"], "anchor": "schema-origin_servers--origin_servers--k8s_service--protocol", "description": "Type of protocol - PROTOCOL_TCP: TCP - PROTOCOL_UDP: UDP.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "protocol"], "syntax": "attribute", "type": "string"}, {"aliases": ["service name"], "anchor": "schema-origin_servers--origin_servers--k8s_service--service_name", "description": "Exclusive with K8s service name of the origin server will be listed, including the namespace and cluster-ID. For vK8s services, you need to enter a string with the format servicename.namespace:example-namespace\"frontend\", namespace is \"speedtest\" and cluster-ID is \"prod\", then you will enter \"frontend.speedtest:prod\".", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "service_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["site locator"], "anchor": "section", "description": "This message defines a reference to a site or virtual site object.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:site_locator", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "site_locator"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:snat_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["vk8s networks"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers:origin_servers:k8s_service:vk8s_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "origin_servers", "k8s_service", "vk8s_networks"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify origin server with K8s service name and site information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.origin_servers.k8s_service

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- origin_servers.origin_servers.k8s_service

<a id="section"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

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

## Direct properties

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/inside_network/): complete subsection reference.

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/outside_network/): complete subsection reference.

<a id="schema-origin_servers--origin_servers--k8s_service--protocol"></a>

### protocol property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

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

<a id="schema-origin_servers--origin_servers--k8s_service--service_name"></a>

### service_name property

Type: `"string"`. Computed.

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

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/): complete subsection reference.

- [vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/vk8s_networks/): complete subsection reference.

## Next pages

- [origin_servers.origin_servers.k8s_service.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/inside_network/)
- [origin_servers.origin_servers.k8s_service.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/outside_network/)
- [origin_servers.origin_servers.k8s_service.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/)
- [origin_servers.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/)
- [origin_servers.origin_servers.k8s_service.vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/vk8s_networks/)
- [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
