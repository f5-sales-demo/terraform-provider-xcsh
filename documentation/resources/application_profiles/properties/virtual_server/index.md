---
page_title: "virtual_server"
subcategory: ""
description: "virtual_server for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 14322, "body_sha256": "sha256:79bbe1589fce61a6e660aab33ca702f46cfecbb414c15a9c9acfaafe8e6b3008", "child_ids": ["xcsh-docs:resources:application_profiles:properties:virtual_server:access_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:address_translation", "xcsh-docs:resources:application_profiles:properties:virtual_server:auto_last_hop", "xcsh-docs:resources:application_profiles:properties:virtual_server:clone_pool_client", "xcsh-docs:resources:application_profiles:properties:virtual_server:clone_pool_server", "xcsh-docs:resources:application_profiles:properties:virtual_server:connection_rate_limit_mode", "xcsh-docs:resources:application_profiles:properties:virtual_server:default_persistence_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:default_pool", "xcsh-docs:resources:application_profiles:properties:virtual_server:fallback_persistence_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:fix_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:http", "xcsh-docs:resources:application_profiles:properties:virtual_server:http3", "xcsh-docs:resources:application_profiles:properties:virtual_server:https", "xcsh-docs:resources:application_profiles:properties:virtual_server:immediate_action_on_service_down", "xcsh-docs:resources:application_profiles:properties:virtual_server:last_hop_pool", "xcsh-docs:resources:application_profiles:properties:virtual_server:nat64", "xcsh-docs:resources:application_profiles:properties:virtual_server:port_translation", "xcsh-docs:resources:application_profiles:properties:virtual_server:request_logging_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:source_port", "xcsh-docs:resources:application_profiles:properties:virtual_server:statistics_profile", "xcsh-docs:resources:application_profiles:properties:virtual_server:tcp", "xcsh-docs:resources:application_profiles:properties:virtual_server:udp", "xcsh-docs:resources:application_profiles:properties:virtual_server:virtual_server_state"], "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:resources:application_profiles:properties:virtual_server", "parent_id": "xcsh-docs:resources:application_profiles:reference", "path": "documentation/resources/application_profiles/properties/virtual_server/index.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["virtual_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/properties/virtual_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "virtual_server for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# virtual_server

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- virtual_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration related to virtual server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http",
    "http3"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "tcp"),
  validators.ConflictingObjectAttributes("http",
    "udp"),
  validators.ConflictingObjectAttributes("http3",
    "https"),
  validators.ConflictingObjectAttributes("http3",
    "tcp"),
  validators.ConflictingObjectAttributes("http3",
    "udp"),
  validators.ConflictingObjectAttributes("https",
    "tcp"),
  validators.ConflictingObjectAttributes("https",
    "udp"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
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
  "x-ves-oneof-field-virtual_server_type": "[\"http\",\"http3\",\"https\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
virtual_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [access_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/access_profile/): complete subsection reference.

- [address_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/address_translation/): complete subsection reference.

- [auto_last_hop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/auto_last_hop/): complete subsection reference.

- [clone_pool_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/clone_pool_client/): complete subsection reference.

- [clone_pool_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/clone_pool_server/): complete subsection reference.

<a id="schema-virtual_server--connection_limit"></a>

### connection_limit property

Type: `"number"`. Optional.

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The.

Upstream description:

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The default is 0.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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

Type: `"number"`. Optional.

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where..

Upstream description:

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where
connection requests flood a virtual server. Setting this to 0 turns off connection limits. The
default is 0.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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

- [connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/): complete subsection reference.

- [default_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/default_persistence_profile/): complete subsection reference.

- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/default_pool/): complete subsection reference.

- [fallback_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/fallback_persistence_profile/): complete subsection reference.

- [fix_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/fix_profile/): complete subsection reference.

- [http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http/): complete subsection reference.

- [http3](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/): complete subsection reference.

- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/): complete subsection reference.

- [immediate_action_on_service_down](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/): complete subsection reference.

- [last_hop_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/last_hop_pool/): complete subsection reference.

- [nat64](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/nat64/): complete subsection reference.

- [port_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/): complete subsection reference.

- [request_logging_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/request_logging_profile/): complete subsection reference.

- [source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/): complete subsection reference.

- [statistics_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/statistics_profile/): complete subsection reference.

- [tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/): complete subsection reference.

- [udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/): complete subsection reference.

- [virtual_server_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/virtual_server_state/): complete subsection reference.

<a id="schema-virtual_server--vs_score"></a>

### vs_score property

Type: `"number"`. Optional.

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The , meaning that no additional metric is applied
for the virtual server.

Upstream description:

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The default is 0, meaning that no additional
metric is applied for the virtual server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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

- [virtual_server.access_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/access_profile/)
- [virtual_server.address_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/address_translation/)
- [virtual_server.auto_last_hop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/auto_last_hop/)
- [virtual_server.clone_pool_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/clone_pool_client/)
- [virtual_server.clone_pool_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/clone_pool_server/)
- [virtual_server.connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/connection_rate_limit_mode/)
- [virtual_server.default_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/default_persistence_profile/)
- [virtual_server.default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/default_pool/)
- [virtual_server.fallback_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/fallback_persistence_profile/)
- [virtual_server.fix_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/fix_profile/)
- [virtual_server.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http/)
- [virtual_server.http3](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/http3/)
- [virtual_server.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/https/)
- [virtual_server.immediate_action_on_service_down](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/immediate_action_on_service_down/)
- [virtual_server.last_hop_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/last_hop_pool/)
- [virtual_server.nat64](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/nat64/)
- [virtual_server.port_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/port_translation/)
- [virtual_server.request_logging_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/request_logging_profile/)
- [virtual_server.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/source_port/)
- [virtual_server.statistics_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/statistics_profile/)
- [virtual_server.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/tcp/)
- [virtual_server.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/udp/)
- [virtual_server.virtual_server_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/virtual_server/virtual_server_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/properties/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
