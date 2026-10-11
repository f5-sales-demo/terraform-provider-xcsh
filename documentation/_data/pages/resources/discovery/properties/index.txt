---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_discovery."
xcsh_docs: {"aliases": ["discovery"], "body_bytes": 39175, "body_sha256": "sha256:2d1b8bee6e038803be5805d93421e7e303f18b70a2139bcd3645db4375732491", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul", "xcsh-docs:resources:discovery:properties:discovery_k8s", "xcsh-docs:resources:discovery:properties:no_cluster_id", "xcsh-docs:resources:discovery:properties:timeouts", "xcsh-docs:resources:discovery:properties:where"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:reference", "parent_id": "xcsh-docs:resources:discovery:fundamentals", "path": "documentation/resources/discovery/properties/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["cluster id"], "anchor": "schema-cluster_id", "description": "Exclusive with Specify identifier for discovery cluster. This identifier can be specified in endpoint object to discover only from this discovery object.", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["discovery consul"], "anchor": "section", "description": "Discovery configuration for Hashicorp Consul.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul"], "syntax": "block", "type": "object"}, {"aliases": ["discovery k8s"], "anchor": "section", "description": "Discovery configuration for K8s.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s:ConflictingObjectAttributes:default_all,namespace_mapping", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:default_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s:ConflictingObjectAttributes:default_all,namespace_mapping", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "type": "conflicts"}], "schema_path": ["discovery_k8s"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:discovery:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["no cluster id"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:no_cluster_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_cluster_id"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:discovery:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["where"], "anchor": "section", "description": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "document_id": "xcsh-docs:resources:discovery:properties:where", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:where:virtual_site", "type": "conflicts"}], "schema_path": ["where"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-cluster_id"></a>

### cluster_id property

Type: `"string"`. Optional, Computed.

\[OneOf: cluster\_id, no\_cluster\_id; Default: no\_cluster\_id\] Exclusive with \[no\_cluster\_id\]
Specify identifier for discovery cluster. This identifier can be specified in endpoint object to
discover only from this discovery object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

OneOf alternatives in this subsection:

- [cluster_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-cluster_id)
- [no_cluster_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/no_cluster_id/#section)

Select alternatives according to the provider validators above.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/): complete subsection reference.

- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Discovery. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Discovery is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_cluster_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/no_cluster_id/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/timeouts/): complete subsection reference.

- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-annotations) |
| `cluster_id` | [cluster_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-cluster_id) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-disable) |
| `discovery_consul` | [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/#section) |
| `discovery_consul.access_info` | [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/#section) |
| `discovery_consul.access_info.connection_info` | [discovery_consul.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/#section) |
| `discovery_consul.access_info.connection_info.api_server` | [discovery_consul.access_info.connection_info.api_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/#schema-discovery_consul--access_info--connection_info--api_server) |
| `discovery_consul.access_info.connection_info.tls_info` | [discovery_consul.access_info.connection_info.tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/#section) |
| `discovery_consul.access_info.connection_info.tls_info.certificate` | [discovery_consul.access_info.connection_info.tls_info.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/#schema-discovery_consul--access_info--connection_info--tls_info--certificate) |
| `discovery_consul.access_info.connection_info.tls_info.key_url` | [discovery_consul.access_info.connection_info.tls_info.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/#section) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#section) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#schema-discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info--decryption_provider) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#schema-discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info--location) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#schema-discovery_consul--access_info--connection_info--tls_info--key_url--blindfold_secret_info--store_provider) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/clear_secret_info/#section) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/clear_secret_info/#schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--provider_ref) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/key_url/clear_secret_info/#schema-discovery_consul--access_info--connection_info--tls_info--key_url--clear_secret_info--url) |
| `discovery_consul.access_info.connection_info.tls_info.server_name` | [discovery_consul.access_info.connection_info.tls_info.server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/#schema-discovery_consul--access_info--connection_info--tls_info--server_name) |
| `discovery_consul.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_consul.access_info.connection_info.tls_info.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/connection_info/tls_info/#schema-discovery_consul--access_info--connection_info--tls_info--trusted_ca_url) |
| `discovery_consul.access_info.http_basic_auth_info` | [discovery_consul.access_info.http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/blindfold_secret_info/#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/blindfold_secret_info/#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info--decryption_provider) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/blindfold_secret_info/#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info--location) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/blindfold_secret_info/#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--blindfold_secret_info--store_provider) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/clear_secret_info/#section) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/clear_secret_info/#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info--provider_ref) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/passwd_url/clear_secret_info/#schema-discovery_consul--access_info--http_basic_auth_info--passwd_url--clear_secret_info--url) |
| `discovery_consul.access_info.http_basic_auth_info.user_name` | [discovery_consul.access_info.http_basic_auth_info.user_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/#schema-discovery_consul--access_info--http_basic_auth_info--user_name) |
| `discovery_consul.publish_info` | [discovery_consul.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/publish_info/#section) |
| `discovery_consul.publish_info.disable_spec` | [discovery_consul.publish_info.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/publish_info/disable_spec/#section) |
| `discovery_consul.publish_info.publish` | [discovery_consul.publish_info.publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/publish_info/publish/#section) |
| `discovery_k8s` | [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/#section) |
| `discovery_k8s.access_info` | [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/#section) |
| `discovery_k8s.access_info.connection_info` | [discovery_k8s.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/#section) |
| `discovery_k8s.access_info.connection_info.api_server` | [discovery_k8s.access_info.connection_info.api_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/#schema-discovery_k8s--access_info--connection_info--api_server) |
| `discovery_k8s.access_info.connection_info.tls_info` | [discovery_k8s.access_info.connection_info.tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/#section) |
| `discovery_k8s.access_info.connection_info.tls_info.certificate` | [discovery_k8s.access_info.connection_info.tls_info.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/#schema-discovery_k8s--access_info--connection_info--tls_info--certificate) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url` | [discovery_k8s.access_info.connection_info.tls_info.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/#section) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#section) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info--decryption_provider) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info--location) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/blindfold_secret_info/#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--blindfold_secret_info--store_provider) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/clear_secret_info/#section) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/clear_secret_info/#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--clear_secret_info--provider_ref) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/clear_secret_info/#schema-discovery_k8s--access_info--connection_info--tls_info--key_url--clear_secret_info--url) |
| `discovery_k8s.access_info.connection_info.tls_info.server_name` | [discovery_k8s.access_info.connection_info.tls_info.server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/#schema-discovery_k8s--access_info--connection_info--tls_info--server_name) |
| `discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/#schema-discovery_k8s--access_info--connection_info--tls_info--trusted_ca_url) |
| `discovery_k8s.access_info.isolated` | [discovery_k8s.access_info.isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/isolated/#section) |
| `discovery_k8s.access_info.kubeconfig_url` | [discovery_k8s.access_info.kubeconfig_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/#section) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/blindfold_secret_info/#section) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/blindfold_secret_info/#schema-discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info--decryption_provider) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/blindfold_secret_info/#schema-discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info--location) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/blindfold_secret_info/#schema-discovery_k8s--access_info--kubeconfig_url--blindfold_secret_info--store_provider) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/clear_secret_info/#section) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/clear_secret_info/#schema-discovery_k8s--access_info--kubeconfig_url--clear_secret_info--provider_ref) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/clear_secret_info/#schema-discovery_k8s--access_info--kubeconfig_url--clear_secret_info--url) |
| `discovery_k8s.access_info.reachable` | [discovery_k8s.access_info.reachable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/reachable/#section) |
| `discovery_k8s.default_all` | [discovery_k8s.default_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/default_all/#section) |
| `discovery_k8s.namespace_mapping` | [discovery_k8s.namespace_mapping](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/#section) |
| `discovery_k8s.namespace_mapping.items` | [discovery_k8s.namespace_mapping.items](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/items/#section) |
| `discovery_k8s.namespace_mapping.items.namespace` | [discovery_k8s.namespace_mapping.items.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/items/#schema-discovery_k8s--namespace_mapping--items--namespace) |
| `discovery_k8s.namespace_mapping.items.namespace_regex` | [discovery_k8s.namespace_mapping.items.namespace_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/items/#schema-discovery_k8s--namespace_mapping--items--namespace_regex) |
| `discovery_k8s.publish_info` | [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/#section) |
| `discovery_k8s.publish_info.disable_spec` | [discovery_k8s.publish_info.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/disable_spec/#section) |
| `discovery_k8s.publish_info.dns_delegation` | [discovery_k8s.publish_info.dns_delegation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/#section) |
| `discovery_k8s.publish_info.dns_delegation.dns_mode` | [discovery_k8s.publish_info.dns_delegation.dns_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/#schema-discovery_k8s--publish_info--dns_delegation--dns_mode) |
| `discovery_k8s.publish_info.dns_delegation.subdomain` | [discovery_k8s.publish_info.dns_delegation.subdomain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/#schema-discovery_k8s--publish_info--dns_delegation--subdomain) |
| `discovery_k8s.publish_info.publish` | [discovery_k8s.publish_info.publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish/#section) |
| `discovery_k8s.publish_info.publish.namespace` | [discovery_k8s.publish_info.publish.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish/#schema-discovery_k8s--publish_info--publish--namespace) |
| `discovery_k8s.publish_info.publish_fqdns` | [discovery_k8s.publish_info.publish_fqdns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/#schema-namespace) |
| `no_cluster_id` | [no_cluster_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/no_cluster_id/#section) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/timeouts/#schema-timeouts--update) |
| `where` | [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/#section) |
| `where.site` | [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/#section) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/disable_internet_vip/#section) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/enable_internet_vip/#section) |
| `where.site.network_type` | [where.site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/#schema-where--site--network_type) |
| `where.site.ref` | [where.site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/ref/#section) |
| `where.site.ref.kind` | [where.site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/ref/#schema-where--site--ref--kind) |
| `where.site.ref.name` | [where.site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/ref/#schema-where--site--ref--name) |
| `where.site.ref.namespace` | [where.site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/ref/#schema-where--site--ref--namespace) |
| `where.site.ref.tenant` | [where.site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/ref/#schema-where--site--ref--tenant) |
| `where.site.ref.uid` | [where.site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/site/ref/#schema-where--site--ref--uid) |
| `where.virtual_network` | [where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/#section) |
| `where.virtual_network.ref` | [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/#section) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--kind) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--name) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--namespace) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--tenant) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_network/ref/#schema-where--virtual_network--ref--uid) |
| `where.virtual_site` | [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/#section) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/disable_internet_vip/#section) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/enable_internet_vip/#section) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/#schema-where--virtual_site--network_type) |
| `where.virtual_site.ref` | [where.virtual_site.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/ref/#section) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--kind) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--name) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--namespace) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--tenant) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/where/virtual_site/ref/#schema-where--virtual_site--ref--uid) |
