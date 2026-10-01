---
page_title: "virtual_server"
subcategory: ""
description: "virtual_server for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 10714, "body_sha256": "sha256:089145aa2fb952d88f52c968d44ee54c271d0468e06479f030dcf8b107d74009", "canonical_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:virtual_server:access_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:address_translation", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:auto_last_hop", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:clone_pool_client", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:clone_pool_server", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:default_persistence_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:default_pool", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:fallback_persistence_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:fix_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:http3", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:https", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:last_hop_pool", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:nat64", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:port_translation", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:request_logging_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:source_port", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:statistics_profile", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:tcp", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:udp", "xcsh-docs:data-sources:application_profiles:properties:virtual_server:virtual_server_state"], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "parent_id": "xcsh-docs:data-sources:application_profiles:reference", "path": "docs/guides/data-sources--application_profiles--properties--virtual_server.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["virtual_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/virtual_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Property reference](data-sources--application_profiles--reference.md)
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

- [access_profile](data-sources--application_profiles--properties--virtual_server--access_profile.md): complete subsection reference.

- [address_translation](data-sources--application_profiles--properties--virtual_server--address_translation.md): complete subsection reference.

- [auto_last_hop](data-sources--application_profiles--properties--virtual_server--auto_last_hop.md): complete subsection reference.

- [clone_pool_client](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md): complete subsection reference.

- [clone_pool_server](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md): complete subsection reference.

<a id="schema-virtual_server--connection_limit"></a>

### connection_limit property

Type: `"number"`. Computed.

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where..

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [connection_rate_limit_mode](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md): complete subsection reference.

- [default_persistence_profile](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md): complete subsection reference.

- [default_pool](data-sources--application_profiles--properties--virtual_server--default_pool.md): complete subsection reference.

- [fallback_persistence_profile](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md): complete subsection reference.

- [fix_profile](data-sources--application_profiles--properties--virtual_server--fix_profile.md): complete subsection reference.

- [http](data-sources--application_profiles--properties--virtual_server--http.md): complete subsection reference.

- [http3](data-sources--application_profiles--properties--virtual_server--http3.md): complete subsection reference.

- [https](data-sources--application_profiles--properties--virtual_server--https.md): complete subsection reference.

- [immediate_action_on_service_down](data-sources--application_profiles--properties--virtual_server--immediate_action_on_service_down.md): complete subsection reference.

- [last_hop_pool](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md): complete subsection reference.

- [nat64](data-sources--application_profiles--properties--virtual_server--nat64.md): complete subsection reference.

- [port_translation](data-sources--application_profiles--properties--virtual_server--port_translation.md): complete subsection reference.

- [request_logging_profile](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md): complete subsection reference.

- [source_port](data-sources--application_profiles--properties--virtual_server--source_port.md): complete subsection reference.

- [statistics_profile](data-sources--application_profiles--properties--virtual_server--statistics_profile.md): complete subsection reference.

- [tcp](data-sources--application_profiles--properties--virtual_server--tcp.md): complete subsection reference.

- [udp](data-sources--application_profiles--properties--virtual_server--udp.md): complete subsection reference.

- [virtual_server_state](data-sources--application_profiles--properties--virtual_server--virtual_server_state.md): complete subsection reference.

<a id="schema-virtual_server--vs_score"></a>

### vs_score property

Type: `"number"`. Computed.

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The , meaning that no additional metric is applied
for the virtual server.

Upstream description:

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The default is 0, meaning that no additional
metric is applied for the virtual server.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [virtual_server.access_profile](data-sources--application_profiles--properties--virtual_server--access_profile.md)
- [virtual_server.address_translation](data-sources--application_profiles--properties--virtual_server--address_translation.md)
- [virtual_server.auto_last_hop](data-sources--application_profiles--properties--virtual_server--auto_last_hop.md)
- [virtual_server.clone_pool_client](data-sources--application_profiles--properties--virtual_server--clone_pool_client.md)
- [virtual_server.clone_pool_server](data-sources--application_profiles--properties--virtual_server--clone_pool_server.md)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--properties--virtual_server--connection_rate_limit_mode.md)
- [virtual_server.default_persistence_profile](data-sources--application_profiles--properties--virtual_server--default_persistence_profile.md)
- [virtual_server.default_pool](data-sources--application_profiles--properties--virtual_server--default_pool.md)
- [virtual_server.fallback_persistence_profile](data-sources--application_profiles--properties--virtual_server--fallback_persistence_profile.md)
- [virtual_server.fix_profile](data-sources--application_profiles--properties--virtual_server--fix_profile.md)
- [virtual_server.http](data-sources--application_profiles--properties--virtual_server--http.md)
- [virtual_server.http3](data-sources--application_profiles--properties--virtual_server--http3.md)
- [virtual_server.https](data-sources--application_profiles--properties--virtual_server--https.md)
- [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--properties--virtual_server--immediate_action_on_service_down.md)
- [virtual_server.last_hop_pool](data-sources--application_profiles--properties--virtual_server--last_hop_pool.md)
- [virtual_server.nat64](data-sources--application_profiles--properties--virtual_server--nat64.md)
- [virtual_server.port_translation](data-sources--application_profiles--properties--virtual_server--port_translation.md)
- [virtual_server.request_logging_profile](data-sources--application_profiles--properties--virtual_server--request_logging_profile.md)
- [virtual_server.source_port](data-sources--application_profiles--properties--virtual_server--source_port.md)
- [virtual_server.statistics_profile](data-sources--application_profiles--properties--virtual_server--statistics_profile.md)
- [virtual_server.tcp](data-sources--application_profiles--properties--virtual_server--tcp.md)
- [virtual_server.udp](data-sources--application_profiles--properties--virtual_server--udp.md)
- [virtual_server.virtual_server_state](data-sources--application_profiles--properties--virtual_server--virtual_server_state.md)
- [Property reference](data-sources--application_profiles--reference.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
