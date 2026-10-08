---
page_title: "virtual_server"
subcategory: ""
description: "Specifies configuration related to virtual server."
xcsh_docs: {"aliases": ["virtual server"], "body_bytes": 8451, "body_sha256": "sha256:001fde46b4e4da968d537d70dafee2d904a8243dbb45a6acaad1d7967137b317", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:access_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:clone_pool_client", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:clone_pool_server", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:default_persistence_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:default_pool", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:fallback_persistence_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:fix_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:https", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:last_hop_pool", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:nat64", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:request_logging_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:statistics_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "parent_id": "xcsh-docs:data-sources:application_profiles:reference", "path": "documentation/data-sources/application_profiles/properties/virtual_server/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["virtual_server"], "schema_version": 1, "sections": [{"aliases": ["virtual server access profile"], "anchor": "section", "description": "Specifies an access policy that determines the authentication rules and access controls applied to user sessions for this virtual server.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:access_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "access_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server address translation"], "anchor": "section", "description": "Specifies, when checked (enabled), that the system translates the address of the virtual server. When cleared (disabled), specifies that the system uses the address without translation. This option is useful when the system is load balancing devices that have the same IP address. The default is enabled.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "address_translation"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server auto last hop"], "anchor": "section", "description": "When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "auto_last_hop"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server clone pool client"], "anchor": "section", "description": "Replicates client-side traffic (that is, prior to address translation) to a member of the specified pool.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:clone_pool_client", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "clone_pool_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server clone pool server"], "anchor": "section", "description": "Replicates server-side traffic (that is, prior to address translation) to a member of the specified pool.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:clone_pool_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "clone_pool_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server connection limit"], "anchor": "schema-virtual_server--connection_limit", "description": "Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this to 0 turns off connection limits. The default is 0.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_limit"], "syntax": "attribute", "type": "number"}, {"aliases": ["virtual server connection rate limit"], "anchor": "schema-virtual_server--connection_rate_limit", "description": "Specifies the maximum number of connections-per-second allowed for a virtual server. When the number of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where connection requests", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit"], "syntax": "attribute", "type": "number"}, {"aliases": ["virtual server connection rate limit mode"], "anchor": "section", "description": "Configuration parameter for connection rate limit mode.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "connection_rate_limit_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server default persistence profile"], "anchor": "section", "description": "Configuration parameter for default persistence profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:default_persistence_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "default_persistence_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server default pool"], "anchor": "section", "description": "Specifies the pool name that you want the virtual server to use as the default pool. A load balancing virtual server sends traffic to this pool automatically, unless an iRule directs the server to send the traffic to another pool instead.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:default_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "default_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server fallback persistence profile"], "anchor": "section", "description": "Configuration parameter for fallback persistence profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:fallback_persistence_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "fallback_persistence_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server fix profile"], "anchor": "section", "description": "Configuration parameter for fix profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:fix_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "fix_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server http"], "anchor": "section", "description": "HTTP profiles.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "http"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server http3"], "anchor": "section", "description": "HTTP/3 profiles.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "http3"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server https"], "anchor": "section", "description": "HTTP profiles.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:https", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "https"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server immediate action on service down"], "anchor": "section", "description": "Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial client's SYN packet, if the availability status of the virtual server is Offline or Unavailable. This is supported for the virtual server of Standard type and TCP protocol. The default is None. None: Specifies that the", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "immediate_action_on_service_down"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server last hop pool"], "anchor": "section", "description": "Directs reply traffic to the last hop router using the specified pool.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:last_hop_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "last_hop_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server nat64"], "anchor": "section", "description": "When enabled, allows the system to send return traffic to the MAC address that transmitted the request, even if the routing table points to a different network or interface. As a result, the system can send return traffic to clients even when there is no matching route. For example, if the system does not have a", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:nat64", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "nat64"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server port translation"], "anchor": "section", "description": "Specifies, when checked (enabled), that the system translates the port of the virtual server. When cleared (disabled), specifies that the system uses the port without translation. Turning off port translation for a virtual server is useful if you want to use the virtual server to load balance connections to any", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "port_translation"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server request logging profile"], "anchor": "section", "description": "Configuration parameter for request logging profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:request_logging_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "request_logging_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server source port"], "anchor": "section", "description": "Specifies whether the system preserves the source port of the connection. The default is Preserve.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "source_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server statistics profile"], "anchor": "section", "description": "Configuration parameter for statistics profile", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:statistics_profile", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["virtual_server", "statistics_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server tcp"], "anchor": "section", "description": "TCP profiles.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "tcp"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server udp"], "anchor": "section", "description": "UDP profiles.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "udp"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server virtual server state"], "anchor": "section", "description": "Displays the current state on the object.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server", "virtual_server_state"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server vs score"], "anchor": "schema-virtual_server--vs_score", "description": "Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value to load balance traffic in a proportional manner. The default is 0, meaning that no additional metric is applied for the virtual server.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server", "vs_score"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Specifies configuration related to virtual server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["application_profilesCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/)
- virtual_server

<a id="section"></a>

Type: `"single"`. Computed.

Specifies configuration related to virtual server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_server_type": "[\"http\",\"http3\",\"https\",\"tcp\",\"udp\"]"
}
```

## Direct properties

- [access_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/access_profile/): complete subsection reference.

- [address_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/): complete subsection reference.

- [auto_last_hop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/): complete subsection reference.

- [clone_pool_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_client/): complete subsection reference.

- [clone_pool_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_server/): complete subsection reference.

<a id="schema-virtual_server--connection_limit"></a>

### connection_limit property

Type: `"number"`. Computed.

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The default is 0.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="schema-virtual_server--connection_rate_limit"></a>

### connection_rate_limit property

Type: `"number"`. Computed.

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where
connection requests flood a virtual server. Setting this to 0 turns off connection limits. The
default is 0.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/): complete subsection reference.

- [default_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_persistence_profile/): complete subsection reference.

- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_pool/): complete subsection reference.

- [fallback_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fallback_persistence_profile/): complete subsection reference.

- [fix_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fix_profile/): complete subsection reference.

- [http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/): complete subsection reference.

- [http3](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/): complete subsection reference.

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/): complete subsection reference.

- [immediate_action_on_service_down](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/): complete subsection reference.

- [last_hop_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/last_hop_pool/): complete subsection reference.

- [nat64](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/nat64/): complete subsection reference.

- [port_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/port_translation/): complete subsection reference.

- [request_logging_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/request_logging_profile/): complete subsection reference.

- [source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/): complete subsection reference.

- [statistics_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/statistics_profile/): complete subsection reference.

- [tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/): complete subsection reference.

- [udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/): complete subsection reference.

- [virtual_server_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/): complete subsection reference.

<a id="schema-virtual_server--vs_score"></a>

### vs_score property

Type: `"number"`. Computed.

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The , meaning that no additional metric is applied
for the virtual server.

Additional upstream details:

The default is 0, meaning that no additional metric is applied for the virtual server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```
